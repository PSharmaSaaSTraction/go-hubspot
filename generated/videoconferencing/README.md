# Go API client for videoconferencing

These APIs allow you to specify URLs that can be used to interact with a video conferencing application, to allow HubSpot to add video conference links to meeting requests with contacts.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./videoconferencing"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SettingsApi* | [**Deletecrmv3extensionsvideoconferencingsettingsappIdArchive**](docs/SettingsApi.md#deletecrmv3extensionsvideoconferencingsettingsappidarchive) | **Delete** /crm/v3/extensions/videoconferencing/settings/{appId} | Delete settings
*SettingsApi* | [**Getcrmv3extensionsvideoconferencingsettingsappIdGetById**](docs/SettingsApi.md#getcrmv3extensionsvideoconferencingsettingsappidgetbyid) | **Get** /crm/v3/extensions/videoconferencing/settings/{appId} | Get settings
*SettingsApi* | [**Putcrmv3extensionsvideoconferencingsettingsappIdReplace**](docs/SettingsApi.md#putcrmv3extensionsvideoconferencingsettingsappidreplace) | **Put** /crm/v3/extensions/videoconferencing/settings/{appId} | Update settings

## Documentation For Models

 - [ErrorDetail](docs/ErrorDetail.md)
 - [ExternalSettings](docs/ExternalSettings.md)
 - [ModelError](docs/ModelError.md)

## Documentation For Authorization

## hapikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

