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

type Token struct {

	// Name of the token.  Should be formatted as \"{{my.name}}\"
	Name string `json:"name"`

	// Value of the token
	Value string `json:"value"`
}
