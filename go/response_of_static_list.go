/*
 * Marketo Rest API
 *
 * Marketo Rest API
 *
 * API version: 1.0
 * Contact: developerfeedback@marketo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseOfStaticList struct {

	// Array of errors that occurred if the request was unsuccessful
	Errors []ModelError `json:"errors"`

	// Paging token given if the result set exceeded the allowed batch size
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Id of the request made
	RequestId string `json:"requestId"`

	// Array of results for individual records in the operation, may be empty
	Result []StaticList `json:"result"`

	// Whether the request succeeded
	Success bool `json:"success"`

	// Array of warnings given for the operation
	Warnings []Warning `json:"warnings"`
}
