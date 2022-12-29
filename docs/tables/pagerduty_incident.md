# Table: pagerduty_incident

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| body | json | X | √ | Describes the additional incident details. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| created_at | timestamp | X | √ | The date/time the incident was first triggered. | 
| last_status_change_at | timestamp | X | √ | The time at which the status of the incident last changed. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| alert_counts | json | X | √ | Describes the count of triggered and resolved alerts. | 
| first_trigger_log_entry | json | X | √ | Specifies the first log entry when the incident was triggered. | 
| teams | json | X | √ | The teams involved in the incident's lifecycle. | 
| priority | json | X | √ | Specifies the priority set for this incident. | 
| html_url | string | X | √ | The API show URL at which the object is accessible. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | An unique identifier of the incident. | 
| is_mergeable | bool | X | √ | Indicates whether the incident's alerts can be merged with another incident, or not. | 
| type | string | X | √ | The type of object being created. | 
| assignments | json | X | √ | A list of all assignments for this incident. This list will be empty if the 'Incident.status' is resolved. | 
| escalation_policy | json | X | √ | Specifies the escalation policy assigned to this incident. | 
| status | string | X | √ | The current status of the incident. | 
| description | string | X | √ | The description of the incident. | 
| incident_key | string | X | √ | The incident's de-duplication key. | 
| conference_bridge | json | X | √ | Specifies the contact information that allows responders to easily connect and collaborate during major incident response. | 
| urgency | string | X | √ | The current urgency of the incident. | 
| last_status_change_by | json | X | √ | The agent (user, service or integration) that created or modified the incident log entry. | 
| incident_number | int | X | √ | The number of the incident. This is unique across your account. | 
| service | json | X | √ | Specifies the information about the impacted service. | 
| acknowledgements | json | X | √ | A list of all acknowledgements for this incident. This list will be empty if the 'Incident.status' is resolved or triggered. | 
| resolve_reason | json | X | √ | Specifies the reason the incident was resolved. | 
| pending_actions | json | X | √ | A list of pending_actions on the incident. A pending_action object contains a type of action which can be escalate, unacknowledge, resolve or urgency_change. A pending_action object contains at, the time at which the action will take place. An urgency_change pending_action will contain to, the urgency that the incident will change to. | 


