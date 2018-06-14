package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// WorkflowsClient is the REST API for Azure Logic Apps.
type WorkflowsClient struct {
	BaseClient
}

// NewWorkflowsClient creates an instance of the WorkflowsClient client.
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return NewWorkflowsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowsClientWithBaseURI creates an instance of the WorkflowsClient client.
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return WorkflowsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. workflow is the workflow.
func (client WorkflowsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workflowName, workflow)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkflowsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", pathParameters),
		autorest.WithJSON(workflow),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) CreateOrUpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name.
func (client WorkflowsClient) Delete(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, workflowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkflowsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disable disables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name.
func (client WorkflowsClient) Disable(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.DisablePreparer(ctx, resourceGroupName, workflowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Disable", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Disable", resp, "Failure sending request")
		return
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client WorkflowsClient) DisablePreparer(ctx context.Context, resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) DisableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable enables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name.
func (client WorkflowsClient) Enable(ctx context.Context, resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.EnablePreparer(ctx, resourceGroupName, workflowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client WorkflowsClient) EnablePreparer(ctx context.Context, resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) EnableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name.
func (client WorkflowsClient) Get(ctx context.Context, resourceGroupName string, workflowName string) (result Workflow, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowsClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) GetResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of workflows by resource group.
//
// resourceGroupName is the resource group name. top is the number of items to be included in the result. filter is
// the filter to apply on the operation.
func (client WorkflowsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32, filter string) (result WorkflowListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkflowsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ListByResourceGroupResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client WorkflowsClient) listByResourceGroupNextResults(lastResults WorkflowListResult) (result WorkflowListResult, err error) {
	req, err := lastResults.workflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32, filter string) (result WorkflowListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top, filter)
	return
}

// ListBySubscription gets a list of workflows by subscription.
//
// top is the number of items to be included in the result. filter is the filter to apply on the operation.
func (client WorkflowsClient) ListBySubscription(ctx context.Context, top *int32, filter string) (result WorkflowListResultPage, err error) {
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client WorkflowsClient) ListBySubscriptionPreparer(ctx context.Context, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ListBySubscriptionResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client WorkflowsClient) listBySubscriptionNextResults(lastResults WorkflowListResult) (result WorkflowListResult, err error) {
	req, err := lastResults.workflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowsClient) ListBySubscriptionComplete(ctx context.Context, top *int32, filter string) (result WorkflowListResultIterator, err error) {
	result.page, err = client.ListBySubscription(ctx, top, filter)
	return
}

// Run runs a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. parameters is the parameters.
func (client WorkflowsClient) Run(ctx context.Context, resourceGroupName string, workflowName string, parameters RunWorkflowParameters) (result WorkflowsRunFuture, err error) {
	req, err := client.RunPreparer(ctx, resourceGroupName, workflowName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Run", nil, "Failure preparing request")
		return
	}

	result, err = client.RunSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Run", result.Response(), "Failure sending request")
		return
	}

	return
}

// RunPreparer prepares the Run request.
func (client WorkflowsClient) RunPreparer(ctx context.Context, resourceGroupName string, workflowName string, parameters RunWorkflowParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/run", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) RunSender(req *http.Request) (future WorkflowsRunFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) RunResponder(resp *http.Response) (result WorkflowRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. workflow is the workflow.
func (client WorkflowsClient) Update(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, workflowName, workflow)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkflowsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", pathParameters),
		autorest.WithJSON(workflow),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) UpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate validates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. workflow is the workflow.
func (client WorkflowsClient) Validate(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (result autorest.Response, err error) {
	req, err := client.ValidatePreparer(ctx, resourceGroupName, workflowName, workflow)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowsClient", "Validate", resp, "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client WorkflowsClient) ValidatePreparer(ctx context.Context, resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/validate", pathParameters),
		autorest.WithJSON(workflow),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ValidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
