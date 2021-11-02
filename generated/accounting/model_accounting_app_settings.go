/*
 * Accounting Extension
 *
 * These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package accounting

// The URL Settings, which defines the URL endpoints that HubSpot will send requests to an external accounting application for certain actions.
type AccountingAppSettings struct {
	// The ID of the accounting app. This is the identifier of the application created in your HubSpot developer portal.
	AppId    int32               `json:"appId"`
	Urls     *AccountingAppUrls  `json:"urls"`
	Features *AccountingFeatures `json:"features,omitempty"`
}