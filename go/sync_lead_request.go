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

type SyncLeadRequest struct {

	// Type of sync operation to perform.  Defaults to createOrUpdate if unset
	Action string `json:"action,omitempty"`

	// If set to true, the call will return immediately
	AsyncProcessing bool `json:"asyncProcessing,omitempty"`

	// List of leads for input
	Input []Lead `json:"input"`

	// Field to deduplicate on.  The field must be present in each lead record of the input.  Defaults to email if unset
	LookupField string `json:"lookupField,omitempty"`

	PartitionCode string `json:"partitionCode,omitempty"`

	// Name of the partition to operate on, if applicable.  Should be set whenver possible, when interacting with an instance where partitions are enabled.
	PartitionName string `json:"partitionName,omitempty"`
}
