package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyRulesetRuleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyRulesetRuleGenerator{}

func (x *TablePagerdutyRulesetRuleGenerator) GetTableName() string {
	return "pagerduty_ruleset_rule"
}

func (x *TablePagerdutyRulesetRuleGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyRulesetRuleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyRulesetRuleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyRulesetRuleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			rulesetData := task.ParentRawResult.(*pagerduty.Ruleset)

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			resp, err := client.ListRulesetRulesPaginated(ctx, rulesetData.ID)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, rules := range resp {
				resultChannel <- rulesetRuleInfo{*rules, rulesetData.ID}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type rulesetRuleInfo = struct {
	pagerduty.RulesetRule
	RulesetID string
}

func (x *TablePagerdutyRulesetRuleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyRulesetRuleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("ruleset_id").ColumnType(schema.ColumnTypeString).Description("The ID of the ruleset.").
			Extractor(column_value_extractor.StructSelector("RulesetID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catch_all").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions").ColumnType(schema.ColumnTypeJSON).Description("A set of actions that defines when an event matches this rule, the actions that will be taken to change the resulting alert and incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).Description("A set of information defined the conditions resulting alert and incident.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The ID of the event rule.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the Event Rule is disabled and would therefore not be evaluated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("position").ColumnType(schema.ColumnTypeInt).Description("Position/index of the Event Rule in the Ruleset. Starting from position 0 (the first rule), rules are evaluated one-by-one until a matching rule is found.").
			Extractor(column_value_extractor.StructSelector("Position")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_frame").ColumnType(schema.ColumnTypeJSON).Description("Time-based conditions for limiting when the rule is active.").Build(),
	}
}

func (x *TablePagerdutyRulesetRuleGenerator) GetSubTables() []*schema.Table {
	return nil
}
