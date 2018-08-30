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

type ExportLeadFilter struct {

	// Date range to filter new leads on
	CreatedAt *DateRange `json:"createdAt"`

	// Id of smart list to retrieve leads from
	SmartListId int32 `json:"smartListId"`

	// Name of smart list to retrieve leads from
	SmartListName string `json:"smartListName"`

	// Id of static list to retrieve leads from
	StaticListId int32 `json:"staticListId"`

	// Name of static list to retrieve leads from
	StaticListName string `json:"staticListName"`

	// Date range to filter updated leads on
	UpdatedAt *DateRange `json:"updatedAt"`
}
