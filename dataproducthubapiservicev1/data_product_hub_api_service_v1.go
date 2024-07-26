/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

// Package dataproducthubapiservicev1 : Operations and models for the DataProductHubApiServiceV1 service
package dataproducthubapiservicev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	common "github.com/IBM/data-product-exchange-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// DataProductHubApiServiceV1 : Data Product Hub API Service
//
// API Version: 0.0.6
type DataProductHubApiServiceV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "data_product_hub_api_service"

// DataProductHubApiServiceV1Options : Service options
type DataProductHubApiServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDataProductHubApiServiceV1UsingExternalConfig : constructs an instance of DataProductHubApiServiceV1 with passed in options and external configuration.
func NewDataProductHubApiServiceV1UsingExternalConfig(options *DataProductHubApiServiceV1Options) (dataProductHubApiService *DataProductHubApiServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dataProductHubApiService, err = NewDataProductHubApiServiceV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = dataProductHubApiService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dataProductHubApiService.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDataProductHubApiServiceV1 : constructs an instance of DataProductHubApiServiceV1 with passed in options.
func NewDataProductHubApiServiceV1(options *DataProductHubApiServiceV1Options) (service *DataProductHubApiServiceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &DataProductHubApiServiceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "dataProductHubApiService" suitable for processing requests.
func (dataProductHubApiService *DataProductHubApiServiceV1) Clone() *DataProductHubApiServiceV1 {
	if core.IsNil(dataProductHubApiService) {
		return nil
	}
	clone := *dataProductHubApiService
	clone.Service = dataProductHubApiService.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dataProductHubApiService *DataProductHubApiServiceV1) SetServiceURL(url string) error {
	err := dataProductHubApiService.Service.SetServiceURL(url)
	return err
}

// GetServiceURL returns the service URL
func (dataProductHubApiService *DataProductHubApiServiceV1) GetServiceURL() string {
	return dataProductHubApiService.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dataProductHubApiService *DataProductHubApiServiceV1) SetDefaultHeaders(headers http.Header) {
	dataProductHubApiService.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dataProductHubApiService *DataProductHubApiServiceV1) SetEnableGzipCompression(enableGzip bool) {
	dataProductHubApiService.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dataProductHubApiService *DataProductHubApiServiceV1) GetEnableGzipCompression() bool {
	return dataProductHubApiService.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dataProductHubApiService *DataProductHubApiServiceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dataProductHubApiService.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dataProductHubApiService *DataProductHubApiServiceV1) DisableRetries() {
	dataProductHubApiService.Service.DisableRetries()
}

// GetInitializeStatus : Get resource initialization status
// Use this API to get the status of resource initialization in Data Product Hub.<br/><br/>If the data product catalog
// exists but has never been initialized, the status will be "not_started".<br/><br/>If the data product catalog exists
// and has been or is being initialized, the response will contain the status of the last or current initialization. If
// the initialization failed, the "errors" and "trace" fields will contain the error(s) encountered during the
// initialization, including the ID to trace the error(s).<br/><br/>If the data product catalog doesn't exist, an HTTP
// 404 response is returned.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetInitializeStatusWithContext(context.Background(), getInitializeStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInitializeStatusWithContext is an alternate form of the GetInitializeStatus method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetInitializeStatusWithContext(ctx context.Context, getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getInitializeStatusOptions, "getInitializeStatusOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize/status`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInitializeStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetInitializeStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getInitializeStatusOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*getInitializeStatusOptions.ContainerID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_initialize_status", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetServiceIdCredentials : Get service id credentials
// Use this API to get the information of service id credentials in Data Product Hub.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetServiceIdCredentials(getServiceIdCredentialsOptions *GetServiceIdCredentialsOptions) (result *ServiceIdCredentials, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetServiceIdCredentialsWithContext(context.Background(), getServiceIdCredentialsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetServiceIdCredentialsWithContext is an alternate form of the GetServiceIdCredentials method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetServiceIdCredentialsWithContext(ctx context.Context, getServiceIdCredentialsOptions *GetServiceIdCredentialsOptions) (result *ServiceIdCredentials, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getServiceIdCredentialsOptions, "getServiceIdCredentialsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/credentials`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceIdCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetServiceIdCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_service_id_credentials", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceIdCredentials)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Initialize : Initialize resources
// Use this API to initialize default assets for data product hub. <br/><br/>You can initialize:
// <br/><ul><li>`delivery_methods` - Methods through which data product parts can be delivered to consumers of the data
// product hub</li><li>`domains_multi_industry` - Taxonomy of domains and use cases applicable to multiple
// industries</li><li>`data_product_samples` - Sample data products used to illustrate capabilities of the data product
// hub</li><li>`workflows` - Workflows to enable restricted data products</li><li>`project` - A default project for
// exporting data assets to files</li></ul><br/><br/>If a resource depends on resources that are not specified in the
// request, these dependent resources will be automatically initialized. E.g., initializing `data_product_samples` will
// also initialize `domains_multi_industry` and `delivery_methods` even if they are not specified in the request because
// it depends on them.<br/><br/>If initializing the data product hub for the first time, do not specify a container. The
// default data product catalog will be created.<br/>For first time initialization, it is recommended that at least
// `delivery_methods` and `domains_multi_industry` is included in the initialize operation.<br/><br/>If the data product
// hub has already been initialized, you may call this API again to initialize new resources, such as new delivery
// methods.In this case, specify the default data product catalog container information.
func (dataProductHubApiService *DataProductHubApiServiceV1) Initialize(initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.InitializeWithContext(context.Background(), initializeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// InitializeWithContext is an alternate form of the Initialize method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) InitializeWithContext(ctx context.Context, initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(initializeOptions, "initializeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(initializeOptions, "initializeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range initializeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "Initialize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if initializeOptions.Container != nil {
		body["container"] = initializeOptions.Container
	}
	if initializeOptions.Include != nil {
		body["include"] = initializeOptions.Include
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "initialize", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ManageApiKeys : Rotate credentials for a Data Product Hub instance
// Use this API to rotate credentials for a Data Product Hub instance.
func (dataProductHubApiService *DataProductHubApiServiceV1) ManageApiKeys(manageApiKeysOptions *ManageApiKeysOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.ManageApiKeysWithContext(context.Background(), manageApiKeysOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ManageApiKeysWithContext is an alternate form of the ManageApiKeys method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) ManageApiKeysWithContext(ctx context.Context, manageApiKeysOptions *ManageApiKeysOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(manageApiKeysOptions, "manageApiKeysOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/rotate_credentials`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range manageApiKeysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ManageApiKeys")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "manage_api_keys", getServiceComponentInfo())
		return
	}

	return
}

// ListDataProducts : Retrieve a list of data products
// Retrieve a list of data products.
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProducts(listDataProductsOptions *ListDataProductsOptions) (result *DataProductSummaryCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductsWithContext(context.Background(), listDataProductsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductsWithContext is an alternate form of the ListDataProducts method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProductsWithContext(ctx context.Context, listDataProductsOptions *ListDataProductsOptions) (result *DataProductSummaryCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductsOptions, "listDataProductsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProducts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDataProductsOptions.Limit))
	}
	if listDataProductsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listDataProductsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_products", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductSummaryCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProduct : Create a new data product
// Use this API to create a new data product.<br/><br/>Provide the initial draft of the data product.<br/><br/>Required
// fields:<br/><br/>- name<br/>- container<br/><br/>If `version` is not specified, the default version **1.0.0** will be
// used.<br/><br/>The `domain` is optional.
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDataProduct(createDataProductOptions *CreateDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductWithContext(context.Background(), createDataProductOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductWithContext is an alternate form of the CreateDataProduct method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDataProductWithContext(ctx context.Context, createDataProductOptions *CreateDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductOptions, "createDataProductOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDataProductOptions, "createDataProductOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDataProductOptions.Drafts != nil {
		body["drafts"] = createDataProductOptions.Drafts
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProduct)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDataProduct : Retrieve a data product identified by id
// Retrieve a data product identified by id.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProduct(getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductWithContext(context.Background(), getDataProductOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductWithContext is an alternate form of the GetDataProduct method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProductWithContext(ctx context.Context, getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductOptions, "getDataProductOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDataProductOptions, "getDataProductOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductOptions.DataProductID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProduct)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CompleteDraftContractTermsDocument : Complete a contract document upload operation
// After uploading a file to the provided signed URL, call this endpoint to mark the upload as complete. After the
// upload operation is marked as complete, the file is available to download.
// - After the upload is marked as complete, the returned URL is displayed in the "url" field. The signed URL is used to
// download the document.
// - Calling complete on referential documents results in an error.
// - Calling complete on attachment documents for which the file has not been uploaded will result in an error.
func (dataProductHubApiService *DataProductHubApiServiceV1) CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CompleteDraftContractTermsDocumentWithContext(context.Background(), completeDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CompleteDraftContractTermsDocumentWithContext is an alternate form of the CompleteDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) CompleteDraftContractTermsDocumentWithContext(ctx context.Context, completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeDraftContractTermsDocumentOptions, "completeDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeDraftContractTermsDocumentOptions, "completeDraftContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *completeDraftContractTermsDocumentOptions.DataProductID,
		"draft_id":          *completeDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *completeDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id":       *completeDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}/complete`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CompleteDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "complete_draft_contract_terms_document", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDataProductDrafts : Retrieve a list of data product drafts
// Retrieve a list of data product drafts.
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) (result *DataProductDraftCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductDraftsWithContext(context.Background(), listDataProductDraftsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductDraftsWithContext is an alternate form of the ListDataProductDrafts method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProductDraftsWithContext(ctx context.Context, listDataProductDraftsOptions *ListDataProductDraftsOptions) (result *DataProductDraftCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDataProductDraftsOptions, "listDataProductDraftsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDataProductDraftsOptions, "listDataProductDraftsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *listDataProductDraftsOptions.DataProductID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDataProductDraftsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductDrafts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductDraftsOptions.AssetContainerID != nil {
		builder.AddQuery("asset.container.id", fmt.Sprint(*listDataProductDraftsOptions.AssetContainerID))
	}
	if listDataProductDraftsOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*listDataProductDraftsOptions.Version))
	}
	if listDataProductDraftsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDataProductDraftsOptions.Limit))
	}
	if listDataProductDraftsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listDataProductDraftsOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_drafts", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDraftCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProductDraft : Create a new draft of an existing data product
// Create a new draft of an existing data product.
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductDraftWithContext(context.Background(), createDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductDraftWithContext is an alternate form of the CreateDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDataProductDraftWithContext(ctx context.Context, createDataProductDraftOptions *CreateDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductDraftOptions, "createDataProductDraftOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDataProductDraftOptions, "createDataProductDraftOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *createDataProductDraftOptions.DataProductID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDataProductDraftOptions.Asset != nil {
		body["asset"] = createDataProductDraftOptions.Asset
	}
	if createDataProductDraftOptions.Version != nil {
		body["version"] = createDataProductDraftOptions.Version
	}
	if createDataProductDraftOptions.State != nil {
		body["state"] = createDataProductDraftOptions.State
	}
	if createDataProductDraftOptions.DataProduct != nil {
		body["data_product"] = createDataProductDraftOptions.DataProduct
	}
	if createDataProductDraftOptions.Name != nil {
		body["name"] = createDataProductDraftOptions.Name
	}
	if createDataProductDraftOptions.Description != nil {
		body["description"] = createDataProductDraftOptions.Description
	}
	if createDataProductDraftOptions.Tags != nil {
		body["tags"] = createDataProductDraftOptions.Tags
	}
	if createDataProductDraftOptions.UseCases != nil {
		body["use_cases"] = createDataProductDraftOptions.UseCases
	}
	if createDataProductDraftOptions.Types != nil {
		body["types"] = createDataProductDraftOptions.Types
	}
	if createDataProductDraftOptions.ContractTerms != nil {
		body["contract_terms"] = createDataProductDraftOptions.ContractTerms
	}
	if createDataProductDraftOptions.IsRestricted != nil {
		body["is_restricted"] = createDataProductDraftOptions.IsRestricted
	}
	if createDataProductDraftOptions.Domain != nil {
		body["domain"] = createDataProductDraftOptions.Domain
	}
	if createDataProductDraftOptions.PartsOut != nil {
		body["parts_out"] = createDataProductDraftOptions.PartsOut
	}
	if createDataProductDraftOptions.Workflows != nil {
		body["workflows"] = createDataProductDraftOptions.Workflows
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product_draft", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDraftContractTermsDocument : Upload a contract document to the data product draft contract terms
// Upload a contract document to the data product draft identified by draft_id.
//
// - If the request object contains a "url" parameter, a referential document is created to store the provided url.
// - If the request object does not contain a "url" parameter, an attachment document is created, and a signed url will
// be returned in an "upload_url" parameter. The data product producer can upload the document using the provided
// "upload_url". After the upload is completed, call "complete_contract_terms_document" for the given document needs to
// be called to mark the upload as completed. After completion of the upload, "get_contract_terms_document" for the
// given document returns a signed "url" parameter that can be used to download the attachment document.
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDraftContractTermsDocumentWithContext(context.Background(), createDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDraftContractTermsDocumentWithContext is an alternate form of the CreateDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) CreateDraftContractTermsDocumentWithContext(ctx context.Context, createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDraftContractTermsDocumentOptions, "createDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDraftContractTermsDocumentOptions, "createDraftContractTermsDocumentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *createDraftContractTermsDocumentOptions.DataProductID,
		"draft_id":          *createDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *createDraftContractTermsDocumentOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDraftContractTermsDocumentOptions.Type != nil {
		body["type"] = createDraftContractTermsDocumentOptions.Type
	}
	if createDraftContractTermsDocumentOptions.Name != nil {
		body["name"] = createDraftContractTermsDocumentOptions.Name
	}
	if createDraftContractTermsDocumentOptions.ID != nil {
		body["id"] = createDraftContractTermsDocumentOptions.ID
	}
	if createDraftContractTermsDocumentOptions.URL != nil {
		body["url"] = createDraftContractTermsDocumentOptions.URL
	}
	if createDraftContractTermsDocumentOptions.Attachment != nil {
		body["attachment"] = createDraftContractTermsDocumentOptions.Attachment
	}
	if createDraftContractTermsDocumentOptions.UploadURL != nil {
		body["upload_url"] = createDraftContractTermsDocumentOptions.UploadURL
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_draft_contract_terms_document", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductDraft : Get a draft of an existing data product
// Get a draft of an existing data product.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductDraftWithContext(context.Background(), getDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductDraftWithContext is an alternate form of the GetDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProductDraftWithContext(ctx context.Context, getDataProductDraftOptions *GetDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductDraftOptions, "getDataProductDraftOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(getDataProductDraftOptions, "getDataProductDraftOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductDraftOptions.DataProductID,
		"draft_id":        *getDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_draft", getServiceComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataProductDraft : Delete a data product draft identified by ID
// Delete a data product draft identified by a valid ID.
func (dataProductHubApiService *DataProductHubApiServiceV1) DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDataProductDraftWithContext(context.Background(), deleteDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDataProductDraftWithContext is an alternate form of the DeleteDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) DeleteDataProductDraftWithContext(ctx context.Context, deleteDataProductDraftOptions *DeleteDataProductDraftOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataProductDraftOptions, "deleteDataProductDraftOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(deleteDataProductDraftOptions, "deleteDataProductDraftOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *deleteDataProductDraftOptions.DataProductID,
		"draft_id":        *deleteDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range deleteDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_data_product_draft", getServiceComponentInfo())

		return
	}

	return
}

// UpdateDataProductDraft : Update the data product draft identified by ID
// Use this API to update the properties of a data product draft identified by a valid ID.<br/><br/>Specify patch
// operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the
// properties of a data product<br/><br/>- Add/Remove parts from a data product (up to 20 parts)<br/><br/>- Add/Remove
// use cases from a data product<br/><br/>- Update the data product state<br/><br/>.
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductDraftWithContext(context.Background(), updateDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductDraftWithContext is an alternate form of the UpdateDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDataProductDraftWithContext(ctx context.Context, updateDataProductDraftOptions *UpdateDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductDraftOptions, "updateDataProductDraftOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(updateDataProductDraftOptions, "updateDataProductDraftOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDataProductDraftOptions.DataProductID,
		"draft_id":        *updateDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range updateDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductDraftOptions.JSONPatchInstructions)
	if err != nil {

		return
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_draft", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// GetDraftContractTermsDocument : Get a contract document
// If a document has a completed attachment, the response contains the `url` which can be used to download the
// attachment. If a document does not have a completed attachment, the response contains the `url` which was submitted
// at document creation. If a document has an attachment that is incomplete, an error is returned to prompt the user to
// upload the document file and complete it.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDraftContractTermsDocumentWithContext(context.Background(), getDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDraftContractTermsDocumentWithContext is an alternate form of the GetDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDraftContractTermsDocumentWithContext(ctx context.Context, getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDraftContractTermsDocumentOptions, "getDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(getDraftContractTermsDocumentOptions, "getDraftContractTermsDocumentOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *getDraftContractTermsDocumentOptions.DataProductID,
		"draft_id":          *getDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *getDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id":       *getDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range getDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_draft_contract_terms_document", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// DeleteDraftContractTermsDocument : Delete a contract document
// Delete an existing contract document.
//
// Contract documents can only be deleted for data product versions that are in DRAFT state.
func (dataProductHubApiService *DataProductHubApiServiceV1) DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDraftContractTermsDocumentWithContext(context.Background(), deleteDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDraftContractTermsDocumentWithContext is an alternate form of the DeleteDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) DeleteDraftContractTermsDocumentWithContext(ctx context.Context, deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDraftContractTermsDocumentOptions, "deleteDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(deleteDraftContractTermsDocumentOptions, "deleteDraftContractTermsDocumentOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *deleteDraftContractTermsDocumentOptions.DataProductID,
		"draft_id":          *deleteDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *deleteDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id":       *deleteDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range deleteDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_draft_contract_terms_document", getServiceComponentInfo())

		return
	}

	return
}

// UpdateDraftContractTermsDocument : Update a contract document
// Use this API to update the properties of a contract document that is identified by a valid ID.
//
// Specify patch operations using http://jsonpatch.com/ syntax.
//
// Supported patch operations include:
// - Update the url of document if it does not have an attachment.
// - Update the type of the document.
// <br/><br/>Contract terms documents can only be updated if the associated data product version is in DRAFT state.
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDraftContractTermsDocumentWithContext(context.Background(), updateDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDraftContractTermsDocumentWithContext is an alternate form of the UpdateDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDraftContractTermsDocumentWithContext(ctx context.Context, updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDraftContractTermsDocumentOptions, "updateDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(updateDraftContractTermsDocumentOptions, "updateDraftContractTermsDocumentOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *updateDraftContractTermsDocumentOptions.DataProductID,
		"draft_id":          *updateDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *updateDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id":       *updateDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range updateDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDraftContractTermsDocumentOptions.JSONPatchInstructions)
	if err != nil {

		return
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_draft_contract_terms_document", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// PublishDataProductDraft : Publish a draft of an existing data product
// Publish a draft of an existing data product.
func (dataProductHubApiService *DataProductHubApiServiceV1) PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.PublishDataProductDraftWithContext(context.Background(), publishDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PublishDataProductDraftWithContext is an alternate form of the PublishDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) PublishDataProductDraftWithContext(ctx context.Context, publishDataProductDraftOptions *PublishDataProductDraftOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(publishDataProductDraftOptions, "publishDataProductDraftOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(publishDataProductDraftOptions, "publishDataProductDraftOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *publishDataProductDraftOptions.DataProductID,
		"draft_id":        *publishDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/publish`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range publishDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "PublishDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "publish_data_product_draft", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// GetDataProductRelease : Get a release of an existing data product
// Get a release of an existing data product.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductReleaseWithContext(context.Background(), getDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductReleaseWithContext is an alternate form of the GetDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetDataProductReleaseWithContext(ctx context.Context, getDataProductReleaseOptions *GetDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductReleaseOptions, "getDataProductReleaseOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(getDataProductReleaseOptions, "getDataProductReleaseOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductReleaseOptions.DataProductID,
		"release_id":      *getDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range getDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_release", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// UpdateDataProductRelease : Update the data product release identified by ID
// Use this API to update the properties of a data product release identified by a valid ID.<br/><br/>Specify patch
// operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the
// properties of a data product<br/><br/>- Add/remove parts from a data product (up to 20 parts)<br/><br/>- Add/remove
// use cases from a data product<br/><br/>.
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductReleaseWithContext(context.Background(), updateDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductReleaseWithContext is an alternate form of the UpdateDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) UpdateDataProductReleaseWithContext(ctx context.Context, updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductReleaseOptions, "updateDataProductReleaseOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(updateDataProductReleaseOptions, "updateDataProductReleaseOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDataProductReleaseOptions.DataProductID,
		"release_id":      *updateDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range updateDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductReleaseOptions.JSONPatchInstructions)
	if err != nil {

		return
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_release", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// GetReleaseContractTermsDocument : Get a contract document
// If the document has a completed attachment, the response contains the `url` to download the attachment.<br/><br/> If
// the document does not have an attachment, the response contains the `url` which was submitted at document
// creation.<br/><br/> If the document has an incomplete attachment, an error is returned to prompt the user to upload
// the document file to complete the attachment.
func (dataProductHubApiService *DataProductHubApiServiceV1) GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetReleaseContractTermsDocumentWithContext(context.Background(), getReleaseContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetReleaseContractTermsDocumentWithContext is an alternate form of the GetReleaseContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) GetReleaseContractTermsDocumentWithContext(ctx context.Context, getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReleaseContractTermsDocumentOptions, "getReleaseContractTermsDocumentOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(getReleaseContractTermsDocumentOptions, "getReleaseContractTermsDocumentOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id":   *getReleaseContractTermsDocumentOptions.DataProductID,
		"release_id":        *getReleaseContractTermsDocumentOptions.ReleaseID,
		"contract_terms_id": *getReleaseContractTermsDocumentOptions.ContractTermsID,
		"document_id":       *getReleaseContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range getReleaseContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetReleaseContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_release_contract_terms_document", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// ListDataProductReleases : Retrieve a list of data product releases
// Retrieve a list of data product releases.
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) (result *DataProductReleaseCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductReleasesWithContext(context.Background(), listDataProductReleasesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductReleasesWithContext is an alternate form of the ListDataProductReleases method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) ListDataProductReleasesWithContext(ctx context.Context, listDataProductReleasesOptions *ListDataProductReleasesOptions) (result *DataProductReleaseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDataProductReleasesOptions, "listDataProductReleasesOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(listDataProductReleasesOptions, "listDataProductReleasesOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *listDataProductReleasesOptions.DataProductID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range listDataProductReleasesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductReleases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductReleasesOptions.AssetContainerID != nil {
		builder.AddQuery("asset.container.id", fmt.Sprint(*listDataProductReleasesOptions.AssetContainerID))
	}
	if listDataProductReleasesOptions.State != nil {
		builder.AddQuery("state", strings.Join(listDataProductReleasesOptions.State, ","))
	}
	if listDataProductReleasesOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*listDataProductReleasesOptions.Version))
	}
	if listDataProductReleasesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDataProductReleasesOptions.Limit))
	}
	if listDataProductReleasesOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listDataProductReleasesOptions.Start))
	}

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_releases", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductReleaseCollection)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}

// RetireDataProductRelease : Retire a release of an existing data product
// Retire a release of an existing data product.
func (dataProductHubApiService *DataProductHubApiServiceV1) RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.RetireDataProductReleaseWithContext(context.Background(), retireDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RetireDataProductReleaseWithContext is an alternate form of the RetireDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubApiServiceV1) RetireDataProductReleaseWithContext(ctx context.Context, retireDataProductReleaseOptions *RetireDataProductReleaseOptions) (result *DataProductVersion, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(retireDataProductReleaseOptions, "retireDataProductReleaseOptions cannot be nil")
	if err != nil {

		return
	}
	err = core.ValidateStruct(retireDataProductReleaseOptions, "retireDataProductReleaseOptions")
	if err != nil {

		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *retireDataProductReleaseOptions.DataProductID,
		"release_id":      *retireDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}/retire`, pathParamsMap)
	if err != nil {

		return
	}

	for headerName, headerValue := range retireDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "RetireDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {

		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "retire_data_product_release", getServiceComponentInfo())

		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersion)
		if err != nil {

			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// AssetPartReference : The asset represented in this part.
type AssetPartReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The type of the asset.
	Type *string `json:"type,omitempty"`
}

// NewAssetPartReference : Instantiate AssetPartReference (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewAssetPartReference(container *ContainerReference) (_model *AssetPartReference, err error) {
	_model = &AssetPartReference{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetPartReference unmarshals an instance of AssetPartReference from the specified map of raw messages.
func UnmarshalAssetPartReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPartReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetPrototype : New asset input properties.
type AssetPrototype struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	Container *ContainerIdentity `json:"container" validate:"required"`
}

// NewAssetPrototype : Instantiate AssetPrototype (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewAssetPrototype(container *ContainerIdentity) (_model *AssetPrototype, err error) {
	_model = &AssetPrototype{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetPrototype unmarshals an instance of AssetPrototype from the specified map of raw messages.
func UnmarshalAssetPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPrototype)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerIdentity)

	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetReference : AssetReference struct
type AssetReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// NewAssetReference : Instantiate AssetReference (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewAssetReference(container *ContainerReference) (_model *AssetReference, err error) {
	_model = &AssetReference{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetReference unmarshals an instance of AssetReference from the specified map of raw messages.
func UnmarshalAssetReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompleteDraftContractTermsDocumentOptions : The CompleteDraftContractTermsDocument options.
type CompleteDraftContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCompleteDraftContractTermsDocumentOptions : Instantiate CompleteDraftContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewCompleteDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *CompleteDraftContractTermsDocumentOptions {
	return &CompleteDraftContractTermsDocumentOptions{
		DataProductID:   core.StringPtr(dataProductID),
		DraftID:         core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID:      core.StringPtr(documentID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *CompleteDraftContractTermsDocumentOptions) SetDataProductID(dataProductID string) *CompleteDraftContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *CompleteDraftContractTermsDocumentOptions) SetDraftID(draftID string) *CompleteDraftContractTermsDocumentOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *CompleteDraftContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *CompleteDraftContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *CompleteDraftContractTermsDocumentOptions) SetDocumentID(documentID string) *CompleteDraftContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *CompleteDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// ContainerIdentity : ContainerIdentity struct
type ContainerIdentity struct {
	// Container identifier.
	ID *string `json:"id" validate:"required"`
}

// NewContainerIdentity : Instantiate ContainerIdentity (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewContainerIdentity(id string) (_model *ContainerIdentity, err error) {
	_model = &ContainerIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContainerIdentity unmarshals an instance of ContainerIdentity from the specified map of raw messages.
func UnmarshalContainerIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContainerReference : Container reference.
type ContainerReference struct {
	// Container identifier.
	ID *string `json:"id" validate:"required"`

	// Container type.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the ContainerReference.Type property.
// Container type.
const (
	ContainerReference_Type_Catalog = "catalog"
	ContainerReference_Type_Project = "project"
)

// NewContainerReference : Instantiate ContainerReference (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewContainerReference(id string, typeVar string) (_model *ContainerReference, err error) {
	_model = &ContainerReference{
		ID:   core.StringPtr(id),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContainerReference unmarshals an instance of ContainerReference from the specified map of raw messages.
func UnmarshalContainerReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTermsDocument : Standard contract terms document, which is used for get and list contract terms responses.
type ContractTermsDocument struct {
	// URL that can be used to retrieve the contract document.
	URL *string `json:"url,omitempty"`

	// Type of the contract document.
	Type *string `json:"type" validate:"required"`

	// Name of the contract document.
	Name *string `json:"name" validate:"required"`

	// Id uniquely identifying this document within the contract terms instance.
	ID *string `json:"id" validate:"required"`

	// Attachment associated witht the document.
	Attachment *ContractTermsDocumentAttachment `json:"attachment,omitempty"`

	// URL which can be used to upload document file.
	UploadURL *string `json:"upload_url,omitempty"`
}

// Constants associated with the ContractTermsDocument.Type property.
// Type of the contract document.
const (
	ContractTermsDocument_Type_Sla                = "sla"
	ContractTermsDocument_Type_TermsAndConditions = "terms_and_conditions"
)

// NewContractTermsDocument : Instantiate ContractTermsDocument (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewContractTermsDocument(typeVar string, name string, id string) (_model *ContractTermsDocument, err error) {
	_model = &ContractTermsDocument{
		Type: core.StringPtr(typeVar),
		Name: core.StringPtr(name),
		ID:   core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContractTermsDocument unmarshals an instance of ContractTermsDocument from the specified map of raw messages.
func UnmarshalContractTermsDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsDocument)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "attachment", &obj.Attachment, UnmarshalContractTermsDocumentAttachment)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "upload_url", &obj.UploadURL)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubApiServiceV1) NewContractTermsDocumentPatch(contractTermsDocument *ContractTermsDocument) (_patch []JSONPatchOperation) {
	if contractTermsDocument.URL != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/url"),
			Value: contractTermsDocument.URL,
		})
	}
	if contractTermsDocument.Type != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/type"),
			Value: contractTermsDocument.Type,
		})
	}
	if contractTermsDocument.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: contractTermsDocument.Name,
		})
	}
	if contractTermsDocument.ID != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/id"),
			Value: contractTermsDocument.ID,
		})
	}
	if contractTermsDocument.Attachment != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/attachment"),
			Value: contractTermsDocument.Attachment,
		})
	}
	if contractTermsDocument.UploadURL != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/upload_url"),
			Value: contractTermsDocument.UploadURL,
		})
	}
	return
}

