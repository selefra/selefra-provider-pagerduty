# Table: pagerduty_schedule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | An unique identifier of a schedule. | 
| users | json | X | √ | A list of the users on the schedule. | 
| type | string | X | √ | The type of object being created. | 
| final_schedule | json | X | √ | Specifies the final schedule. | 
| schedule_layers | json | X | √ | A list of schedule layers. | 
| name | string | X | √ | The name of the schedule. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| description | string | X | √ | The description of the schedule. | 
| timezone | string | X | √ | The time zone of the schedule. | 
| title | string | X | √ | Title of the resource. | 
| teams | json | X | √ | A list of the teams on the schedule. | 
| html_url | string | X | √ | The API show URL at which the object is accessible. | 
| escalation_policies | json | X | √ | A list of the escalation policies that uses this schedule. | 
| override_sub_schedule | json | X | √ | Specifies schedule overrides for a given time range. | 


