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
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"AddCustomActivityUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external.json",
		AddCustomActivityUsingPOST,
	},

	Route{
		"ApproveCustomActivityTypeUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/approve.json",
		ApproveCustomActivityTypeUsingPOST,
	},

	Route{
		"CreateCustomActivityTypeAttributesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/attributes/create.json",
		CreateCustomActivityTypeAttributesUsingPOST,
	},

	Route{
		"CreateCustomActivityTypeUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type.json",
		CreateCustomActivityTypeUsingPOST,
	},

	Route{
		"DeleteCustomActivityTypeAttributesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/attributes/delete.json",
		DeleteCustomActivityTypeAttributesUsingPOST,
	},

	Route{
		"DeleteCustomActivityTypeUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/delete.json",
		DeleteCustomActivityTypeUsingPOST,
	},

	Route{
		"DescribeCustomActivityTypeUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/external/type/{apiName}/describe.json",
		DescribeCustomActivityTypeUsingGET,
	},

	Route{
		"DiscardDraftofCustomActivityTypeUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/discardDraft.json",
		DiscardDraftofCustomActivityTypeUsingPOST,
	},

	Route{
		"GetActivitiesPagingTokenUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/pagingtoken.json",
		GetActivitiesPagingTokenUsingGET,
	},

	Route{
		"GetAllActivityTypesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/types.json",
		GetAllActivityTypesUsingGET,
	},

	Route{
		"GetCustomActivityTypeUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/external/types.json",
		GetCustomActivityTypeUsingGET,
	},

	Route{
		"GetDeletedLeadsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/deletedleads.json",
		GetDeletedLeadsUsingGET,
	},

	Route{
		"GetLeadActivitiesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities.json",
		GetLeadActivitiesUsingGET,
	},

	Route{
		"GetLeadChangesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/activities/leadchanges.json",
		GetLeadChangesUsingGET,
	},

	Route{
		"UpdateCustomActivityTypeAttributesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}/attributes/update.json",
		UpdateCustomActivityTypeAttributesUsingPOST,
	},

	Route{
		"UpdateCustomActivityTypeUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/activities/external/type/{apiName}.json",
		UpdateCustomActivityTypeUsingPOST,
	},

	Route{
		"GetImportCustomObjectFailuresUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/customobjects/{apiName}/import/{id}/failures.json",
		GetImportCustomObjectFailuresUsingGET,
	},

	Route{
		"GetImportCustomObjectStatusUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/customobjects/{apiName}/import/{id}/status.json",
		GetImportCustomObjectStatusUsingGET,
	},

	Route{
		"GetImportCustomObjectWarningsUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/customobjects/{apiName}/import/{id}/warnings.json",
		GetImportCustomObjectWarningsUsingGET,
	},

	Route{
		"ImportCustomObjectUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/customobjects/{apiName}/import.json",
		ImportCustomObjectUsingPOST,
	},

	Route{
		"CancelExportActivitiesUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/activities/export/{exportId}/cancel.json",
		CancelExportActivitiesUsingPOST,
	},

	Route{
		"CreateExportActivitiesUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/activities/export/create.json",
		CreateExportActivitiesUsingPOST,
	},

	Route{
		"EnqueueExportActivitiesUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/activities/export/{exportId}/enqueue.json",
		EnqueueExportActivitiesUsingPOST,
	},

	Route{
		"GetExportActivitiesFileUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/activities/export/{exportId}/file.json",
		GetExportActivitiesFileUsingGET,
	},

	Route{
		"GetExportActivitiesStatusUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/activities/export/{exportId}/status.json",
		GetExportActivitiesStatusUsingGET,
	},

	Route{
		"GetExportActivitiesUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/activities/export.json",
		GetExportActivitiesUsingGET,
	},

	Route{
		"CancelExportLeadsUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/leads/export/{exportId}/cancel.json",
		CancelExportLeadsUsingPOST,
	},

	Route{
		"CreateExportLeadsUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/leads/export/create.json",
		CreateExportLeadsUsingPOST,
	},

	Route{
		"EnqueueExportLeadsUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/leads/export/{exportId}/enqueue.json",
		EnqueueExportLeadsUsingPOST,
	},

	Route{
		"GetExportLeadsFileUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/export/{exportId}/file.json",
		GetExportLeadsFileUsingGET,
	},

	Route{
		"GetExportLeadsStatusUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/export/{exportId}/status.json",
		GetExportLeadsStatusUsingGET,
	},

	Route{
		"GetExportLeadsUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/export.json",
		GetExportLeadsUsingGET,
	},

	Route{
		"GetImportLeadFailuresUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/batch/{batchId}/failures.json",
		GetImportLeadFailuresUsingGET,
	},

	Route{
		"GetImportLeadStatusUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/batch/{batchId}.json",
		GetImportLeadStatusUsingGET,
	},

	Route{
		"GetImportLeadWarningsUsingGET",
		strings.ToUpper("Get"),
		"//bulk/v1/leads/batch/{batchId}/warnings.json",
		GetImportLeadWarningsUsingGET,
	},

	Route{
		"ImportLeadUsingPOST",
		strings.ToUpper("Post"),
		"//bulk/v1/leads.json",
		ImportLeadUsingPOST,
	},

	Route{
		"GetCampaignByIdUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/campaigns/{campaignId}.json",
		GetCampaignByIdUsingGET,
	},

	Route{
		"GetCampaignsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/campaigns.json",
		GetCampaignsUsingGET,
	},

	Route{
		"ScheduleCampaignUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/campaigns/{campaignId}/schedule.json",
		ScheduleCampaignUsingPOST,
	},

	Route{
		"TriggerCampaignUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/campaigns/{campaignId}/trigger.json",
		TriggerCampaignUsingPOST,
	},

	Route{
		"DeleteCompaniesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/companies/delete.json",
		DeleteCompaniesUsingPOST,
	},

	Route{
		"DescribeUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/companies/describe.json",
		DescribeUsingGET,
	},

	Route{
		"GetCompaniesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/companies.json",
		GetCompaniesUsingGET,
	},

	Route{
		"SyncCompaniesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/companies.json",
		SyncCompaniesUsingPOST,
	},

	Route{
		"DeleteCustomObjectsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/customobjects/{customObjectName}/delete.json",
		DeleteCustomObjectsUsingPOST,
	},

	Route{
		"DescribeUsingGET1",
		strings.ToUpper("Get"),
		"//rest/v1/customobjects/{customObjectName}/describe.json",
		DescribeUsingGET1,
	},

	Route{
		"GetCustomObjectsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/customobjects/{customObjectName}.json",
		GetCustomObjectsUsingGET,
	},

	Route{
		"ListCustomObjectsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/customobjects.json",
		ListCustomObjectsUsingGET,
	},

	Route{
		"SyncCustomObjectsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/customobjects/{customObjectName}.json",
		SyncCustomObjectsUsingPOST,
	},

	Route{
		"AssociateLeadUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/{leadId}/associate.json",
		AssociateLeadUsingPOST,
	},

	Route{
		"ChangeLeadProgramStatusUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/programs/{programId}/status.json",
		ChangeLeadProgramStatusUsingPOST,
	},

	Route{
		"DeleteLeadsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/delete.json",
		DeleteLeadsUsingPOST,
	},

	Route{
		"DescribeUsingGET2",
		strings.ToUpper("Get"),
		"//rest/v1/leads/describe.json",
		DescribeUsingGET2,
	},

	Route{
		"GetLeadByIdUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/lead/{leadId}.json",
		GetLeadByIdUsingGET,
	},

	Route{
		"GetLeadByIdUsingGET1",
		strings.ToUpper("Get"),
		"//rest/v1/leads/{leadId}.json",
		GetLeadByIdUsingGET1,
	},

	Route{
		"GetLeadPartitionsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/leads/partitions.json",
		GetLeadPartitionsUsingGET,
	},

	Route{
		"GetLeadsByFilterUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/leads.json",
		GetLeadsByFilterUsingGET,
	},

	Route{
		"GetLeadsByProgramIdUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/leads/programs/{programId}.json",
		GetLeadsByProgramIdUsingGET,
	},

	Route{
		"MergeLeadsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/{leadId}/merge.json",
		MergeLeadsUsingPOST,
	},

	Route{
		"PushToMarketoUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/push.json",
		PushToMarketoUsingPOST,
	},

	Route{
		"SyncLeadUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads.json",
		SyncLeadUsingPOST,
	},

	Route{
		"UpdatePartitionsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/leads/partitions.json",
		UpdatePartitionsUsingPOST,
	},

	Route{
		"AddNamedAccountListMembersUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedAccountList/{id}/namedAccounts.json",
		AddNamedAccountListMembersUsingPOST,
	},

	Route{
		"DeleteNamedAccountListsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedAccountLists/delete.json",
		DeleteNamedAccountListsUsingPOST,
	},

	Route{
		"GetNamedAccountListMembersUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/namedAccountList/{id}/namedAccounts.json",
		GetNamedAccountListMembersUsingGET,
	},

	Route{
		"GetNamedAccountListsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/namedAccountLists.json",
		GetNamedAccountListsUsingGET,
	},

	Route{
		"RemoveNamedAccountListMembersUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedAccountList/{id}/namedAccounts/remove.json",
		RemoveNamedAccountListMembersUsingPOST,
	},

	Route{
		"SyncNamedAccountListsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedAccountLists.json",
		SyncNamedAccountListsUsingPOST,
	},

	Route{
		"DeleteNamedAccountsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedaccounts/delete.json",
		DeleteNamedAccountsUsingPOST,
	},

	Route{
		"DescribeUsingGET3",
		strings.ToUpper("Get"),
		"//rest/v1/namedaccounts/describe.json",
		DescribeUsingGET3,
	},

	Route{
		"GetNamedAccountsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/namedaccounts.json",
		GetNamedAccountsUsingGET,
	},

	Route{
		"SyncNamedAccountsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/namedaccounts.json",
		SyncNamedAccountsUsingPOST,
	},

	Route{
		"DeleteOpportunitiesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/opportunities/delete.json",
		DeleteOpportunitiesUsingPOST,
	},

	Route{
		"DeleteOpportunityRolesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/opportunities/roles/delete.json",
		DeleteOpportunityRolesUsingPOST,
	},

	Route{
		"DescribeOpportunityRoleUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/opportunities/roles/describe.json",
		DescribeOpportunityRoleUsingGET,
	},

	Route{
		"DescribeUsingGET4",
		strings.ToUpper("Get"),
		"//rest/v1/opportunities/describe.json",
		DescribeUsingGET4,
	},

	Route{
		"GetOpportunitiesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/opportunities.json",
		GetOpportunitiesUsingGET,
	},

	Route{
		"GetOpportunityRolesUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/opportunities/roles.json",
		GetOpportunityRolesUsingGET,
	},

	Route{
		"SyncOpportunitiesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/opportunities.json",
		SyncOpportunitiesUsingPOST,
	},

	Route{
		"SyncOpportunityRolesUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/opportunities/roles.json",
		SyncOpportunityRolesUsingPOST,
	},

	Route{
		"DeleteSalesPersonUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/salespersons/delete.json",
		DeleteSalesPersonUsingPOST,
	},

	Route{
		"DescribeUsingGET5",
		strings.ToUpper("Get"),
		"//rest/v1/salespersons/describe.json",
		DescribeUsingGET5,
	},

	Route{
		"GetSalesPersonUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/salespersons.json",
		GetSalesPersonUsingGET,
	},

	Route{
		"SyncSalesPersonsUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/salespersons.json",
		SyncSalesPersonsUsingPOST,
	},

	Route{
		"AddLeadsToListUsingPOST",
		strings.ToUpper("Post"),
		"//rest/v1/lists/{listId}/leads.json",
		AddLeadsToListUsingPOST,
	},

	Route{
		"AreLeadsMemberOfListUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/lists/{listId}/leads/ismember.json",
		AreLeadsMemberOfListUsingGET,
	},

	Route{
		"GetLeadsByListIdUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/list/{listId}/leads.json",
		GetLeadsByListIdUsingGET,
	},

	Route{
		"GetLeadsByListIdUsingGET1",
		strings.ToUpper("Get"),
		"//rest/v1/lists/{listId}/leads.json",
		GetLeadsByListIdUsingGET1,
	},

	Route{
		"GetListByIdUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/lists/{listId}.json",
		GetListByIdUsingGET,
	},

	Route{
		"GetListsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/lists.json",
		GetListsUsingGET,
	},

	Route{
		"RemoveLeadsFromListUsingDELETE",
		strings.ToUpper("Delete"),
		"//rest/v1/lists/{listId}/leads.json",
		RemoveLeadsFromListUsingDELETE,
	},

	Route{
		"GetDailyErrorsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/stats/errors.json",
		GetDailyErrorsUsingGET,
	},

	Route{
		"GetDailyUsageUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/stats/usage.json",
		GetDailyUsageUsingGET,
	},

	Route{
		"GetLast7DaysErrorsUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/stats/errors/last7days.json",
		GetLast7DaysErrorsUsingGET,
	},

	Route{
		"GetLast7DaysUsageUsingGET",
		strings.ToUpper("Get"),
		"//rest/v1/stats/usage/last7days.json",
		GetLast7DaysUsageUsingGET,
	},
}
