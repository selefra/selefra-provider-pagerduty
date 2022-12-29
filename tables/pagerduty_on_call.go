package tables

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyOnCallGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyOnCallGenerator{}

func (x *TablePagerdutyOnCallGenerator) GetTableName() string {
	return "pagerduty_on_call"
}

func (x *TablePagerdutyOnCallGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyOnCallGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyOnCallGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyOnCallGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListOnCallOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				data, err := client.ListOnCallsWithContext(ctx, req)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				listResponse := data

				for _, oncall := range listResponse.OnCalls {
					resultChannel <- oncall
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

func isNotFoundError(err error) bool {
	var aerr pagerduty.APIError

	if errors.As(err, &aerr) {
		return aerr.NotFound()
	}
	return false
}

func (x *TablePagerdutyOnCallGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyOnCallGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_level").ColumnType(schema.ColumnTypeInt).Description("The escalation level for the on-call.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start").ColumnType(schema.ColumnTypeTimestamp).Description("The start of the on-call. If null, the on-call is a permanent user on-call.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end").ColumnType(schema.ColumnTypeTimestamp).Description("The end of the on-call. If null, the user does not go off-call.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_policy").ColumnType(schema.ColumnTypeJSON).Description("The escalation_policy object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule").ColumnType(schema.ColumnTypeJSON).Description("The schedule object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_on_call").ColumnType(schema.ColumnTypeJSON).Description("The user object.").
			Extractor(column_value_extractor.StructSelector("User")).Build(),
	}
}

func (x *TablePagerdutyOnCallGenerator) GetSubTables() []*schema.Table {
	return nil
}
