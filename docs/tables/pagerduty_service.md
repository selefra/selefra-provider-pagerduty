# Table: pagerduty_service

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| scheduled_actions | json | X | √ | An array containing scheduled actions for the service. | 
| support_hours | json | X | √ | Defines the service's support hours | 
| id | string | X | √ | An unique identifier of a service. | 
| acknowledgement_timeout | int | X | √ | Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled. | 
| alert_creation | string | X | √ | Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging. | 
| last_incident_timestamp | timestamp | X | √ | The date/time when the most recent incident was created for this service. | 
| teams | json | X | √ | The set of teams associated with this service. | 
| integrations | json | X | √ | An array containing integrations that belong to this service. If integrations is passed as an argument, these are full objects - otherwise, these are references. | 
| name | string | X | √ | The name of the service. | 
| description | string | X | √ | The user-provided description of the service. | 
| auto_resolve_timeout | int | X | √ | Time in seconds that an incident is automatically resolved if left open for that long. Value is null if the feature is disabled. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| type | string | X | √ | The type of object being created. | 
| alert_grouping_parameters | json | X | √ | Defines how alerts on this service will be automatically grouped into incidents. Note that the alert grouping features are available only on certain plans. | 
| incident_urgency_rule | json | X | √ | A list of incident urgency rules. | 
| title | string | X | √ | Title of the resource. | 
| status | string | X | √ | The current state of the service. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| create_at | timestamp | X | √ | The date/time when this service was created. | 
| html_url | string | X | √ | An URL at which the entity is uniquely displayed in the Web app. | 
| escalation_policy | json | X | √ | Escalation policy associated with the service. | 


