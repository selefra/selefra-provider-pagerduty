package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyServiceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyServiceGenerator{}

func (x *TablePagerdutyServiceGenerator) GetTableName() string {
	return "pagerduty_service"
}

func (x *TablePagerdutyServiceGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyServiceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyServiceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyServiceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListServiceOptions{
				Includes: []string{"integrations"},
			}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			givenColumns := []string{}
			includeFields := buildServiceRequestFields(ctx, givenColumns)
			if len(includeFields) > 0 {
				req.Includes = includeFields
			}

			resp, err := client.ListServicesPaginated(ctx, req)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, service := range resp {
				resultChannel <- service

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func buildServiceRequestFields(ctx context.Context, queryColumns []string) []string {
	var fields []string
	for _, columnName := range queryColumns {
		switch columnName {
		case "escalation_policy":
			fields = append(fields, "escalation_policies")
		case "teams":
			fields = append(fields, columnName)
		}
	}
	return fields
}

func (x *TablePagerdutyServiceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyServiceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("scheduled_actions").ColumnType(schema.ColumnTypeJSON).Description("An array containing scheduled actions for the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("support_hours").ColumnType(schema.ColumnTypeJSON).Description("Defines the service's support hours").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of a service.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acknowledgement_timeout").ColumnType(schema.ColumnTypeInt).Description("Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_creation").ColumnType(schema.ColumnTypeString).Description("Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_incident_timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The date/time when the most recent incident was created for this service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Description("The set of teams associated with this service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integrations").ColumnType(schema.ColumnTypeJSON).Description("An array containing integrations that belong to this service. If integrations is passed as an argument, these are full objects - otherwise, these are references.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The user-provided description of the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_resolve_timeout").ColumnType(schema.ColumnTypeInt).Description("Time in seconds that an incident is automatically resolved if left open for that long. Value is null if the feature is disabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_grouping_parameters").ColumnType(schema.ColumnTypeJSON).Description("Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("incident_urgency_rule").ColumnType(schema.ColumnTypeJSON).Description("A list of incident urgency rules.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The current state of the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date/time when this service was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_policy").ColumnType(schema.ColumnTypeJSON).Description("Escalation policy associated with the service.").Build(),
	}
}

func (x *TablePagerdutyServiceGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TablePagerdutyServiceIntegrationGenerator{}),
	}
}
