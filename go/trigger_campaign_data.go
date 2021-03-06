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

type TriggerCampaignData struct {

	// List of leads for input
	Leads []InputLead `json:"leads"`

	// List of my tokens to replace during the run of the target campaign.  The tokens must be available in a parent program or folder to be replaced during the run
	Tokens []Token `json:"tokens,omitempty"`
}
