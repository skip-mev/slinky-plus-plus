<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [cosmwasm/slpp/v1/slpp.proto](#cosmwasm/slpp/v1/slpp.proto)
    - [AVS](#cosmwasm.slpp.v1.AVS)
    - [GetAVSRequest](#cosmwasm.slpp.v1.GetAVSRequest)
    - [GetAVSResponse](#cosmwasm.slpp.v1.GetAVSResponse)
    - [MsgRegisterAVS](#cosmwasm.slpp.v1.MsgRegisterAVS)
    - [MsgRegisterAVSResponse](#cosmwasm.slpp.v1.MsgRegisterAVSResponse)
  
    - [Msg](#cosmwasm.slpp.v1.Msg)
    - [Query](#cosmwasm.slpp.v1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cosmwasm/slpp/v1/slpp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmwasm/slpp/v1/slpp.proto



<a name="cosmwasm.slpp.v1.AVS"></a>

### AVS
AVS represents the state-ful information stored per AVS


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `sidecar_docker_image` | [string](#string) |  |  |






<a name="cosmwasm.slpp.v1.GetAVSRequest"></a>

### GetAVSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |






<a name="cosmwasm.slpp.v1.GetAVSResponse"></a>

### GetAVSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `avs` | [AVS](#cosmwasm.slpp.v1.AVS) |  |  |






<a name="cosmwasm.slpp.v1.MsgRegisterAVS"></a>

### MsgRegisterAVS
MsgRegisterAVS defines a message-type handled by the x/slpp module for ingressing a new AVS.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_bin` | [bytes](#bytes) |  |  |
| `sidecar_docker_image` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `instantiate_msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on instantiation |






<a name="cosmwasm.slpp.v1.MsgRegisterAVSResponse"></a>

### MsgRegisterAVSResponse
MsgRegisterAVSResponse defines the Msg/RegisterAVS response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmwasm.slpp.v1.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RegisterAVS` | [MsgRegisterAVS](#cosmwasm.slpp.v1.MsgRegisterAVS) | [MsgRegisterAVSResponse](#cosmwasm.slpp.v1.MsgRegisterAVSResponse) |  | |


<a name="cosmwasm.slpp.v1.Query"></a>

### Query
Query is the query service for the x/slpp module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetAVS` | [GetAVSRequest](#cosmwasm.slpp.v1.GetAVSRequest) | [GetAVSResponse](#cosmwasm.slpp.v1.GetAVSResponse) |  | GET|/slpp/v1/get_avs/{id}|

 <!-- end services -->



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

