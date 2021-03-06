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

type ErrorCount struct {

	// Number of occurences of the error
	Count int32 `json:"count"`

	// Integer error code of the error
	ErrorCode string `json:"errorCode"`
}
