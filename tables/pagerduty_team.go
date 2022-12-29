package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyTeamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyTeamGenerator{}

func (x *TablePagerdutyTeamGenerator) GetTableName() string {
	return "pagerduty_team"
}

func (x *TablePagerdutyTeamGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyTeamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyTeamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyTeamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListTeamOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				teams, err := client.ListTeamsWithContext(ctx, req)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				listResponse := teams

				for _, team := range listResponse.Teams {
					resultChannel <- team
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

func listPagerDutyTeamMembers(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	info := result.(pagerduty.Team)

	client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
	if err != nil {

		return nil, err
	}

	data, err := client.ListMembersPaginated(ctx, info.ID)
	if err != nil {
		return nil, err
	}
	members := data

	return members, nil
}

func listPagerDutyTeamTags(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	info := result.(pagerduty.Team)

	client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
	if err != nil {

		return nil, err
	}

	data, err := client.GetTagsForEntityPaginated(ctx, "teams", info.ID, pagerduty.ListTagOptions{})

	if err != nil {
		return nil, err
	}
	getResp := data
	return getResp, nil
}

func (x *TablePagerdutyTeamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyTeamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("members").ColumnType(schema.ColumnTypeJSON).Description("A list of members of a team.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listPagerDutyTeamMembers(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A list of tags applied on team.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listPagerDutyTeamTags(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of a team.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TablePagerdutyTeamGenerator) GetSubTables() []*schema.Table {
	return nil
}
