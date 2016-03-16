package resources

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// GroupsClient is the client for the Groups methods of the Resources service.
type GroupsClient struct {
	ManagementClient
}

// NewGroupsClient creates an instance of the GroupsClient client.
func NewGroupsClient(subscriptionID string) GroupsClient {
	return NewGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupsClientWithBaseURI creates an instance of the GroupsClient client.
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return GroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckExistence checks whether resource group exists.
//
// resourceGroupName is the name of the resource group to check. The name is
// case insensitive.
func (client GroupsClient) CheckExistence(resourceGroupName string) (result autorest.Response, err error) {
	req, err := client.CheckExistencePreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "CheckExistence", nil, "Failure preparing request")
	}

	resp, err := client.CheckExistenceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "CheckExistence", resp, "Failure sending request")
	}

	result, err = client.CheckExistenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "CheckExistence", resp, "Failure responding to request")
	}

	return
}

// CheckExistencePreparer prepares the CheckExistence request.
func (client GroupsClient) CheckExistencePreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CheckExistenceSender sends the CheckExistence request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) CheckExistenceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckExistenceResponder handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (client GroupsClient) CheckExistenceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create a resource group.
//
// resourceGroupName is the name of the resource group to be created or
// updated. parameters is parameters supplied to the create or update
// resource group service operation.
func (client GroupsClient) CreateOrUpdate(resourceGroupName string, parameters ResourceGroup) (result ResourceGroup, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GroupsClient) CreateOrUpdatePreparer(resourceGroupName string, parameters ResourceGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ResourceGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete begin deleting resource group.To determine whether the operation has
// finished processing the request, call GetLongRunningOperationStatus.
//
// resourceGroupName is the name of the resource group to be deleted. The name
// is case insensitive.
func (client GroupsClient) Delete(resourceGroupName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupsClient) DeletePreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(autorest.DefaultPollingDuration,
			autorest.DefaultPollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a resource group.
//
// resourceGroupName is the name of the resource group to get. The name is
// case insensitive.
func (client GroupsClient) Get(resourceGroupName string) (result ResourceGroup, err error) {
	req, err := client.GetPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupsClient) GetPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupsClient) GetResponder(resp *http.Response) (result ResourceGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a collection of resource groups.
//
// filter is the filter to apply on the operation. top is query parameters. If
// null is passed returns all resource groups.
func (client GroupsClient) List(filter string, top *int32) (result ResourceGroupListResult, err error) {
	req, err := client.ListPreparer(filter, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupsClient) ListPreparer(filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupsClient) ListResponder(resp *http.Response) (result ResourceGroupListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client GroupsClient) ListNextResults(lastResults ResourceGroupListResult) (result ResourceGroupListResult, err error) {
	req, err := lastResults.ResourceGroupListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListResources get all of the resources under a subscription.
//
// resourceGroupName is query parameters. If null is passed returns all
// resource groups. filter is the filter to apply on the operation. top is
// query parameters. If null is passed returns all resource groups.
func (client GroupsClient) ListResources(resourceGroupName string, filter string, top *int32) (result ResourceListResult, err error) {
	req, err := client.ListResourcesPreparer(resourceGroupName, filter, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", nil, "Failure preparing request")
	}

	resp, err := client.ListResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", resp, "Failure sending request")
	}

	result, err = client.ListResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", resp, "Failure responding to request")
	}

	return
}

// ListResourcesPreparer prepares the ListResources request.
func (client GroupsClient) ListResourcesPreparer(resourceGroupName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/resources"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListResourcesSender sends the ListResources request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) ListResourcesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResourcesResponder handles the response to the ListResources request. The method always
// closes the http.Response Body.
func (client GroupsClient) ListResourcesResponder(resp *http.Response) (result ResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListResourcesNextResults retrieves the next set of results, if any.
func (client GroupsClient) ListResourcesNextResults(lastResults ResourceListResult) (result ResourceListResult, err error) {
	req, err := lastResults.ResourceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", resp, "Failure sending next results request request")
	}

	result, err = client.ListResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "ListResources", resp, "Failure responding to next results request request")
	}

	return
}

// Patch resource groups can be updated through a simple PATCH operation to a
// group address. The format of the request is the same as that for creating
// a resource groups, though if a field is unspecified current value will be
// carried over.
//
// resourceGroupName is the name of the resource group to be created or
// updated. The name is case insensitive. parameters is parameters supplied
// to the update state resource group service operation.
func (client GroupsClient) Patch(resourceGroupName string, parameters ResourceGroup) (result ResourceGroup, err error) {
	req, err := client.PatchPreparer(resourceGroupName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Patch", nil, "Failure preparing request")
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/GroupsClient", "Patch", resp, "Failure sending request")
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/GroupsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client GroupsClient) PatchPreparer(resourceGroupName string, parameters ResourceGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client GroupsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client GroupsClient) PatchResponder(resp *http.Response) (result ResourceGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
