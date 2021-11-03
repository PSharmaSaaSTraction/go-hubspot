# Go API client for timeline

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./timeline"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EventsApi* | [**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdDetailGetDetailById**](docs/EventsApi.md#getintegratorstimelinev3eventseventtemplateideventiddetailgetdetailbyid) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
*EventsApi* | [**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdGetById**](docs/EventsApi.md#getintegratorstimelinev3eventseventtemplateideventidgetbyid) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
*EventsApi* | [**GetIntegratorsTimelineV3EventsEventTemplateIdEventIdRenderGetRenderById**](docs/EventsApi.md#getintegratorstimelinev3eventseventtemplateideventidrendergetrenderbyid) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML
*EventsApi* | [**PostIntegratorsTimelineV3EventsBatchCreateCreateBatch**](docs/EventsApi.md#postintegratorstimelinev3eventsbatchcreatecreatebatch) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events
*EventsApi* | [**PostIntegratorsTimelineV3EventsCreate**](docs/EventsApi.md#postintegratorstimelinev3eventscreate) | **Post** /crm/v3/timeline/events | Create a single event
*TemplatesApi* | [**DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdArchive**](docs/TemplatesApi.md#deleteintegratorstimelinev3appideventtemplateseventtemplateidarchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
*TemplatesApi* | [**GetIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdGetById**](docs/TemplatesApi.md#getintegratorstimelinev3appideventtemplateseventtemplateidgetbyid) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
*TemplatesApi* | [**GetIntegratorsTimelineV3AppIdEventTemplatesGetAll**](docs/TemplatesApi.md#getintegratorstimelinev3appideventtemplatesgetall) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
*TemplatesApi* | [**PostIntegratorsTimelineV3AppIdEventTemplatesCreate**](docs/TemplatesApi.md#postintegratorstimelinev3appideventtemplatescreate) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
*TemplatesApi* | [**PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdUpdate**](docs/TemplatesApi.md#putintegratorstimelinev3appideventtemplateseventtemplateidupdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template
*TokensApi* | [**DeleteIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameArchive**](docs/TokensApi.md#deleteintegratorstimelinev3appideventtemplateseventtemplateidtokenstokennamearchive) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Removes a token from the event template
*TokensApi* | [**PostIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensCreate**](docs/TokensApi.md#postintegratorstimelinev3appideventtemplateseventtemplateidtokenscreate) | **Post** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens | Adds a token to an existing event template
*TokensApi* | [**PutIntegratorsTimelineV3AppIdEventTemplatesEventTemplateIdTokensTokenNameUpdate**](docs/TokensApi.md#putintegratorstimelinev3appideventtemplateseventtemplateidtokenstokennameupdate) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Updates an existing token on an event template


## Documentation For Models

 - [BatchInputTimelineEvent](docs/BatchInputTimelineEvent.md)
 - [BatchResponseTimelineEventResponse](docs/BatchResponseTimelineEventResponse.md)
 - [BatchResponseTimelineEventResponseWithErrors](docs/BatchResponseTimelineEventResponseWithErrors.md)
 - [CollectionResponseTimelineEventTemplateNoPaging](docs/CollectionResponseTimelineEventTemplateNoPaging.md)
 - [Error](docs/Error.md)
 - [ErrorCategory](docs/ErrorCategory.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [EventDetail](docs/EventDetail.md)
 - [StandardError](docs/StandardError.md)
 - [TimelineEvent](docs/TimelineEvent.md)
 - [TimelineEventIFrame](docs/TimelineEventIFrame.md)
 - [TimelineEventResponse](docs/TimelineEventResponse.md)
 - [TimelineEventTemplate](docs/TimelineEventTemplate.md)
 - [TimelineEventTemplateCreateRequest](docs/TimelineEventTemplateCreateRequest.md)
 - [TimelineEventTemplateToken](docs/TimelineEventTemplateToken.md)
 - [TimelineEventTemplateTokenOption](docs/TimelineEventTemplateTokenOption.md)
 - [TimelineEventTemplateTokenUpdateRequest](docs/TimelineEventTemplateTokenUpdateRequest.md)
 - [TimelineEventTemplateUpdateRequest](docs/TimelineEventTemplateUpdateRequest.md)


## Documentation For Authorization



### developer_hapikey

- **Type**: API key
- **API key parameter name**: hapikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: hapikey and passed in as the auth context for each request.


### oauth2_legacy


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **contacts**: Read from and write to my Contacts
 - **tickets**: Read and write tickets
 - **timeline**: Create timeline events

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


