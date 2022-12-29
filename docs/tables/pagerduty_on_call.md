# Table: pagerduty_on_call

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| escalation_level | int | X | √ | The escalation level for the on-call. | 
| start | timestamp | X | √ | The start of the on-call. If null, the on-call is a permanent user on-call. | 
| end | timestamp | X | √ | The end of the on-call. If null, the user does not go off-call. | 
| escalation_policy | json | X | √ | The escalation_policy object. | 
| schedule | json | X | √ | The schedule object. | 
| user_on_call | json | X | √ | The user object. | 


