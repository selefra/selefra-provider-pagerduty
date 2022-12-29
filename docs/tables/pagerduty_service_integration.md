# Table: pagerduty_service_integration

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| integration_key | string | X | √ | Specify the integration key for the service integration. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| type | string | X | √ | The type of object being created. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name of this integration. | 
| id | string | X | √ | An unique identifier of the integration. | 
| service_id | string | X | √ | An unique identifier of the queried service. | 
| created_at | timestamp | X | √ | The date/time when this integration was created. | 
| vendor | json | X | √ | Describes the information about a specific type of integration. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| html_url | string | X | √ | An URL at which the entity is uniquely displayed in the Web app. | 
| integration_email | string | X | √ | Specify for generic_email_inbound_integration. Must be set to an email address @your-subdomain.pagerduty.com. | 
| service | json | X | √ | Describes the information about the queried service. | 


