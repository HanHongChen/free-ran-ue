# GNBApi

All URIs are relative to *http://127.0.0.1:40104*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiConsoleGnbRegistrationPost**](#apiconsolegnbregistrationpost) | **POST** /api/console/gnb/registration | Register gNB|

# **apiConsoleGnbRegistrationPost**
> ApiConsoleGnbRegistrationPost200Response apiConsoleGnbRegistrationPost(apiConsoleGnbRegistrationPostRequest)

Register a new gNB with specified IP and port

### Example

```typescript
import {
    GNBApi,
    Configuration,
    ApiConsoleGnbRegistrationPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new GNBApi(configuration);

let apiConsoleGnbRegistrationPostRequest: ApiConsoleGnbRegistrationPostRequest; //

const { status, data } = await apiInstance.apiConsoleGnbRegistrationPost(
    apiConsoleGnbRegistrationPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiConsoleGnbRegistrationPostRequest** | **ApiConsoleGnbRegistrationPostRequest**|  | |


### Return type

**ApiConsoleGnbRegistrationPost200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Registration successful |  -  |
|**400** | Invalid request format |  -  |
|**401** | Authentication failed |  -  |
|**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

