# Table: pagerduty_vendor

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| html_url | string | X | √ | The API show URL at which the object is accessible. | 
| summary | string | X | √ | A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier. | 
| self | string | X | √ | The API show URL at which the object is accessible. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name of the vendor. | 
| id | string | X | √ | An unique identifier of the vendor. | 
| integration_guide_url | string | X | √ | Specifies the URL of an integration guide for this vendor. | 
| is_pdcef | bool | X | √ | Indicates the PagerDuty Common Event Format(PD-CEF). | 
| logo_url | string | X | √ | Specifies the URL of a logo identifying the vendor. | 
| thumbnail_url | string | X | √ | Specifies the URL of a small thumbnail image identifying the vendor. | 
| type | string | X | √ | The type of object being created. | 
| long_name | string | X | √ | The full name of the vendor. | 
| alert_creation_default | string | X | √ | Specifies the default method for the alert creation. | 
| generic_service_type | string | X | √ | Specifies the generic service type. | 
| description | string | X | √ | The description of the vendor. | 
| website_url | string | X | √ | The description of the vendor. | 
| alert_creation_editable | bool | X | √ | Indicates whether the default alert creation method can be editable, or not. | 


