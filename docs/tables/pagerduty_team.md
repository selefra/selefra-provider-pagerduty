# Table: pagerduty_team

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| members | json | X | √ | A list of members of a team. | 
| tags | json | X | √ | A list of tags applied on team. | 
| id | string | X | √ | An unique identifier of a team. | 
| description | string | X | √ | The description of the team. | 
| html_url | string | X | √ | The API show URL at which the object is accessible. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| type | string | X | √ | The type of object being created. | 
| name | string | X | √ | The name of the team. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| title | string | X | √ | Title of the resource. | 


