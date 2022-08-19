/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// PipelineAuditsApiService PipelineAuditsApi service
type PipelineAuditsApiService service

type ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest struct {
	ctx        context.Context
	ApiService *PipelineAuditsApiService
	objectType string
	pipelineId string
}

func (r ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest) Execute() (*CollectionResponsePublicAuditInfoNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditExecute(r)
}

/*
GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit Return an audit of all changes to the pipeline

Return a reverse chronological list of all mutations that have occurred on the pipeline identified by `{pipelineId}`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param pipelineId
 @return ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest
*/
func (a *PipelineAuditsApiService) GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit(ctx context.Context, objectType string, pipelineId string) ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest {
	return ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//  @return CollectionResponsePublicAuditInfoNoPaging
func (a *PipelineAuditsApiService) GetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditExecute(r ApiGetCrmV3PipelinesObjectTypePipelineIdAuditGetAuditRequest) (*CollectionResponsePublicAuditInfoNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicAuditInfoNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelineAuditsApiService.GetCrmV3PipelinesObjectTypePipelineIdAuditGetAudit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}/audit"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
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
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
