# Table: pagerduty_ruleset_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ruleset_id | string | X | √ | The ID of the ruleset. | 
| catch_all | bool | X | √ | Indicates whether the Event Rule is the last Event Rule of the Ruleset that serves as a catch-all. It has limited functionality compared to other rules and always matches. | 
| actions | json | X | √ | A set of actions that defines when an event matches this rule, the actions that will be taken to change the resulting alert and incident. | 
| conditions | json | X | √ | A set of information defined the conditions resulting alert and incident. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The ID of the event rule. | 
| disabled | bool | X | √ | Indicates whether the Event Rule is disabled and would therefore not be evaluated. | 
| position | int | X | √ | Position/index of the Event Rule in the Ruleset. Starting from position 0 (the first rule), rules are evaluated one-by-one until a matching rule is found. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| time_frame | json | X | √ | Time-based conditions for limiting when the rule is active. | 


