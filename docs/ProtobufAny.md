# ProtobufAny

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | یک URL/نام منبع که نوع پیام protocol buffer سریالایز شده را به طور منحصر به فرد شناسایی می‌کند. این رشته باید حداقل یک کاراکتر \&quot;/\&quot; داشته باشد. آخرین بخش مسیر URL باید نام کاملاً واجد شرایط نوع را نمایش دهد (مانند &#x60;path/google.protobuf.Duration&#x60;). نام باید در فرم کانونیکال باشد (مثلاً، \&quot;.\&quot; ابتدایی پذیرفته نیست).  در عمل، تیم‌ها معمولاً تمام انواعی که انتظار دارند در زمینه Any استفاده شود را به صورت پیش‌کامپایل در باینری قرار می‌دهند. با این حال، برای URL هایی که از طرح &#x60;http&#x60;، &#x60;https&#x60; یا بدون طرح استفاده می‌کنند، می‌توان به طور اختیاری یک سرور نوع تنظیم کرد که URL های نوع را به تعاریف پیام نگاشت می‌کند:  * اگر طرحی ارائه نشود، &#x60;https&#x60; فرض می‌شود. * یک HTTP GET روی URL باید مقدار [google.protobuf.Type][] را   به فرمت باینری ارائه دهد، یا خطا تولید کند. * برنامه‌ها مجاز هستند نتایج جستجو را بر اساس   URL کش کنند، یا آنها را در باینری پیش‌کامپایل کنند تا از هر   جستجویی جلوگیری کنند. بنابراین، سازگاری باینری باید در   تغییرات انواع حفظ شود. (از نام‌های نوع نسخه‌دار برای مدیریت   تغییرات شکست‌آور استفاده کنید.)  توجه: این قابلیت در حال حاضر در نسخه رسمی protobuf در دسترس نیست، و برای URL های نوع که با type.googleapis.com شروع می‌شوند استفاده نمی‌شود.  طرح‌های غیر از &#x60;http&#x60;، &#x60;https&#x60; (یا طرح خالی) ممکن است با معناشناسی خاص پیاده‌سازی استفاده شوند. | [optional] 

## Methods

### NewProtobufAny

`func NewProtobufAny() *ProtobufAny`

NewProtobufAny instantiates a new ProtobufAny object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtobufAnyWithDefaults

`func NewProtobufAnyWithDefaults() *ProtobufAny`

NewProtobufAnyWithDefaults instantiates a new ProtobufAny object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProtobufAny) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtobufAny) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtobufAny) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtobufAny) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


