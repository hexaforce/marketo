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

type LeadPartition struct {

	// Description of the partition
	Description string `json:"description,omitempty"`

	// Unique integer id of the partition
	Id int32 `json:"id"`

	// Name of the partition
	Name string `json:"name"`
}
