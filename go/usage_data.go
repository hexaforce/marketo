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

import (
	"time"
)

type UsageData struct {

	// Number of calls made in the time period
	Count int32 `json:"count,omitempty"`

	// Date of the collected calls
	Date time.Time `json:"date"`

	// Counts for individual error codes
	Errors []ErrorCount `json:"errors,omitempty"`

	// Total number of errors in the time period
	Total int32 `json:"total,omitempty"`

	// Counts for individual users
	Users []UserCount `json:"users,omitempty"`
}
