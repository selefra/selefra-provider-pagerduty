package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyTagGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyTagGenerator{}

func (x *TablePagerdutyTagGenerator) GetTableName() string {
	return "pagerduty_tag"
}

func (x *TablePagerdutyTagGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyTagGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyTagGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyTagGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListTagOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			data, err := client.ListTagsPaginated(ctx, req)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			tags := data

			for _, tag := range tags {
				resultChannel <- tag

			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TablePagerdutyTagGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyTagGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Label")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of a tag.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("The label of the tag.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
	}
}

func (x *TablePagerdutyTagGenerator) GetSubTables() []*schema.Table {
	return nil
}
