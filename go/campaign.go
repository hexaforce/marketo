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

// Record of a Marketo Smart Campaign
type Campaign struct {

	// Whether the campaign is active.  Only applicable to trigger campaigns
	Active bool `json:"active,omitempty"`

	// Datetime when the campaign was created
	CreatedAt string `json:"createdAt"`

	// Description of the Smart Campaign
	Description string `json:"description,omitempty"`

	// Unique integer id of the Smart Campaign
	Id int32 `json:"id"`

	// Name of the Smart Campaign
	Name string `json:"name"`

	// Id of the parent program if applicable
	ProgramId int32 `json:"programId,omitempty"`

	// Name of the parent program if applicable
	ProgramName string `json:"programName,omitempty"`

	// Type of the Smart Campaign
	Type_ string `json:"type"`

	// Datetime when the campaign was most recently updated
	UpdatedAt string `json:"updatedAt"`

	// Name of the parent workspace if applicable
	WorkspaceName string `json:"workspaceName,omitempty"`
}
