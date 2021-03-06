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

type CustomActivityTypeAttribute struct {

	// API Name of the attribute
	ApiName string `json:"apiName"`

	// Data type of the attribute
	DataType string `json:"dataType,omitempty"`

	// Description of the attribute
	Description string `json:"description,omitempty"`

	// Whether the attribute is the primary attribute of the activity type.  There may only be one primary attribute at a time
	IsPrimary bool `json:"isPrimary,omitempty"`

	// Human-readable display name of the attribute
	Name string `json:"name"`
}
