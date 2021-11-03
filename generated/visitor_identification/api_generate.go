/*
Visitor Identification

The Visitor Identification API allows you to pass identification information to the HubSpot chat widget for otherwise unknown visitors that were verified by your own authentication system.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package visitor_identification

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// GenerateApiService GenerateApi service
type GenerateApiService service

type ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest struct {
	ctx                                  _context.Context
	ApiService                           *GenerateApiService
	identificationTokenGenerationRequest *IdentificationTokenGenerationRequest
}

func (r ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest) IdentificationTokenGenerationRequest(identificationTokenGenerationRequest IdentificationTokenGenerationRequest) ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest {
	r.identificationTokenGenerationRequest = &identificationTokenGenerationRequest
	return r
}

func (r ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest) Execute() (IdentificationTokenResponse, *_nethttp.Response, error) {
	return r.ApiService.PostVisitorIdentificationV3TokensCreateGenerateTokenExecute(r)
}

/*
PostVisitorIdentificationV3TokensCreateGenerateToken Generate a token

Generates a new visitor identification token. This token will be unique every time this endpoint is called, even if called with the same email address. This token is temporary and will expire after 12 hours

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest
*/
func (a *GenerateApiService) PostVisitorIdentificationV3TokensCreateGenerateToken(ctx _context.Context) ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest {
	return ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return IdentificationTokenResponse
func (a *GenerateApiService) PostVisitorIdentificationV3TokensCreateGenerateTokenExecute(r ApiPostVisitorIdentificationV3TokensCreateGenerateTokenRequest) (IdentificationTokenResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IdentificationTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenerateApiService.PostVisitorIdentificationV3TokensCreateGenerateToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversations/v3/visitor-identification/tokens/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.identificationTokenGenerationRequest == nil {
		return localVarReturnValue, nil, reportError("identificationTokenGenerationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.identificationTokenGenerationRequest
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