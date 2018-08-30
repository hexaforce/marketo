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

// Response containing import status information
type ImportCustomObjectResponse struct {

	// Unique integer id of the import batch
	BatchId int32 `json:"batchId"`

	// Time spent on the batch
	ImportTime string `json:"importTime,omitempty"`

	// Status message of the batch
	Message string `json:"message,omitempty"`

	// Number of rows processed so far
	NumOfObjectsProcessed int32 `json:"numOfObjectsProcessed,omitempty"`

	// Number of rows failed so far
	NumOfRowsFailed int32 `json:"numOfRowsFailed,omitempty"`

	// Number of rows with a warning so far
	NumOfRowsWithWarning int32 `json:"numOfRowsWithWarning,omitempty"`

	// Object API Name
	ObjectApiName string `json:"objectApiName"`

	// Bulk operation type. Can be import or export
	Operation string `json:"operation"`

	// Status of the batch
	Status string `json:"status"`
}