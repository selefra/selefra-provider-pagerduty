package tables

import (
	"context"
	"errors"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyEscalationPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyEscalationPolicyGenerator{}

func (x *TablePagerdutyEscalationPolicyGenerator) GetTableName() string {
	return "pagerduty_escalation_policy"
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListEscalationPoliciesOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				policies, err := client.ListEscalationPoliciesWithContext(ctx, req)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				listResponse := policies

				for _, policy := range listResponse.EscalationPolicies {
					resultChannel <- policy
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

func listPagerDutyEscalationPolicyTags(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	info := result.(pagerduty.EscalationPolicy)

	client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
	if err != nil {
		return nil, err
	}

	data, err := client.GetTagsForEntityPaginated(ctx, "escalation_policies", info.ID, pagerduty.ListTagOptions{})

	if err != nil {
		return nil, err
	}

	getResp := data

	return getResp, nil
}

func shouldRetryError(err error) bool {
	var aerr pagerduty.APIError

	if errors.As(err, &aerr) {
		return aerr.RateLimited()
	}
	return false
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("services").ColumnType(schema.ColumnTypeJSON).Description("A list of services associated with the policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("teams").ColumnType(schema.ColumnTypeJSON).Description("A list of teams associated with the policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_loops").ColumnType(schema.ColumnTypeInt).Description("The number of times the escalation policy will repeat after reaching the end of its escalation.").
			Extractor(column_value_extractor.StructSelector("NumLoops")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A shortened description of escalation policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("An URL at which the entity is uniquely displayed in the Web app.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalation_rules").ColumnType(schema.ColumnTypeJSON).Description("A list of escalation rules.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A list of tags applied on escalation policy.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listPagerDutyEscalationPolicyTags(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the escalation policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of an escalation policy.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
	}
}

func (x *TablePagerdutyEscalationPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
