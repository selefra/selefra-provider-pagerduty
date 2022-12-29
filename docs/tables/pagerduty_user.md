# Table: pagerduty_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to 'name', though it is not intended to be an identifier. | 
| timezone | string | X | √ | The preferred time zone name. If null, the account's time zone will be used. | 
| tags | json | X | √ | A list of tags applied on user. | 
| id | string | X | √ | An unique identifier of an user. | 
| email | string | X | √ | The user's email address. | 
| role | string | X | √ | The user role. Account must have the 'read_only_users' ability to set a user as a 'read_only_user' or a 'read_only_limited_user', and must have advanced permissions abilities to set a user as 'observer' or 'restricted_access'. | 
| html_url | string | X | √ | An URL at which the entity is uniquely displayed in the Web app. | 
| job_title | string | X | √ | The user's job title. | 
| notification_rules | json | X | √ | A list of notification rules for the user. | 
| title | string | X | √ | Title of the resource. | 
| invitation_sent | bool | X | √ | If true, the user has an outstanding invitation. | 
| color | string | X | √ | The schedule color. | 
| description | string | X | √ | The user's bio. | 
| type | string | X | √ | The type of object being created. | 
| contact_methods | json | X | √ | A list of contact methods for the user. | 
| name | string | X | √ | The name of the user. | 
| avatar_url | string | X | √ | The URL of the user's avatar. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| teams | json | X | √ | A list of teams to which the user belongs. Account must have the teams ability to set this. | 


