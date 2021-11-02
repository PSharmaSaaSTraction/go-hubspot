/*
 * Marketing Events Extension
 *
 * These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package marketing_events_beta

type ErrorDetail struct {
	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`
	// The name of the field or parameter in which the error was found.
	In string `json:"in,omitempty"`
	// The status code associated with the error detail
	Code string `json:"code,omitempty"`
	// A specific category that contains more specific detail about the error
	SubCategory string `json:"subCategory,omitempty"`
	// Context about the error condition
	Context map[string][]string `json:"context,omitempty"`
}