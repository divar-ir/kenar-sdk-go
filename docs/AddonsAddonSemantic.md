# AddonsAddonSemantic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | Pointer to [**AddonsAddonSemanticPaymentMethod**](AddonsAddonSemanticPaymentMethod.md) |  | [optional] 
**IdentityVerificationResult** | Pointer to [**AddonSemanticIdentityVerificationResult**](AddonSemanticIdentityVerificationResult.md) |  | [optional] 
**PostVerificationResult** | Pointer to [**AddonSemanticPostVerificationResult**](AddonSemanticPostVerificationResult.md) |  | [optional] 
**Status** | Pointer to [**AddonsAddonSemanticStatus**](AddonsAddonSemanticStatus.md) |  | [optional] 
**IdentityVerificationLastSuccessfulStage** | Pointer to [**AddonSemanticIdentityVerificationStage**](AddonSemanticIdentityVerificationStage.md) |  | [optional] 
**IdentityVerificationFailureReason** | Pointer to [**AddonSemanticIdentityVerificationStage**](AddonSemanticIdentityVerificationStage.md) |  | [optional] 
**CarVerificationLastSuccessfulStage** | Pointer to [**AddonSemanticCarVerificationStage**](AddonSemanticCarVerificationStage.md) |  | [optional] 
**CarVerificationFailureReason** | Pointer to [**AddonSemanticCarVerificationStage**](AddonSemanticCarVerificationStage.md) |  | [optional] 
**OwnershipResult** | Pointer to [**AddonSemanticOwnershipResult**](AddonSemanticOwnershipResult.md) |  | [optional] 
**InspectionResult** | Pointer to [**AddonSemanticInspectionResult**](AddonSemanticInspectionResult.md) |  | [optional] 
**NewFaceVerificationResult** | Pointer to [**AddonSemanticNewFaceVerificationResult**](AddonSemanticNewFaceVerificationResult.md) |  | [optional] 

## Methods

### NewAddonsAddonSemantic

`func NewAddonsAddonSemantic() *AddonsAddonSemantic`

NewAddonsAddonSemantic instantiates a new AddonsAddonSemantic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsAddonSemanticWithDefaults

`func NewAddonsAddonSemanticWithDefaults() *AddonsAddonSemantic`

NewAddonsAddonSemanticWithDefaults instantiates a new AddonsAddonSemantic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *AddonsAddonSemantic) GetPaymentMethod() AddonsAddonSemanticPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AddonsAddonSemantic) GetPaymentMethodOk() (*AddonsAddonSemanticPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AddonsAddonSemantic) SetPaymentMethod(v AddonsAddonSemanticPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *AddonsAddonSemantic) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetIdentityVerificationResult

`func (o *AddonsAddonSemantic) GetIdentityVerificationResult() AddonSemanticIdentityVerificationResult`

GetIdentityVerificationResult returns the IdentityVerificationResult field if non-nil, zero value otherwise.

### GetIdentityVerificationResultOk

`func (o *AddonsAddonSemantic) GetIdentityVerificationResultOk() (*AddonSemanticIdentityVerificationResult, bool)`

GetIdentityVerificationResultOk returns a tuple with the IdentityVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVerificationResult

`func (o *AddonsAddonSemantic) SetIdentityVerificationResult(v AddonSemanticIdentityVerificationResult)`

SetIdentityVerificationResult sets IdentityVerificationResult field to given value.

### HasIdentityVerificationResult

`func (o *AddonsAddonSemantic) HasIdentityVerificationResult() bool`

HasIdentityVerificationResult returns a boolean if a field has been set.

### GetPostVerificationResult

`func (o *AddonsAddonSemantic) GetPostVerificationResult() AddonSemanticPostVerificationResult`

GetPostVerificationResult returns the PostVerificationResult field if non-nil, zero value otherwise.

### GetPostVerificationResultOk

`func (o *AddonsAddonSemantic) GetPostVerificationResultOk() (*AddonSemanticPostVerificationResult, bool)`

GetPostVerificationResultOk returns a tuple with the PostVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostVerificationResult

`func (o *AddonsAddonSemantic) SetPostVerificationResult(v AddonSemanticPostVerificationResult)`

SetPostVerificationResult sets PostVerificationResult field to given value.

### HasPostVerificationResult

`func (o *AddonsAddonSemantic) HasPostVerificationResult() bool`

HasPostVerificationResult returns a boolean if a field has been set.

### GetStatus

`func (o *AddonsAddonSemantic) GetStatus() AddonsAddonSemanticStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddonsAddonSemantic) GetStatusOk() (*AddonsAddonSemanticStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddonsAddonSemantic) SetStatus(v AddonsAddonSemanticStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddonsAddonSemantic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIdentityVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) GetIdentityVerificationLastSuccessfulStage() AddonSemanticIdentityVerificationStage`

GetIdentityVerificationLastSuccessfulStage returns the IdentityVerificationLastSuccessfulStage field if non-nil, zero value otherwise.

### GetIdentityVerificationLastSuccessfulStageOk

`func (o *AddonsAddonSemantic) GetIdentityVerificationLastSuccessfulStageOk() (*AddonSemanticIdentityVerificationStage, bool)`

GetIdentityVerificationLastSuccessfulStageOk returns a tuple with the IdentityVerificationLastSuccessfulStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) SetIdentityVerificationLastSuccessfulStage(v AddonSemanticIdentityVerificationStage)`

SetIdentityVerificationLastSuccessfulStage sets IdentityVerificationLastSuccessfulStage field to given value.

### HasIdentityVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) HasIdentityVerificationLastSuccessfulStage() bool`