// ContractTermsDocumentAttachment : Attachment associated witht the document.
type ContractTermsDocumentAttachment struct {
	// Id representing the corresponding attachment.
	ID *string `json:"id,omitempty"`
}

// UnmarshalContractTermsDocumentAttachment unmarshals an instance of ContractTermsDocumentAttachment from the specified map of raw messages.
func UnmarshalContractTermsDocumentAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsDocumentAttachment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDataProductDraftOptions : The CreateDataProductDraft options.
type CreateDataProductDraftOptions struct {
	// Data product ID.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// New asset input properties.
	Asset *AssetPrototype `json:"asset" validate:"required"`

	// The data product version number.
	Version *string `json:"version,omitempty"`

	// The state of the data product version. If not specified, the data product version will be created in `draft` state.
	State *string `json:"state,omitempty"`

	// Data product identifier.
	DataProduct *DataProductIdentity `json:"data_product,omitempty"`

	// The name that refers to the new data product version. If this is a new data product, this value must be specified.
	// If this is a new version of an existing data product, the name will default to the name of the previous data product
	// version. A name can contain letters, numbers, understores, dashes, spaces or periods. A name must contain at least
	// one non-space character.
	Name *string `json:"name,omitempty"`

	// Description of the data product version. If this is a new version of an existing data product, the description will
	// default to the description of the previous version of the data product.
	Description *string `json:"description,omitempty"`

	// Tags on the data product.
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types,omitempty"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []DataProductContractTerms `json:"contract_terms,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted,omitempty"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain,omitempty"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out,omitempty"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDataProductDraftOptions.State property.
// The state of the data product version. If not specified, the data product version will be created in `draft` state.
const (
	CreateDataProductDraftOptions_State_Available = "available"
	CreateDataProductDraftOptions_State_Draft     = "draft"
	CreateDataProductDraftOptions_State_Retired   = "retired"
)

// Constants associated with the CreateDataProductDraftOptions.Types property.
const (
	CreateDataProductDraftOptions_Types_Code = "code"
	CreateDataProductDraftOptions_Types_Data = "data"
)

// NewCreateDataProductDraftOptions : Instantiate CreateDataProductDraftOptions
func (*DataProductHubApiServiceV1) NewCreateDataProductDraftOptions(dataProductID string, asset *AssetPrototype) *CreateDataProductDraftOptions {
	return &CreateDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		Asset:         asset,
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *CreateDataProductDraftOptions) SetDataProductID(dataProductID string) *CreateDataProductDraftOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetAsset : Allow user to set Asset
func (_options *CreateDataProductDraftOptions) SetAsset(asset *AssetPrototype) *CreateDataProductDraftOptions {
	_options.Asset = asset
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CreateDataProductDraftOptions) SetVersion(version string) *CreateDataProductDraftOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetState : Allow user to set State
func (_options *CreateDataProductDraftOptions) SetState(state string) *CreateDataProductDraftOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetDataProduct : Allow user to set DataProduct
func (_options *CreateDataProductDraftOptions) SetDataProduct(dataProduct *DataProductIdentity) *CreateDataProductDraftOptions {
	_options.DataProduct = dataProduct
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDataProductDraftOptions) SetName(name string) *CreateDataProductDraftOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDataProductDraftOptions) SetDescription(description string) *CreateDataProductDraftOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *CreateDataProductDraftOptions) SetTags(tags []string) *CreateDataProductDraftOptions {
	_options.Tags = tags
	return _options
}

