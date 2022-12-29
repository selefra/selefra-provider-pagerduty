package tables

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyServiceIntegrationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyServiceIntegrationGenerator{}

func (x *TablePagerdutyServiceIntegrationGenerator) GetTableName() string {
	return "pagerduty_service_integration"
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			serviceData := task.ParentRawResult.(pagerduty.Service)

			for _, integration := range serviceData.Integrations {
				resultChannel <- integration

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("integration_key").ColumnType(schema.ColumnTypeString).Description("Specify the integration key for the service integration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of this integration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the integration.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the queried service.").
			Extractor(column_value_extractor.StructSelector("Service.ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date/time when this integration was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vendor").ColumnType(schema.ColumnTypeJSON).Description("Describes the information about a specific type of integration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_email").ColumnType(schema.ColumnTypeString).Description("Specify for generic_email_inbound_integration. Must be set to an email address @your-subdomain.pagerduty.com.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeJSON).Description("Describes the information about the queried service.").Build(),
	}
}

func (x *TablePagerdutyServiceIntegrationGenerator) GetSubTables() []*schema.Table {
	return nil
}
