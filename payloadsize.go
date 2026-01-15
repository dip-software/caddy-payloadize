package payloadsize

import (
	"context"
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/dip-software/monitizer"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"payloadsize/client/tenantmapper"
	"strconv"
	"time"
)

func init() {
	caddy.RegisterModule(&PayloadSize{})
	httpcaddyfile.RegisterHandlerDirective("payloadsize", parseCaddyfile)
}

type PayloadSize struct {
	MaxPayloadSize  int `json:"max_payload_size,omitempty"` // Optional: Maximum allowed payload size in bytes
	logger          *zap.Logger
	dataChan        chan monitizer.DataPoint
	cancelFunc      context.CancelFunc
	TenantMapperURL string `json:"tenant_mapper_url,omitempty"`
	tm              *tenantmapper.APIClient
}

// UnmarshalCaddyfile Caddyfile syntax:
//
//		payloadsize {
//			max_payload_size <size>
//	        tenant_mapper_url http://tenant-mapper:8088
//		}
func (j *PayloadSize) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "max_payload_size":
				if !d.NextArg() {
					return d.ArgErr()
				}
				j.MaxPayloadSize, _ = strconv.Atoi(d.Val())
			case "tenant_mapper_url":
				if !d.NextArg() {
					return d.ArgErr()
				}
				j.TenantMapperURL = d.Val()
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	return nil
}

func (j *PayloadSize) Validate() error {
	return nil
}

func (j *PayloadSize) Provision(context caddy.Context) error {
	var err error
	j.logger = context.Logger()
	j.logger.Info("provisioned payloadize middleware")
	j.dataChan = make(chan monitizer.DataPoint)
	j.cancelFunc, err = monitizer.Start(j.dataChan, j.storeEvents, time.Minute)
	if err != nil {
		return err
	}
	// Set up tenant mapper client
	if j.TenantMapperURL != "" {
		tmURL, err := url.Parse(j.TenantMapperURL)
		if err != nil {
			return fmt.Errorf("error parsing tenant mapper URL: %w", err)
		}

		tmcfg := tenantmapper.NewConfiguration()
		tmcfg.Host = tmURL.Host
		tmcfg.Scheme = tmURL.Scheme
		j.tm = tenantmapper.NewAPIClient(tmcfg)
	}
	return err
}

func (j *PayloadSize) storeEvents(data monitizer.AggregateDataPoint) error {
	msg := "per minute billing event"
	if j.tm != nil {
		source := "payloadsize"
		_, err := j.tm.TenantAPI.TenantBillingEvent(context.Background()).XScopeOrgId(data.Tenant).InputBillingEvent2(tenantmapper.InputBillingEvent2{
			Count:     data.Count,
			Value:     data.Sum,
			Service:   data.Service,
			Timestamp: time.Now().Format(time.RFC3339),
			Source:    &source,
		}).Execute()
		if err != nil {
			return fmt.Errorf("error storing billing event: %w", err)
		}
		msg = "pushed billing event to tenant mapper"
	}
	j.logger.Info(msg,
		zap.String("tenant", data.Tenant),
		zap.String("service", data.Service),
		zap.Int64("sum", data.Sum),
		zap.Int64("count", data.Count))
	return nil
}

// CaddyModule returns the Caddy module information
func (j *PayloadSize) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.payloadsize",
		New: func() caddy.Module { return new(PayloadSize) },
	}
}

func (j *PayloadSize) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	// Read content length from the request header
	cl := r.Header.Get("Content-Length")
	tenant := r.Header.Get("X-Scope-OrgID")
	if cl == "" { // If no CL then skip
		return next.ServeHTTP(w, r)
	}
	contentLength, err := strconv.Atoi(cl)
	if err != nil {
		return caddyhttp.Error(http.StatusBadRequest, fmt.Errorf("invalid content length: %v", err))
	}
	// Path detection
	path := r.URL.Path

	j.dataChan <- monitizer.DataPoint{
		Tenant: tenant + "|" + path,
		Value:  int64(contentLength),
	}

	// Proceed to the next handler
	return next.ServeHTTP(w, r)
}

// parseCaddyfile will unmarshal tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := &PayloadSize{}
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

var (
	_ caddy.Module                = (*PayloadSize)(nil)
	_ caddyhttp.MiddlewareHandler = (*PayloadSize)(nil)
	_ caddy.Provisioner           = (*PayloadSize)(nil)
	_ caddyfile.Unmarshaler       = (*PayloadSize)(nil)
	_ caddy.Validator             = (*PayloadSize)(nil)
)