HasIdentityVerificationLastSuccessfulStage returns a boolean if a field has been set.

### GetIdentityVerificationFailureReason

`func (o *AddonsAddonSemantic) GetIdentityVerificationFailureReason() AddonSemanticIdentityVerificationStage`

GetIdentityVerificationFailureReason returns the IdentityVerificationFailureReason field if non-nil, zero value otherwise.

### GetIdentityVerificationFailureReasonOk

`func (o *AddonsAddonSemantic) GetIdentityVerificationFailureReasonOk() (*AddonSemanticIdentityVerificationStage, bool)`

GetIdentityVerificationFailureReasonOk returns a tuple with the IdentityVerificationFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVerificationFailureReason

`func (o *AddonsAddonSemantic) SetIdentityVerificationFailureReason(v AddonSemanticIdentityVerificationStage)`

SetIdentityVerificationFailureReason sets IdentityVerificationFailureReason field to given value.

### HasIdentityVerificationFailureReason

`func (o *AddonsAddonSemantic) HasIdentityVerificationFailureReason() bool`

HasIdentityVerificationFailureReason returns a boolean if a field has been set.

### GetCarVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) GetCarVerificationLastSuccessfulStage() AddonSemanticCarVerificationStage`

GetCarVerificationLastSuccessfulStage returns the CarVerificationLastSuccessfulStage field if non-nil, zero value otherwise.

### GetCarVerificationLastSuccessfulStageOk

`func (o *AddonsAddonSemantic) GetCarVerificationLastSuccessfulStageOk() (*AddonSemanticCarVerificationStage, bool)`

GetCarVerificationLastSuccessfulStageOk returns a tuple with the CarVerificationLastSuccessfulStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) SetCarVerificationLastSuccessfulStage(v AddonSemanticCarVerificationStage)`

SetCarVerificationLastSuccessfulStage sets CarVerificationLastSuccessfulStage field to given value.

### HasCarVerificationLastSuccessfulStage

`func (o *AddonsAddonSemantic) HasCarVerificationLastSuccessfulStage() bool`

HasCarVerificationLastSuccessfulStage returns a boolean if a field has been set.

### GetCarVerificationFailureReason

`func (o *AddonsAddonSemantic) GetCarVerificationFailureReason() AddonSemanticCarVerificationStage`

GetCarVerificationFailureReason returns the CarVerificationFailureReason field if non-nil, zero value otherwise.

### GetCarVerificationFailureReasonOk

`func (o *AddonsAddonSemantic) GetCarVerificationFailureReasonOk() (*AddonSemanticCarVerificationStage, bool)`

GetCarVerificationFailureReasonOk returns a tuple with the CarVerificationFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarVerificationFailureReason

`func (o *AddonsAddonSemantic) SetCarVerificationFailureReason(v AddonSemanticCarVerificationStage)`

SetCarVerificationFailureReason sets CarVerificationFailureReason field to given value.

### HasCarVerificationFailureReason

`func (o *AddonsAddonSemantic) HasCarVerificationFailureReason() bool`

HasCarVerificationFailureReason returns a boolean if a field has been set.

### GetOwnershipResult

`func (o *AddonsAddonSemantic) GetOwnershipResult() AddonSemanticOwnershipResult`

GetOwnershipResult returns the OwnershipResult field if non-nil, zero value otherwise.

### GetOwnershipResultOk

`func (o *AddonsAddonSemantic) GetOwnershipResultOk() (*AddonSemanticOwnershipResult, bool)`

GetOwnershipResultOk returns a tuple with the OwnershipResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipResult

`func (o *AddonsAddonSemantic) SetOwnershipResult(v AddonSemanticOwnershipResult)`

SetOwnershipResult sets OwnershipResult field to given value.

### HasOwnershipResult

`func (o *AddonsAddonSemantic) HasOwnershipResult() bool`

HasOwnershipResult returns a boolean if a field has been set.

### GetInspectionResult

`func (o *AddonsAddonSemantic) GetInspectionResult() AddonSemanticInspectionResult`

GetInspectionResult returns the InspectionResult field if non-nil, zero value otherwise.

### GetInspectionResultOk

`func (o *AddonsAddonSemantic) GetInspectionResultOk() (*AddonSemanticInspectionResult, bool)`

GetInspectionResultOk returns a tuple with the InspectionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectionResult

`func (o *AddonsAddonSemantic) SetInspectionResult(v AddonSemanticInspectionResult)`

SetInspectionResult sets InspectionResult field to given value.

### HasInspectionResult

`func (o *AddonsAddonSemantic) HasInspectionResult() bool`

HasInspectionResult returns a boolean if a field has been set.

### GetNewFaceVerificationResult

`func (o *AddonsAddonSemantic) GetNewFaceVerificationResult() AddonSemanticNewFaceVerificationResult`

GetNewFaceVerificationResult returns the NewFaceVerificationResult field if non-nil, zero value otherwise.

### GetNewFaceVerificationResultOk

`func (o *AddonsAddonSemantic) GetNewFaceVerificationResultOk() (*AddonSemanticNewFaceVerificationResult, bool)`

GetNewFaceVerificationResultOk returns a tuple with the NewFaceVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFaceVerificationResult

`func (o *AddonsAddonSemantic) SetNewFaceVerificationResult(v AddonSemanticNewFaceVerificationResult)`

SetNewFaceVerificationResult sets NewFaceVerificationResult field to given value.

### HasNewFaceVerificationResult

`func (o *AddonsAddonSemantic) HasNewFaceVerificationResult() bool`

HasNewFaceVerificationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


