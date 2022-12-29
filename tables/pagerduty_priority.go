package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyPriorityGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyPriorityGenerator{}

func (x *TablePagerdutyPriorityGenerator) GetTableName() string {
	return "pagerduty_priority"
}

func (x *TablePagerdutyPriorityGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyPriorityGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyPriorityGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyPriorityGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			data, err := client.ListPrioritiesWithContext(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp := data

			for _, priority := range resp.Priorities {
				resultChannel <- priority
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TablePagerdutyPriorityGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyPriorityGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user-provided short name of the priority.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the priority.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The user-provided description of the priority.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
	}
}

func (x *TablePagerdutyPriorityGenerator) GetSubTables() []*schema.Table {
	return nil
}
