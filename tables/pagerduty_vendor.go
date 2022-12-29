package tables

import (
	"context"
	"github.com/selefra/selefra-provider-pagerduty/pagerduty_client"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePagerdutyVendorGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePagerdutyVendorGenerator{}

func (x *TablePagerdutyVendorGenerator) GetTableName() string {
	return "pagerduty_vendor"
}

func (x *TablePagerdutyVendorGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePagerdutyVendorGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePagerdutyVendorGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePagerdutyVendorGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := pagerduty_client.GetSessionConfig(ctx, taskClient.(*pagerduty_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			req := pagerduty.ListVendorOptions{}

			maxResult := uint(100)

			req.APIListObject.Limit = maxResult

			for {
				vendors, err := client.ListVendorsWithContext(ctx, req)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				listResponse := vendors

				for _, vendor := range listResponse.Vendors {
					resultChannel <- vendor

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

func (x *TablePagerdutyVendorGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePagerdutyVendorGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("html_url").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").
			Extractor(column_value_extractor.StructSelector("HTMLURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("summary").ColumnType(schema.ColumnTypeString).Description("A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The API show URL at which the object is accessible.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the vendor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("An unique identifier of the vendor.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_guide_url").ColumnType(schema.ColumnTypeString).Description("Specifies the URL of an integration guide for this vendor.").
			Extractor(column_value_extractor.StructSelector("IntegrationGuideURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_pdcef").ColumnType(schema.ColumnTypeBool).Description("Indicates the PagerDuty Common Event Format(PD-CEF).").
			Extractor(column_value_extractor.StructSelector("IsPDCEF")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logo_url").ColumnType(schema.ColumnTypeString).Description("Specifies the URL of a logo identifying the vendor.").
			Extractor(column_value_extractor.StructSelector("LogoURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thumbnail_url").ColumnType(schema.ColumnTypeString).Description("Specifies the URL of a small thumbnail image identifying the vendor.").
			Extractor(column_value_extractor.StructSelector("ThumbnailURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of object being created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("long_name").ColumnType(schema.ColumnTypeString).Description("The full name of the vendor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_creation_default").ColumnType(schema.ColumnTypeString).Description("Specifies the default method for the alert creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generic_service_type").ColumnType(schema.ColumnTypeString).Description("Specifies the generic service type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the vendor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("website_url").ColumnType(schema.ColumnTypeString).Description("The description of the vendor.").
			Extractor(column_value_extractor.StructSelector("WebsiteURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_creation_editable").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the default alert creation method can be editable, or not.").Build(),
	}
}

func (x *TablePagerdutyVendorGenerator) GetSubTables() []*schema.Table {
	return nil
}
