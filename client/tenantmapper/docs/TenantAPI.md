# \TenantAPI

All URIs are relative to *http://localhost:8088*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantBillingEvent**](TenantAPI.md#TenantBillingEvent) | **Post** /tenants/billing_event | billing_event tenant
[**TenantCreate**](TenantAPI.md#TenantCreate) | **Post** /tenants | create tenant
[**TenantDelete**](TenantAPI.md#TenantDelete) | **Delete** /tenants/{tenantId} | delete tenant
[**TenantGet**](TenantAPI.md#TenantGet) | **Get** /tenants/{tenantId} | get tenant
[**TenantList**](TenantAPI.md#TenantList) | **Get** /tenants | list tenant
[**TenantOnboard**](TenantAPI.md#TenantOnboard) | **Post** /tenants/onboard | onboard tenant
[**TenantSync**](TenantAPI.md#TenantSync) | **Post** /tenants/sync | sync tenant
[**TenantUpdate**](TenantAPI.md#TenantUpdate) | **Put** /tenants/{tenantId} | update tenant



## TenantBillingEvent

> TenantBillingEvent(ctx).XScopeOrgId(xScopeOrgId).InputBillingEvent2(inputBillingEvent2).Execute()

billing_event tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xScopeOrgId := "homelab" // string | Organization, also known as tenant
	inputBillingEvent2 := *openapiclient.NewInputBillingEvent2(int64(10), "loki", "2021-07-01T00:00:00Z", int64(12345)) // InputBillingEvent2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.TenantBillingEvent(context.Background()).XScopeOrgId(xScopeOrgId).InputBillingEvent2(inputBillingEvent2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantBillingEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScopeOrgId** | **string** | Organization, also known as tenant | 
 **inputBillingEvent2** | [**InputBillingEvent2**](InputBillingEvent2.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantCreate

> Tenant TenantCreate(ctx).CreateTenantPayload(createTenantPayload).Execute()

create tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createTenantPayload := *openapiclient.NewCreateTenantPayload("aac48e8e-78f0-4fab-962e-901113d8fcef", "foo") // CreateTenantPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TenantCreate(context.Background()).CreateTenantPayload(createTenantPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantCreate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TenantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTenantPayload** | [**CreateTenantPayload**](CreateTenantPayload.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantDelete

> TenantDelete(ctx, tenantId).Execute()

delete tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "18ee082f-1d61-40d3-b8a2-f4eee67cefff" // string | Tenant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.TenantDelete(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantGet

> Tenant TenantGet(ctx, tenantId).Execute()

get tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "18ee082f-1d61-40d3-b8a2-f4eee67cefff" // string | Tenant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TenantGet(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantGet`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TenantGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantList

> Tenantlist TenantList(ctx).Execute()

list tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TenantList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantList`: Tenantlist
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TenantList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantListRequest struct via the builder pattern


### Return type

[**Tenantlist**](Tenantlist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantOnboard

> Tenant TenantOnboard(ctx).XIdToken(xIdToken).OnboardTenantPayload2(onboardTenantPayload2).Execute()

onboard tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xIdToken := "deadbeef" // string | ID token
	onboardTenantPayload2 := *openapiclient.NewOnboardTenantPayload2("eu-west-1", "my-tenant") // OnboardTenantPayload2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TenantOnboard(context.Background()).XIdToken(xIdToken).OnboardTenantPayload2(onboardTenantPayload2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantOnboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantOnboard`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TenantOnboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantOnboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xIdToken** | **string** | ID token | 
 **onboardTenantPayload2** | [**OnboardTenantPayload2**](OnboardTenantPayload2.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantSync

> TenantSync(ctx).Execute()

sync tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.TenantSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantSyncRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantUpdate

> Tenant TenantUpdate(ctx, tenantId).UpdateTenantPayload2(updateTenantPayload2).Execute()

update tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "18ee082f-1d61-40d3-b8a2-f4eee67cefff" // string | Tenant ID
	updateTenantPayload2 := *openapiclient.NewUpdateTenantPayload2() // UpdateTenantPayload2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.TenantUpdate(context.Background(), tenantId).UpdateTenantPayload2(updateTenantPayload2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.TenantUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantUpdate`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.TenantUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTenantPayload2** | [**UpdateTenantPayload2**](UpdateTenantPayload2.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

