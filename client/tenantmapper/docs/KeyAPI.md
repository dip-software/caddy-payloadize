# \KeyAPI

All URIs are relative to *http://localhost:8088*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeyCreate**](KeyAPI.md#KeyCreate) | **Post** /key | create key
[**KeyGet**](KeyAPI.md#KeyGet) | **Get** /key/{keyId} | get key
[**KeyRevoke**](KeyAPI.md#KeyRevoke) | **Put** /key/$revoke | revoke key
[**KeyVerify**](KeyAPI.md#KeyVerify) | **Post** /key/$verify | verify key



## KeyCreate

> GeneratedKey KeyCreate(ctx).XScopeOrgID(xScopeOrgID).XIdToken(xIdToken).CreateKeyPayload2(createKeyPayload2).Execute()

create key



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
	xScopeOrgID := "foo" // string | Name of the tenant
	xIdToken := "deadbeef" // string | ID token
	createKeyPayload2 := *openapiclient.NewCreateKeyPayload2("admin", "8760h", "k8s-secret") // CreateKeyPayload2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.KeyCreate(context.Background()).XScopeOrgID(xScopeOrgID).XIdToken(xIdToken).CreateKeyPayload2(createKeyPayload2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyCreate`: GeneratedKey
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.KeyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScopeOrgID** | **string** | Name of the tenant | 
 **xIdToken** | **string** | ID token | 
 **createKeyPayload2** | [**CreateKeyPayload2**](CreateKeyPayload2.md) |  | 

### Return type

[**GeneratedKey**](GeneratedKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyGet

> TenantAPIKey KeyGet(ctx, keyId).Execute()

get key



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
	keyId := "a87b5407-dd02-42af-ac60-80bdf13f9696" // string | Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.KeyGet(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeyGet`: TenantAPIKey
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.KeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantAPIKey**](TenantAPIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.goa.error

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeyRevoke

> KeyRevoke(ctx).RevokeKeyPayload(revokeKeyPayload).Execute()

revoke key



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
	revokeKeyPayload := *openapiclient.NewRevokeKeyPayload("deadbeef") // RevokeKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeyAPI.KeyRevoke(context.Background()).RevokeKeyPayload(revokeKeyPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeKeyPayload** | [**RevokeKeyPayload**](RevokeKeyPayload.md) |  | 

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


## KeyVerify

> KeyVerify(ctx).InputKeyPayload(inputKeyPayload).Execute()

verify key



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
	inputKeyPayload := *openapiclient.NewInputKeyPayload("deadbeef") // InputKeyPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeyAPI.KeyVerify(context.Background()).InputKeyPayload(inputKeyPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.KeyVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeyVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputKeyPayload** | [**InputKeyPayload**](InputKeyPayload.md) |  | 

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

