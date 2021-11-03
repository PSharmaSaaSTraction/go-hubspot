/*
Associations

Associations define the relationships between objects in HubSpot. These endpoints allow you to create, read, and remove associations.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package associations

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TypesApiService TypesApi service
type TypesApiService service

type ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest struct {
	ctx            _context.Context
	ApiService     *TypesApiService
	fromObjectType string
	toObjectType   string
}

func (r ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest) Execute() (CollectionResponsePublicAssociationDefiniton, *_nethttp.Response, error) {
	return r.ApiService.GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllExecute(r)
}

/*
GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll List association types

List all the valid association types available between two object types

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromObjectType
 @param toObjectType
 @return ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest
*/
func (a *TypesApiService) GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll(ctx _context.Context, fromObjectType string, toObjectType string) ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest {
	return ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicAssociationDefiniton
func (a *TypesApiService) GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllExecute(r ApiGetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAllRequest) (CollectionResponsePublicAssociationDefiniton, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionResponsePublicAssociationDefiniton
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TypesApiService.GetCrmV3AssociationsFromObjectTypeToObjectTypeTypesGetAll")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/associations/{fromObjectType}/{toObjectType}/types"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", _neturl.PathEscape(parameterToString(r.fromObjectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", _neturl.PathEscape(parameterToString(r.toObjectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}