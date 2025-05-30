/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// PaymentTicketAPIService PaymentTicketAPI service
type PaymentTicketAPIService service

type ApiPaymentTicketValidateRequest struct {
	ctx context.Context
	ApiService *PaymentTicketAPIService
	paymentTicketValidateRequest *PaymentTicketValidateRequest
}

func (r ApiPaymentTicketValidateRequest) PaymentTicketValidateRequest(paymentTicketValidateRequest PaymentTicketValidateRequest) ApiPaymentTicketValidateRequest {
	r.paymentTicketValidateRequest = &paymentTicketValidateRequest
	return r
}

func (r ApiPaymentTicketValidateRequest) Execute() (*PaymentTicketValidateResponse, *http.Response, error) {
	return r.ApiService.PaymentTicketValidateExecute(r)
}

/*
PaymentTicketValidate اعتبارسنجی بلیط پرداخت

بلیط‌های پرداخت برای یکپارچه‌سازی اپلیکیشن‌های کنار با تیم‌های داخلی دیوار طراحی شده‌اند.
با استفاده از این API می‌توانید بلیط پرداخت را اعتبارسنجی کرده و payload صادرکننده را دریافت کنید.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaymentTicketValidateRequest
*/
func (a *PaymentTicketAPIService) PaymentTicketValidate(ctx context.Context) ApiPaymentTicketValidateRequest {
	return ApiPaymentTicketValidateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentTicketValidateResponse
func (a *PaymentTicketAPIService) PaymentTicketValidateExecute(r ApiPaymentTicketValidateRequest) (*PaymentTicketValidateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentTicketValidateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentTicketAPIService.PaymentTicketValidate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/open-platform/payment-ticket/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paymentTicketValidateRequest == nil {
		return localVarReturnValue, nil, reportError("paymentTicketValidateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentTicketValidateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GooglerpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
