# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [notification/v1/notification.proto](#notification_v1_notification-proto)
    - [EmailAddress](#notification-v1-EmailAddress)
    - [KeyValue](#notification-v1-KeyValue)
  
    - [DeliveryStatus](#notification-v1-DeliveryStatus)
    - [MessageType](#notification-v1-MessageType)
  
- [notification/v1/notification_api.proto](#notification_v1_notification_api-proto)
    - [EmailNotifyRequest](#notification-v1-EmailNotifyRequest)
    - [EmailNotifyResponse](#notification-v1-EmailNotifyResponse)
    - [SmsNotifyRequest](#notification-v1-SmsNotifyRequest)
    - [SmsNotifyResponse](#notification-v1-SmsNotifyResponse)
  
    - [NotificationAPI](#notification-v1-NotificationAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="notification_v1_notification-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/v1/notification.proto



<a name="notification-v1-EmailAddress"></a>

### EmailAddress
===== Email =====


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 可选显示名 |
| email | [string](#string) |  | RFC 5322 格式 |






<a name="notification-v1-KeyValue"></a>

### KeyValue
===== Common =====


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="notification-v1-DeliveryStatus"></a>

### DeliveryStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| DELIVERY_STATUS_UNSPECIFIED | 0 |  |
| DELIVERY_QUEUED | 1 |  |
| DELIVERY_SENT | 2 |  |
| DELIVERY_DELIVERED | 3 |  |
| DELIVERY_UNDELIVERED | 4 |  |
| DELIVERY_FAILED | 5 |  |
| DELIVERY_PARTIAL | 6 |  |



<a name="notification-v1-MessageType"></a>

### MessageType


| Name | Number | Description |
| ---- | ------ | ----------- |
| MESSAGE_TYPE_UNSPECIFIED | 0 |  |
| MESSAGE_TYPE_TRANSACTIONAL | 1 | 验证码/通知 |
| MESSAGE_TYPE_PROMOTIONAL | 2 | 营销 |


 

 

 



<a name="notification_v1_notification_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notification/v1/notification_api.proto



<a name="notification-v1-EmailNotifyRequest"></a>

### EmailNotifyRequest







<a name="notification-v1-EmailNotifyResponse"></a>

### EmailNotifyResponse







<a name="notification-v1-SmsNotifyRequest"></a>

### SmsNotifyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | E.164 |
| attempt_id | [string](#string) |  | 我方单目标尝试ID |
| provider_message_id | [string](#string) |  | 供应商消息ID（若有） |
| status | [DeliveryStatus](#notification-v1-DeliveryStatus) |  |  |
| error_code | [string](#string) |  |  |
| error_message | [string](#string) |  |  |






<a name="notification-v1-SmsNotifyResponse"></a>

### SmsNotifyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_region | [string](#string) |  |  |
| attempt_id | [string](#string) |  | 本次请求的聚合ID（可选） |
| error_code | [string](#string) |  |  |
| error_message | [string](#string) |  |  |





 

 

 


<a name="notification-v1-NotificationAPI"></a>

### NotificationAPI
Security config api.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NotifyEmail | [EmailNotifyRequest](#notification-v1-EmailNotifyRequest) | [EmailNotifyResponse](#notification-v1-EmailNotifyResponse) | Send an email notification. |
| NotifySms | [SmsNotifyRequest](#notification-v1-SmsNotifyRequest) | [SmsNotifyResponse](#notification-v1-SmsNotifyResponse) | Send an SMS notification. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

