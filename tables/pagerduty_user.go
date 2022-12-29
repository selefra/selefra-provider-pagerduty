package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyUserGenerator{}

func (x *TablePagerdutyUserGenerator) GetTableName() string {
	return "pagerduty_user"
}

func (x *TablePagerdutyUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListUsersOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			givenColumns := []string{}
			includeFields := buildUserRequestFields(ctx, givenColumns)

			if len(includeFields) > 0 {
				req.Includes = includeFields
			}

			for {
				users, err := client.ListUsersWithContext(ctx, req)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				listResponse := users

				for _, user := range listResponse.Users {
					resultChannel <- user
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

func buildUserRequestFields(ctx context.Context, queryColumns []string) []string {
	var fields []string
	for _, columnName := range queryColumns {
		if columnName == "contact_methods" || columnName == "notification_rules" || columnName == "teams" {
			fields = append(fields, columnName)
		}
	}
	return fields
}
func listPagerDutyUserTags(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	info := result.(pagerduty.User)

	client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
	if err != nil {

		return nil, err
	}

	data, err := client.GetTagsForEntityPaginated(ctx, "users", info.ID, pagerduty.ListTagOptions{})

	if err != nil {
		return nil, err
	}

	getResp := data

	return getResp, nil
}

func (x *TablePagerdutyUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to 'name', though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The preferred time zone name. If null, the account's time zone will be used.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A list of tags applied on user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listPagerDutyUserTags(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of an user.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The user's email address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("The user role. Account must have the 'read_only_users' ability to set a user as a 'read_only_user' or a 'read_only_limited_user', and must have advanced permissions abilities to set a user as 'observer' or 'restricted_access'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_title").ColumnType(schema.ColumnTypeString).Description("The user's job title.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_rules").ColumnType(schema.ColumnTypeJSON).Description("A list of notification rules for the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invitation_sent").ColumnType(schema.ColumnTypeBool).Description("If true, the user has an outstanding invitation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("color").ColumnType(schema.ColumnTypeString).Description("The schedule color.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The user's bio.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contact_methods").ColumnType(schema.ColumnTypeJSON).Description("A list of contact methods for the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("avatar_url").ColumnType(schema.ColumnTypeString).Description("The URL of the user's avatar.").
			Extractor(column_value_extractor.StructSelector("AvatarURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Description("A list of teams to which the user belongs. Account must have the teams ability to set this.").Build(),
	}
}

func (x *TablePagerdutyUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
