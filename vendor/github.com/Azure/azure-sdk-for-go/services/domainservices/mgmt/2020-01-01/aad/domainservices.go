package aad

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DomainServicesClient is the the AAD Domain Services API.
type DomainServicesClient struct {
	BaseClient
}

// NewDomainServicesClient creates an instance of the DomainServicesClient client.
func NewDomainServicesClient(subscriptionID string) DomainServicesClient {
	return NewDomainServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDomainServicesClientWithBaseURI creates an instance of the DomainServicesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDomainServicesClientWithBaseURI(baseURI string, subscriptionID string) DomainServicesClient {
	return DomainServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Create Domain Service operation creates a new domain service with the specified parameters. If
// the specific service already exists, then any patchable properties will be updated and any immutable properties will
// remain unchanged.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// domainService - properties supplied to the Create or Update a Domain Service operation.
func (client DomainServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService) (result DomainServicesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.DomainServicesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, domainServiceName, domainService)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DomainServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", pathParameters),
		autorest.WithJSON(domainService),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) CreateOrUpdateSender(req *http.Request) (future DomainServicesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DomainServicesClient) (ds DomainService, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("aad.DomainServicesCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ds.Response.Response, err = future.GetResult(sender)
		if ds.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
			ds, err = client.CreateOrUpdateResponder(ds.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "aad.DomainServicesCreateOrUpdateFuture", "Result", ds.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) CreateOrUpdateResponder(resp *http.Response) (result DomainService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete Domain Service operation deletes an existing Domain Service.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
func (client DomainServicesClient) Delete(ctx context.Context, resourceGroupName string, domainServiceName string) (result DomainServicesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.DomainServicesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, domainServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DomainServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, domainServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) DeleteSender(req *http.Request) (future DomainServicesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DomainServicesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("aad.DomainServicesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get Domain Service operation retrieves a json representation of the Domain Service.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
func (client DomainServicesClient) Get(ctx context.Context, resourceGroupName string, domainServiceName string) (result DomainService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.DomainServicesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, domainServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DomainServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, domainServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) GetResponder(resp *http.Response) (result DomainService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List Domain Services in Subscription operation lists all the domain services available under the given
// subscription (and across all resource groups within that subscription).
func (client DomainServicesClient) List(ctx context.Context) (result DomainServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.List")
		defer func() {
			sc := -1
			if result.dslr.Response.Response != nil {
				sc = result.dslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "List", resp, "Failure sending request")
		return
	}

	result.dslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dslr.hasNextLink() && result.dslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DomainServicesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) ListResponder(resp *http.Response) (result DomainServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DomainServicesClient) listNextResults(ctx context.Context, lastResults DomainServiceListResult) (result DomainServiceListResult, err error) {
	req, err := lastResults.domainServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DomainServicesClient) ListComplete(ctx context.Context) (result DomainServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup the List Domain Services in Resource Group operation lists all the domain services available
// under the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
func (client DomainServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DomainServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dslr.Response.Response != nil {
				sc = result.dslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.DomainServicesClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.dslr.hasNextLink() && result.dslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DomainServicesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) ListByResourceGroupResponder(resp *http.Response) (result DomainServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DomainServicesClient) listByResourceGroupNextResults(ctx context.Context, lastResults DomainServiceListResult) (result DomainServiceListResult, err error) {
	req, err := lastResults.domainServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DomainServicesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result DomainServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update the Update Domain Service operation can be used to update the existing deployment. The update call only
// supports the properties listed in the PATCH body.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// domainServiceName - the name of the domain service.
// domainService - properties supplied to the Update a Domain Service operation.
func (client DomainServicesClient) Update(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService) (result DomainServicesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServicesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("aad.DomainServicesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, domainServiceName, domainService)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServicesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DomainServicesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainServiceName": autorest.Encode("path", domainServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", pathParameters),
		autorest.WithJSON(domainService),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServicesClient) UpdateSender(req *http.Request) (future DomainServicesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DomainServicesClient) (ds DomainService, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("aad.DomainServicesUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ds.Response.Response, err = future.GetResult(sender)
		if ds.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "aad.DomainServicesUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ds.Response.Response.StatusCode != http.StatusNoContent {
			ds, err = client.UpdateResponder(ds.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "aad.DomainServicesUpdateFuture", "Result", ds.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DomainServicesClient) UpdateResponder(resp *http.Response) (result DomainService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}