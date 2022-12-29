package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyRulesetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyRulesetGenerator{}

func (x *TablePagerdutyRulesetGenerator) GetTableName() string {
	return "pagerduty_ruleset"
}

func (x *TablePagerdutyRulesetGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyRulesetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyRulesetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyRulesetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp, err := client.ListRulesetsPaginated(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, ruleset := range resp {
				resultChannel <- ruleset
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TablePagerdutyRulesetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyRulesetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the ruleset.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of a ruleset.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the ruleset. Allowed values are: 'global' and 'default_global'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeJSON).Description("A set of information about the user who created the ruleset.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("routing_keys").ColumnType(schema.ColumnTypeJSON).Description("A list of routing keys for this ruleset.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team").ColumnType(schema.ColumnTypeJSON).Description("A set of information about the team that owns the ruleset.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updater").ColumnType(schema.ColumnTypeJSON).Description("A set information about the user that has updated the ruleset.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TablePagerdutyRulesetGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TablePagerdutyRulesetRuleGenerator{}),
	}
}
