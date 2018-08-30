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
type ImportLeadResponse struct {

	// Unique integer id of the import batch
	BatchId int32 `json:"batchId"`

	ImportId string `json:"importId,omitempty"`

	Message string `json:"message,omitempty"`

	// Number of rows processed so far
	NumOfLeadsProcessed int32 `json:"numOfLeadsProcessed"`

	// Number of rows failed so far
	NumOfRowsFailed int32 `json:"numOfRowsFailed,omitempty"`

	// Number of rows with a warning so far
	NumOfRowsWithWarning int32 `json:"numOfRowsWithWarning,omitempty"`

	// Status of the batch
	Status string `json:"status"`
}
