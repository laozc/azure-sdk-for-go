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

// IotSecuritySolutionClient contains the methods for the IotSecuritySolution group.
// Don't use this type directly, use NewIotSecuritySolutionClient() instead.
type IotSecuritySolutionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIotSecuritySolutionClient creates a new instance of IotSecuritySolutionClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIotSecuritySolutionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IotSecuritySolutionClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &IotSecuritySolutionClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Use this method to create or update yours IoT Security solution
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// solutionName - The name of the IoT Security solution.
// iotSecuritySolutionData - The security solution data
// options - IotSecuritySolutionClientCreateOrUpdateOptions contains the optional parameters for the IotSecuritySolutionClient.CreateOrUpdate
// method.
func (client *IotSecuritySolutionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, solutionName string, iotSecuritySolutionData IoTSecuritySolutionModel, options *IotSecuritySolutionClientCreateOrUpdateOptions) (IotSecuritySolutionClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, solutionName, iotSecuritySolutionData, options)
	if err != nil {
		return IotSecuritySolutionClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IotSecuritySolutionClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IotSecuritySolutionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, iotSecuritySolutionData IoTSecuritySolutionModel, options *IotSecuritySolutionClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iotSecuritySolutionData)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IotSecuritySolutionClient) createOrUpdateHandleResponse(resp *http.Response) (IotSecuritySolutionClientCreateOrUpdateResponse, error) {
	result := IotSecuritySolutionClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionModel); err != nil {
		return IotSecuritySolutionClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Use this method to delete yours IoT Security solution
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// solutionName - The name of the IoT Security solution.
// options - IotSecuritySolutionClientDeleteOptions contains the optional parameters for the IotSecuritySolutionClient.Delete
// method.
func (client *IotSecuritySolutionClient) Delete(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionClientDeleteOptions) (IotSecuritySolutionClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IotSecuritySolutionClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IotSecuritySolutionClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IotSecuritySolutionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - User this method to get details of a specific IoT Security solution based on solution name
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// solutionName - The name of the IoT Security solution.
// options - IotSecuritySolutionClientGetOptions contains the optional parameters for the IotSecuritySolutionClient.Get method.
func (client *IotSecuritySolutionClient) Get(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionClientGetOptions) (IotSecuritySolutionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotSecuritySolutionClient) getCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotSecuritySolutionClient) getHandleResponse(resp *http.Response) (IotSecuritySolutionClientGetResponse, error) {
	result := IotSecuritySolutionClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionModel); err != nil {
		return IotSecuritySolutionClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Use this method to get the list IoT Security solutions organized by resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// options - IotSecuritySolutionClientListByResourceGroupOptions contains the optional parameters for the IotSecuritySolutionClient.ListByResourceGroup
// method.
func (client *IotSecuritySolutionClient) ListByResourceGroup(resourceGroupName string, options *IotSecuritySolutionClientListByResourceGroupOptions) *IotSecuritySolutionClientListByResourceGroupPager {
	return &IotSecuritySolutionClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp IotSecuritySolutionClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IoTSecuritySolutionsList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IotSecuritySolutionClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IotSecuritySolutionClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IotSecuritySolutionClient) listByResourceGroupHandleResponse(resp *http.Response) (IotSecuritySolutionClientListByResourceGroupResponse, error) {
	result := IotSecuritySolutionClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionsList); err != nil {
		return IotSecuritySolutionClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Use this method to get the list of IoT Security solutions by subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - IotSecuritySolutionClientListBySubscriptionOptions contains the optional parameters for the IotSecuritySolutionClient.ListBySubscription
// method.
func (client *IotSecuritySolutionClient) ListBySubscription(options *IotSecuritySolutionClientListBySubscriptionOptions) *IotSecuritySolutionClientListBySubscriptionPager {
	return &IotSecuritySolutionClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IotSecuritySolutionClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IoTSecuritySolutionsList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IotSecuritySolutionClient) listBySubscriptionCreateRequest(ctx context.Context, options *IotSecuritySolutionClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IotSecuritySolutionClient) listBySubscriptionHandleResponse(resp *http.Response) (IotSecuritySolutionClientListBySubscriptionResponse, error) {
	result := IotSecuritySolutionClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionsList); err != nil {
		return IotSecuritySolutionClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Use this method to update existing IoT Security solution tags or user defined resources. To update other fields
// use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// solutionName - The name of the IoT Security solution.
// updateIotSecuritySolutionData - The security solution data
// options - IotSecuritySolutionClientUpdateOptions contains the optional parameters for the IotSecuritySolutionClient.Update
// method.
func (client *IotSecuritySolutionClient) Update(ctx context.Context, resourceGroupName string, solutionName string, updateIotSecuritySolutionData UpdateIotSecuritySolutionData, options *IotSecuritySolutionClientUpdateOptions) (IotSecuritySolutionClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, solutionName, updateIotSecuritySolutionData, options)
	if err != nil {
		return IotSecuritySolutionClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *IotSecuritySolutionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, updateIotSecuritySolutionData UpdateIotSecuritySolutionData, options *IotSecuritySolutionClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, updateIotSecuritySolutionData)
}

// updateHandleResponse handles the Update response.
func (client *IotSecuritySolutionClient) updateHandleResponse(resp *http.Response) (IotSecuritySolutionClientUpdateResponse, error) {
	result := IotSecuritySolutionClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionModel); err != nil {
		return IotSecuritySolutionClientUpdateResponse{}, err
	}
	return result, nil
}
