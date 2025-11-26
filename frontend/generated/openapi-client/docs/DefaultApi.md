# DefaultApi

All URIs are relative to *http://localhost:8080*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1AuthLoginPost**](DefaultApi.md#v1authloginpostoperation) | **POST** /v1/auth/login | Login with email + password |
| [**v1AuthRegisterPost**](DefaultApi.md#v1authregisterpostoperation) | **POST** /v1/auth/register | Login with email + password |



## v1AuthLoginPost

> User v1AuthLoginPost(v1AuthLoginPostRequest)

Login with email + password

### Example

```ts
import {
  Configuration,
  DefaultApi,
} from '';
import type { V1AuthLoginPostOperationRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DefaultApi();

  const body = {
    // V1AuthLoginPostRequest
    v1AuthLoginPostRequest: ...,
  } satisfies V1AuthLoginPostOperationRequest;

  try {
    const data = await api.v1AuthLoginPost(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **v1AuthLoginPostRequest** | [V1AuthLoginPostRequest](V1AuthLoginPostRequest.md) |  | |

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | return token and user information |  -  |
| **401** | Authentication failed or missing credentials |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## v1AuthRegisterPost

> V1AuthRegisterPost200Response v1AuthRegisterPost(v1AuthRegisterPostRequest)

Login with email + password

### Example

```ts
import {
  Configuration,
  DefaultApi,
} from '';
import type { V1AuthRegisterPostOperationRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new DefaultApi();

  const body = {
    // V1AuthRegisterPostRequest
    v1AuthRegisterPostRequest: ...,
  } satisfies V1AuthRegisterPostOperationRequest;

  try {
    const data = await api.v1AuthRegisterPost(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **v1AuthRegisterPostRequest** | [V1AuthRegisterPostRequest](V1AuthRegisterPostRequest.md) |  | |

### Return type

[**V1AuthRegisterPost200Response**](V1AuthRegisterPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | return register success message |  -  |
| **400** | User is authenticated but not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

