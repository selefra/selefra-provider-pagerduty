# Table: pagerduty_escalation_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| services | json | X | √ | A list of services associated with the policy. | 
| teams | json | X | √ | A list of teams associated with the policy. | 
| title | string | X | √ | Title of the resource. | 
| num_loops | int | X | √ | The number of times the escalation policy will repeat after reaching the end of its escalation. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| description | string | X | √ | A shortened description of escalation policy. | 
| html_url | string | X | √ | An URL at which the entity is uniquely displayed in the Web app. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| type | string | X | √ | The type of object being created. | 
| escalation_rules | json | X | √ | A list of escalation rules. | 
| tags | json | X | √ | A list of tags applied on escalation policy. | 
| name | string | X | √ | The name of the escalation policy. | 
| id | string | X | √ | An unique identifier of an escalation policy. | 