// SetUseCases : Allow user to set UseCases
func (_options *CreateDataProductDraftOptions) SetUseCases(useCases []UseCase) *CreateDataProductDraftOptions {
	_options.UseCases = useCases
	return _options
}

// SetTypes : Allow user to set Types
func (_options *CreateDataProductDraftOptions) SetTypes(types []string) *CreateDataProductDraftOptions {
	_options.Types = types
	return _options
}

// SetContractTerms : Allow user to set ContractTerms
func (_options *CreateDataProductDraftOptions) SetContractTerms(contractTerms []DataProductContractTerms) *CreateDataProductDraftOptions {
	_options.ContractTerms = contractTerms
	return _options
}

// SetIsRestricted : Allow user to set IsRestricted
func (_options *CreateDataProductDraftOptions) SetIsRestricted(isRestricted bool) *CreateDataProductDraftOptions {
	_options.IsRestricted = core.BoolPtr(isRestricted)
	return _options
}

// SetDomain : Allow user to set Domain
func (_options *CreateDataProductDraftOptions) SetDomain(domain *Domain) *CreateDataProductDraftOptions {
	_options.Domain = domain
	return _options
}

// SetPartsOut : Allow user to set PartsOut
func (_options *CreateDataProductDraftOptions) SetPartsOut(partsOut []DataProductPart) *CreateDataProductDraftOptions {
	_options.PartsOut = partsOut
	return _options
}

