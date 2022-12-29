package provider

import (
	"github.com/selefra/selefra-provider-pagerduty/table_schema_generator"
	"github.com/selefra/selefra-provider-pagerduty/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyEscalationPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyTagGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyServiceGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyVendorGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyPriorityGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyIncidentGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyOnCallGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyTeamGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyScheduleGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePagerdutyRulesetGenerator{}),
	}
}
