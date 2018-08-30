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

type StaticList struct {

	// Datetime when the list was created
	CreatedAt string `json:"createdAt"`

	// Description of the static list
	Description string `json:"description,omitempty"`

	// Unique integer id of the static list
	Id int32 `json:"id"`

	// Name of the static list
	Name string `json:"name"`

	// Name of the program
	ProgramName string `json:"programName,omitempty"`

	// Datetime when the list was most recently updated
	UpdatedAt string `json:"updatedAt"`

	// Name of the parent workspace, if applicable
	WorkspaceName string `json:"workspaceName,omitempty"`
}