// SetWorkflows : Allow user to set Workflows
func (_options *CreateDataProductDraftOptions) SetWorkflows(workflows *DataProductWorkflows) *CreateDataProductDraftOptions {
	_options.Workflows = workflows
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductDraftOptions) SetHeaders(param map[string]string) *CreateDataProductDraftOptions {
	options.Headers = param
	return options
}

// CreateDataProductOptions : The CreateDataProduct options.
type CreateDataProductOptions struct {
	// Collection of data products drafts to add to data product.
	Drafts []DataProductVersionPrototype `json:"drafts" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDataProductOptions : Instantiate CreateDataProductOptions
func (*DataProductHubApiServiceV1) NewCreateDataProductOptions(drafts []DataProductVersionPrototype) *CreateDataProductOptions {
	return &CreateDataProductOptions{
		Drafts: drafts,
	}
}

// SetDrafts : Allow user to set Drafts
func (_options *CreateDataProductOptions) SetDrafts(drafts []DataProductVersionPrototype) *CreateDataProductOptions {
	_options.Drafts = drafts
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductOptions) SetHeaders(param map[string]string) *CreateDataProductOptions {
	options.Headers = param
	return options
}

// CreateDraftContractTermsDocumentOptions : The CreateDraftContractTermsDocument options.
type CreateDraftContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Type of the contract document.
	Type *string `json:"type" validate:"required"`

	// Name of the contract document.
	Name *string `json:"name" validate:"required"`

	// Id uniquely identifying this document within the contract terms instance.
	ID *string `json:"id" validate:"required"`

	// URL that can be used to retrieve the contract document.
	URL *string `json:"url,omitempty"`

	// Attachment associated witht the document.
	Attachment *ContractTermsDocumentAttachment `json:"attachment,omitempty"`

	// URL which can be used to upload document file.
	UploadURL *string `json:"upload_url,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDraftContractTermsDocumentOptions.Type property.
// Type of the contract document.
const (
	CreateDraftContractTermsDocumentOptions_Type_Sla                = "sla"
	CreateDraftContractTermsDocumentOptions_Type_TermsAndConditions = "terms_and_conditions"
)

// NewCreateDraftContractTermsDocumentOptions : Instantiate CreateDraftContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewCreateDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, typeVar string, name string, id string) *CreateDraftContractTermsDocumentOptions {
	return &CreateDraftContractTermsDocumentOptions{
		DataProductID:   core.StringPtr(dataProductID),
		DraftID:         core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		Type:            core.StringPtr(typeVar),
		Name:            core.StringPtr(name),
		ID:              core.StringPtr(id),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *CreateDraftContractTermsDocumentOptions) SetDataProductID(dataProductID string) *CreateDraftContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *CreateDraftContractTermsDocumentOptions) SetDraftID(draftID string) *CreateDraftContractTermsDocumentOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *CreateDraftContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *CreateDraftContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateDraftContractTermsDocumentOptions) SetType(typeVar string) *CreateDraftContractTermsDocumentOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDraftContractTermsDocumentOptions) SetName(name string) *CreateDraftContractTermsDocumentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateDraftContractTermsDocumentOptions) SetID(id string) *CreateDraftContractTermsDocumentOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetURL : Allow user to set URL
func (_options *CreateDraftContractTermsDocumentOptions) SetURL(url string) *CreateDraftContractTermsDocumentOptions {
	_options.URL = core.StringPtr(url)
	return _options
}

// SetAttachment : Allow user to set Attachment
func (_options *CreateDraftContractTermsDocumentOptions) SetAttachment(attachment *ContractTermsDocumentAttachment) *CreateDraftContractTermsDocumentOptions {
	_options.Attachment = attachment
	return _options
}

// SetUploadURL : Allow user to set UploadURL
func (_options *CreateDraftContractTermsDocumentOptions) SetUploadURL(uploadURL string) *CreateDraftContractTermsDocumentOptions {
	_options.UploadURL = core.StringPtr(uploadURL)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *CreateDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// DataProduct : Data Product.
type DataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// Summary of Data Product Version object.
	LatestRelease *DataProductVersionSummary `json:"latest_release,omitempty"`

	// List of draft summaries of this data product.
	Drafts []DataProductVersionSummary `json:"drafts,omitempty"`
}

// UnmarshalDataProduct unmarshals an instance of DataProduct from the specified map of raw messages.
func UnmarshalDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "latest_release", &obj.LatestRelease, UnmarshalDataProductVersionSummary)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "drafts", &obj.Drafts, UnmarshalDataProductVersionSummary)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductContractTerms : DataProductContractTerms struct
type DataProductContractTerms struct {
	Asset *AssetReference `json:"asset,omitempty"`

	// ID of the contract terms.
	ID *string `json:"id,omitempty"`

	// Collection of contract terms documents.
	Documents []ContractTermsDocument `json:"documents,omitempty"`
}

