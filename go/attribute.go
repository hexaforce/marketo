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

// Name-Value pair
type Attribute struct {

	ApiName string `json:"apiName,omitempty"`

	// Name of the attribute
	Name string `json:"name"`

	// Value of the attribute
	Value *interface{} `json:"value"`
}