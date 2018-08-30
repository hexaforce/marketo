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

type ObjectField struct {

	// Datatype of the field
	DataType string `json:"dataType,omitempty"`

	// UI display-name of the field
	DisplayName string `json:"displayName,omitempty"`

	// Max length of the field.  Only applicable to text, string, and text area.
	Length int32 `json:"length,omitempty"`

	// Name of the field
	Name string `json:"name,omitempty"`

	// Whether the field is updateable
	Updateable bool `json:"updateable,omitempty"`
}