// UnmarshalDataProductContractTerms unmarshals an instance of DataProductContractTerms from the specified map of raw messages.
func UnmarshalDataProductContractTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductContractTerms)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalContractTermsDocument)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraftCollection : A collection of data product draft summaries.
type DataProductDraftCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Collection of data product drafts.
	Drafts []DataProductVersionSummary `json:"drafts" validate:"required"`
}

// UnmarshalDataProductDraftCollection unmarshals an instance of DataProductDraftCollection from the specified map of raw messages.
func UnmarshalDataProductDraftCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "drafts", &obj.Drafts, UnmarshalDataProductVersionSummary)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductDraftCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductIdentity : Data product identifier.
type DataProductIdentity struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`
}

// NewDataProductIdentity : Instantiate DataProductIdentity (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewDataProductIdentity(id string) (_model *DataProductIdentity, err error) {
	_model = &DataProductIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataProductIdentity unmarshals an instance of DataProductIdentity from the specified map of raw messages.
func UnmarshalDataProductIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductOrderAccessRequest : The approval workflows associated with the data product version.
type DataProductOrderAccessRequest struct {
	// The workflow approvers associated with the data product version.
	TaskAssigneeUsers []string `json:"task_assignee_users,omitempty"`
}

// UnmarshalDataProductOrderAccessRequest unmarshals an instance of DataProductOrderAccessRequest from the specified map of raw messages.
func UnmarshalDataProductOrderAccessRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductOrderAccessRequest)
	err = core.UnmarshalPrimitive(m, "task_assignee_users", &obj.TaskAssigneeUsers)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductPart : Data Product Part.
type DataProductPart struct {
	// The asset represented in this part.
	Asset *AssetPartReference `json:"asset" validate:"required"`

	// Delivery methods describing the delivery options available for this part.
	DeliveryMethods []DeliveryMethod `json:"delivery_methods,omitempty"`
}

// NewDataProductPart : Instantiate DataProductPart (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewDataProductPart(asset *AssetPartReference) (_model *DataProductPart, err error) {
	_model = &DataProductPart{
		Asset: asset,
	}
	err = core.ValidateStruct(_model, "required parameters")

	return
}

// UnmarshalDataProductPart unmarshals an instance of DataProductPart from the specified map of raw messages.
func UnmarshalDataProductPart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductPart)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetPartReference)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "delivery_methods", &obj.DeliveryMethods, UnmarshalDeliveryMethod)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductReleaseCollection : A collection of data product release summaries.
type DataProductReleaseCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Collection of data product releases.
	Releases []DataProductVersionSummary `json:"releases" validate:"required"`
}

// UnmarshalDataProductReleaseCollection unmarshals an instance of DataProductReleaseCollection from the specified map of raw messages.
func UnmarshalDataProductReleaseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductReleaseCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "releases", &obj.Releases, UnmarshalDataProductVersionSummary)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductReleaseCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductSummary : Data Product Summary.
type DataProductSummary struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductSummary unmarshals an instance of DataProductSummary from the specified map of raw messages.
func UnmarshalDataProductSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductSummaryCollection : A collection of data product summaries.
type DataProductSummaryCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Collection of data product summaries.
	DataProducts []DataProductSummary `json:"data_products" validate:"required"`
}

// UnmarshalDataProductSummaryCollection unmarshals an instance of DataProductSummaryCollection from the specified map of raw messages.
func UnmarshalDataProductSummaryCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductSummaryCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "data_products", &obj.DataProducts, UnmarshalDataProductSummary)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductSummaryCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductVersion : Data Product version.
type DataProductVersion struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductVersionDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags" validate:"required"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases" validate:"required"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []DataProductContractTerms `json:"contract_terms" validate:"required"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	Asset *AssetReference `json:"asset" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain" validate:"required"`

	// Outgoing parts of a data product used to deliver the data product to consumers.
	PartsOut []DataProductPart `json:"parts_out" validate:"required"`

	// The user who published this data product version.
	PublishedBy *string `json:"published_by,omitempty"`

	// The time when this data product version was published.
	PublishedAt *strfmt.DateTime `json:"published_at,omitempty"`

	// The creator of this data product version.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The time when this data product version was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`
}

// Constants associated with the DataProductVersion.State property.
// The state of the data product version.
const (
	DataProductVersion_State_Available = "available"
	DataProductVersion_State_Draft     = "draft"
	DataProductVersion_State_Retired   = "retired"
)

// Constants associated with the DataProductVersion.Types property.
const (
	DataProductVersion_Types_Code = "code"
	DataProductVersion_Types_Data = "data"
)

// UnmarshalDataProductVersion unmarshals an instance of DataProductVersion from the specified map of raw messages.
func UnmarshalDataProductVersion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersion)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductVersionDataProduct)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalDataProductContractTerms)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "published_by", &obj.PublishedBy)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "published_at", &obj.PublishedAt)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubApiServiceV1) NewDataProductVersionPatch(dataProductVersion *DataProductVersion) (_patch []JSONPatchOperation) {
	if dataProductVersion.Version != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/version"),
			Value: dataProductVersion.Version,
		})
	}
	if dataProductVersion.State != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/state"),
			Value: dataProductVersion.State,
		})
	}
	if dataProductVersion.DataProduct != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/data_product"),
			Value: dataProductVersion.DataProduct,
		})
	}
	if dataProductVersion.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: dataProductVersion.Name,
		})
	}
	if dataProductVersion.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: dataProductVersion.Description,
		})
	}
	if dataProductVersion.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: dataProductVersion.Tags,
		})
	}
	if dataProductVersion.UseCases != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/use_cases"),
			Value: dataProductVersion.UseCases,
		})
	}
	if dataProductVersion.Types != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/types"),
			Value: dataProductVersion.Types,
		})
	}
	if dataProductVersion.ContractTerms != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/contract_terms"),
			Value: dataProductVersion.ContractTerms,
		})
	}
	if dataProductVersion.IsRestricted != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/is_restricted"),
			Value: dataProductVersion.IsRestricted,
		})
	}
	if dataProductVersion.ID != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/id"),
			Value: dataProductVersion.ID,
		})
	}
	if dataProductVersion.Asset != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/asset"),
			Value: dataProductVersion.Asset,
		})
	}
	if dataProductVersion.Domain != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/domain"),
			Value: dataProductVersion.Domain,
		})
	}
	if dataProductVersion.PartsOut != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/parts_out"),
			Value: dataProductVersion.PartsOut,
		})
	}
	if dataProductVersion.PublishedBy != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/published_by"),
			Value: dataProductVersion.PublishedBy,
		})
	}
	if dataProductVersion.PublishedAt != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/published_at"),
			Value: dataProductVersion.PublishedAt,
		})
	}
	if dataProductVersion.CreatedBy != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/created_by"),
			Value: dataProductVersion.CreatedBy,
		})
	}
	if dataProductVersion.CreatedAt != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/created_at"),
			Value: dataProductVersion.CreatedAt,
		})
	}
	if dataProductVersion.Workflows != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/workflows"),
			Value: dataProductVersion.Workflows,
		})
	}
	return
}

// DataProductVersionDataProduct : Data product reference.
type DataProductVersionDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductVersionDataProduct unmarshals an instance of DataProductVersionDataProduct from the specified map of raw messages.
func UnmarshalDataProductVersionDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersionPrototype : New data product version input properties.
type DataProductVersionPrototype struct {
	// The data product version number.
	Version *string `json:"version,omitempty"`

	// The state of the data product version. If not specified, the data product version will be created in `draft` state.
	State *string `json:"state,omitempty"`

	// Data product identifier.
	DataProduct *DataProductIdentity `json:"data_product,omitempty"`

	// The name that refers to the new data product version. If this is a new data product, this value must be specified.
	// If this is a new version of an existing data product, the name will default to the name of the previous data product
	// version. A name can contain letters, numbers, understores, dashes, spaces or periods. A name must contain at least
	// one non-space character.
	Name *string `json:"name,omitempty"`

	// Description of the data product version. If this is a new version of an existing data product, the description will
	// default to the description of the previous version of the data product.
	Description *string `json:"description,omitempty"`

	// Tags on the data product.
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types,omitempty"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []DataProductContractTerms `json:"contract_terms,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted,omitempty"`

	// New asset input properties.
	Asset *AssetPrototype `json:"asset" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain,omitempty"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out,omitempty"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`
}

// Constants associated with the DataProductVersionPrototype.State property.
// The state of the data product version. If not specified, the data product version will be created in `draft` state.
const (
	DataProductVersionPrototype_State_Available = "available"
	DataProductVersionPrototype_State_Draft     = "draft"
	DataProductVersionPrototype_State_Retired   = "retired"
)

// Constants associated with the DataProductVersionPrototype.Types property.
const (
	DataProductVersionPrototype_Types_Code = "code"
	DataProductVersionPrototype_Types_Data = "data"
)

