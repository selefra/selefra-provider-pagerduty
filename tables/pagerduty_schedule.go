package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyScheduleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyScheduleGenerator{}

func (x *TablePagerdutyScheduleGenerator) GetTableName() string {
	return "pagerduty_schedule"
}

func (x *TablePagerdutyScheduleGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyScheduleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyScheduleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyScheduleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListSchedulesOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				schedules, err := client.ListSchedulesWithContext(ctx, req)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				listResponse := schedules

				for _, schedule := range listResponse.Schedules {
					resultChannel <- schedule
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

func getPagerDutySchedule(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
	if err != nil {
		return nil, err
	}

	var id string
	if result != nil {
		id = result.(pagerduty.Schedule).ID
	}

	if id == "" {
		return nil, nil
	}

	data, err := client.GetScheduleWithContext(ctx, id, pagerduty.GetScheduleOptions{})

	getResp := data

	return *getResp, nil
}

func (x *TablePagerdutyScheduleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyScheduleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of a schedule.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).Description("A list of the users on the schedule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("final_schedule").ColumnType(schema.ColumnTypeJSON).Description("Specifies the final schedule.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 004
				r, err := getPagerDutySchedule(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_layers").ColumnType(schema.ColumnTypeJSON).Description("A list of schedule layers.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 004
				r, err := getPagerDutySchedule(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return column_value_extractor.DefaultColumnValueExtractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the schedule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the schedule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The time zone of the schedule.").
			Extractor(column_value_extractor.StructSelector("TimeZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Description("A list of the teams on the schedule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_policies").ColumnType(schema.ColumnTypeJSON).Description("A list of the escalation policies that uses this schedule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("override_sub_schedule").ColumnType(schema.ColumnTypeJSON).Description("Specifies schedule overrides for a given time range.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := getPagerDutySchedule(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("OverrideSubschedule")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TablePagerdutyScheduleGenerator) GetSubTables() []*schema.Table {
	return nil
}
