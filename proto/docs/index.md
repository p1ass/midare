# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth.proto](#auth-proto)
    - [ExchangeCodeReq](#-ExchangeCodeReq)
    - [StartAuthorizationResp](#-StartAuthorizationResp)
  
    - [AuthService](#-AuthService)
  
- [midare.proto](#midare-proto)
    - [GetAwakePeriodsResp](#-GetAwakePeriodsResp)
    - [GetMeResp](#-GetMeResp)
    - [Period](#-Period)
    - [Tweet](#-Tweet)
    - [User](#-User)
  
    - [MidareService](#-MidareService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth.proto



<a name="-ExchangeCodeReq"></a>

### ExchangeCodeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | 認可コード |
| state | [string](#string) |  | OAuthのstate |






<a name="-StartAuthorizationResp"></a>

### StartAuthorizationResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorization_url | [string](#string) |  | 認可エンドポイントのURL |





 

 

 


<a name="-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartAuthorization | [.google.protobuf.Empty](#google-protobuf-Empty) | [.StartAuthorizationResp](#StartAuthorizationResp) | StartAuthorization はTwitter APIのOAuth 2.0 Authorization Code Flowを開始する |
| ExchangeCode | [.ExchangeCodeReq](#ExchangeCodeReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | ExchangeCode は認可レスポンスに含まれる認可コードを利用して、アクセストークンをサーバー再度に保存する。 |

 



<a name="midare-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## midare.proto



<a name="-GetAwakePeriodsResp"></a>

### GetAwakePeriodsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| awake_periods | [Period](#Period) | repeated |  |
| share_url | [string](#string) |  | 動的OGP画像が表示されるシェア用のURL |






<a name="-GetMeResp"></a>

### GetMeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#User) |  | ログインユーザの情報 |






<a name="-Period"></a>

### Period
Periods は起きている期間を表す


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oki_tweet | [Tweet](#Tweet) |  | 起きたときのツイート |
| ne_tweet | [Tweet](#Tweet) |  | 寝たときのツイート |






<a name="-Tweet"></a>

### Tweet
Tweet はTwitterのツイートを表す


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 例: &#34;1521682005137977345&#34; |
| text | [string](#string) |  | 例: &#34;おはよう&#34; |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="-User"></a>

### User
User はユーザ情報を表す


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 例: &#34;1032935958964973568&#34; |
| name | [string](#string) |  | 例: &#34;ぷらす&#34; |
| screen_name | [string](#string) |  | 例: &#34;p1ass&#34; |
| image_url | [string](#string) |  | プロフィール画像のURL |





 

 

 


<a name="-MidareService"></a>

### MidareService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetMe | [.google.protobuf.Empty](#google-protobuf-Empty) | [.GetMeResp](#GetMeResp) | GetMe はログインユーザ情報を取得する。 |
| GetAwakePeriods | [.google.protobuf.Empty](#google-protobuf-Empty) | [.GetAwakePeriodsResp](#GetAwakePeriodsResp) | GetAwakePeriods はログインユーザの起きている期間を取得する。 |

 



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

