# Table: pagerduty_ruleset

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the ruleset. | 
| id | string | X | √ | An unique identifier of a ruleset. | 
| type | string | X | √ | The type of the ruleset. Allowed values are: 'global' and 'default_global'. | 
| creator | json | X | √ | A set of information about the user who created the ruleset. | 
| routing_keys | json | X | √ | A list of routing keys for this ruleset. | 
| team | json | X | √ | A set of information about the team that owns the ruleset. | 
| updater | json | X | √ | A set information about the user that has updated the ruleset. | 
| title | string | X | √ | Title of the resource. | 


