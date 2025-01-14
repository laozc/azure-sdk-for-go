//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AssessmentsMetadataClient contains the methods for the AssessmentsMetadata group.
// Don't use this type directly, use NewAssessmentsMetadataClient() instead.
type AssessmentsMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAssessmentsMetadataClient creates a new instance of AssessmentsMetadataClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssessmentsMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AssessmentsMetadataClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AssessmentsMetadataClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateInSubscription - Create metadata information on an assessment type in a specific subscription
// If the operation fails it returns an *azcore.ResponseError type.
// assessmentMetadataName - The Assessment Key - Unique key for the assessment type
// assessmentMetadata - AssessmentMetadata object
// options - AssessmentsMetadataClientCreateInSubscriptionOptions contains the optional parameters for the AssessmentsMetadataClient.CreateInSubscription
// method.
func (client *AssessmentsMetadataClient) CreateInSubscription(ctx context.Context, assessmentMetadataName string, assessmentMetadata AssessmentMetadataResponse, options *AssessmentsMetadataClientCreateInSubscriptionOptions) (AssessmentsMetadataClientCreateInSubscriptionResponse, error) {
	req, err := client.createInSubscriptionCreateRequest(ctx, assessmentMetadataName, assessmentMetadata, options)
	if err != nil {
		return AssessmentsMetadataClientCreateInSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsMetadataClientCreateInSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsMetadataClientCreateInSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.createInSubscriptionHandleResponse(resp)
}

// createInSubscriptionCreateRequest creates the CreateInSubscription request.
func (client *AssessmentsMetadataClient) createInSubscriptionCreateRequest(ctx context.Context, assessmentMetadataName string, assessmentMetadata AssessmentMetadataResponse, options *AssessmentsMetadataClientCreateInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}"
	if assessmentMetadataName == "" {
		return nil, errors.New("parameter assessmentMetadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentMetadataName}", url.PathEscape(assessmentMetadataName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, assessmentMetadata)
}

// createInSubscriptionHandleResponse handles the CreateInSubscription response.
func (client *AssessmentsMetadataClient) createInSubscriptionHandleResponse(resp *http.Response) (AssessmentsMetadataClientCreateInSubscriptionResponse, error) {
	result := AssessmentsMetadataClientCreateInSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentMetadataResponse); err != nil {
		return AssessmentsMetadataClientCreateInSubscriptionResponse{}, err
	}
	return result, nil
}

// DeleteInSubscription - Delete metadata information on an assessment type in a specific subscription, will cause the deletion
// of all the assessments of that type in that subscription
// If the operation fails it returns an *azcore.ResponseError type.
// assessmentMetadataName - The Assessment Key - Unique key for the assessment type
// options - AssessmentsMetadataClientDeleteInSubscriptionOptions contains the optional parameters for the AssessmentsMetadataClient.DeleteInSubscription
// method.
func (client *AssessmentsMetadataClient) DeleteInSubscription(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientDeleteInSubscriptionOptions) (AssessmentsMetadataClientDeleteInSubscriptionResponse, error) {
	req, err := client.deleteInSubscriptionCreateRequest(ctx, assessmentMetadataName, options)
	if err != nil {
		return AssessmentsMetadataClientDeleteInSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsMetadataClientDeleteInSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsMetadataClientDeleteInSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return AssessmentsMetadataClientDeleteInSubscriptionResponse{RawResponse: resp}, nil
}

// deleteInSubscriptionCreateRequest creates the DeleteInSubscription request.
func (client *AssessmentsMetadataClient) deleteInSubscriptionCreateRequest(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientDeleteInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}"
	if assessmentMetadataName == "" {
		return nil, errors.New("parameter assessmentMetadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentMetadataName}", url.PathEscape(assessmentMetadataName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get metadata information on an assessment type
// If the operation fails it returns an *azcore.ResponseError type.
// assessmentMetadataName - The Assessment Key - Unique key for the assessment type
// options - AssessmentsMetadataClientGetOptions contains the optional parameters for the AssessmentsMetadataClient.Get method.
func (client *AssessmentsMetadataClient) Get(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientGetOptions) (AssessmentsMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, assessmentMetadataName, options)
	if err != nil {
		return AssessmentsMetadataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssessmentsMetadataClient) getCreateRequest(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}"
	if assessmentMetadataName == "" {
		return nil, errors.New("parameter assessmentMetadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentMetadataName}", url.PathEscape(assessmentMetadataName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessmentsMetadataClient) getHandleResponse(resp *http.Response) (AssessmentsMetadataClientGetResponse, error) {
	result := AssessmentsMetadataClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentMetadataResponse); err != nil {
		return AssessmentsMetadataClientGetResponse{}, err
	}
	return result, nil
}

// GetInSubscription - Get metadata information on an assessment type in a specific subscription
// If the operation fails it returns an *azcore.ResponseError type.
// assessmentMetadataName - The Assessment Key - Unique key for the assessment type
// options - AssessmentsMetadataClientGetInSubscriptionOptions contains the optional parameters for the AssessmentsMetadataClient.GetInSubscription
// method.
func (client *AssessmentsMetadataClient) GetInSubscription(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientGetInSubscriptionOptions) (AssessmentsMetadataClientGetInSubscriptionResponse, error) {
	req, err := client.getInSubscriptionCreateRequest(ctx, assessmentMetadataName, options)
	if err != nil {
		return AssessmentsMetadataClientGetInSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssessmentsMetadataClientGetInSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssessmentsMetadataClientGetInSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInSubscriptionHandleResponse(resp)
}

// getInSubscriptionCreateRequest creates the GetInSubscription request.
func (client *AssessmentsMetadataClient) getInSubscriptionCreateRequest(ctx context.Context, assessmentMetadataName string, options *AssessmentsMetadataClientGetInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}"
	if assessmentMetadataName == "" {
		return nil, errors.New("parameter assessmentMetadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentMetadataName}", url.PathEscape(assessmentMetadataName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInSubscriptionHandleResponse handles the GetInSubscription response.
func (client *AssessmentsMetadataClient) getInSubscriptionHandleResponse(resp *http.Response) (AssessmentsMetadataClientGetInSubscriptionResponse, error) {
	result := AssessmentsMetadataClientGetInSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentMetadataResponse); err != nil {
		return AssessmentsMetadataClientGetInSubscriptionResponse{}, err
	}
	return result, nil
}

// List - Get metadata information on all assessment types
// If the operation fails it returns an *azcore.ResponseError type.
// options - AssessmentsMetadataClientListOptions contains the optional parameters for the AssessmentsMetadataClient.List
// method.
func (client *AssessmentsMetadataClient) List(options *AssessmentsMetadataClientListOptions) *AssessmentsMetadataClientListPager {
	return &AssessmentsMetadataClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AssessmentsMetadataClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssessmentMetadataResponseList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AssessmentsMetadataClient) listCreateRequest(ctx context.Context, options *AssessmentsMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Security/assessmentMetadata"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssessmentsMetadataClient) listHandleResponse(resp *http.Response) (AssessmentsMetadataClientListResponse, error) {
	result := AssessmentsMetadataClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentMetadataResponseList); err != nil {
		return AssessmentsMetadataClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get metadata information on all assessment types in a specific subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - AssessmentsMetadataClientListBySubscriptionOptions contains the optional parameters for the AssessmentsMetadataClient.ListBySubscription
// method.
func (client *AssessmentsMetadataClient) ListBySubscription(options *AssessmentsMetadataClientListBySubscriptionOptions) *AssessmentsMetadataClientListBySubscriptionPager {
	return &AssessmentsMetadataClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AssessmentsMetadataClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssessmentMetadataResponseList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AssessmentsMetadataClient) listBySubscriptionCreateRequest(ctx context.Context, options *AssessmentsMetadataClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AssessmentsMetadataClient) listBySubscriptionHandleResponse(resp *http.Response) (AssessmentsMetadataClientListBySubscriptionResponse, error) {
	result := AssessmentsMetadataClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentMetadataResponseList); err != nil {
		return AssessmentsMetadataClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
