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

type CustomActivity struct {

	// Datetime of the activity
	ActivityDate string `json:"activityDate"`

	// Id of the activity type
	ActivityTypeId int32 `json:"activityTypeId"`

	ApiName string `json:"apiName,omitempty"`

	// List of secondary attributes
	Attributes []Attribute `json:"attributes"`

	// Array of errors that occurred if the request was unsuccessful
	Errors []ModelError `json:"errors"`

	// Integer id of the activity.  For instances which have been migrated to Activity Service, this field may not be present, and should not be treated as unique.
	Id int64 `json:"id"`

	// Id of the lead associated to the activity
	LeadId int64 `json:"leadId"`

	// Unique id of the activity (128 character string)
	MarketoGUID string `json:"marketoGUID,omitempty"`

	// Value of the primary attribute
	PrimaryAttributeValue string `json:"primaryAttributeValue"`

	// Status of the operation performed on the record
	Status string `json:"status,omitempty"`
}
