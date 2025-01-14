//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// BackupEnginesClientListPager provides operations for iterating over paged responses.
type BackupEnginesClientListPager struct {
	client    *BackupEnginesClient
	current   BackupEnginesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupEnginesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupEnginesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupEnginesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.BackupEngineBaseResourceList.NextLink == nil || len(*p.current.BackupEngineBaseResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupEnginesClientListResponse page.
func (p *BackupEnginesClientListPager) PageResponse() BackupEnginesClientListResponse {
	return p.current
}

// BackupJobsClientListPager provides operations for iterating over paged responses.
type BackupJobsClientListPager struct {
	client    *BackupJobsClient
	current   BackupJobsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupJobsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupJobsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupJobsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobResourceList.NextLink == nil || len(*p.current.JobResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupJobsClientListResponse page.
func (p *BackupJobsClientListPager) PageResponse() BackupJobsClientListResponse {
	return p.current
}

// BackupPoliciesClientListPager provides operations for iterating over paged responses.
type BackupPoliciesClientListPager struct {
	client    *BackupPoliciesClient
	current   BackupPoliciesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupPoliciesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupPoliciesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupPoliciesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProtectionPolicyResourceList.NextLink == nil || len(*p.current.ProtectionPolicyResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupPoliciesClientListResponse page.
func (p *BackupPoliciesClientListPager) PageResponse() BackupPoliciesClientListResponse {
	return p.current
}

// BackupProtectableItemsClientListPager provides operations for iterating over paged responses.
type BackupProtectableItemsClientListPager struct {
	client    *BackupProtectableItemsClient
	current   BackupProtectableItemsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupProtectableItemsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupProtectableItemsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupProtectableItemsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadProtectableItemResourceList.NextLink == nil || len(*p.current.WorkloadProtectableItemResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupProtectableItemsClientListResponse page.
func (p *BackupProtectableItemsClientListPager) PageResponse() BackupProtectableItemsClientListResponse {
	return p.current
}

// BackupProtectedItemsClientListPager provides operations for iterating over paged responses.
type BackupProtectedItemsClientListPager struct {
	client    *BackupProtectedItemsClient
	current   BackupProtectedItemsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupProtectedItemsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupProtectedItemsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupProtectedItemsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProtectedItemResourceList.NextLink == nil || len(*p.current.ProtectedItemResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupProtectedItemsClientListResponse page.
func (p *BackupProtectedItemsClientListPager) PageResponse() BackupProtectedItemsClientListResponse {
	return p.current
}

// BackupProtectionContainersClientListPager provides operations for iterating over paged responses.
type BackupProtectionContainersClientListPager struct {
	client    *BackupProtectionContainersClient
	current   BackupProtectionContainersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupProtectionContainersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupProtectionContainersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupProtectionContainersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProtectionContainerResourceList.NextLink == nil || len(*p.current.ProtectionContainerResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupProtectionContainersClientListResponse page.
func (p *BackupProtectionContainersClientListPager) PageResponse() BackupProtectionContainersClientListResponse {
	return p.current
}

// BackupProtectionIntentClientListPager provides operations for iterating over paged responses.
type BackupProtectionIntentClientListPager struct {
	client    *BackupProtectionIntentClient
	current   BackupProtectionIntentClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupProtectionIntentClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupProtectionIntentClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupProtectionIntentClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProtectionIntentResourceList.NextLink == nil || len(*p.current.ProtectionIntentResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupProtectionIntentClientListResponse page.
func (p *BackupProtectionIntentClientListPager) PageResponse() BackupProtectionIntentClientListResponse {
	return p.current
}

// BackupWorkloadItemsClientListPager provides operations for iterating over paged responses.
type BackupWorkloadItemsClientListPager struct {
	client    *BackupWorkloadItemsClient
	current   BackupWorkloadItemsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, BackupWorkloadItemsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *BackupWorkloadItemsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *BackupWorkloadItemsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WorkloadItemResourceList.NextLink == nil || len(*p.current.WorkloadItemResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current BackupWorkloadItemsClientListResponse page.
func (p *BackupWorkloadItemsClientListPager) PageResponse() BackupWorkloadItemsClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ClientDiscoveryResponse.NextLink == nil || len(*p.current.ClientDiscoveryResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// ProtectableContainersClientListPager provides operations for iterating over paged responses.
type ProtectableContainersClientListPager struct {
	client    *ProtectableContainersClient
	current   ProtectableContainersClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ProtectableContainersClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ProtectableContainersClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ProtectableContainersClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProtectableContainerResourceList.NextLink == nil || len(*p.current.ProtectableContainerResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ProtectableContainersClientListResponse page.
func (p *ProtectableContainersClientListPager) PageResponse() ProtectableContainersClientListResponse {
	return p.current
}

// RecoveryPointsClientListPager provides operations for iterating over paged responses.
type RecoveryPointsClientListPager struct {
	client    *RecoveryPointsClient
	current   RecoveryPointsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecoveryPointsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecoveryPointsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecoveryPointsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecoveryPointResourceList.NextLink == nil || len(*p.current.RecoveryPointResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecoveryPointsClientListResponse page.
func (p *RecoveryPointsClientListPager) PageResponse() RecoveryPointsClientListResponse {
	return p.current
}

// RecoveryPointsRecommendedForMoveClientListPager provides operations for iterating over paged responses.
type RecoveryPointsRecommendedForMoveClientListPager struct {
	client    *RecoveryPointsRecommendedForMoveClient
	current   RecoveryPointsRecommendedForMoveClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecoveryPointsRecommendedForMoveClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecoveryPointsRecommendedForMoveClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecoveryPointsRecommendedForMoveClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecoveryPointResourceList.NextLink == nil || len(*p.current.RecoveryPointResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecoveryPointsRecommendedForMoveClientListResponse page.
func (p *RecoveryPointsRecommendedForMoveClientListPager) PageResponse() RecoveryPointsRecommendedForMoveClientListResponse {
	return p.current
}

// ResourceGuardProxiesClientGetPager provides operations for iterating over paged responses.
type ResourceGuardProxiesClientGetPager struct {
	client    *ResourceGuardProxiesClient
	current   ResourceGuardProxiesClientGetResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ResourceGuardProxiesClientGetResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ResourceGuardProxiesClientGetPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ResourceGuardProxiesClientGetPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceGuardProxyBaseResourceList.NextLink == nil || len(*p.current.ResourceGuardProxyBaseResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.getHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ResourceGuardProxiesClientGetResponse page.
func (p *ResourceGuardProxiesClientGetPager) PageResponse() ResourceGuardProxiesClientGetResponse {
	return p.current
}
