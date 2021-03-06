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

type PushLeadToMarketoRequest struct {

	Input []Lead `json:"input,omitempty"`

	LookupField string `json:"lookupField,omitempty"`

	PartitionName string `json:"partitionName,omitempty"`

	ProgramName string `json:"programName,omitempty"`

	ProgramStatus string `json:"programStatus,omitempty"`

	Reason string `json:"reason,omitempty"`

	Source string `json:"source,omitempty"`
}