// NewDataProductVersionPrototype : Instantiate DataProductVersionPrototype (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewDataProductVersionPrototype(asset *AssetPrototype) (_model *DataProductVersionPrototype, err error) {
	_model = &DataProductVersionPrototype{
		Asset: asset,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataProductVersionPrototype unmarshals an instance of DataProductVersionPrototype from the specified map of raw messages.
func UnmarshalDataProductVersionPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionPrototype)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductIdentity)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalDataProductContractTerms)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetPrototype)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersionSummary : Summary of Data Product Version object.
type DataProductVersionSummary struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductVersionSummaryDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags" validate:"required"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases" validate:"required"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []DataProductContractTerms `json:"contract_terms" validate:"required"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	Asset *AssetReference `json:"asset" validate:"required"`
}

// Constants associated with the DataProductVersionSummary.State property.
// The state of the data product version.
const (
	DataProductVersionSummary_State_Available = "available"
	DataProductVersionSummary_State_Draft     = "draft"
	DataProductVersionSummary_State_Retired   = "retired"
)

// Constants associated with the DataProductVersionSummary.Types property.
const (
	DataProductVersionSummary_Types_Code = "code"
	DataProductVersionSummary_Types_Data = "data"
)

// UnmarshalDataProductVersionSummary unmarshals an instance of DataProductVersionSummary from the specified map of raw messages.
func UnmarshalDataProductVersionSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductVersionSummaryDataProduct)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalDataProductContractTerms)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersionSummaryDataProduct : Data product reference.
type DataProductVersionSummaryDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductVersionSummaryDataProduct unmarshals an instance of DataProductVersionSummaryDataProduct from the specified map of raw messages.
func UnmarshalDataProductVersionSummaryDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionSummaryDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductWorkflows : The workflows associated with the data product version.
type DataProductWorkflows struct {
	// The approval workflows associated with the data product version.
	OrderAccessRequest *DataProductOrderAccessRequest `json:"order_access_request,omitempty"`
}

// UnmarshalDataProductWorkflows unmarshals an instance of DataProductWorkflows from the specified map of raw messages.
func UnmarshalDataProductWorkflows(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductWorkflows)
	err = core.UnmarshalModel(m, "order_access_request", &obj.OrderAccessRequest, UnmarshalDataProductOrderAccessRequest)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDataProductDraftOptions : The DeleteDataProductDraft options.
type DeleteDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDataProductDraftOptions : Instantiate DeleteDataProductDraftOptions
func (*DataProductHubApiServiceV1) NewDeleteDataProductDraftOptions(dataProductID string, draftID string) *DeleteDataProductDraftOptions {
	return &DeleteDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID:       core.StringPtr(draftID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *DeleteDataProductDraftOptions) SetDataProductID(dataProductID string) *DeleteDataProductDraftOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *DeleteDataProductDraftOptions) SetDraftID(draftID string) *DeleteDataProductDraftOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDataProductDraftOptions) SetHeaders(param map[string]string) *DeleteDataProductDraftOptions {
	options.Headers = param
	return options
}

// DeleteDraftContractTermsDocumentOptions : The DeleteDraftContractTermsDocument options.
type DeleteDraftContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDraftContractTermsDocumentOptions : Instantiate DeleteDraftContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewDeleteDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *DeleteDraftContractTermsDocumentOptions {
	return &DeleteDraftContractTermsDocumentOptions{
		DataProductID:   core.StringPtr(dataProductID),
		DraftID:         core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID:      core.StringPtr(documentID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *DeleteDraftContractTermsDocumentOptions) SetDataProductID(dataProductID string) *DeleteDraftContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *DeleteDraftContractTermsDocumentOptions) SetDraftID(draftID string) *DeleteDraftContractTermsDocumentOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *DeleteDraftContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *DeleteDraftContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *DeleteDraftContractTermsDocumentOptions) SetDocumentID(documentID string) *DeleteDraftContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *DeleteDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// DeliveryMethod : DeliveryMethod struct
type DeliveryMethod struct {
	// The ID of the delivery method.
	ID *string `json:"id" validate:"required"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// NewDeliveryMethod : Instantiate DeliveryMethod (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewDeliveryMethod(id string, container *ContainerReference) (_model *DeliveryMethod, err error) {
	_model = &DeliveryMethod{
		ID:        core.StringPtr(id),
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDeliveryMethod unmarshals an instance of DeliveryMethod from the specified map of raw messages.
func UnmarshalDeliveryMethod(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeliveryMethod)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Domain : Domain that the data product version belongs to. If this is the first version of a data product, this field is
// required. If this is a new version of an existing data product, the domain will default to the domain of the previous
// version of the data product.
type Domain struct {
	// The ID of the domain.
	ID *string `json:"id" validate:"required"`

	// The display name of the domain.
	Name *string `json:"name,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewDomain : Instantiate Domain (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewDomain(id string) (_model *Domain, err error) {
	_model = &Domain{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDomain unmarshals an instance of Domain from the specified map of raw messages.
func UnmarshalDomain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Domain)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorModelResource : Detailed error information.
type ErrorModelResource struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message.
	Message *string `json:"message,omitempty"`

	// Extra information about the error.
	Extra map[string]interface{} `json:"extra,omitempty"`

	// More info message.
	MoreInfo *string `json:"more_info,omitempty"`
}

// Constants associated with the ErrorModelResource.Code property.
// Error code.
const (
	ErrorModelResource_Code_AlreadyExists          = "already_exists"
	ErrorModelResource_Code_ConfigurationError     = "configuration_error"
	ErrorModelResource_Code_Conflict               = "conflict"
	ErrorModelResource_Code_ConstraintViolation    = "constraint_violation"
	ErrorModelResource_Code_CreateError            = "create_error"
	ErrorModelResource_Code_DataError              = "data_error"
	ErrorModelResource_Code_DatabaseError          = "database_error"
	ErrorModelResource_Code_DatabaseQueryError     = "database_query_error"
	ErrorModelResource_Code_DatabaseUsageLimits    = "database_usage_limits"
	ErrorModelResource_Code_DeleteError            = "delete_error"
	ErrorModelResource_Code_Deleted                = "deleted"
	ErrorModelResource_Code_DependentServiceError  = "dependent_service_error"
	ErrorModelResource_Code_DoesNotExist           = "does_not_exist"
	ErrorModelResource_Code_EntitlementEnforcement = "entitlement_enforcement"
	ErrorModelResource_Code_FeatureNotEnabled      = "feature_not_enabled"
	ErrorModelResource_Code_FetchError             = "fetch_error"
	ErrorModelResource_Code_Forbidden              = "forbidden"
	ErrorModelResource_Code_GovernancePolicyDenial = "governance_policy_denial"
	ErrorModelResource_Code_InactiveUser           = "inactive_user"
	ErrorModelResource_Code_InvalidParameter       = "invalid_parameter"
	ErrorModelResource_Code_MissingRequiredValue   = "missing_required_value"
	ErrorModelResource_Code_NotAuthenticated       = "not_authenticated"
	ErrorModelResource_Code_NotAuthorized          = "not_authorized"
	ErrorModelResource_Code_NotImplemented         = "not_implemented"
	ErrorModelResource_Code_RequestBodyError       = "request_body_error"
	ErrorModelResource_Code_TooManyRequests        = "too_many_requests"
	ErrorModelResource_Code_UnableToPerform        = "unable_to_perform"
	ErrorModelResource_Code_UnexpectedException    = "unexpected_exception"
	ErrorModelResource_Code_UpdateError            = "update_error"
)

// UnmarshalErrorModelResource unmarshals an instance of ErrorModelResource from the specified map of raw messages.
func UnmarshalErrorModelResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorModelResource)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "extra", &obj.Extra)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FirstPage : First page in the collection.
type FirstPage struct {
	// Link to the first page in the collection.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalFirstPage unmarshals an instance of FirstPage from the specified map of raw messages.
func UnmarshalFirstPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FirstPage)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDataProductDraftOptions : The GetDataProductDraft options.
type GetDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductDraftOptions : Instantiate GetDataProductDraftOptions
func (*DataProductHubApiServiceV1) NewGetDataProductDraftOptions(dataProductID string, draftID string) *GetDataProductDraftOptions {
	return &GetDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID:       core.StringPtr(draftID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetDataProductDraftOptions) SetDataProductID(dataProductID string) *GetDataProductDraftOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *GetDataProductDraftOptions) SetDraftID(draftID string) *GetDataProductDraftOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductDraftOptions) SetHeaders(param map[string]string) *GetDataProductDraftOptions {
	options.Headers = param
	return options
}

// GetDataProductOptions : The GetDataProduct options.
type GetDataProductOptions struct {
	// Data product ID.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductOptions : Instantiate GetDataProductOptions
func (*DataProductHubApiServiceV1) NewGetDataProductOptions(dataProductID string) *GetDataProductOptions {
	return &GetDataProductOptions{
		DataProductID: core.StringPtr(dataProductID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetDataProductOptions) SetDataProductID(dataProductID string) *GetDataProductOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductOptions) SetHeaders(param map[string]string) *GetDataProductOptions {
	options.Headers = param
	return options
}

// GetDataProductReleaseOptions : The GetDataProductRelease options.
type GetDataProductReleaseOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDataProductReleaseOptions : Instantiate GetDataProductReleaseOptions
func (*DataProductHubApiServiceV1) NewGetDataProductReleaseOptions(dataProductID string, releaseID string) *GetDataProductReleaseOptions {
	return &GetDataProductReleaseOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID:     core.StringPtr(releaseID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetDataProductReleaseOptions) SetDataProductID(dataProductID string) *GetDataProductReleaseOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetReleaseID : Allow user to set ReleaseID
func (_options *GetDataProductReleaseOptions) SetReleaseID(releaseID string) *GetDataProductReleaseOptions {
	_options.ReleaseID = core.StringPtr(releaseID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductReleaseOptions) SetHeaders(param map[string]string) *GetDataProductReleaseOptions {
	options.Headers = param
	return options
}

// GetDraftContractTermsDocumentOptions : The GetDraftContractTermsDocument options.
type GetDraftContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDraftContractTermsDocumentOptions : Instantiate GetDraftContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewGetDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *GetDraftContractTermsDocumentOptions {
	return &GetDraftContractTermsDocumentOptions{
		DataProductID:   core.StringPtr(dataProductID),
		DraftID:         core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID:      core.StringPtr(documentID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetDraftContractTermsDocumentOptions) SetDataProductID(dataProductID string) *GetDraftContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *GetDraftContractTermsDocumentOptions) SetDraftID(draftID string) *GetDraftContractTermsDocumentOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *GetDraftContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *GetDraftContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetDraftContractTermsDocumentOptions) SetDocumentID(documentID string) *GetDraftContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *GetDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// GetInitializeStatusOptions : The GetInitializeStatus options.
type GetInitializeStatusOptions struct {
	// Container ID of the data product catalog. If not supplied, the data product catalog is looked up by using the uid of
	// the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetInitializeStatusOptions : Instantiate GetInitializeStatusOptions
func (*DataProductHubApiServiceV1) NewGetInitializeStatusOptions() *GetInitializeStatusOptions {
	return &GetInitializeStatusOptions{}
}

// SetContainerID : Allow user to set ContainerID
func (_options *GetInitializeStatusOptions) SetContainerID(containerID string) *GetInitializeStatusOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInitializeStatusOptions) SetHeaders(param map[string]string) *GetInitializeStatusOptions {
	options.Headers = param
	return options
}

// GetReleaseContractTermsDocumentOptions : The GetReleaseContractTermsDocument options.
type GetReleaseContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReleaseContractTermsDocumentOptions : Instantiate GetReleaseContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewGetReleaseContractTermsDocumentOptions(dataProductID string, releaseID string, contractTermsID string, documentID string) *GetReleaseContractTermsDocumentOptions {
	return &GetReleaseContractTermsDocumentOptions{
		DataProductID:   core.StringPtr(dataProductID),
		ReleaseID:       core.StringPtr(releaseID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID:      core.StringPtr(documentID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetReleaseContractTermsDocumentOptions) SetDataProductID(dataProductID string) *GetReleaseContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetReleaseID : Allow user to set ReleaseID
func (_options *GetReleaseContractTermsDocumentOptions) SetReleaseID(releaseID string) *GetReleaseContractTermsDocumentOptions {
	_options.ReleaseID = core.StringPtr(releaseID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *GetReleaseContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *GetReleaseContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *GetReleaseContractTermsDocumentOptions) SetDocumentID(documentID string) *GetReleaseContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetReleaseContractTermsDocumentOptions) SetHeaders(param map[string]string) *GetReleaseContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// GetServiceIdCredentialsOptions : The GetServiceIdCredentials options.
type GetServiceIdCredentialsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceIdCredentialsOptions : Instantiate GetServiceIdCredentialsOptions
func (*DataProductHubApiServiceV1) NewGetServiceIdCredentialsOptions() *GetServiceIdCredentialsOptions {
	return &GetServiceIdCredentialsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceIdCredentialsOptions) SetHeaders(param map[string]string) *GetServiceIdCredentialsOptions {
	options.Headers = param
	return options
}

// InitializeOptions : The Initialize options.
type InitializeOptions struct {
	// Container reference.
	Container *ContainerReference `json:"container,omitempty"`

	// List of configuration options to (re-)initialize.
	Include []string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the InitializeOptions.Include property.
const (
	InitializeOptions_Include_DataProductSamples   = "data_product_samples"
	InitializeOptions_Include_DeliveryMethods      = "delivery_methods"
	InitializeOptions_Include_DomainsMultiIndustry = "domains_multi_industry"
	InitializeOptions_Include_Project              = "project"
	InitializeOptions_Include_Workflows            = "workflows"
)

// NewInitializeOptions : Instantiate InitializeOptions
func (*DataProductHubApiServiceV1) NewInitializeOptions() *InitializeOptions {
	return &InitializeOptions{}
}

// SetContainer : Allow user to set Container
func (_options *InitializeOptions) SetContainer(container *ContainerReference) *InitializeOptions {
	_options.Container = container
	return _options
}

// SetInclude : Allow user to set Include
func (_options *InitializeOptions) SetInclude(include []string) *InitializeOptions {
	_options.Include = include
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InitializeOptions) SetHeaders(param map[string]string) *InitializeOptions {
	options.Headers = param
	return options
}

// InitializeResource : Resource defining initialization parameters.
type InitializeResource struct {
	// Container reference.
	Container *ContainerReference `json:"container,omitempty"`

	// Link to monitor the status of the initialize operation.
	Href *string `json:"href,omitempty"`

	// Status of the initialize operation.
	Status *string `json:"status,omitempty"`

	// The id to trace the failed initialization operation.
	Trace *string `json:"trace,omitempty"`

	// Set of errors on the latest initialization request.
	Errors []ErrorModelResource `json:"errors,omitempty"`

	// Start time of the last initialization.
	LastStartedAt *strfmt.DateTime `json:"last_started_at,omitempty"`

	// End time of the last initialization.
	LastFinishedAt *strfmt.DateTime `json:"last_finished_at,omitempty"`

	// Initialized options.
	InitializedOptions []InitializedOption `json:"initialized_options,omitempty"`
}

// Constants associated with the InitializeResource.Status property.
// Status of the initialize operation.
const (
	InitializeResource_Status_Failed     = "failed"
	InitializeResource_Status_InProgress = "in_progress"
	InitializeResource_Status_NotStarted = "not_started"
	InitializeResource_Status_Succeeded  = "succeeded"
)

// UnmarshalInitializeResource unmarshals an instance of InitializeResource from the specified map of raw messages.
func UnmarshalInitializeResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializeResource)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModelResource)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "last_started_at", &obj.LastStartedAt)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "last_finished_at", &obj.LastFinishedAt)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "initialized_options", &obj.InitializedOptions, UnmarshalInitializedOption)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InitializedOption : List of options successfully initialized.
type InitializedOption struct {
	// The name of the option.
	Name *string `json:"name,omitempty"`

	// The version of the option.
	Version *int64 `json:"version,omitempty"`
}

// UnmarshalInitializedOption unmarshals an instance of InitializedOption from the specified map of raw messages.
func UnmarshalInitializedOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializedOption)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add     = "add"
	JSONPatchOperation_Op_Copy    = "copy"
	JSONPatchOperation_Op_Move    = "move"
	JSONPatchOperation_Op_Remove  = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test    = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op:   core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDataProductDraftsOptions : The ListDataProductDrafts options.
type ListDataProductDraftsOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Filter the list of data product drafts by container id.
	AssetContainerID *string `json:"asset.container.id,omitempty"`

	// Filter the list of data product drafts by version number.
	Version *string `json:"version,omitempty"`

	// Limit the number of data product drafts in the results. The maximum limit is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDataProductDraftsOptions : Instantiate ListDataProductDraftsOptions
func (*DataProductHubApiServiceV1) NewListDataProductDraftsOptions(dataProductID string) *ListDataProductDraftsOptions {
	return &ListDataProductDraftsOptions{
		DataProductID: core.StringPtr(dataProductID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *ListDataProductDraftsOptions) SetDataProductID(dataProductID string) *ListDataProductDraftsOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetAssetContainerID : Allow user to set AssetContainerID
func (_options *ListDataProductDraftsOptions) SetAssetContainerID(assetContainerID string) *ListDataProductDraftsOptions {
	_options.AssetContainerID = core.StringPtr(assetContainerID)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ListDataProductDraftsOptions) SetVersion(version string) *ListDataProductDraftsOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListDataProductDraftsOptions) SetLimit(limit int64) *ListDataProductDraftsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListDataProductDraftsOptions) SetStart(start string) *ListDataProductDraftsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductDraftsOptions) SetHeaders(param map[string]string) *ListDataProductDraftsOptions {
	options.Headers = param
	return options
}

// ListDataProductReleasesOptions : The ListDataProductReleases options.
type ListDataProductReleasesOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Filter the list of data product releases by container id.
	AssetContainerID *string `json:"asset.container.id,omitempty"`

	// Filter the list of data product versions by state. States are: available and retired. Default is
	// "available","retired".
	State []string `json:"state,omitempty"`

	// Filter the list of data product releases by version number.
	Version *string `json:"version,omitempty"`

	// Limit the number of data product releases in the results. The maximum is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListDataProductReleasesOptions.State property.
const (
	ListDataProductReleasesOptions_State_Available = "available"
	ListDataProductReleasesOptions_State_Retired   = "retired"
)

// NewListDataProductReleasesOptions : Instantiate ListDataProductReleasesOptions
func (*DataProductHubApiServiceV1) NewListDataProductReleasesOptions(dataProductID string) *ListDataProductReleasesOptions {
	return &ListDataProductReleasesOptions{
		DataProductID: core.StringPtr(dataProductID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *ListDataProductReleasesOptions) SetDataProductID(dataProductID string) *ListDataProductReleasesOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetAssetContainerID : Allow user to set AssetContainerID
func (_options *ListDataProductReleasesOptions) SetAssetContainerID(assetContainerID string) *ListDataProductReleasesOptions {
	_options.AssetContainerID = core.StringPtr(assetContainerID)
	return _options
}

// SetState : Allow user to set State
func (_options *ListDataProductReleasesOptions) SetState(state []string) *ListDataProductReleasesOptions {
	_options.State = state
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ListDataProductReleasesOptions) SetVersion(version string) *ListDataProductReleasesOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListDataProductReleasesOptions) SetLimit(limit int64) *ListDataProductReleasesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListDataProductReleasesOptions) SetStart(start string) *ListDataProductReleasesOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductReleasesOptions) SetHeaders(param map[string]string) *ListDataProductReleasesOptions {
	options.Headers = param
	return options
}

// ListDataProductsOptions : The ListDataProducts options.
type ListDataProductsOptions struct {
	// Limit the number of data products in the results. The maximum limit is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDataProductsOptions : Instantiate ListDataProductsOptions
func (*DataProductHubApiServiceV1) NewListDataProductsOptions() *ListDataProductsOptions {
	return &ListDataProductsOptions{}
}

// SetLimit : Allow user to set Limit
func (_options *ListDataProductsOptions) SetLimit(limit int64) *ListDataProductsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListDataProductsOptions) SetStart(start string) *ListDataProductsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductsOptions) SetHeaders(param map[string]string) *ListDataProductsOptions {
	options.Headers = param
	return options
}

// ManageApiKeysOptions : The ManageApiKeys options.
type ManageApiKeysOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewManageApiKeysOptions : Instantiate ManageApiKeysOptions
func (*DataProductHubApiServiceV1) NewManageApiKeysOptions() *ManageApiKeysOptions {
	return &ManageApiKeysOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ManageApiKeysOptions) SetHeaders(param map[string]string) *ManageApiKeysOptions {
	options.Headers = param
	return options
}

// NextPage : Next page in the collection.
type NextPage struct {
	// Link to the next page in the collection.
	Href *string `json:"href" validate:"required"`

	// Start token for pagination to the next page in the collection.
	Start *string `json:"start" validate:"required"`
}

// UnmarshalNextPage unmarshals an instance of NextPage from the specified map of raw messages.
func UnmarshalNextPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NextPage)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PublishDataProductDraftOptions : The PublishDataProductDraft options.
type PublishDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPublishDataProductDraftOptions : Instantiate PublishDataProductDraftOptions
func (*DataProductHubApiServiceV1) NewPublishDataProductDraftOptions(dataProductID string, draftID string) *PublishDataProductDraftOptions {
	return &PublishDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID:       core.StringPtr(draftID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *PublishDataProductDraftOptions) SetDataProductID(dataProductID string) *PublishDataProductDraftOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *PublishDataProductDraftOptions) SetDraftID(draftID string) *PublishDataProductDraftOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PublishDataProductDraftOptions) SetHeaders(param map[string]string) *PublishDataProductDraftOptions {
	options.Headers = param
	return options
}

// RetireDataProductReleaseOptions : The RetireDataProductRelease options.
type RetireDataProductReleaseOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRetireDataProductReleaseOptions : Instantiate RetireDataProductReleaseOptions
func (*DataProductHubApiServiceV1) NewRetireDataProductReleaseOptions(dataProductID string, releaseID string) *RetireDataProductReleaseOptions {
	return &RetireDataProductReleaseOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID:     core.StringPtr(releaseID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *RetireDataProductReleaseOptions) SetDataProductID(dataProductID string) *RetireDataProductReleaseOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetReleaseID : Allow user to set ReleaseID
func (_options *RetireDataProductReleaseOptions) SetReleaseID(releaseID string) *RetireDataProductReleaseOptions {
	_options.ReleaseID = core.StringPtr(releaseID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RetireDataProductReleaseOptions) SetHeaders(param map[string]string) *RetireDataProductReleaseOptions {
	options.Headers = param
	return options
}

// ServiceIdCredentials : Service id credentials.
type ServiceIdCredentials struct {
	// Name of the api key of the service id.
	Name *string `json:"name,omitempty"`

	// Created date of the api key of the service id.
	CreatedAt *string `json:"created_at,omitempty"`
}

// UnmarshalServiceIdCredentials unmarshals an instance of ServiceIdCredentials from the specified map of raw messages.
func UnmarshalServiceIdCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceIdCredentials)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateDataProductDraftOptions : The UpdateDataProductDraft options.
type UpdateDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDataProductDraftOptions : Instantiate UpdateDataProductDraftOptions
func (*DataProductHubApiServiceV1) NewUpdateDataProductDraftOptions(dataProductID string, draftID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDraftOptions {
	return &UpdateDataProductDraftOptions{
		DataProductID:         core.StringPtr(dataProductID),
		DraftID:               core.StringPtr(draftID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *UpdateDataProductDraftOptions) SetDataProductID(dataProductID string) *UpdateDataProductDraftOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *UpdateDataProductDraftOptions) SetDraftID(draftID string) *UpdateDataProductDraftOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductDraftOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDraftOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductDraftOptions) SetHeaders(param map[string]string) *UpdateDataProductDraftOptions {
	options.Headers = param
	return options
}

// UpdateDataProductReleaseOptions : The UpdateDataProductRelease options.
type UpdateDataProductReleaseOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDataProductReleaseOptions : Instantiate UpdateDataProductReleaseOptions
func (*DataProductHubApiServiceV1) NewUpdateDataProductReleaseOptions(dataProductID string, releaseID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductReleaseOptions {
	return &UpdateDataProductReleaseOptions{
		DataProductID:         core.StringPtr(dataProductID),
		ReleaseID:             core.StringPtr(releaseID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *UpdateDataProductReleaseOptions) SetDataProductID(dataProductID string) *UpdateDataProductReleaseOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetReleaseID : Allow user to set ReleaseID
func (_options *UpdateDataProductReleaseOptions) SetReleaseID(releaseID string) *UpdateDataProductReleaseOptions {
	_options.ReleaseID = core.StringPtr(releaseID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductReleaseOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductReleaseOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductReleaseOptions) SetHeaders(param map[string]string) *UpdateDataProductReleaseOptions {
	options.Headers = param
	return options
}

// UpdateDraftContractTermsDocumentOptions : The UpdateDraftContractTermsDocument options.
type UpdateDraftContractTermsDocumentOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// Document id.
	DocumentID *string `json:"document_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDraftContractTermsDocumentOptions : Instantiate UpdateDraftContractTermsDocumentOptions
func (*DataProductHubApiServiceV1) NewUpdateDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDraftContractTermsDocumentOptions {
	return &UpdateDraftContractTermsDocumentOptions{
		DataProductID:         core.StringPtr(dataProductID),
		DraftID:               core.StringPtr(draftID),
		ContractTermsID:       core.StringPtr(contractTermsID),
		DocumentID:            core.StringPtr(documentID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *UpdateDraftContractTermsDocumentOptions) SetDataProductID(dataProductID string) *UpdateDraftContractTermsDocumentOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *UpdateDraftContractTermsDocumentOptions) SetDraftID(draftID string) *UpdateDraftContractTermsDocumentOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *UpdateDraftContractTermsDocumentOptions) SetContractTermsID(contractTermsID string) *UpdateDraftContractTermsDocumentOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetDocumentID : Allow user to set DocumentID
func (_options *UpdateDraftContractTermsDocumentOptions) SetDocumentID(documentID string) *UpdateDraftContractTermsDocumentOptions {
	_options.DocumentID = core.StringPtr(documentID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDraftContractTermsDocumentOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDraftContractTermsDocumentOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *UpdateDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// UseCase : UseCase struct
type UseCase struct {
	// The id of the use case associated with the data product.
	ID *string `json:"id" validate:"required"`

	// The display name of the use case associated with the data product.
	Name *string `json:"name,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container,omitempty"`
}

// NewUseCase : Instantiate UseCase (Generic Model Constructor)
func (*DataProductHubApiServiceV1) NewUseCase(id string) (_model *UseCase, err error) {
	_model = &UseCase{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalUseCase unmarshals an instance of UseCase from the specified map of raw messages.
func UnmarshalUseCase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UseCase)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {

		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {

		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {

		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductsPager can be used to simplify the use of the "ListDataProducts" method.
type DataProductsPager struct {
	hasNext     bool
	options     *ListDataProductsOptions
	client      *DataProductHubApiServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductsPager returns a new DataProductsPager instance.
func (dataProductHubApiService *DataProductHubApiServiceV1) NewDataProductsPager(options *ListDataProductsOptions) (pager *DataProductsPager, err error) {
	if options.Start != nil && *options.Start != "" {

		return
	}

	var optionsCopy ListDataProductsOptions = *options
	pager = &DataProductsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dataProductHubApiService,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DataProductsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DataProductsPager) GetNextWithContext(ctx context.Context) (page []DataProductSummary, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListDataProductsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.DataProducts

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DataProductsPager) GetAllWithContext(ctx context.Context) (allItems []DataProductSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductSummary
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DataProductsPager) GetNext() (page []DataProductSummary, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductsPager) GetAll() (allItems []DataProductSummary, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DataProductDraftsPager can be used to simplify the use of the "ListDataProductDrafts" method.
type DataProductDraftsPager struct {
	hasNext     bool
	options     *ListDataProductDraftsOptions
	client      *DataProductHubApiServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductDraftsPager returns a new DataProductDraftsPager instance.
func (dataProductHubApiService *DataProductHubApiServiceV1) NewDataProductDraftsPager(options *ListDataProductDraftsOptions) (pager *DataProductDraftsPager, err error) {
	if options.Start != nil && *options.Start != "" {

		return
	}

	var optionsCopy ListDataProductDraftsOptions = *options
	pager = &DataProductDraftsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dataProductHubApiService,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DataProductDraftsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DataProductDraftsPager) GetNextWithContext(ctx context.Context) (page []DataProductVersionSummary, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListDataProductDraftsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Drafts

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DataProductDraftsPager) GetAllWithContext(ctx context.Context) (allItems []DataProductVersionSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductVersionSummary
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DataProductDraftsPager) GetNext() (page []DataProductVersionSummary, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductDraftsPager) GetAll() (allItems []DataProductVersionSummary, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DataProductReleasesPager can be used to simplify the use of the "ListDataProductReleases" method.
type DataProductReleasesPager struct {
	hasNext     bool
	options     *ListDataProductReleasesOptions
	client      *DataProductHubApiServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductReleasesPager returns a new DataProductReleasesPager instance.
func (dataProductHubApiService *DataProductHubApiServiceV1) NewDataProductReleasesPager(options *ListDataProductReleasesOptions) (pager *DataProductReleasesPager, err error) {
	if options.Start != nil && *options.Start != "" {

		return
	}

	var optionsCopy ListDataProductReleasesOptions = *options
	pager = &DataProductReleasesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  dataProductHubApiService,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DataProductReleasesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DataProductReleasesPager) GetNextWithContext(ctx context.Context) (page []DataProductVersionSummary, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListDataProductReleasesWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Releases

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DataProductReleasesPager) GetAllWithContext(ctx context.Context) (allItems []DataProductVersionSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductVersionSummary
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DataProductReleasesPager) GetNext() (page []DataProductVersionSummary, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductReleasesPager) GetAll() (allItems []DataProductVersionSummary, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
