/*
 * CRM Imports
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package imports

type ImportRowCore struct {
	FileId     int32  `json:"fileId"`
	PageName   string `json:"pageName,omitempty"`
	LineNumber int32  `json:"lineNumber"`
}