package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyIncidentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyIncidentGenerator{}

func (x *TablePagerdutyIncidentGenerator) GetTableName() string {
	return "pagerduty_incident"
}

func (x *TablePagerdutyIncidentGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyIncidentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyIncidentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyIncidentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListIncidentsOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				incidents, err := client.ListIncidentsWithContext(ctx, req)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				listResponse := incidents

				for _, incident := range listResponse.Incidents {
					resultChannel <- incident
				}

				if !listResponse.APIListObject.More {
					break
				}

				req.APIListObject.Offset = listResponse.APIListObject.Offset + listResponse.APIListObject.Limit
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TablePagerdutyIncidentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyIncidentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeJSON).Description("Describes the additional incident details.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date/time the incident was first triggered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_status_change_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which the status of the incident last changed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_counts").ColumnType(schema.ColumnTypeJSON).Description("Describes the count of triggered and resolved alerts.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_trigger_log_entry").ColumnType(schema.ColumnTypeJSON).Description("Specifies the first log entry when the incident was triggered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Description("The teams involved in the incident's lifecycle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeJSON).Description("Specifies the priority set for this incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Summary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_mergeable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the incident's alerts can be merged with another incident, or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assignments").ColumnType(schema.ColumnTypeJSON).Description("A list of all assignments for this incident. This list will be empty if the 'Incident.status' is resolved.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_policy").ColumnType(schema.ColumnTypeJSON).Description("Specifies the escalation policy assigned to this incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The current status of the incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("incident_key").ColumnType(schema.ColumnTypeString).Description("The incident's de-duplication key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conference_bridge").ColumnType(schema.ColumnTypeJSON).Description("Specifies the contact information that allows responders to easily connect and collaborate during major incident response.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("urgency").ColumnType(schema.ColumnTypeString).Description("The current urgency of the incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_status_change_by").ColumnType(schema.ColumnTypeJSON).Description("The agent (user, service or integration) that created or modified the incident log entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("incident_number").ColumnType(schema.ColumnTypeInt).Description("The number of the incident. This is unique across your account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeJSON).Description("Specifies the information about the impacted service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acknowledgements").ColumnType(schema.ColumnTypeJSON).Description("A list of all acknowledgements for this incident. This list will be empty if the 'Incident.status' is resolved or triggered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resolve_reason").ColumnType(schema.ColumnTypeJSON).Description("Specifies the reason the incident was resolved.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_actions").ColumnType(schema.ColumnTypeJSON).Description("A list of pending_actions on the incident. A pending_action object contains a type of action which can be escalate, unacknowledge, resolve or urgency_change. A pending_action object contains at, the time at which the action will take place. An urgency_change pending_action will contain to, the urgency that the incident will change to.").Build(),
	}
}

func (x *TablePagerdutyIncidentGenerator) GetSubTables() []*schema.Table {
	return nil
}
