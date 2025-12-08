/**
 * (C) Copyright IBM Corp. 2025.
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
 * IBM OpenAPI SDK Code Generator Version: 3.105.1-067d600b-20250616-154447
 */

// Package dataproducthubapiservicev1 : Operations and models for the DataProductHubAPIServiceV1 service
package dataproducthubapiservicev1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	common "github.com/IBM/data-product-exchange-go-sdk/v4/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"io"

)

// DataProductHubAPIServiceV1 : Data Product Hub API Service
//
// API Version: 1
type DataProductHubAPIServiceV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "data_product_hub_api_service"

// DataProductHubAPIServiceV1Options : Service options
type DataProductHubAPIServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDataProductHubAPIServiceV1UsingExternalConfig : constructs an instance of DataProductHubAPIServiceV1 with passed in options and external configuration.
func NewDataProductHubAPIServiceV1UsingExternalConfig(options *DataProductHubAPIServiceV1Options) (dataProductHubApiService *DataProductHubAPIServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	dataProductHubApiService, err = NewDataProductHubAPIServiceV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = dataProductHubApiService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = dataProductHubApiService.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDataProductHubAPIServiceV1 : constructs an instance of DataProductHubAPIServiceV1 with passed in options.
func NewDataProductHubAPIServiceV1(options *DataProductHubAPIServiceV1Options) (service *DataProductHubAPIServiceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &DataProductHubAPIServiceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "dataProductHubApiService" suitable for processing requests.
func (dataProductHubApiService *DataProductHubAPIServiceV1) Clone() *DataProductHubAPIServiceV1 {
	if core.IsNil(dataProductHubApiService) {
		return nil
	}
	clone := *dataProductHubApiService
	clone.Service = dataProductHubApiService.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dataProductHubApiService *DataProductHubAPIServiceV1) SetServiceURL(url string) error {
	err := dataProductHubApiService.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetServiceURL() string {
	return dataProductHubApiService.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dataProductHubApiService *DataProductHubAPIServiceV1) SetDefaultHeaders(headers http.Header) {
	dataProductHubApiService.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dataProductHubApiService *DataProductHubAPIServiceV1) SetEnableGzipCompression(enableGzip bool) {
	dataProductHubApiService.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetEnableGzipCompression() bool {
	return dataProductHubApiService.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dataProductHubApiService *DataProductHubAPIServiceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dataProductHubApiService.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dataProductHubApiService *DataProductHubAPIServiceV1) DisableRetries() {
	dataProductHubApiService.Service.DisableRetries()
}

// GetInitializeStatus : Get resource initialization status
// Use this API to get the status of resource initialization in Data Product Hub.<br/><br/>If the data product catalog
// exists but has never been initialized, the status will be "not_started".<br/><br/>If the data product catalog exists
// and has been or is being initialized, the response will contain the status of the last or current initialization. If
// the initialization failed, the "errors" and "trace" fields will contain the error(s) encountered during the
// initialization, including the ID to trace the error(s).<br/><br/>If the data product catalog doesn't exist, an HTTP
// 404 response is returned.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetInitializeStatusWithContext(context.Background(), getInitializeStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetInitializeStatusWithContext is an alternate form of the GetInitializeStatus method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetInitializeStatusWithContext(ctx context.Context, getInitializeStatusOptions *GetInitializeStatusOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getInitializeStatusOptions, "getInitializeStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize/status`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetInitializeStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getInitializeStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getInitializeStatusOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*getInitializeStatusOptions.ContainerID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_initialize_status", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetServiceIDCredentials : Get service id credentials
// Use this API to get the information of service id credentials in Data Product Hub.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetServiceIDCredentials(getServiceIDCredentialsOptions *GetServiceIDCredentialsOptions) (result *ServiceIDCredentials, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetServiceIDCredentialsWithContext(context.Background(), getServiceIDCredentialsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetServiceIDCredentialsWithContext is an alternate form of the GetServiceIDCredentials method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetServiceIDCredentialsWithContext(ctx context.Context, getServiceIDCredentialsOptions *GetServiceIDCredentialsOptions) (result *ServiceIDCredentials, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getServiceIDCredentialsOptions, "getServiceIDCredentialsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/credentials`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetServiceIDCredentials")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getServiceIDCredentialsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_service_id_credentials", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceIDCredentials)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
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
// exporting data assets to files</li><li>`catalog_configurations` - Catalog configurations for the default data product
// catalog</li></ul><br/><br/>If a resource depends on resources that are not specified in the request, these dependent
// resources will be automatically initialized. E.g., initializing `data_product_samples` will also initialize
// `domains_multi_industry` and `delivery_methods` even if they are not specified in the request because it depends on
// them.<br/><br/>If initializing the data product hub for the first time, do not specify a container. The default data
// product catalog will be created.<br/>For first time initialization, it is recommended that at least
// `delivery_methods` and `domains_multi_industry` is included in the initialize operation.<br/><br/>If the data product
// hub has already been initialized, you may call this API again to initialize new resources, such as new delivery
// methods. In this case, specify the default data product catalog container information.
func (dataProductHubApiService *DataProductHubAPIServiceV1) Initialize(initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.InitializeWithContext(context.Background(), initializeOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// InitializeWithContext is an alternate form of the Initialize method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) InitializeWithContext(ctx context.Context, initializeOptions *InitializeOptions) (result *InitializeResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(initializeOptions, "initializeOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(initializeOptions, "initializeOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/initialize`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "Initialize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range initializeOptions.Headers {
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
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "initialize", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeResource)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ManageAPIKeys : Rotate credentials for a Data Product Hub instance
// Use this API to rotate credentials for a Data Product Hub instance.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ManageAPIKeys(manageAPIKeysOptions *ManageAPIKeysOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.ManageAPIKeysWithContext(context.Background(), manageAPIKeysOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ManageAPIKeysWithContext is an alternate form of the ManageAPIKeys method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ManageAPIKeysWithContext(ctx context.Context, manageAPIKeysOptions *ManageAPIKeysOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(manageAPIKeysOptions, "manageAPIKeysOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/configuration/rotate_credentials`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ManageAPIKeys")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range manageAPIKeysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "manage_api_keys", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateDataAssetVisualization : Create visualization asset and initialize profiling for the provided data assets
// Use this API to create visualization asset and initialize profiling for the provided data assets<br/><br/>Provide the
// below required fields<br/><br/>Required fields:<br/><br/>- catalog_id<br/>- Collection of assetId with it's related
// asset id<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataAssetVisualization(createDataAssetVisualizationOptions *CreateDataAssetVisualizationOptions) (result *DataAssetVisualizationRes, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataAssetVisualizationWithContext(context.Background(), createDataAssetVisualizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataAssetVisualizationWithContext is an alternate form of the CreateDataAssetVisualization method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataAssetVisualizationWithContext(ctx context.Context, createDataAssetVisualizationOptions *CreateDataAssetVisualizationOptions) (result *DataAssetVisualizationRes, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataAssetVisualizationOptions, "createDataAssetVisualizationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDataAssetVisualizationOptions, "createDataAssetVisualizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_asset/visualization`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataAssetVisualization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDataAssetVisualizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDataAssetVisualizationOptions.Assets != nil {
		body["assets"] = createDataAssetVisualizationOptions.Assets
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_asset_visualization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataAssetVisualizationRes)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReinitiateDataAssetVisualization : Reinitiate visualization for an asset
// Use this API to Reinitiate visualization for an asset which is in below scenarios<br/><br/>- Previous bucket got
// deleted and new bucket is created.<br/>- Data visualization attachment is missing in asset details.<br/>-
// Visualization asset reference is missing in related asset details.<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ReinitiateDataAssetVisualization(reinitiateDataAssetVisualizationOptions *ReinitiateDataAssetVisualizationOptions) (result *DataAssetVisualizationRes, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ReinitiateDataAssetVisualizationWithContext(context.Background(), reinitiateDataAssetVisualizationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReinitiateDataAssetVisualizationWithContext is an alternate form of the ReinitiateDataAssetVisualization method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ReinitiateDataAssetVisualizationWithContext(ctx context.Context, reinitiateDataAssetVisualizationOptions *ReinitiateDataAssetVisualizationOptions) (result *DataAssetVisualizationRes, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reinitiateDataAssetVisualizationOptions, "reinitiateDataAssetVisualizationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(reinitiateDataAssetVisualizationOptions, "reinitiateDataAssetVisualizationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_asset/visualization/reinitiate`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ReinitiateDataAssetVisualization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range reinitiateDataAssetVisualizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if reinitiateDataAssetVisualizationOptions.Assets != nil {
		body["assets"] = reinitiateDataAssetVisualizationOptions.Assets
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "reinitiate_data_asset_visualization", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataAssetVisualizationRes)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListDataProducts : Retrieve a list of data products
// Retrieve a list of data products.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProducts(listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductsWithContext(context.Background(), listDataProductsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductsWithContext is an alternate form of the ListDataProducts method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductsWithContext(ctx context.Context, listDataProductsOptions *ListDataProductsOptions) (result *DataProductCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductsOptions, "listDataProductsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProducts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDataProductsOptions.Headers {
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
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_products", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
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
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProduct(createDataProductOptions *CreateDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductWithContext(context.Background(), createDataProductOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductWithContext is an alternate form of the CreateDataProduct method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductWithContext(ctx context.Context, createDataProductOptions *CreateDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductOptions, "createDataProductOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDataProductOptions, "createDataProductOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createDataProductOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*createDataProductOptions.Limit))
	}
	if createDataProductOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*createDataProductOptions.Start))
	}

	body := make(map[string]interface{})
	if createDataProductOptions.Drafts != nil {
		body["drafts"] = createDataProductOptions.Drafts
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProduct)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDataProduct : Retrieve a data product identified by id
// Retrieve a data product identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProduct(getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductWithContext(context.Background(), getDataProductOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductWithContext is an alternate form of the GetDataProduct method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductWithContext(ctx context.Context, getDataProductOptions *GetDataProductOptions) (result *DataProduct, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductOptions, "getDataProductOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDataProductOptions, "getDataProductOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProduct")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDataProductOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProduct)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CompleteDraftContractTermsDocument : Complete a contract document upload operation
// After uploading a file to the provided signed URL, call this endpoint to mark the upload as complete. After the
// upload operation is marked as complete, the file is available to download. Use '-' for the `data_product_id` to skip
// specifying the data product ID explicitly.
// - After the upload is marked as complete, the returned URL is displayed in the "url" field. The signed URL is used to
// download the document.
// - Calling complete on referential documents results in an error.
// - Calling complete on attachment documents for which the file has not been uploaded will result in an error.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CompleteDraftContractTermsDocumentWithContext(context.Background(), completeDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CompleteDraftContractTermsDocumentWithContext is an alternate form of the CompleteDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CompleteDraftContractTermsDocumentWithContext(ctx context.Context, completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeDraftContractTermsDocumentOptions, "completeDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(completeDraftContractTermsDocumentOptions, "completeDraftContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *completeDraftContractTermsDocumentOptions.DataProductID,
		"draft_id": *completeDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *completeDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id": *completeDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}/complete`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CompleteDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range completeDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "complete_draft_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListDataProductDrafts : Retrieve a list of data product drafts
// Retrieve a list of data product drafts. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) (result *DataProductDraftCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductDraftsWithContext(context.Background(), listDataProductDraftsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductDraftsWithContext is an alternate form of the ListDataProductDrafts method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductDraftsWithContext(ctx context.Context, listDataProductDraftsOptions *ListDataProductDraftsOptions) (result *DataProductDraftCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDataProductDraftsOptions, "listDataProductDraftsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listDataProductDraftsOptions, "listDataProductDraftsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductDrafts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDataProductDraftsOptions.Headers {
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
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_drafts", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDraftCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProductDraft : Create a new draft of an existing data product
// Create a new draft of an existing data product.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductDraftWithContext(context.Background(), createDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductDraftWithContext is an alternate form of the CreateDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductDraftWithContext(ctx context.Context, createDataProductDraftOptions *CreateDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductDraftOptions, "createDataProductDraftOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDataProductDraftOptions, "createDataProductDraftOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDataProductDraftOptions.Headers {
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
	if createDataProductDraftOptions.Domain != nil {
		body["domain"] = createDataProductDraftOptions.Domain
	}
	if createDataProductDraftOptions.PartsOut != nil {
		body["parts_out"] = createDataProductDraftOptions.PartsOut
	}
	if createDataProductDraftOptions.Workflows != nil {
		body["workflows"] = createDataProductDraftOptions.Workflows
	}
	if createDataProductDraftOptions.DataviewEnabled != nil {
		body["dataview_enabled"] = createDataProductDraftOptions.DataviewEnabled
	}
	if createDataProductDraftOptions.Comments != nil {
		body["comments"] = createDataProductDraftOptions.Comments
	}
	if createDataProductDraftOptions.AccessControl != nil {
		body["access_control"] = createDataProductDraftOptions.AccessControl
	}
	if createDataProductDraftOptions.LastUpdatedAt != nil {
		body["last_updated_at"] = createDataProductDraftOptions.LastUpdatedAt
	}
	if createDataProductDraftOptions.IsRestricted != nil {
		body["is_restricted"] = createDataProductDraftOptions.IsRestricted
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product_draft", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDraft)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDraftContractTermsDocument : Upload a contract document to the data product draft contract terms
// Upload a contract document to the data product draft identified by draft_id. Use '-' for the `data_product_id` to
// skip specifying the data product ID explicitly.
//
// - If the request object contains a "url" parameter, a referential document is created to store the provided url.
// - If the request object does not contain a "url" parameter, an attachment document is created, and a signed url will
// be returned in an "upload_url" parameter. The data product producer can upload the document using the provided
// "upload_url". After the upload is completed, call "complete_contract_terms_document" for the given document needs to
// be called to mark the upload as completed. After completion of the upload, "get_contract_terms_document" for the
// given document returns a signed "url" parameter that can be used to download the attachment document.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDraftContractTermsDocumentWithContext(context.Background(), createDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDraftContractTermsDocumentWithContext is an alternate form of the CreateDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDraftContractTermsDocumentWithContext(ctx context.Context, createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDraftContractTermsDocumentOptions, "createDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDraftContractTermsDocumentOptions, "createDraftContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *createDraftContractTermsDocumentOptions.DataProductID,
		"draft_id": *createDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *createDraftContractTermsDocumentOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDraftContractTermsDocumentOptions.Headers {
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
	if createDraftContractTermsDocumentOptions.URL != nil {
		body["url"] = createDraftContractTermsDocumentOptions.URL
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_draft_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductDraft : Get a draft of an existing data product
// Get a draft of an existing data product. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductDraftWithContext(context.Background(), getDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductDraftWithContext is an alternate form of the GetDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductDraftWithContext(ctx context.Context, getDataProductDraftOptions *GetDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductDraftOptions, "getDataProductDraftOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDataProductDraftOptions, "getDataProductDraftOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductDraftOptions.DataProductID,
		"draft_id": *getDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_draft", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDraft)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataProductDraft : Delete a data product draft identified by ID
// Delete a data product draft identified by a valid ID. Use '-' for the `data_product_id` to skip specifying the data
// product ID explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDataProductDraftWithContext(context.Background(), deleteDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDataProductDraftWithContext is an alternate form of the DeleteDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDataProductDraftWithContext(ctx context.Context, deleteDataProductDraftOptions *DeleteDataProductDraftOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataProductDraftOptions, "deleteDataProductDraftOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDataProductDraftOptions, "deleteDataProductDraftOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *deleteDataProductDraftOptions.DataProductID,
		"draft_id": *deleteDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_data_product_draft", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateDataProductDraft : Update the data product draft identified by ID
// Use this API to update the properties of a data product draft identified by a valid ID. Use '-' for the
// `data_product_id` to skip specifying the data product ID explicitly.<br/><br/>Specify patch operations using
// http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the properties of a data
// product<br/><br/>- Add/Remove parts from a data product (up to 20 parts)<br/><br/>- Add/Remove use cases from a data
// product<br/><br/>- Update the data product state<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductDraftWithContext(context.Background(), updateDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductDraftWithContext is an alternate form of the UpdateDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDraftWithContext(ctx context.Context, updateDataProductDraftOptions *UpdateDataProductDraftOptions) (result *DataProductDraft, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductDraftOptions, "updateDataProductDraftOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDataProductDraftOptions, "updateDataProductDraftOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDataProductDraftOptions.DataProductID,
		"draft_id": *updateDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductDraftOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_draft", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDraft)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
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
// upload the document file and complete it. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDraftContractTermsDocumentWithContext(context.Background(), getDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDraftContractTermsDocumentWithContext is an alternate form of the GetDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDraftContractTermsDocumentWithContext(ctx context.Context, getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDraftContractTermsDocumentOptions, "getDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDraftContractTermsDocumentOptions, "getDraftContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDraftContractTermsDocumentOptions.DataProductID,
		"draft_id": *getDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *getDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id": *getDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_draft_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDraftContractTermsDocument : Delete a contract document
// Delete an existing contract document.
//
// Contract documents can only be deleted for data product versions that are in DRAFT state. Use '-' for the
// `data_product_id` to skip specifying the data product ID explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDraftContractTermsDocumentWithContext(context.Background(), deleteDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDraftContractTermsDocumentWithContext is an alternate form of the DeleteDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDraftContractTermsDocumentWithContext(ctx context.Context, deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDraftContractTermsDocumentOptions, "deleteDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDraftContractTermsDocumentOptions, "deleteDraftContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *deleteDraftContractTermsDocumentOptions.DataProductID,
		"draft_id": *deleteDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *deleteDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id": *deleteDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_draft_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
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
// <br/><br/>Contract terms documents can only be updated if the associated data product version is in DRAFT state. Use
// '-' for the `data_product_id` to skip specifying the data product ID explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDraftContractTermsDocumentWithContext(context.Background(), updateDraftContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDraftContractTermsDocumentWithContext is an alternate form of the UpdateDraftContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDraftContractTermsDocumentWithContext(ctx context.Context, updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDraftContractTermsDocumentOptions, "updateDraftContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDraftContractTermsDocumentOptions, "updateDraftContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDraftContractTermsDocumentOptions.DataProductID,
		"draft_id": *updateDraftContractTermsDocumentOptions.DraftID,
		"contract_terms_id": *updateDraftContractTermsDocumentOptions.ContractTermsID,
		"document_id": *updateDraftContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDraftContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDraftContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDraftContractTermsDocumentOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_draft_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductDraftContractTerms : Retrieve a data product contract terms identified by id
// Retrieve a data product contract terms identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductDraftContractTerms(getDataProductDraftContractTermsOptions *GetDataProductDraftContractTermsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductDraftContractTermsWithContext(context.Background(), getDataProductDraftContractTermsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductDraftContractTermsWithContext is an alternate form of the GetDataProductDraftContractTerms method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductDraftContractTermsWithContext(ctx context.Context, getDataProductDraftContractTermsOptions *GetDataProductDraftContractTermsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductDraftContractTermsOptions, "getDataProductDraftContractTermsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDataProductDraftContractTermsOptions, "getDataProductDraftContractTermsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductDraftContractTermsOptions.DataProductID,
		"draft_id": *getDataProductDraftContractTermsOptions.DraftID,
		"contract_terms_id": *getDataProductDraftContractTermsOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductDraftContractTerms")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDataProductDraftContractTermsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/odcs+yaml")
	if getDataProductDraftContractTermsOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getDataProductDraftContractTermsOptions.Accept))
	}

	if getDataProductDraftContractTermsOptions.IncludeContractDocuments != nil {
		builder.AddQuery("include_contract_documents", fmt.Sprint(*getDataProductDraftContractTermsOptions.IncludeContractDocuments))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_draft_contract_terms", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ReplaceDataProductDraftContractTerms : Update a data product contract terms identified by id
// Update a data product contract terms identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ReplaceDataProductDraftContractTerms(replaceDataProductDraftContractTermsOptions *ReplaceDataProductDraftContractTermsOptions) (result *ContractTerms, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ReplaceDataProductDraftContractTermsWithContext(context.Background(), replaceDataProductDraftContractTermsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceDataProductDraftContractTermsWithContext is an alternate form of the ReplaceDataProductDraftContractTerms method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ReplaceDataProductDraftContractTermsWithContext(ctx context.Context, replaceDataProductDraftContractTermsOptions *ReplaceDataProductDraftContractTermsOptions) (result *ContractTerms, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceDataProductDraftContractTermsOptions, "replaceDataProductDraftContractTermsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceDataProductDraftContractTermsOptions, "replaceDataProductDraftContractTermsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *replaceDataProductDraftContractTermsOptions.DataProductID,
		"draft_id": *replaceDataProductDraftContractTermsOptions.DraftID,
		"contract_terms_id": *replaceDataProductDraftContractTermsOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ReplaceDataProductDraftContractTerms")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceDataProductDraftContractTermsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceDataProductDraftContractTermsOptions.Asset != nil {
		body["asset"] = replaceDataProductDraftContractTermsOptions.Asset
	}
	if replaceDataProductDraftContractTermsOptions.ID != nil {
		body["id"] = replaceDataProductDraftContractTermsOptions.ID
	}
	if replaceDataProductDraftContractTermsOptions.Documents != nil {
		body["documents"] = replaceDataProductDraftContractTermsOptions.Documents
	}
	if replaceDataProductDraftContractTermsOptions.ErrorMsg != nil {
		body["error_msg"] = replaceDataProductDraftContractTermsOptions.ErrorMsg
	}
	if replaceDataProductDraftContractTermsOptions.Overview != nil {
		body["overview"] = replaceDataProductDraftContractTermsOptions.Overview
	}
	if replaceDataProductDraftContractTermsOptions.Description != nil {
		body["description"] = replaceDataProductDraftContractTermsOptions.Description
	}
	if replaceDataProductDraftContractTermsOptions.Organization != nil {
		body["organization"] = replaceDataProductDraftContractTermsOptions.Organization
	}
	if replaceDataProductDraftContractTermsOptions.Roles != nil {
		body["roles"] = replaceDataProductDraftContractTermsOptions.Roles
	}
	if replaceDataProductDraftContractTermsOptions.Price != nil {
		body["price"] = replaceDataProductDraftContractTermsOptions.Price
	}
	if replaceDataProductDraftContractTermsOptions.SLA != nil {
		body["sla"] = replaceDataProductDraftContractTermsOptions.SLA
	}
	if replaceDataProductDraftContractTermsOptions.SupportAndCommunication != nil {
		body["support_and_communication"] = replaceDataProductDraftContractTermsOptions.SupportAndCommunication
	}
	if replaceDataProductDraftContractTermsOptions.CustomProperties != nil {
		body["custom_properties"] = replaceDataProductDraftContractTermsOptions.CustomProperties
	}
	if replaceDataProductDraftContractTermsOptions.ContractTest != nil {
		body["contract_test"] = replaceDataProductDraftContractTermsOptions.ContractTest
	}
	if replaceDataProductDraftContractTermsOptions.Schema != nil {
		body["schema"] = replaceDataProductDraftContractTermsOptions.Schema
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "replace_data_product_draft_contract_terms", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTerms)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateDataProductDraftContractTerms : Update a contract terms property
// Use this API to update the properties of a contract terms that is identified by a valid ID.
//
// Specify patch operations using http://jsonpatch.com/ syntax.
//
// Supported patch operations include:
// - Update the contract terms properties.
// <br/><br/>Contract terms can only be updated if the associated data product version is in DRAFT state.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDraftContractTerms(updateDataProductDraftContractTermsOptions *UpdateDataProductDraftContractTermsOptions) (result *ContractTerms, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductDraftContractTermsWithContext(context.Background(), updateDataProductDraftContractTermsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductDraftContractTermsWithContext is an alternate form of the UpdateDataProductDraftContractTerms method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDraftContractTermsWithContext(ctx context.Context, updateDataProductDraftContractTermsOptions *UpdateDataProductDraftContractTermsOptions) (result *ContractTerms, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductDraftContractTermsOptions, "updateDataProductDraftContractTermsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDataProductDraftContractTermsOptions, "updateDataProductDraftContractTermsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDataProductDraftContractTermsOptions.DataProductID,
		"draft_id": *updateDataProductDraftContractTermsOptions.DraftID,
		"contract_terms_id": *updateDataProductDraftContractTermsOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/contract_terms/{contract_terms_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductDraftContractTerms")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDataProductDraftContractTermsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductDraftContractTermsOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_draft_contract_terms", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTerms)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// PublishDataProductDraft : Publish a draft of an existing data product
// Publish a draft of an existing data product. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.PublishDataProductDraftWithContext(context.Background(), publishDataProductDraftOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// PublishDataProductDraftWithContext is an alternate form of the PublishDataProductDraft method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) PublishDataProductDraftWithContext(ctx context.Context, publishDataProductDraftOptions *PublishDataProductDraftOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(publishDataProductDraftOptions, "publishDataProductDraftOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(publishDataProductDraftOptions, "publishDataProductDraftOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *publishDataProductDraftOptions.DataProductID,
		"draft_id": *publishDataProductDraftOptions.DraftID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/drafts/{draft_id}/publish`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "PublishDataProductDraft")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range publishDataProductDraftOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "publish_data_product_draft", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductRelease)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductRelease : Get a release of an existing data product
// Get a release of an existing data product. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductReleaseWithContext(context.Background(), getDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductReleaseWithContext is an alternate form of the GetDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductReleaseWithContext(ctx context.Context, getDataProductReleaseOptions *GetDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductReleaseOptions, "getDataProductReleaseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDataProductReleaseOptions, "getDataProductReleaseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getDataProductReleaseOptions.DataProductID,
		"release_id": *getDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getDataProductReleaseOptions.CheckCallerApproval != nil {
		builder.AddQuery("check_caller_approval", fmt.Sprint(*getDataProductReleaseOptions.CheckCallerApproval))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_release", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductRelease)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateDataProductRelease : Update the data product release identified by ID
// Use this API to update the properties of a data product release identified by a valid ID. Use '-' for the
// `data_product_id` to skip specifying the data product ID explicitly.<br/><br/>Specify patch operations using
// http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the properties of a data
// product<br/><br/>- Add/remove parts from a data product (up to 20 parts)<br/><br/>- Add/remove use cases from a data
// product<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductReleaseWithContext(context.Background(), updateDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductReleaseWithContext is an alternate form of the UpdateDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductReleaseWithContext(ctx context.Context, updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductReleaseOptions, "updateDataProductReleaseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDataProductReleaseOptions, "updateDataProductReleaseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *updateDataProductReleaseOptions.DataProductID,
		"release_id": *updateDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateDataProductReleaseOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_release", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductRelease)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
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
// the document file to complete the attachment. Use '-' for the `data_product_id` to skip specifying the data product
// ID explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetReleaseContractTermsDocumentWithContext(context.Background(), getReleaseContractTermsDocumentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetReleaseContractTermsDocumentWithContext is an alternate form of the GetReleaseContractTermsDocument method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetReleaseContractTermsDocumentWithContext(ctx context.Context, getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) (result *ContractTermsDocument, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReleaseContractTermsDocumentOptions, "getReleaseContractTermsDocumentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getReleaseContractTermsDocumentOptions, "getReleaseContractTermsDocumentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getReleaseContractTermsDocumentOptions.DataProductID,
		"release_id": *getReleaseContractTermsDocumentOptions.ReleaseID,
		"contract_terms_id": *getReleaseContractTermsDocumentOptions.ContractTermsID,
		"document_id": *getReleaseContractTermsDocumentOptions.DocumentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}/contract_terms/{contract_terms_id}/documents/{document_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetReleaseContractTermsDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getReleaseContractTermsDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_release_contract_terms_document", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContractTermsDocument)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPublishedDataProductDraftContractTerms : Retrieve a published data product contract terms identified by id
// Retrieve a published data product contract terms identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetPublishedDataProductDraftContractTerms(getPublishedDataProductDraftContractTermsOptions *GetPublishedDataProductDraftContractTermsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetPublishedDataProductDraftContractTermsWithContext(context.Background(), getPublishedDataProductDraftContractTermsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPublishedDataProductDraftContractTermsWithContext is an alternate form of the GetPublishedDataProductDraftContractTerms method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetPublishedDataProductDraftContractTermsWithContext(ctx context.Context, getPublishedDataProductDraftContractTermsOptions *GetPublishedDataProductDraftContractTermsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPublishedDataProductDraftContractTermsOptions, "getPublishedDataProductDraftContractTermsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPublishedDataProductDraftContractTermsOptions, "getPublishedDataProductDraftContractTermsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *getPublishedDataProductDraftContractTermsOptions.DataProductID,
		"release_id": *getPublishedDataProductDraftContractTermsOptions.ReleaseID,
		"contract_terms_id": *getPublishedDataProductDraftContractTermsOptions.ContractTermsID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}/contract_terms/{contract_terms_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetPublishedDataProductDraftContractTerms")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getPublishedDataProductDraftContractTermsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/odcs+yaml")
	if getPublishedDataProductDraftContractTermsOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getPublishedDataProductDraftContractTermsOptions.Accept))
	}

	if getPublishedDataProductDraftContractTermsOptions.IncludeContractDocuments != nil {
		builder.AddQuery("include_contract_documents", fmt.Sprint(*getPublishedDataProductDraftContractTermsOptions.IncludeContractDocuments))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_published_data_product_draft_contract_terms", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListDataProductReleases : Retrieve a list of data product releases
// Retrieve a list of data product releases. Use '-' for the `data_product_id` to skip specifying the data product ID
// explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) (result *DataProductReleaseCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductReleasesWithContext(context.Background(), listDataProductReleasesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductReleasesWithContext is an alternate form of the ListDataProductReleases method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductReleasesWithContext(ctx context.Context, listDataProductReleasesOptions *ListDataProductReleasesOptions) (result *DataProductReleaseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDataProductReleasesOptions, "listDataProductReleasesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listDataProductReleasesOptions, "listDataProductReleasesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductReleases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDataProductReleasesOptions.Headers {
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
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_releases", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductReleaseCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// RetireDataProductRelease : Retire a release of an existing data product
// Retire a release of an existing data product. Use '-' for the `data_product_id` to skip specifying the data product
// ID explicitly.
func (dataProductHubApiService *DataProductHubAPIServiceV1) RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.RetireDataProductReleaseWithContext(context.Background(), retireDataProductReleaseOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// RetireDataProductReleaseWithContext is an alternate form of the RetireDataProductRelease method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) RetireDataProductReleaseWithContext(ctx context.Context, retireDataProductReleaseOptions *RetireDataProductReleaseOptions) (result *DataProductRelease, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(retireDataProductReleaseOptions, "retireDataProductReleaseOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(retireDataProductReleaseOptions, "retireDataProductReleaseOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"data_product_id": *retireDataProductReleaseOptions.DataProductID,
		"release_id": *retireDataProductReleaseOptions.ReleaseID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/data_products/{data_product_id}/releases/{release_id}/retire`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "RetireDataProductRelease")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range retireDataProductReleaseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if retireDataProductReleaseOptions.RevokeAccess != nil {
		builder.AddQuery("revoke_access", fmt.Sprint(*retireDataProductReleaseOptions.RevokeAccess))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "retire_data_product_release", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductRelease)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListDataProductContractTemplate : Retrieve a list of data product contract templates
// Retrieve a list of data product contract templates.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductContractTemplate(listDataProductContractTemplateOptions *ListDataProductContractTemplateOptions) (result *DataProductContractTemplateCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductContractTemplateWithContext(context.Background(), listDataProductContractTemplateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductContractTemplateWithContext is an alternate form of the ListDataProductContractTemplate method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductContractTemplateWithContext(ctx context.Context, listDataProductContractTemplateOptions *ListDataProductContractTemplateOptions) (result *DataProductContractTemplateCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductContractTemplateOptions, "listDataProductContractTemplateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/contract_templates`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductContractTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDataProductContractTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductContractTemplateOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*listDataProductContractTemplateOptions.ContainerID))
	}
	if listDataProductContractTemplateOptions.ContractTemplateName != nil {
		builder.AddQuery("contract_template.name", fmt.Sprint(*listDataProductContractTemplateOptions.ContractTemplateName))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_contract_template", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductContractTemplateCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateContractTemplate : Create new data product contract template
// Create new data product contract template.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateContractTemplate(createContractTemplateOptions *CreateContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateContractTemplateWithContext(context.Background(), createContractTemplateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateContractTemplateWithContext is an alternate form of the CreateContractTemplate method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateContractTemplateWithContext(ctx context.Context, createContractTemplateOptions *CreateContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createContractTemplateOptions, "createContractTemplateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createContractTemplateOptions, "createContractTemplateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/contract_templates`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateContractTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createContractTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createContractTemplateOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*createContractTemplateOptions.ContainerID))
	}
	if createContractTemplateOptions.ContractTemplateName != nil {
		builder.AddQuery("contract_template.name", fmt.Sprint(*createContractTemplateOptions.ContractTemplateName))
	}

	body := make(map[string]interface{})
	if createContractTemplateOptions.Container != nil {
		body["container"] = createContractTemplateOptions.Container
	}
	if createContractTemplateOptions.ID != nil {
		body["id"] = createContractTemplateOptions.ID
	}
	if createContractTemplateOptions.Name != nil {
		body["name"] = createContractTemplateOptions.Name
	}
	if createContractTemplateOptions.Error != nil {
		body["error"] = createContractTemplateOptions.Error
	}
	if createContractTemplateOptions.ContractTerms != nil {
		body["contract_terms"] = createContractTemplateOptions.ContractTerms
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_contract_template", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductContractTemplate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetContractTemplate : Retrieve a data product contract template identified by id
// Retrieve a data product contract template identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetContractTemplate(getContractTemplateOptions *GetContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetContractTemplateWithContext(context.Background(), getContractTemplateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetContractTemplateWithContext is an alternate form of the GetContractTemplate method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetContractTemplateWithContext(ctx context.Context, getContractTemplateOptions *GetContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getContractTemplateOptions, "getContractTemplateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getContractTemplateOptions, "getContractTemplateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"contract_template_id": *getContractTemplateOptions.ContractTemplateID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/contract_templates/{contract_template_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetContractTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getContractTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("container.id", fmt.Sprint(*getContractTemplateOptions.ContainerID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_contract_template", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductContractTemplate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDataProductContractTemplate : Delete a data product contract template identified by id
// Delete a data product contract template identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDataProductContractTemplate(deleteDataProductContractTemplateOptions *DeleteDataProductContractTemplateOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDataProductContractTemplateWithContext(context.Background(), deleteDataProductContractTemplateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDataProductContractTemplateWithContext is an alternate form of the DeleteDataProductContractTemplate method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDataProductContractTemplateWithContext(ctx context.Context, deleteDataProductContractTemplateOptions *DeleteDataProductContractTemplateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDataProductContractTemplateOptions, "deleteDataProductContractTemplateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDataProductContractTemplateOptions, "deleteDataProductContractTemplateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"contract_template_id": *deleteDataProductContractTemplateOptions.ContractTemplateID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/contract_templates/{contract_template_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDataProductContractTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDataProductContractTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("container.id", fmt.Sprint(*deleteDataProductContractTemplateOptions.ContainerID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_data_product_contract_template", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateDataProductContractTemplate : Update the data product contract template identified by ID
// Use this API to update the properties of a data product contract template identified by a valid ID.<br/><br/>Specify
// patch operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update
// the name of a data product contract template<br/><br/>- Update the contract terms of data product contract
// template<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductContractTemplate(updateDataProductContractTemplateOptions *UpdateDataProductContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductContractTemplateWithContext(context.Background(), updateDataProductContractTemplateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductContractTemplateWithContext is an alternate form of the UpdateDataProductContractTemplate method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductContractTemplateWithContext(ctx context.Context, updateDataProductContractTemplateOptions *UpdateDataProductContractTemplateOptions) (result *DataProductContractTemplate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductContractTemplateOptions, "updateDataProductContractTemplateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDataProductContractTemplateOptions, "updateDataProductContractTemplateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"contract_template_id": *updateDataProductContractTemplateOptions.ContractTemplateID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/contract_templates/{contract_template_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductContractTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDataProductContractTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	builder.AddQuery("container.id", fmt.Sprint(*updateDataProductContractTemplateOptions.ContainerID))

	_, err = builder.SetBodyContentJSON(updateDataProductContractTemplateOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_contract_template", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductContractTemplate)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListDataProductDomains : Retrieve a list of data product domains
// Retrieve a list of data product domains.
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductDomains(listDataProductDomainsOptions *ListDataProductDomainsOptions) (result *DataProductDomainCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.ListDataProductDomainsWithContext(context.Background(), listDataProductDomainsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDataProductDomainsWithContext is an alternate form of the ListDataProductDomains method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) ListDataProductDomainsWithContext(ctx context.Context, listDataProductDomainsOptions *ListDataProductDomainsOptions) (result *DataProductDomainCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDataProductDomainsOptions, "listDataProductDomainsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "ListDataProductDomains")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDataProductDomainsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listDataProductDomainsOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*listDataProductDomainsOptions.ContainerID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_data_product_domains", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDomainCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProductDomain : Create new data product domain
// Create new data product domain.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductDomain(createDataProductDomainOptions *CreateDataProductDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductDomainWithContext(context.Background(), createDataProductDomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductDomainWithContext is an alternate form of the CreateDataProductDomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductDomainWithContext(ctx context.Context, createDataProductDomainOptions *CreateDataProductDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductDomainOptions, "createDataProductDomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDataProductDomainOptions, "createDataProductDomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProductDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDataProductDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createDataProductDomainOptions.ContainerID != nil {
		builder.AddQuery("container.id", fmt.Sprint(*createDataProductDomainOptions.ContainerID))
	}

	body := make(map[string]interface{})
	if createDataProductDomainOptions.Container != nil {
		body["container"] = createDataProductDomainOptions.Container
	}
	if createDataProductDomainOptions.Trace != nil {
		body["trace"] = createDataProductDomainOptions.Trace
	}
	if createDataProductDomainOptions.Errors != nil {
		body["errors"] = createDataProductDomainOptions.Errors
	}
	if createDataProductDomainOptions.Name != nil {
		body["name"] = createDataProductDomainOptions.Name
	}
	if createDataProductDomainOptions.Description != nil {
		body["description"] = createDataProductDomainOptions.Description
	}
	if createDataProductDomainOptions.ID != nil {
		body["id"] = createDataProductDomainOptions.ID
	}
	if createDataProductDomainOptions.MemberRoles != nil {
		body["member_roles"] = createDataProductDomainOptions.MemberRoles
	}
	if createDataProductDomainOptions.Properties != nil {
		body["properties"] = createDataProductDomainOptions.Properties
	}
	if createDataProductDomainOptions.SubDomains != nil {
		body["sub_domains"] = createDataProductDomainOptions.SubDomains
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product_domain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDomain)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDataProductSubdomain : Create data product subdomains for a specific domain identified by id
// Create data product subdomains for a specific domain identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductSubdomain(createDataProductSubdomainOptions *CreateDataProductSubdomainOptions) (result *InitializeSubDomain, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateDataProductSubdomainWithContext(context.Background(), createDataProductSubdomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDataProductSubdomainWithContext is an alternate form of the CreateDataProductSubdomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateDataProductSubdomainWithContext(ctx context.Context, createDataProductSubdomainOptions *CreateDataProductSubdomainOptions) (result *InitializeSubDomain, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDataProductSubdomainOptions, "createDataProductSubdomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDataProductSubdomainOptions, "createDataProductSubdomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"domain_id": *createDataProductSubdomainOptions.DomainID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains/{domain_id}/subdomains`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateDataProductSubdomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDataProductSubdomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("container.id", fmt.Sprint(*createDataProductSubdomainOptions.ContainerID))

	body := make(map[string]interface{})
	if createDataProductSubdomainOptions.Name != nil {
		body["name"] = createDataProductSubdomainOptions.Name
	}
	if createDataProductSubdomainOptions.ID != nil {
		body["id"] = createDataProductSubdomainOptions.ID
	}
	if createDataProductSubdomainOptions.Description != nil {
		body["description"] = createDataProductSubdomainOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_data_product_subdomain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInitializeSubDomain)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDomain : Retrieve a data product domain or subdomain identified by id
// Retrieve a data product domain or subdomain identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDomain(getDomainOptions *GetDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDomainWithContext(context.Background(), getDomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDomainWithContext is an alternate form of the GetDomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDomainWithContext(ctx context.Context, getDomainOptions *GetDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDomainOptions, "getDomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDomainOptions, "getDomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"domain_id": *getDomainOptions.DomainID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains/{domain_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_domain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDomain)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDomain : Delete a data product domain identified by id
// Delete a data product domain identified by id.
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDomain(deleteDomainOptions *DeleteDomainOptions) (response *core.DetailedResponse, err error) {
	response, err = dataProductHubApiService.DeleteDomainWithContext(context.Background(), deleteDomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDomainWithContext is an alternate form of the DeleteDomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) DeleteDomainWithContext(ctx context.Context, deleteDomainOptions *DeleteDomainOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDomainOptions, "deleteDomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDomainOptions, "deleteDomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"domain_id": *deleteDomainOptions.DomainID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains/{domain_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "DeleteDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = dataProductHubApiService.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_domain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateDataProductDomain : Update the data product domain identified by ID
// Use this API to update the properties of a data product domain identified by a valid ID.<br/><br/>Specify patch
// operations using http://jsonpatch.com/ syntax.<br/><br/>Supported patch operations include:<br/><br/>- Update the
// name of a data product domain<br/><br/>- Update the description of a data product domain<br/><br/>- Update the rov of
// a data product domain<br/><br/>.
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDomain(updateDataProductDomainOptions *UpdateDataProductDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.UpdateDataProductDomainWithContext(context.Background(), updateDataProductDomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateDataProductDomainWithContext is an alternate form of the UpdateDataProductDomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) UpdateDataProductDomainWithContext(ctx context.Context, updateDataProductDomainOptions *UpdateDataProductDomainOptions) (result *DataProductDomain, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDataProductDomainOptions, "updateDataProductDomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateDataProductDomainOptions, "updateDataProductDomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"domain_id": *updateDataProductDomainOptions.DomainID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains/{domain_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "UpdateDataProductDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range updateDataProductDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	builder.AddQuery("container.id", fmt.Sprint(*updateDataProductDomainOptions.ContainerID))

	_, err = builder.SetBodyContentJSON(updateDataProductDomainOptions.JSONPatchInstructions)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_data_product_domain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductDomain)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDataProductByDomain : Retrieve all data products in a domain specified by id or any of it's subdomains
// Retrieve all the data products tagged to the domain identified by id or any of it's subdomains.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductByDomain(getDataProductByDomainOptions *GetDataProductByDomainOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetDataProductByDomainWithContext(context.Background(), getDataProductByDomainOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDataProductByDomainWithContext is an alternate form of the GetDataProductByDomain method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetDataProductByDomainWithContext(ctx context.Context, getDataProductByDomainOptions *GetDataProductByDomainOptions) (result *DataProductVersionCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDataProductByDomainOptions, "getDataProductByDomainOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDataProductByDomainOptions, "getDataProductByDomainOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"domain_id": *getDataProductByDomainOptions.DomainID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/domains/{domain_id}/data_products`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetDataProductByDomain")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDataProductByDomainOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("container.id", fmt.Sprint(*getDataProductByDomainOptions.ContainerID))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_data_product_by_domain", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataProductVersionCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateS3Bucket : Create a new Bucket
// Use this API to create a new S3 Bucket on an AWS Hosting.
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateS3Bucket(createS3BucketOptions *CreateS3BucketOptions) (result *BucketResponse, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.CreateS3BucketWithContext(context.Background(), createS3BucketOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateS3BucketWithContext is an alternate form of the CreateS3Bucket method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) CreateS3BucketWithContext(ctx context.Context, createS3BucketOptions *CreateS3BucketOptions) (result *BucketResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createS3BucketOptions, "createS3BucketOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createS3BucketOptions, "createS3BucketOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/bucket`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "CreateS3Bucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createS3BucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("is_shared", fmt.Sprint(*createS3BucketOptions.IsShared))

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_s3_bucket", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetS3BucketValidation : Validate the Bucket Existence
// Use this API to validate the bucket existence on an AWS hosting.
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetS3BucketValidation(getS3BucketValidationOptions *GetS3BucketValidationOptions) (result *BucketValidationResponse, response *core.DetailedResponse, err error) {
	result, response, err = dataProductHubApiService.GetS3BucketValidationWithContext(context.Background(), getS3BucketValidationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetS3BucketValidationWithContext is an alternate form of the GetS3BucketValidation method which supports a Context parameter
func (dataProductHubApiService *DataProductHubAPIServiceV1) GetS3BucketValidationWithContext(ctx context.Context, getS3BucketValidationOptions *GetS3BucketValidationOptions) (result *BucketValidationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getS3BucketValidationOptions, "getS3BucketValidationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getS3BucketValidationOptions, "getS3BucketValidationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"bucket_name": *getS3BucketValidationOptions.BucketName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dataProductHubApiService.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dataProductHubApiService.Service.Options.URL, `/data_product_exchange/v1/bucket/validate/{bucket_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("data_product_hub_api_service", "V1", "GetS3BucketValidation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getS3BucketValidationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dataProductHubApiService.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_s3_bucket_validation", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketValidationResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1")
}

// AssetListAccessControl : Access control object.
type AssetListAccessControl struct {
	// The owner of the asset.
	Owner *string `json:"owner,omitempty"`
}

// UnmarshalAssetListAccessControl unmarshals an instance of AssetListAccessControl from the specified map of raw messages.
func UnmarshalAssetListAccessControl(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetListAccessControl)
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		err = core.SDKErrorf(err, "", "owner-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetPartReference : The asset represented in this part.
type AssetPartReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	// Asset name.
	Name *string `json:"name,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The type of the asset.
	Type *string `json:"type,omitempty"`
}

// NewAssetPartReference : Instantiate AssetPartReference (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewAssetPartReference(container *ContainerReference) (_model *AssetPartReference, err error) {
	_model = &AssetPartReference{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalAssetPartReference unmarshals an instance of AssetPartReference from the specified map of raw messages.
func UnmarshalAssetPartReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPartReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetPrototype : New asset input properties.
type AssetPrototype struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	// The identity schema for a IBM knowledge catalog container (catalog/project/space).
	Container *ContainerIdentity `json:"container" validate:"required"`
}

// NewAssetPrototype : Instantiate AssetPrototype (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewAssetPrototype(container *ContainerIdentity) (_model *AssetPrototype, err error) {
	_model = &AssetPrototype{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalAssetPrototype unmarshals an instance of AssetPrototype from the specified map of raw messages.
func UnmarshalAssetPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPrototype)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerIdentity)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetReference : The reference schema for a asset in a container.
type AssetReference struct {
	// The unique identifier of the asset.
	ID *string `json:"id,omitempty"`

	// Asset name.
	Name *string `json:"name,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// NewAssetReference : Instantiate AssetReference (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewAssetReference(container *ContainerReference) (_model *AssetReference, err error) {
	_model = &AssetReference{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalAssetReference unmarshals an instance of AssetReference from the specified map of raw messages.
func UnmarshalAssetReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketResponse : BucketResponse to hold the Bucket response data.
type BucketResponse struct {
	// Name of the Bucket.
	BucketName *string `json:"bucket_name,omitempty"`

	// Location of the Bucket stored.
	BucketLocation *string `json:"bucket_location,omitempty"`

	// Role ARN.
	RoleArn *string `json:"role_arn,omitempty"`

	// Bucket Type.
	BucketType *string `json:"bucket_type,omitempty"`

	// Is Shared Bucket.
	Shared *bool `json:"shared,omitempty"`
}

// UnmarshalBucketResponse unmarshals an instance of BucketResponse from the specified map of raw messages.
func UnmarshalBucketResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketResponse)
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_location", &obj.BucketLocation)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_location-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role_arn", &obj.RoleArn)
	if err != nil {
		err = core.SDKErrorf(err, "", "role_arn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_type", &obj.BucketType)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "shared", &obj.Shared)
	if err != nil {
		err = core.SDKErrorf(err, "", "shared-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketValidationResponse : BucketValidationResponse to hold the bucket validation data.
type BucketValidationResponse struct {
	// Flag of bucket existence.
	BucketExists *bool `json:"bucket_exists,omitempty"`
}

// UnmarshalBucketValidationResponse unmarshals an instance of BucketValidationResponse from the specified map of raw messages.
func UnmarshalBucketValidationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketValidationResponse)
	err = core.UnmarshalPrimitive(m, "bucket_exists", &obj.BucketExists)
	if err != nil {
		err = core.SDKErrorf(err, "", "bucket_exists-error", common.GetComponentInfo())
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCompleteDraftContractTermsDocumentOptions : Instantiate CompleteDraftContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewCompleteDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *CompleteDraftContractTermsDocumentOptions {
	return &CompleteDraftContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
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

// ContainerIdentity : The identity schema for a IBM knowledge catalog container (catalog/project/space).
type ContainerIdentity struct {
	// Container identifier.
	ID *string `json:"id" validate:"required"`
}

// NewContainerIdentity : Instantiate ContainerIdentity (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContainerIdentity(id string) (_model *ContainerIdentity, err error) {
	_model = &ContainerIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContainerIdentity unmarshals an instance of ContainerIdentity from the specified map of raw messages.
func UnmarshalContainerIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
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
	ContainerReferenceTypeCatalogConst = "catalog"
	ContainerReferenceTypeProjectConst = "project"
)

// NewContainerReference : Instantiate ContainerReference (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContainerReference(id string, typeVar string) (_model *ContainerReference, err error) {
	_model = &ContainerReference{
		ID: core.StringPtr(id),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContainerReference unmarshals an instance of ContainerReference from the specified map of raw messages.
func UnmarshalContainerReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractSchema : Schema definition of the data asset.
type ContractSchema struct {
	// Name of the schema or data asset part.
	Name *string `json:"name,omitempty"`

	// Description of the schema.
	Description *string `json:"description,omitempty"`

	// MIME type or physical type.
	PhysicalType *string `json:"physical_type,omitempty"`

	// List of properties.
	Properties []ContractSchemaProperty `json:"properties,omitempty"`
}

// UnmarshalContractSchema unmarshals an instance of ContractSchema from the specified map of raw messages.
func UnmarshalContractSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractSchema)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "physical_type", &obj.PhysicalType)
	if err != nil {
		err = core.SDKErrorf(err, "", "physical_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalContractSchemaProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "properties-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractSchemaProperty : Defines a property inside the schema.
type ContractSchemaProperty struct {
	// Property name.
	Name *string `json:"name" validate:"required"`

	// Detailed type definition of a schema property.
	Type *ContractSchemaPropertyType `json:"type,omitempty"`
}

// NewContractSchemaProperty : Instantiate ContractSchemaProperty (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractSchemaProperty(name string) (_model *ContractSchemaProperty, err error) {
	_model = &ContractSchemaProperty{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractSchemaProperty unmarshals an instance of ContractSchemaProperty from the specified map of raw messages.
func UnmarshalContractSchemaProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractSchemaProperty)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "type", &obj.Type, UnmarshalContractSchemaPropertyType)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractSchemaPropertyType : Detailed type definition of a schema property.
type ContractSchemaPropertyType struct {
	// Type of the field.
	Type *string `json:"type,omitempty"`

	// Length of the field as string.
	Length *string `json:"length,omitempty"`

	// Scale of the field as string.
	Scale *string `json:"scale,omitempty"`

	// Is field nullable? true/false as string.
	Nullable *string `json:"nullable,omitempty"`

	// Is field signed? true/false as string.
	Signed *string `json:"signed,omitempty"`

	// Native type of the field.
	NativeType *string `json:"native_type,omitempty"`
}

// UnmarshalContractSchemaPropertyType unmarshals an instance of ContractSchemaPropertyType from the specified map of raw messages.
func UnmarshalContractSchemaPropertyType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractSchemaPropertyType)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "length", &obj.Length)
	if err != nil {
		err = core.SDKErrorf(err, "", "length-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "scale", &obj.Scale)
	if err != nil {
		err = core.SDKErrorf(err, "", "scale-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "nullable", &obj.Nullable)
	if err != nil {
		err = core.SDKErrorf(err, "", "nullable-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "signed", &obj.Signed)
	if err != nil {
		err = core.SDKErrorf(err, "", "signed-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "native_type", &obj.NativeType)
	if err != nil {
		err = core.SDKErrorf(err, "", "native_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTemplateCustomProperty : Represents a custom property within the contract.
type ContractTemplateCustomProperty struct {
	// The name of the key. Names should be in camel case–the same as if they were permanent properties in the contract.
	Key *string `json:"key" validate:"required"`

	// The value of the key.
	Value *string `json:"value" validate:"required"`
}

// NewContractTemplateCustomProperty : Instantiate ContractTemplateCustomProperty (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTemplateCustomProperty(key string, value string) (_model *ContractTemplateCustomProperty, err error) {
	_model = &ContractTemplateCustomProperty{
		Key: core.StringPtr(key),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTemplateCustomProperty unmarshals an instance of ContractTemplateCustomProperty from the specified map of raw messages.
func UnmarshalContractTemplateCustomProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTemplateCustomProperty)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		err = core.SDKErrorf(err, "", "key-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTemplateOrganization : Represents a user and their role in the contract.
type ContractTemplateOrganization struct {
	// The user ID associated with the contract.
	UserID *string `json:"user_id" validate:"required"`

	// The role of the user in the contract.
	Role *string `json:"role" validate:"required"`
}

// NewContractTemplateOrganization : Instantiate ContractTemplateOrganization (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTemplateOrganization(userID string, role string) (_model *ContractTemplateOrganization, err error) {
	_model = &ContractTemplateOrganization{
		UserID: core.StringPtr(userID),
		Role: core.StringPtr(role),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTemplateOrganization unmarshals an instance of ContractTemplateOrganization from the specified map of raw messages.
func UnmarshalContractTemplateOrganization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTemplateOrganization)
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		err = core.SDKErrorf(err, "", "user_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTemplateSLA : Represents the SLA details of the contract.
type ContractTemplateSLA struct {
	// The default SLA element.
	DefaultElement *string `json:"default_element,omitempty"`

	// List of SLA properties and their values.
	Properties []ContractTemplateSLAProperty `json:"properties,omitempty"`
}

// UnmarshalContractTemplateSLA unmarshals an instance of ContractTemplateSLA from the specified map of raw messages.
func UnmarshalContractTemplateSLA(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTemplateSLA)
	err = core.UnmarshalPrimitive(m, "default_element", &obj.DefaultElement)
	if err != nil {
		err = core.SDKErrorf(err, "", "default_element-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalContractTemplateSLAProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "properties-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTemplateSLAProperty : Represents an SLA property and its value.
type ContractTemplateSLAProperty struct {
	// The SLA property name.
	Property *string `json:"property" validate:"required"`

	// The value associated with the SLA property.
	Value *string `json:"value" validate:"required"`
}

// NewContractTemplateSLAProperty : Instantiate ContractTemplateSLAProperty (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTemplateSLAProperty(property string, value string) (_model *ContractTemplateSLAProperty, err error) {
	_model = &ContractTemplateSLAProperty{
		Property: core.StringPtr(property),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTemplateSLAProperty unmarshals an instance of ContractTemplateSLAProperty from the specified map of raw messages.
func UnmarshalContractTemplateSLAProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTemplateSLAProperty)
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
	if err != nil {
		err = core.SDKErrorf(err, "", "property-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTemplateSupportAndCommunication : Represents a support and communication channel for the contract.
type ContractTemplateSupportAndCommunication struct {
	// The communication channel.
	Channel *string `json:"channel" validate:"required"`

	// The URL associated with the communication channel.
	URL *string `json:"url" validate:"required"`
}

// NewContractTemplateSupportAndCommunication : Instantiate ContractTemplateSupportAndCommunication (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTemplateSupportAndCommunication(channel string, url string) (_model *ContractTemplateSupportAndCommunication, err error) {
	_model = &ContractTemplateSupportAndCommunication{
		Channel: core.StringPtr(channel),
		URL: core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTemplateSupportAndCommunication unmarshals an instance of ContractTemplateSupportAndCommunication from the specified map of raw messages.
func UnmarshalContractTemplateSupportAndCommunication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTemplateSupportAndCommunication)
	err = core.UnmarshalPrimitive(m, "channel", &obj.Channel)
	if err != nil {
		err = core.SDKErrorf(err, "", "channel-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTerms : Defines the complete structure of a contract terms.
type ContractTerms struct {
	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset,omitempty"`

	// ID of the contract terms.
	ID *string `json:"id,omitempty"`

	// Collection of contract terms documents.
	Documents []ContractTermsDocument `json:"documents,omitempty"`

	// An error message, if existing, relating to the contract terms.
	ErrorMsg *string `json:"error_msg,omitempty"`

	// Overview details of a data contract.
	Overview *Overview `json:"overview,omitempty"`

	// Description details of a data contract.
	Description *Description `json:"description,omitempty"`

	// List of sub domains to be added within a domain.
	Organization []ContractTemplateOrganization `json:"organization,omitempty"`

	// List of roles associated with the contract.
	Roles []Roles `json:"roles,omitempty"`

	// Represents the pricing details of the contract.
	Price *Pricing `json:"price,omitempty"`

	// Service Level Agreement details.
	SLA []ContractTemplateSLA `json:"sla,omitempty"`

	// Support and communication details for the contract.
	SupportAndCommunication []ContractTemplateSupportAndCommunication `json:"support_and_communication,omitempty"`

	// Custom properties that are not part of the standard contract.
	CustomProperties []ContractTemplateCustomProperty `json:"custom_properties,omitempty"`

	// Contains the contract test status and related metadata.
	ContractTest *ContractTest `json:"contract_test,omitempty"`

	// Schema details of the data asset.
	Schema []ContractSchema `json:"schema,omitempty"`
}

// UnmarshalContractTerms unmarshals an instance of ContractTerms from the specified map of raw messages.
func UnmarshalContractTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTerms)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalContractTermsDocument)
	if err != nil {
		err = core.SDKErrorf(err, "", "documents-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "error_msg", &obj.ErrorMsg)
	if err != nil {
		err = core.SDKErrorf(err, "", "error_msg-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "overview", &obj.Overview, UnmarshalOverview)
	if err != nil {
		err = core.SDKErrorf(err, "", "overview-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "description", &obj.Description, UnmarshalDescription)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "organization", &obj.Organization, UnmarshalContractTemplateOrganization)
	if err != nil {
		err = core.SDKErrorf(err, "", "organization-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalRoles)
	if err != nil {
		err = core.SDKErrorf(err, "", "roles-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "price", &obj.Price, UnmarshalPricing)
	if err != nil {
		err = core.SDKErrorf(err, "", "price-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "sla", &obj.SLA, UnmarshalContractTemplateSLA)
	if err != nil {
		err = core.SDKErrorf(err, "", "sla-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "support_and_communication", &obj.SupportAndCommunication, UnmarshalContractTemplateSupportAndCommunication)
	if err != nil {
		err = core.SDKErrorf(err, "", "support_and_communication-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "custom_properties", &obj.CustomProperties, UnmarshalContractTemplateCustomProperty)
	if err != nil {
		err = core.SDKErrorf(err, "", "custom_properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_test", &obj.ContractTest, UnmarshalContractTest)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_test-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "schema", &obj.Schema, UnmarshalContractSchema)
	if err != nil {
		err = core.SDKErrorf(err, "", "schema-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewContractTermsPatch(contractTerms *ContractTerms) (_patch []JSONPatchOperation) {
	if (contractTerms.Asset != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/asset"),
			Value: contractTerms.Asset,
		})
	}
	if (contractTerms.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: contractTerms.ID,
		})
	}
	if (contractTerms.Documents != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/documents"),
			Value: contractTerms.Documents,
		})
	}
	if (contractTerms.ErrorMsg != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/error_msg"),
			Value: contractTerms.ErrorMsg,
		})
	}
	if (contractTerms.Overview != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/overview"),
			Value: contractTerms.Overview,
		})
	}
	if (contractTerms.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/description"),
			Value: contractTerms.Description,
		})
	}
	if (contractTerms.Organization != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/organization"),
			Value: contractTerms.Organization,
		})
	}
	if (contractTerms.Roles != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/roles"),
			Value: contractTerms.Roles,
		})
	}
	if (contractTerms.Price != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/price"),
			Value: contractTerms.Price,
		})
	}
	if (contractTerms.SLA != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/sla"),
			Value: contractTerms.SLA,
		})
	}
	if (contractTerms.SupportAndCommunication != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/support_and_communication"),
			Value: contractTerms.SupportAndCommunication,
		})
	}
	if (contractTerms.CustomProperties != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/custom_properties"),
			Value: contractTerms.CustomProperties,
		})
	}
	if (contractTerms.ContractTest != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/contract_test"),
			Value: contractTerms.ContractTest,
		})
	}
	if (contractTerms.Schema != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/schema"),
			Value: contractTerms.Schema,
		})
	}
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
	ContractTermsDocumentTypeSLAConst = "sla"
	ContractTermsDocumentTypeTermsAndConditionsConst = "terms_and_conditions"
)

// NewContractTermsDocument : Instantiate ContractTermsDocument (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTermsDocument(typeVar string, name string, id string) (_model *ContractTermsDocument, err error) {
	_model = &ContractTermsDocument{
		Type: core.StringPtr(typeVar),
		Name: core.StringPtr(name),
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTermsDocument unmarshals an instance of ContractTermsDocument from the specified map of raw messages.
func UnmarshalContractTermsDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsDocument)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "attachment", &obj.Attachment, UnmarshalContractTermsDocumentAttachment)
	if err != nil {
		err = core.SDKErrorf(err, "", "attachment-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "upload_url", &obj.UploadURL)
	if err != nil {
		err = core.SDKErrorf(err, "", "upload_url-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewContractTermsDocumentPatch(contractTermsDocument *ContractTermsDocument) (_patch []JSONPatchOperation) {
	if (contractTermsDocument.URL != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/url"),
			Value: contractTermsDocument.URL,
		})
	}
	if (contractTermsDocument.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/type"),
			Value: contractTermsDocument.Type,
		})
	}
	if (contractTermsDocument.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/name"),
			Value: contractTermsDocument.Name,
		})
	}
	if (contractTermsDocument.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: contractTermsDocument.ID,
		})
	}
	if (contractTermsDocument.Attachment != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/attachment"),
			Value: contractTermsDocument.Attachment,
		})
	}
	if (contractTermsDocument.UploadURL != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/upload_url"),
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
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTermsMoreInfo : List of links to sources that provide more details on the dataset.
type ContractTermsMoreInfo struct {
	// Type of Source Link.
	Type *string `json:"type" validate:"required"`

	// Link to source that provide more details on the dataset.
	URL *string `json:"url" validate:"required"`
}

// NewContractTermsMoreInfo : Instantiate ContractTermsMoreInfo (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTermsMoreInfo(typeVar string, url string) (_model *ContractTermsMoreInfo, err error) {
	_model = &ContractTermsMoreInfo{
		Type: core.StringPtr(typeVar),
		URL: core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTermsMoreInfo unmarshals an instance of ContractTermsMoreInfo from the specified map of raw messages.
func UnmarshalContractTermsMoreInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTermsMoreInfo)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContractTest : Contains the contract test status and related metadata.
type ContractTest struct {
	// Status of the contract test (pass or fail).
	Status *string `json:"status" validate:"required"`

	// Timestamp of when the contract was last tested.
	LastTestedTime *string `json:"last_tested_time" validate:"required"`

	// Optional message or details about the contract test.
	Message *string `json:"message,omitempty"`
}

// Constants associated with the ContractTest.Status property.
// Status of the contract test (pass or fail).
const (
	ContractTestStatusFailConst = "fail"
	ContractTestStatusPassConst = "pass"
)

// NewContractTest : Instantiate ContractTest (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewContractTest(status string, lastTestedTime string) (_model *ContractTest, err error) {
	_model = &ContractTest{
		Status: core.StringPtr(status),
		LastTestedTime: core.StringPtr(lastTestedTime),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalContractTest unmarshals an instance of ContractTest from the specified map of raw messages.
func UnmarshalContractTest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContractTest)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_tested_time", &obj.LastTestedTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_tested_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateContractTemplateOptions : The CreateContractTemplate options.
type CreateContractTemplateOptions struct {
	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The identifier of the data product contract template.
	ID *string `json:"id,omitempty"`

	// The name of the contract template.
	Name *string `json:"name,omitempty"`

	// Contains the code and details.
	Error *ErrorMessage `json:"error,omitempty"`

	// Defines the complete structure of a contract terms.
	ContractTerms *ContractTerms `json:"contract_terms,omitempty"`

	// Container ID of the data product catalog. If not supplied, the data product catalog is looked up by using the uid of
	// the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Name of the data product contract template. If not supplied, the data product templates within the catalog will
	// returned.
	ContractTemplateName *string `json:"contract_template.name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateContractTemplateOptions : Instantiate CreateContractTemplateOptions
func (*DataProductHubAPIServiceV1) NewCreateContractTemplateOptions(container *ContainerReference) *CreateContractTemplateOptions {
	return &CreateContractTemplateOptions{
		Container: container,
	}
}

// SetContainer : Allow user to set Container
func (_options *CreateContractTemplateOptions) SetContainer(container *ContainerReference) *CreateContractTemplateOptions {
	_options.Container = container
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateContractTemplateOptions) SetID(id string) *CreateContractTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateContractTemplateOptions) SetName(name string) *CreateContractTemplateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetError : Allow user to set Error
func (_options *CreateContractTemplateOptions) SetError(error *ErrorMessage) *CreateContractTemplateOptions {
	_options.Error = error
	return _options
}

// SetContractTerms : Allow user to set ContractTerms
func (_options *CreateContractTemplateOptions) SetContractTerms(contractTerms *ContractTerms) *CreateContractTemplateOptions {
	_options.ContractTerms = contractTerms
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *CreateContractTemplateOptions) SetContainerID(containerID string) *CreateContractTemplateOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetContractTemplateName : Allow user to set ContractTemplateName
func (_options *CreateContractTemplateOptions) SetContractTemplateName(contractTemplateName string) *CreateContractTemplateOptions {
	_options.ContractTemplateName = core.StringPtr(contractTemplateName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateContractTemplateOptions) SetHeaders(param map[string]string) *CreateContractTemplateOptions {
	options.Headers = param
	return options
}

// CreateDataAssetVisualizationOptions : The CreateDataAssetVisualization options.
type CreateDataAssetVisualizationOptions struct {
	// Data product hub asset and it's related part asset.
	Assets []DataAssetRelationship `json:"assets,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDataAssetVisualizationOptions : Instantiate CreateDataAssetVisualizationOptions
func (*DataProductHubAPIServiceV1) NewCreateDataAssetVisualizationOptions() *CreateDataAssetVisualizationOptions {
	return &CreateDataAssetVisualizationOptions{}
}

// SetAssets : Allow user to set Assets
func (_options *CreateDataAssetVisualizationOptions) SetAssets(assets []DataAssetRelationship) *CreateDataAssetVisualizationOptions {
	_options.Assets = assets
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataAssetVisualizationOptions) SetHeaders(param map[string]string) *CreateDataAssetVisualizationOptions {
	options.Headers = param
	return options
}

// CreateDataProductDomainOptions : The CreateDataProductDomain options.
type CreateDataProductDomainOptions struct {
	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The id to trace the failed domain creations.
	Trace *string `json:"trace,omitempty"`

	// Set of errors on the sub domain creation.
	Errors []ErrorModelResource `json:"errors,omitempty"`

	// The name of the data product domain.
	Name *string `json:"name,omitempty"`

	// The description of the data product domain.
	Description *string `json:"description,omitempty"`

	// The identifier of the data product domain.
	ID *string `json:"id,omitempty"`

	// Member roles of a corresponding asset.
	MemberRoles *MemberRolesSchema `json:"member_roles,omitempty"`

	// Properties of the corresponding asset.
	Properties *PropertiesSchema `json:"properties,omitempty"`

	// List of sub domains to be added within a domain.
	SubDomains []InitializeSubDomain `json:"sub_domains,omitempty"`

	// Container ID of the data product catalog. If not supplied, the data product catalog is looked up by using the uid of
	// the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDataProductDomainOptions : Instantiate CreateDataProductDomainOptions
func (*DataProductHubAPIServiceV1) NewCreateDataProductDomainOptions(container *ContainerReference) *CreateDataProductDomainOptions {
	return &CreateDataProductDomainOptions{
		Container: container,
	}
}

// SetContainer : Allow user to set Container
func (_options *CreateDataProductDomainOptions) SetContainer(container *ContainerReference) *CreateDataProductDomainOptions {
	_options.Container = container
	return _options
}

// SetTrace : Allow user to set Trace
func (_options *CreateDataProductDomainOptions) SetTrace(trace string) *CreateDataProductDomainOptions {
	_options.Trace = core.StringPtr(trace)
	return _options
}

// SetErrors : Allow user to set Errors
func (_options *CreateDataProductDomainOptions) SetErrors(errors []ErrorModelResource) *CreateDataProductDomainOptions {
	_options.Errors = errors
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDataProductDomainOptions) SetName(name string) *CreateDataProductDomainOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDataProductDomainOptions) SetDescription(description string) *CreateDataProductDomainOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateDataProductDomainOptions) SetID(id string) *CreateDataProductDomainOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetMemberRoles : Allow user to set MemberRoles
func (_options *CreateDataProductDomainOptions) SetMemberRoles(memberRoles *MemberRolesSchema) *CreateDataProductDomainOptions {
	_options.MemberRoles = memberRoles
	return _options
}

// SetProperties : Allow user to set Properties
func (_options *CreateDataProductDomainOptions) SetProperties(properties *PropertiesSchema) *CreateDataProductDomainOptions {
	_options.Properties = properties
	return _options
}

// SetSubDomains : Allow user to set SubDomains
func (_options *CreateDataProductDomainOptions) SetSubDomains(subDomains []InitializeSubDomain) *CreateDataProductDomainOptions {
	_options.SubDomains = subDomains
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *CreateDataProductDomainOptions) SetContainerID(containerID string) *CreateDataProductDomainOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductDomainOptions) SetHeaders(param map[string]string) *CreateDataProductDomainOptions {
	options.Headers = param
	return options
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
	ContractTerms []ContractTerms `json:"contract_terms,omitempty"`

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

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateDataProductDraftOptions.State property.
// The state of the data product version. If not specified, the data product version will be created in `draft` state.
const (
	CreateDataProductDraftOptionsStateAvailableConst = "available"
	CreateDataProductDraftOptionsStateDraftConst = "draft"
	CreateDataProductDraftOptionsStateRetiredConst = "retired"
)

// Constants associated with the CreateDataProductDraftOptions.Types property.
const (
	CreateDataProductDraftOptionsTypesCodeConst = "code"
	CreateDataProductDraftOptionsTypesDataConst = "data"
)

// NewCreateDataProductDraftOptions : Instantiate CreateDataProductDraftOptions
func (*DataProductHubAPIServiceV1) NewCreateDataProductDraftOptions(dataProductID string, asset *AssetPrototype) *CreateDataProductDraftOptions {
	return &CreateDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		Asset: asset,
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
func (_options *CreateDataProductDraftOptions) SetContractTerms(contractTerms []ContractTerms) *CreateDataProductDraftOptions {
	_options.ContractTerms = contractTerms
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

// SetDataviewEnabled : Allow user to set DataviewEnabled
func (_options *CreateDataProductDraftOptions) SetDataviewEnabled(dataviewEnabled bool) *CreateDataProductDraftOptions {
	_options.DataviewEnabled = core.BoolPtr(dataviewEnabled)
	return _options
}

// SetComments : Allow user to set Comments
func (_options *CreateDataProductDraftOptions) SetComments(comments string) *CreateDataProductDraftOptions {
	_options.Comments = core.StringPtr(comments)
	return _options
}

// SetAccessControl : Allow user to set AccessControl
func (_options *CreateDataProductDraftOptions) SetAccessControl(accessControl *AssetListAccessControl) *CreateDataProductDraftOptions {
	_options.AccessControl = accessControl
	return _options
}

// SetLastUpdatedAt : Allow user to set LastUpdatedAt
func (_options *CreateDataProductDraftOptions) SetLastUpdatedAt(lastUpdatedAt *strfmt.DateTime) *CreateDataProductDraftOptions {
	_options.LastUpdatedAt = lastUpdatedAt
	return _options
}

// SetIsRestricted : Allow user to set IsRestricted
func (_options *CreateDataProductDraftOptions) SetIsRestricted(isRestricted bool) *CreateDataProductDraftOptions {
	_options.IsRestricted = core.BoolPtr(isRestricted)
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
	Drafts []DataProductDraftPrototype `json:"drafts" validate:"required"`

	// Limit the number of data products in the results. The maximum limit is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Start token for pagination.
	Start *string `json:"start,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDataProductOptions : Instantiate CreateDataProductOptions
func (*DataProductHubAPIServiceV1) NewCreateDataProductOptions(drafts []DataProductDraftPrototype) *CreateDataProductOptions {
	return &CreateDataProductOptions{
		Drafts: drafts,
	}
}

// SetDrafts : Allow user to set Drafts
func (_options *CreateDataProductOptions) SetDrafts(drafts []DataProductDraftPrototype) *CreateDataProductOptions {
	_options.Drafts = drafts
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *CreateDataProductOptions) SetLimit(limit int64) *CreateDataProductOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetStart : Allow user to set Start
func (_options *CreateDataProductOptions) SetStart(start string) *CreateDataProductOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductOptions) SetHeaders(param map[string]string) *CreateDataProductOptions {
	options.Headers = param
	return options
}

// CreateDataProductSubdomainOptions : The CreateDataProductSubdomain options.
type CreateDataProductSubdomainOptions struct {
	// Domain id.
	DomainID *string `json:"domain_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// The name of the data product subdomain.
	Name *string `json:"name,omitempty"`

	// The identifier of the data product subdomain.
	ID *string `json:"id,omitempty"`

	// The description of the data product subdomain.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDataProductSubdomainOptions : Instantiate CreateDataProductSubdomainOptions
func (*DataProductHubAPIServiceV1) NewCreateDataProductSubdomainOptions(domainID string, containerID string) *CreateDataProductSubdomainOptions {
	return &CreateDataProductSubdomainOptions{
		DomainID: core.StringPtr(domainID),
		ContainerID: core.StringPtr(containerID),
	}
}

// SetDomainID : Allow user to set DomainID
func (_options *CreateDataProductSubdomainOptions) SetDomainID(domainID string) *CreateDataProductSubdomainOptions {
	_options.DomainID = core.StringPtr(domainID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *CreateDataProductSubdomainOptions) SetContainerID(containerID string) *CreateDataProductSubdomainOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDataProductSubdomainOptions) SetName(name string) *CreateDataProductSubdomainOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetID : Allow user to set ID
func (_options *CreateDataProductSubdomainOptions) SetID(id string) *CreateDataProductSubdomainOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDataProductSubdomainOptions) SetDescription(description string) *CreateDataProductSubdomainOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDataProductSubdomainOptions) SetHeaders(param map[string]string) *CreateDataProductSubdomainOptions {
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

	// URL that can be used to retrieve the contract document.
	URL *string `json:"url,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateDraftContractTermsDocumentOptions.Type property.
// Type of the contract document.
const (
	CreateDraftContractTermsDocumentOptionsTypeSLAConst = "sla"
	CreateDraftContractTermsDocumentOptionsTypeTermsAndConditionsConst = "terms_and_conditions"
)

// NewCreateDraftContractTermsDocumentOptions : Instantiate CreateDraftContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewCreateDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, typeVar string, name string) *CreateDraftContractTermsDocumentOptions {
	return &CreateDraftContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		Type: core.StringPtr(typeVar),
		Name: core.StringPtr(name),
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

// SetURL : Allow user to set URL
func (_options *CreateDraftContractTermsDocumentOptions) SetURL(url string) *CreateDraftContractTermsDocumentOptions {
	_options.URL = core.StringPtr(url)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDraftContractTermsDocumentOptions) SetHeaders(param map[string]string) *CreateDraftContractTermsDocumentOptions {
	options.Headers = param
	return options
}

// CreateS3BucketOptions : The CreateS3Bucket options.
type CreateS3BucketOptions struct {
	// Flag to specify whether the bucket is dedicated or shared.
	IsShared *bool `json:"is_shared" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateS3BucketOptions : Instantiate CreateS3BucketOptions
func (*DataProductHubAPIServiceV1) NewCreateS3BucketOptions(isShared bool) *CreateS3BucketOptions {
	return &CreateS3BucketOptions{
		IsShared: core.BoolPtr(isShared),
	}
}

// SetIsShared : Allow user to set IsShared
func (_options *CreateS3BucketOptions) SetIsShared(isShared bool) *CreateS3BucketOptions {
	_options.IsShared = core.BoolPtr(isShared)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateS3BucketOptions) SetHeaders(param map[string]string) *CreateS3BucketOptions {
	options.Headers = param
	return options
}

// DataAssetRelationship : Data members for visualization process.
type DataAssetRelationship struct {
	// Data members for visualization.
	Visualization *Visualization `json:"visualization,omitempty"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`

	// The reference schema for a asset in a container.
	RelatedAsset *AssetReference `json:"related_asset" validate:"required"`

	// Contains the code and details.
	Error *ErrorMessage `json:"error,omitempty"`
}

// NewDataAssetRelationship : Instantiate DataAssetRelationship (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDataAssetRelationship(asset *AssetReference, relatedAsset *AssetReference) (_model *DataAssetRelationship, err error) {
	_model = &DataAssetRelationship{
		Asset: asset,
		RelatedAsset: relatedAsset,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataAssetRelationship unmarshals an instance of DataAssetRelationship from the specified map of raw messages.
func UnmarshalDataAssetRelationship(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataAssetRelationship)
	err = core.UnmarshalModel(m, "visualization", &obj.Visualization, UnmarshalVisualization)
	if err != nil {
		err = core.SDKErrorf(err, "", "visualization-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "related_asset", &obj.RelatedAsset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "related_asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalErrorMessage)
	if err != nil {
		err = core.SDKErrorf(err, "", "error-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataAssetVisualizationRes : Data relationships for the visualization process response.
type DataAssetVisualizationRes struct {
	// Data asset Ids and their related asset Ids.
	Results []DataAssetRelationship `json:"results,omitempty"`
}

// UnmarshalDataAssetVisualizationRes unmarshals an instance of DataAssetVisualizationRes from the specified map of raw messages.
func UnmarshalDataAssetVisualizationRes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataAssetVisualizationRes)
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalDataAssetRelationship)
	if err != nil {
		err = core.SDKErrorf(err, "", "results-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProduct : Data Product.
type DataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// Data product name.
	Name *string `json:"name,omitempty"`

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
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "latest_release", &obj.LatestRelease, UnmarshalDataProductVersionSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "latest_release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "drafts", &obj.Drafts, UnmarshalDataProductVersionSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "drafts-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductCollection : A collection of data product summaries.
type DataProductCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Indicates the total number of results returned.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Collection of data product summaries.
	DataProducts []DataProductSummary `json:"data_products" validate:"required"`
}

// UnmarshalDataProductCollection unmarshals an instance of DataProductCollection from the specified map of raw messages.
func UnmarshalDataProductCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_results-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_products", &obj.DataProducts, UnmarshalDataProductSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_products-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *DataProductCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// DataProductContractTemplate : Defines the complete structure of a contract template.
type DataProductContractTemplate struct {
	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The identifier of the data product contract template.
	ID *string `json:"id,omitempty"`

	// The name of the contract template.
	Name *string `json:"name,omitempty"`

	// Contains the code and details.
	Error *ErrorMessage `json:"error,omitempty"`

	// Defines the complete structure of a contract terms.
	ContractTerms *ContractTerms `json:"contract_terms,omitempty"`
}

// NewDataProductContractTemplate : Instantiate DataProductContractTemplate (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDataProductContractTemplate(container *ContainerReference) (_model *DataProductContractTemplate, err error) {
	_model = &DataProductContractTemplate{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataProductContractTemplate unmarshals an instance of DataProductContractTemplate from the specified map of raw messages.
func UnmarshalDataProductContractTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductContractTemplate)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalErrorMessage)
	if err != nil {
		err = core.SDKErrorf(err, "", "error-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewDataProductContractTemplatePatch(dataProductContractTemplate *DataProductContractTemplate) (_patch []JSONPatchOperation) {
	if (dataProductContractTemplate.Container != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/container"),
			Value: dataProductContractTemplate.Container,
		})
	}
	if (dataProductContractTemplate.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: dataProductContractTemplate.ID,
		})
	}
	if (dataProductContractTemplate.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/name"),
			Value: dataProductContractTemplate.Name,
		})
	}
	if (dataProductContractTemplate.Error != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/error"),
			Value: dataProductContractTemplate.Error,
		})
	}
	if (dataProductContractTemplate.ContractTerms != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/contract_terms"),
			Value: dataProductContractTemplate.ContractTerms,
		})
	}
	return
}

// DataProductContractTemplateCollection : A collection of data product contract templates.
type DataProductContractTemplateCollection struct {
	// Collection of data product contract templates.
	ContractTemplates []DataProductContractTemplate `json:"contract_templates" validate:"required"`
}

// UnmarshalDataProductContractTemplateCollection unmarshals an instance of DataProductContractTemplateCollection from the specified map of raw messages.
func UnmarshalDataProductContractTemplateCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductContractTemplateCollection)
	err = core.UnmarshalModel(m, "contract_templates", &obj.ContractTemplates, UnmarshalDataProductContractTemplate)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_templates-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductCustomWorkflowDefinition : A custom workflow definition to be used to create a workflow to approve a data product subscription.
type DataProductCustomWorkflowDefinition struct {
	// ID of a workflow definition.
	ID *string `json:"id,omitempty"`
}

// UnmarshalDataProductCustomWorkflowDefinition unmarshals an instance of DataProductCustomWorkflowDefinition from the specified map of raw messages.
func UnmarshalDataProductCustomWorkflowDefinition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductCustomWorkflowDefinition)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDomain : The data product domain.
type DataProductDomain struct {
	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// The id to trace the failed domain creations.
	Trace *string `json:"trace,omitempty"`

	// Set of errors on the sub domain creation.
	Errors []ErrorModelResource `json:"errors,omitempty"`

	// The name of the data product domain.
	Name *string `json:"name,omitempty"`

	// The description of the data product domain.
	Description *string `json:"description,omitempty"`

	// The identifier of the data product domain.
	ID *string `json:"id,omitempty"`

	// Member roles of a corresponding asset.
	MemberRoles *MemberRolesSchema `json:"member_roles,omitempty"`

	// Properties of the corresponding asset.
	Properties *PropertiesSchema `json:"properties,omitempty"`

	// List of sub domains to be added within a domain.
	SubDomains []InitializeSubDomain `json:"sub_domains,omitempty"`
}

// NewDataProductDomain : Instantiate DataProductDomain (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDataProductDomain(container *ContainerReference) (_model *DataProductDomain, err error) {
	_model = &DataProductDomain{
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataProductDomain unmarshals an instance of DataProductDomain from the specified map of raw messages.
func UnmarshalDataProductDomain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDomain)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		err = core.SDKErrorf(err, "", "trace-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModelResource)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "member_roles", &obj.MemberRoles, UnmarshalMemberRolesSchema)
	if err != nil {
		err = core.SDKErrorf(err, "", "member_roles-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalPropertiesSchema)
	if err != nil {
		err = core.SDKErrorf(err, "", "properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "sub_domains", &obj.SubDomains, UnmarshalInitializeSubDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "sub_domains-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewDataProductDomainPatch(dataProductDomain *DataProductDomain) (_patch []JSONPatchOperation) {
	if (dataProductDomain.Container != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/container"),
			Value: dataProductDomain.Container,
		})
	}
	if (dataProductDomain.Trace != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/trace"),
			Value: dataProductDomain.Trace,
		})
	}
	if (dataProductDomain.Errors != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/errors"),
			Value: dataProductDomain.Errors,
		})
	}
	if (dataProductDomain.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/name"),
			Value: dataProductDomain.Name,
		})
	}
	if (dataProductDomain.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/description"),
			Value: dataProductDomain.Description,
		})
	}
	if (dataProductDomain.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: dataProductDomain.ID,
		})
	}
	if (dataProductDomain.MemberRoles != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/member_roles"),
			Value: dataProductDomain.MemberRoles,
		})
	}
	if (dataProductDomain.Properties != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/properties"),
			Value: dataProductDomain.Properties,
		})
	}
	if (dataProductDomain.SubDomains != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/sub_domains"),
			Value: dataProductDomain.SubDomains,
		})
	}
	return
}

// DataProductDomainCollection : A collection of data product domains.
type DataProductDomainCollection struct {
	// Collection of data product domains.
	Domains []DataProductDomain `json:"domains" validate:"required"`
}

// UnmarshalDataProductDomainCollection unmarshals an instance of DataProductDomainCollection from the specified map of raw messages.
func UnmarshalDataProductDomainCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDomainCollection)
	err = core.UnmarshalModel(m, "domains", &obj.Domains, UnmarshalDataProductDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domains-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraft : Data Product version draft.
type DataProductDraft struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductDraftDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags" validate:"required"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []ContractTerms `json:"contract_terms" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain" validate:"required"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out" validate:"required"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`

	// The user who published this data product version.
	PublishedBy *string `json:"published_by,omitempty"`

	// The time when this data product version was published.
	PublishedAt *strfmt.DateTime `json:"published_at,omitempty"`

	// The creator of this data product version.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The time when this data product version was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Metadata properties on data products.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Errors encountered during the visualization creation process.
	VisualizationErrors []DataAssetRelationship `json:"visualization_errors,omitempty"`
}

// Constants associated with the DataProductDraft.State property.
// The state of the data product version.
const (
	DataProductDraftStateAvailableConst = "available"
	DataProductDraftStateDraftConst = "draft"
	DataProductDraftStateRetiredConst = "retired"
)

// Constants associated with the DataProductDraft.Types property.
const (
	DataProductDraftTypesCodeConst = "code"
	DataProductDraftTypesDataConst = "data"
)

// UnmarshalDataProductDraft unmarshals an instance of DataProductDraft from the specified map of raw messages.
func UnmarshalDataProductDraft(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraft)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductDraftDataProduct)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "published_by", &obj.PublishedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "published_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "published_at", &obj.PublishedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "published_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "properties", &obj.Properties)
	if err != nil {
		err = core.SDKErrorf(err, "", "properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "visualization_errors", &obj.VisualizationErrors, UnmarshalDataAssetRelationship)
	if err != nil {
		err = core.SDKErrorf(err, "", "visualization_errors-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewDataProductDraftPatch(dataProductDraft *DataProductDraft) (_patch []JSONPatchOperation) {
	if (dataProductDraft.Version != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/version"),
			Value: dataProductDraft.Version,
		})
	}
	if (dataProductDraft.State != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/state"),
			Value: dataProductDraft.State,
		})
	}
	if (dataProductDraft.DataProduct != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/data_product"),
			Value: dataProductDraft.DataProduct,
		})
	}
	if (dataProductDraft.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/name"),
			Value: dataProductDraft.Name,
		})
	}
	if (dataProductDraft.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/description"),
			Value: dataProductDraft.Description,
		})
	}
	if (dataProductDraft.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/tags"),
			Value: dataProductDraft.Tags,
		})
	}
	if (dataProductDraft.UseCases != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/use_cases"),
			Value: dataProductDraft.UseCases,
		})
	}
	if (dataProductDraft.Types != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/types"),
			Value: dataProductDraft.Types,
		})
	}
	if (dataProductDraft.ContractTerms != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/contract_terms"),
			Value: dataProductDraft.ContractTerms,
		})
	}
	if (dataProductDraft.Domain != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/domain"),
			Value: dataProductDraft.Domain,
		})
	}
	if (dataProductDraft.PartsOut != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/parts_out"),
			Value: dataProductDraft.PartsOut,
		})
	}
	if (dataProductDraft.Workflows != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/workflows"),
			Value: dataProductDraft.Workflows,
		})
	}
	if (dataProductDraft.DataviewEnabled != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/dataview_enabled"),
			Value: dataProductDraft.DataviewEnabled,
		})
	}
	if (dataProductDraft.Comments != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/comments"),
			Value: dataProductDraft.Comments,
		})
	}
	if (dataProductDraft.AccessControl != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/access_control"),
			Value: dataProductDraft.AccessControl,
		})
	}
	if (dataProductDraft.LastUpdatedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/last_updated_at"),
			Value: dataProductDraft.LastUpdatedAt,
		})
	}
	if (dataProductDraft.IsRestricted != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/is_restricted"),
			Value: dataProductDraft.IsRestricted,
		})
	}
	if (dataProductDraft.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: dataProductDraft.ID,
		})
	}
	if (dataProductDraft.Asset != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/asset"),
			Value: dataProductDraft.Asset,
		})
	}
	if (dataProductDraft.PublishedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/published_by"),
			Value: dataProductDraft.PublishedBy,
		})
	}
	if (dataProductDraft.PublishedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/published_at"),
			Value: dataProductDraft.PublishedAt,
		})
	}
	if (dataProductDraft.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/created_by"),
			Value: dataProductDraft.CreatedBy,
		})
	}
	if (dataProductDraft.CreatedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/created_at"),
			Value: dataProductDraft.CreatedAt,
		})
	}
	if (dataProductDraft.Properties != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/properties"),
			Value: dataProductDraft.Properties,
		})
	}
	if (dataProductDraft.VisualizationErrors != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/visualization_errors"),
			Value: dataProductDraft.VisualizationErrors,
		})
	}
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

	// Indicates the total number of results returned.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Collection of data product drafts.
	Drafts []DataProductDraftSummary `json:"drafts" validate:"required"`
}

// UnmarshalDataProductDraftCollection unmarshals an instance of DataProductDraftCollection from the specified map of raw messages.
func UnmarshalDataProductDraftCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_results-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "drafts", &obj.Drafts, UnmarshalDataProductDraftSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "drafts-error", common.GetComponentInfo())
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

// DataProductDraftDataProduct : Data product reference.
type DataProductDraftDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductDraftDataProduct unmarshals an instance of DataProductDraftDataProduct from the specified map of raw messages.
func UnmarshalDataProductDraftDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraftPrototype : New data product version input properties.
type DataProductDraftPrototype struct {
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
	ContractTerms []ContractTerms `json:"contract_terms,omitempty"`

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

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted,omitempty"`

	// New asset input properties.
	Asset *AssetPrototype `json:"asset" validate:"required"`
}

// Constants associated with the DataProductDraftPrototype.State property.
// The state of the data product version. If not specified, the data product version will be created in `draft` state.
const (
	DataProductDraftPrototypeStateAvailableConst = "available"
	DataProductDraftPrototypeStateDraftConst = "draft"
	DataProductDraftPrototypeStateRetiredConst = "retired"
)

// Constants associated with the DataProductDraftPrototype.Types property.
const (
	DataProductDraftPrototypeTypesCodeConst = "code"
	DataProductDraftPrototypeTypesDataConst = "data"
)

// NewDataProductDraftPrototype : Instantiate DataProductDraftPrototype (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDataProductDraftPrototype(asset *AssetPrototype) (_model *DataProductDraftPrototype, err error) {
	_model = &DataProductDraftPrototype{
		Asset: asset,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataProductDraftPrototype unmarshals an instance of DataProductDraftPrototype from the specified map of raw messages.
func UnmarshalDataProductDraftPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftPrototype)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductIdentity)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraftSummary : Summary of Data Product Version object.
type DataProductDraftSummary struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductDraftSummaryDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags" validate:"required"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []ContractTerms `json:"contract_terms" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain" validate:"required"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out" validate:"required"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`
}

// Constants associated with the DataProductDraftSummary.State property.
// The state of the data product version.
const (
	DataProductDraftSummaryStateAvailableConst = "available"
	DataProductDraftSummaryStateDraftConst = "draft"
	DataProductDraftSummaryStateRetiredConst = "retired"
)

// Constants associated with the DataProductDraftSummary.Types property.
const (
	DataProductDraftSummaryTypesCodeConst = "code"
	DataProductDraftSummaryTypesDataConst = "data"
)

// UnmarshalDataProductDraftSummary unmarshals an instance of DataProductDraftSummary from the specified map of raw messages.
func UnmarshalDataProductDraftSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductDraftSummaryDataProduct)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraftSummaryDataProduct : Data product reference.
type DataProductDraftSummaryDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductDraftSummaryDataProduct unmarshals an instance of DataProductDraftSummaryDataProduct from the specified map of raw messages.
func UnmarshalDataProductDraftSummaryDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftSummaryDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductDraftVersionRelease : A data product draft version object.
type DataProductDraftVersionRelease struct {
	// ID of a draft version of data product.
	ID *string `json:"id,omitempty"`
}

// UnmarshalDataProductDraftVersionRelease unmarshals an instance of DataProductDraftVersionRelease from the specified map of raw messages.
func UnmarshalDataProductDraftVersionRelease(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductDraftVersionRelease)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductIdentity : Data product identifier.
type DataProductIdentity struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`
}

// NewDataProductIdentity : Instantiate DataProductIdentity (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDataProductIdentity(id string) (_model *DataProductIdentity, err error) {
	_model = &DataProductIdentity{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataProductIdentity unmarshals an instance of DataProductIdentity from the specified map of raw messages.
func UnmarshalDataProductIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductOrderAccessRequest : The approval workflows associated with the data product version.
type DataProductOrderAccessRequest struct {
	// The workflow approvers associated with the data product version.
	TaskAssigneeUsers []string `json:"task_assignee_users,omitempty"`

	// The list of users or groups whose request will get pre-approved associated with the data product version.
	PreApprovedUsers []string `json:"pre_approved_users,omitempty"`

	// A custom workflow definition to be used to create a workflow to approve a data product subscription.
	CustomWorkflowDefinition *DataProductCustomWorkflowDefinition `json:"custom_workflow_definition,omitempty"`
}

// UnmarshalDataProductOrderAccessRequest unmarshals an instance of DataProductOrderAccessRequest from the specified map of raw messages.
func UnmarshalDataProductOrderAccessRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductOrderAccessRequest)
	err = core.UnmarshalPrimitive(m, "task_assignee_users", &obj.TaskAssigneeUsers)
	if err != nil {
		err = core.SDKErrorf(err, "", "task_assignee_users-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "pre_approved_users", &obj.PreApprovedUsers)
	if err != nil {
		err = core.SDKErrorf(err, "", "pre_approved_users-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "custom_workflow_definition", &obj.CustomWorkflowDefinition, UnmarshalDataProductCustomWorkflowDefinition)
	if err != nil {
		err = core.SDKErrorf(err, "", "custom_workflow_definition-error", common.GetComponentInfo())
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
func (*DataProductHubAPIServiceV1) NewDataProductPart(asset *AssetPartReference) (_model *DataProductPart, err error) {
	_model = &DataProductPart{
		Asset: asset,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDataProductPart unmarshals an instance of DataProductPart from the specified map of raw messages.
func UnmarshalDataProductPart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductPart)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetPartReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "delivery_methods", &obj.DeliveryMethods, UnmarshalDeliveryMethod)
	if err != nil {
		err = core.SDKErrorf(err, "", "delivery_methods-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductRelease : Data Product version release.
type DataProductRelease struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductReleaseDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags" validate:"required"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []ContractTerms `json:"contract_terms" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain" validate:"required"`

	// The outgoing parts of this data product version to be delivered to consumers. If this is the first version of a data
	// product, this field defaults to an empty list. If this is a new version of an existing data product, the data
	// product parts will default to the parts list from the previous version of the data product.
	PartsOut []DataProductPart `json:"parts_out" validate:"required"`

	// The workflows associated with the data product version.
	Workflows *DataProductWorkflows `json:"workflows,omitempty"`

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`

	// The user who published this data product version.
	PublishedBy *string `json:"published_by,omitempty"`

	// The time when this data product version was published.
	PublishedAt *strfmt.DateTime `json:"published_at,omitempty"`

	// The creator of this data product version.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The time when this data product version was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// Metadata properties on data products.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Errors encountered during the visualization creation process.
	VisualizationErrors []DataAssetRelationship `json:"visualization_errors,omitempty"`
}

// Constants associated with the DataProductRelease.State property.
// The state of the data product version.
const (
	DataProductReleaseStateAvailableConst = "available"
	DataProductReleaseStateDraftConst = "draft"
	DataProductReleaseStateRetiredConst = "retired"
)

// Constants associated with the DataProductRelease.Types property.
const (
	DataProductReleaseTypesCodeConst = "code"
	DataProductReleaseTypesDataConst = "data"
)

// UnmarshalDataProductRelease unmarshals an instance of DataProductRelease from the specified map of raw messages.
func UnmarshalDataProductRelease(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductRelease)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductReleaseDataProduct)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "published_by", &obj.PublishedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "published_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "published_at", &obj.PublishedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "published_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_by-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "properties", &obj.Properties)
	if err != nil {
		err = core.SDKErrorf(err, "", "properties-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "visualization_errors", &obj.VisualizationErrors, UnmarshalDataAssetRelationship)
	if err != nil {
		err = core.SDKErrorf(err, "", "visualization_errors-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*DataProductHubAPIServiceV1) NewDataProductReleasePatch(dataProductRelease *DataProductRelease) (_patch []JSONPatchOperation) {
	if (dataProductRelease.Version != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/version"),
			Value: dataProductRelease.Version,
		})
	}
	if (dataProductRelease.State != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/state"),
			Value: dataProductRelease.State,
		})
	}
	if (dataProductRelease.DataProduct != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/data_product"),
			Value: dataProductRelease.DataProduct,
		})
	}
	if (dataProductRelease.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/name"),
			Value: dataProductRelease.Name,
		})
	}
	if (dataProductRelease.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/description"),
			Value: dataProductRelease.Description,
		})
	}
	if (dataProductRelease.Tags != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/tags"),
			Value: dataProductRelease.Tags,
		})
	}
	if (dataProductRelease.UseCases != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/use_cases"),
			Value: dataProductRelease.UseCases,
		})
	}
	if (dataProductRelease.Types != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/types"),
			Value: dataProductRelease.Types,
		})
	}
	if (dataProductRelease.ContractTerms != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/contract_terms"),
			Value: dataProductRelease.ContractTerms,
		})
	}
	if (dataProductRelease.Domain != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/domain"),
			Value: dataProductRelease.Domain,
		})
	}
	if (dataProductRelease.PartsOut != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/parts_out"),
			Value: dataProductRelease.PartsOut,
		})
	}
	if (dataProductRelease.Workflows != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/workflows"),
			Value: dataProductRelease.Workflows,
		})
	}
	if (dataProductRelease.DataviewEnabled != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/dataview_enabled"),
			Value: dataProductRelease.DataviewEnabled,
		})
	}
	if (dataProductRelease.Comments != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/comments"),
			Value: dataProductRelease.Comments,
		})
	}
	if (dataProductRelease.AccessControl != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/access_control"),
			Value: dataProductRelease.AccessControl,
		})
	}
	if (dataProductRelease.LastUpdatedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/last_updated_at"),
			Value: dataProductRelease.LastUpdatedAt,
		})
	}
	if (dataProductRelease.IsRestricted != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/is_restricted"),
			Value: dataProductRelease.IsRestricted,
		})
	}
	if (dataProductRelease.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/id"),
			Value: dataProductRelease.ID,
		})
	}
	if (dataProductRelease.Asset != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/asset"),
			Value: dataProductRelease.Asset,
		})
	}
	if (dataProductRelease.PublishedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/published_by"),
			Value: dataProductRelease.PublishedBy,
		})
	}
	if (dataProductRelease.PublishedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/published_at"),
			Value: dataProductRelease.PublishedAt,
		})
	}
	if (dataProductRelease.CreatedBy != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/created_by"),
			Value: dataProductRelease.CreatedBy,
		})
	}
	if (dataProductRelease.CreatedAt != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/created_at"),
			Value: dataProductRelease.CreatedAt,
		})
	}
	if (dataProductRelease.Properties != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/properties"),
			Value: dataProductRelease.Properties,
		})
	}
	if (dataProductRelease.VisualizationErrors != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperationOpAddConst),
			Path: core.StringPtr("/visualization_errors"),
			Value: dataProductRelease.VisualizationErrors,
		})
	}
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

	// Indicates the total number of results returned.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Collection of data product releases.
	Releases []DataProductReleaseSummary `json:"releases" validate:"required"`
}

// UnmarshalDataProductReleaseCollection unmarshals an instance of DataProductReleaseCollection from the specified map of raw messages.
func UnmarshalDataProductReleaseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductReleaseCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_results-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "releases", &obj.Releases, UnmarshalDataProductReleaseSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "releases-error", common.GetComponentInfo())
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

// DataProductReleaseDataProduct : Data product reference.
type DataProductReleaseDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductReleaseDataProduct unmarshals an instance of DataProductReleaseDataProduct from the specified map of raw messages.
func UnmarshalDataProductReleaseDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductReleaseDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductReleaseSummary : Summary of Data Product Version object.
type DataProductReleaseSummary struct {
	// The data product version number.
	Version *string `json:"version" validate:"required"`

	// The state of the data product version.
	State *string `json:"state" validate:"required"`

	// Data product reference.
	DataProduct *DataProductReleaseSummaryDataProduct `json:"data_product" validate:"required"`

	// The name of the data product version. A name can contain letters, numbers, understores, dashes, spaces or periods.
	// Names are mutable and reusable.
	Name *string `json:"name" validate:"required"`

	// The description of the data product version.
	Description *string `json:"description" validate:"required"`

	// Tags on the data product.
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []ContractTerms `json:"contract_terms" validate:"required"`

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

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`
}

// Constants associated with the DataProductReleaseSummary.State property.
// The state of the data product version.
const (
	DataProductReleaseSummaryStateAvailableConst = "available"
	DataProductReleaseSummaryStateDraftConst = "draft"
	DataProductReleaseSummaryStateRetiredConst = "retired"
)

// Constants associated with the DataProductReleaseSummary.Types property.
const (
	DataProductReleaseSummaryTypesCodeConst = "code"
	DataProductReleaseSummaryTypesDataConst = "data"
)

// UnmarshalDataProductReleaseSummary unmarshals an instance of DataProductReleaseSummary from the specified map of raw messages.
func UnmarshalDataProductReleaseSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductReleaseSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductReleaseSummaryDataProduct)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductReleaseSummaryDataProduct : Data product reference.
type DataProductReleaseSummaryDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductReleaseSummaryDataProduct unmarshals an instance of DataProductReleaseSummaryDataProduct from the specified map of raw messages.
func UnmarshalDataProductReleaseSummaryDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductReleaseSummaryDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductSummary : Data Product Summary.
type DataProductSummary struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`

	// Data product name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalDataProductSummary unmarshals an instance of DataProductSummary from the specified map of raw messages.
func UnmarshalDataProductSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersionCollection : A collection of data product version summaries.
type DataProductVersionCollection struct {
	// Set a limit on the number of results returned.
	Limit *int64 `json:"limit" validate:"required"`

	// First page in the collection.
	First *FirstPage `json:"first" validate:"required"`

	// Next page in the collection.
	Next *NextPage `json:"next,omitempty"`

	// Indicates the total number of results returned.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Collection of data product versions.
	DataProductVersions []DataProductVersionSummary `json:"data_product_versions" validate:"required"`
}

// UnmarshalDataProductVersionCollection unmarshals an instance of DataProductVersionCollection from the specified map of raw messages.
func UnmarshalDataProductVersionCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextPage)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_results-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product_versions", &obj.DataProductVersions, UnmarshalDataProductVersionSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product_versions-error", common.GetComponentInfo())
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
	Tags []string `json:"tags,omitempty"`

	// A list of use cases associated with the data product version.
	UseCases []UseCase `json:"use_cases,omitempty"`

	// Types of parts on the data product.
	Types []string `json:"types" validate:"required"`

	// Contract terms binding various aspects of the data product.
	ContractTerms []ContractTerms `json:"contract_terms" validate:"required"`

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

	// Indicates whether the dataView has enabled for data product.
	DataviewEnabled *bool `json:"dataview_enabled,omitempty"`

	// Comments by a producer that are provided either at the time of data product version creation or retiring.
	Comments *string `json:"comments,omitempty"`

	// Access control object.
	AccessControl *AssetListAccessControl `json:"access_control,omitempty"`

	// Timestamp of last asset update.
	LastUpdatedAt *strfmt.DateTime `json:"last_updated_at,omitempty"`

	// Indicates whether the data product is restricted or not. A restricted data product indicates that orders of the data
	// product requires explicit approval before data is delivered.
	IsRestricted *bool `json:"is_restricted" validate:"required"`

	// The identifier of the data product version.
	ID *string `json:"id" validate:"required"`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset" validate:"required"`
}

// Constants associated with the DataProductVersionSummary.State property.
// The state of the data product version.
const (
	DataProductVersionSummaryStateAvailableConst = "available"
	DataProductVersionSummaryStateDraftConst = "draft"
	DataProductVersionSummaryStateRetiredConst = "retired"
)

// Constants associated with the DataProductVersionSummary.Types property.
const (
	DataProductVersionSummaryTypesCodeConst = "code"
	DataProductVersionSummaryTypesDataConst = "data"
)

// UnmarshalDataProductVersionSummary unmarshals an instance of DataProductVersionSummary from the specified map of raw messages.
func UnmarshalDataProductVersionSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "data_product", &obj.DataProduct, UnmarshalDataProductVersionSummaryDataProduct)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_product-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		err = core.SDKErrorf(err, "", "tags-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "use_cases", &obj.UseCases, UnmarshalUseCase)
	if err != nil {
		err = core.SDKErrorf(err, "", "use_cases-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		err = core.SDKErrorf(err, "", "types-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "contract_terms", &obj.ContractTerms, UnmarshalContractTerms)
	if err != nil {
		err = core.SDKErrorf(err, "", "contract_terms-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parts_out", &obj.PartsOut, UnmarshalDataProductPart)
	if err != nil {
		err = core.SDKErrorf(err, "", "parts_out-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalDataProductWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "dataview_enabled", &obj.DataviewEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "dataview_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "comments", &obj.Comments)
	if err != nil {
		err = core.SDKErrorf(err, "", "comments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "access_control", &obj.AccessControl, UnmarshalAssetListAccessControl)
	if err != nil {
		err = core.SDKErrorf(err, "", "access_control-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_updated_at", &obj.LastUpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restricted", &obj.IsRestricted)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_restricted-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAssetReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "asset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataProductVersionSummaryDataProduct : Data product reference.
type DataProductVersionSummaryDataProduct struct {
	// Data product identifier.
	ID *string `json:"id" validate:"required"`

	// A data product draft version object.
	Release *DataProductDraftVersionRelease `json:"release,omitempty"`

	// Container reference.
	Container *ContainerReference `json:"container" validate:"required"`
}

// UnmarshalDataProductVersionSummaryDataProduct unmarshals an instance of DataProductVersionSummaryDataProduct from the specified map of raw messages.
func UnmarshalDataProductVersionSummaryDataProduct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataProductVersionSummaryDataProduct)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalDataProductDraftVersionRelease)
	if err != nil {
		err = core.SDKErrorf(err, "", "release-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "order_access_request-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDataProductContractTemplateOptions : The DeleteDataProductContractTemplate options.
type DeleteDataProductContractTemplateOptions struct {
	// Data Product Contract Template id.
	ContractTemplateID *string `json:"contract_template_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDataProductContractTemplateOptions : Instantiate DeleteDataProductContractTemplateOptions
func (*DataProductHubAPIServiceV1) NewDeleteDataProductContractTemplateOptions(contractTemplateID string, containerID string) *DeleteDataProductContractTemplateOptions {
	return &DeleteDataProductContractTemplateOptions{
		ContractTemplateID: core.StringPtr(contractTemplateID),
		ContainerID: core.StringPtr(containerID),
	}
}

// SetContractTemplateID : Allow user to set ContractTemplateID
func (_options *DeleteDataProductContractTemplateOptions) SetContractTemplateID(contractTemplateID string) *DeleteDataProductContractTemplateOptions {
	_options.ContractTemplateID = core.StringPtr(contractTemplateID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *DeleteDataProductContractTemplateOptions) SetContainerID(containerID string) *DeleteDataProductContractTemplateOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDataProductContractTemplateOptions) SetHeaders(param map[string]string) *DeleteDataProductContractTemplateOptions {
	options.Headers = param
	return options
}

// DeleteDataProductDraftOptions : The DeleteDataProductDraft options.
type DeleteDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDataProductDraftOptions : Instantiate DeleteDataProductDraftOptions
func (*DataProductHubAPIServiceV1) NewDeleteDataProductDraftOptions(dataProductID string, draftID string) *DeleteDataProductDraftOptions {
	return &DeleteDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
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

// DeleteDomainOptions : The DeleteDomain options.
type DeleteDomainOptions struct {
	// Domain id.
	DomainID *string `json:"domain_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDomainOptions : Instantiate DeleteDomainOptions
func (*DataProductHubAPIServiceV1) NewDeleteDomainOptions(domainID string) *DeleteDomainOptions {
	return &DeleteDomainOptions{
		DomainID: core.StringPtr(domainID),
	}
}

// SetDomainID : Allow user to set DomainID
func (_options *DeleteDomainOptions) SetDomainID(domainID string) *DeleteDomainOptions {
	_options.DomainID = core.StringPtr(domainID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDomainOptions) SetHeaders(param map[string]string) *DeleteDomainOptions {
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDraftContractTermsDocumentOptions : Instantiate DeleteDraftContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewDeleteDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *DeleteDraftContractTermsDocumentOptions {
	return &DeleteDraftContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
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

	// The propertiess of the delivery method.
	Getproperties *DeliveryMethodPropertiesModel `json:"getproperties,omitempty"`
}

// NewDeliveryMethod : Instantiate DeliveryMethod (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewDeliveryMethod(id string, container *ContainerReference) (_model *DeliveryMethod, err error) {
	_model = &DeliveryMethod{
		ID: core.StringPtr(id),
		Container: container,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDeliveryMethod unmarshals an instance of DeliveryMethod from the specified map of raw messages.
func UnmarshalDeliveryMethod(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeliveryMethod)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "getproperties", &obj.Getproperties, UnmarshalDeliveryMethodPropertiesModel)
	if err != nil {
		err = core.SDKErrorf(err, "", "getproperties-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeliveryMethodPropertiesModel : The propertiess of the delivery method.
type DeliveryMethodPropertiesModel struct {
	// Parameters for delivery that are set by a data product producer.
	ProducerInput *ProducerInputModel `json:"producer_input,omitempty"`
}

// UnmarshalDeliveryMethodPropertiesModel unmarshals an instance of DeliveryMethodPropertiesModel from the specified map of raw messages.
func UnmarshalDeliveryMethodPropertiesModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeliveryMethodPropertiesModel)
	err = core.UnmarshalModel(m, "producer_input", &obj.ProducerInput, UnmarshalProducerInputModel)
	if err != nil {
		err = core.SDKErrorf(err, "", "producer_input-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Description : Description details of a data contract.
type Description struct {
	// Intended purpose for the provided data.
	Purpose *string `json:"purpose,omitempty"`

	// Technical, compliance, and legal limitations for data use.
	Limitations *string `json:"limitations,omitempty"`

	// Recommended usage of the data.
	Usage *string `json:"usage,omitempty"`

	// List of links to sources that provide more details on the dataset.
	MoreInfo []ContractTermsMoreInfo `json:"more_info,omitempty"`

	// Custom properties that are not part of the standard.
	CustomProperties *string `json:"custom_properties,omitempty"`
}

// UnmarshalDescription unmarshals an instance of Description from the specified map of raw messages.
func UnmarshalDescription(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Description)
	err = core.UnmarshalPrimitive(m, "purpose", &obj.Purpose)
	if err != nil {
		err = core.SDKErrorf(err, "", "purpose-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "limitations", &obj.Limitations)
	if err != nil {
		err = core.SDKErrorf(err, "", "limitations-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "usage", &obj.Usage)
	if err != nil {
		err = core.SDKErrorf(err, "", "usage-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "more_info", &obj.MoreInfo, UnmarshalContractTermsMoreInfo)
	if err != nil {
		err = core.SDKErrorf(err, "", "more_info-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "custom_properties", &obj.CustomProperties)
	if err != nil {
		err = core.SDKErrorf(err, "", "custom_properties-error", common.GetComponentInfo())
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
func (*DataProductHubAPIServiceV1) NewDomain(id string) (_model *Domain, err error) {
	_model = &Domain{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalDomain unmarshals an instance of Domain from the specified map of raw messages.
func UnmarshalDomain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Domain)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EngineDetailsModel : Engine details as defined by the data product producer.
type EngineDetailsModel struct {
	// The name of the engine defined by the data product producer.
	DisplayName *string `json:"display_name,omitempty"`

	// The id of the engine defined by the data product producer.
	EngineID *string `json:"engine_id,omitempty"`

	// The port of the engine defined by the data product producer.
	EnginePort *string `json:"engine_port,omitempty"`

	// The host of the engine defined by the data product producer.
	EngineHost *string `json:"engine_host,omitempty"`

	// The list of associated catalogs.
	AssociatedCatalogs []string `json:"associated_catalogs,omitempty"`
}

// UnmarshalEngineDetailsModel unmarshals an instance of EngineDetailsModel from the specified map of raw messages.
func UnmarshalEngineDetailsModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EngineDetailsModel)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		err = core.SDKErrorf(err, "", "display_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_id", &obj.EngineID)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_port", &obj.EnginePort)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_host", &obj.EngineHost)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "associated_catalogs", &obj.AssociatedCatalogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "associated_catalogs-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorExtraResource : Detailed error information.
type ErrorExtraResource struct {
	// Error id.
	ID *string `json:"id,omitempty"`

	// Timestamp of the error.
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// Environment where the error occurred.
	EnvironmentName *string `json:"environment_name,omitempty"`

	// Http status code.
	HTTPStatus *int64 `json:"http_status,omitempty"`

	// Source cluster of the error.
	SourceCluster *int64 `json:"source_cluster,omitempty"`

	// Source component of the error.
	SourceComponent *int64 `json:"source_component,omitempty"`

	// Transaction id of the request.
	TransactionID *int64 `json:"transaction_id,omitempty"`
}

// UnmarshalErrorExtraResource unmarshals an instance of ErrorExtraResource from the specified map of raw messages.
func UnmarshalErrorExtraResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorExtraResource)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		err = core.SDKErrorf(err, "", "timestamp-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_name", &obj.EnvironmentName)
	if err != nil {
		err = core.SDKErrorf(err, "", "environment_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "http_status", &obj.HTTPStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "http_status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_cluster", &obj.SourceCluster)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_cluster-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_component", &obj.SourceComponent)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_component-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_id", &obj.TransactionID)
	if err != nil {
		err = core.SDKErrorf(err, "", "transaction_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorMessage : Contains the code and details.
type ErrorMessage struct {
	// The error code.
	Code *string `json:"code" validate:"required"`

	// The error details.
	Message *string `json:"message" validate:"required"`
}

// NewErrorMessage : Instantiate ErrorMessage (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewErrorMessage(code string, message string) (_model *ErrorMessage, err error) {
	_model = &ErrorMessage{
		Code: core.StringPtr(code),
		Message: core.StringPtr(message),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalErrorMessage unmarshals an instance of ErrorMessage from the specified map of raw messages.
func UnmarshalErrorMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorMessage)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		err = core.SDKErrorf(err, "", "code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorModelResource : Detailed error information.
type ErrorModelResource struct {
	// Error code.
	Code *string `json:"code" validate:"required"`

	// Error message.
	Message *string `json:"message,omitempty"`

	// Detailed error information.
	Extra *ErrorExtraResource `json:"extra,omitempty"`

	// More info message.
	MoreInfo *string `json:"more_info,omitempty"`
}

// Constants associated with the ErrorModelResource.Code property.
// Error code.
const (
	ErrorModelResourceCodeAlreadyExistsConst = "already_exists"
	ErrorModelResourceCodeConfigurationErrorConst = "configuration_error"
	ErrorModelResourceCodeConflictConst = "conflict"
	ErrorModelResourceCodeConstraintViolationConst = "constraint_violation"
	ErrorModelResourceCodeCreateErrorConst = "create_error"
	ErrorModelResourceCodeDataErrorConst = "data_error"
	ErrorModelResourceCodeDatabaseErrorConst = "database_error"
	ErrorModelResourceCodeDatabaseQueryErrorConst = "database_query_error"
	ErrorModelResourceCodeDatabaseUsageLimitsConst = "database_usage_limits"
	ErrorModelResourceCodeDeleteErrorConst = "delete_error"
	ErrorModelResourceCodeDeletedConst = "deleted"
	ErrorModelResourceCodeDependentServiceErrorConst = "dependent_service_error"
	ErrorModelResourceCodeDoesNotExistConst = "does_not_exist"
	ErrorModelResourceCodeEntitlementEnforcementConst = "entitlement_enforcement"
	ErrorModelResourceCodeFeatureNotEnabledConst = "feature_not_enabled"
	ErrorModelResourceCodeFetchErrorConst = "fetch_error"
	ErrorModelResourceCodeForbiddenConst = "forbidden"
	ErrorModelResourceCodeGovernancePolicyDenialConst = "governance_policy_denial"
	ErrorModelResourceCodeInactiveUserConst = "inactive_user"
	ErrorModelResourceCodeInvalidParameterConst = "invalid_parameter"
	ErrorModelResourceCodeMissingAssetDetailsConst = "missing_asset_details"
	ErrorModelResourceCodeMissingRequiredValueConst = "missing_required_value"
	ErrorModelResourceCodeNotAuthenticatedConst = "not_authenticated"
	ErrorModelResourceCodeNotAuthorizedConst = "not_authorized"
	ErrorModelResourceCodeNotImplementedConst = "not_implemented"
	ErrorModelResourceCodePatchErrorConst = "patch_error"
	ErrorModelResourceCodeRequestBodyErrorConst = "request_body_error"
	ErrorModelResourceCodeTooManyRequestsConst = "too_many_requests"
	ErrorModelResourceCodeUnableToPerformConst = "unable_to_perform"
	ErrorModelResourceCodeUnexpectedExceptionConst = "unexpected_exception"
	ErrorModelResourceCodeUpdateErrorConst = "update_error"
)

// NewErrorModelResource : Instantiate ErrorModelResource (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewErrorModelResource(code string) (_model *ErrorModelResource, err error) {
	_model = &ErrorModelResource{
		Code: core.StringPtr(code),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalErrorModelResource unmarshals an instance of ErrorModelResource from the specified map of raw messages.
func UnmarshalErrorModelResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorModelResource)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		err = core.SDKErrorf(err, "", "code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "extra", &obj.Extra, UnmarshalErrorExtraResource)
	if err != nil {
		err = core.SDKErrorf(err, "", "extra-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		err = core.SDKErrorf(err, "", "more_info-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetContractTemplateOptions : The GetContractTemplate options.
type GetContractTemplateOptions struct {
	// Data Product Contract Template id.
	ContractTemplateID *string `json:"contract_template_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetContractTemplateOptions : Instantiate GetContractTemplateOptions
func (*DataProductHubAPIServiceV1) NewGetContractTemplateOptions(contractTemplateID string, containerID string) *GetContractTemplateOptions {
	return &GetContractTemplateOptions{
		ContractTemplateID: core.StringPtr(contractTemplateID),
		ContainerID: core.StringPtr(containerID),
	}
}

// SetContractTemplateID : Allow user to set ContractTemplateID
func (_options *GetContractTemplateOptions) SetContractTemplateID(contractTemplateID string) *GetContractTemplateOptions {
	_options.ContractTemplateID = core.StringPtr(contractTemplateID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *GetContractTemplateOptions) SetContainerID(containerID string) *GetContractTemplateOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetContractTemplateOptions) SetHeaders(param map[string]string) *GetContractTemplateOptions {
	options.Headers = param
	return options
}

// GetDataProductByDomainOptions : The GetDataProductByDomain options.
type GetDataProductByDomainOptions struct {
	// Domain id.
	DomainID *string `json:"domain_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDataProductByDomainOptions : Instantiate GetDataProductByDomainOptions
func (*DataProductHubAPIServiceV1) NewGetDataProductByDomainOptions(domainID string, containerID string) *GetDataProductByDomainOptions {
	return &GetDataProductByDomainOptions{
		DomainID: core.StringPtr(domainID),
		ContainerID: core.StringPtr(containerID),
	}
}

// SetDomainID : Allow user to set DomainID
func (_options *GetDataProductByDomainOptions) SetDomainID(domainID string) *GetDataProductByDomainOptions {
	_options.DomainID = core.StringPtr(domainID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *GetDataProductByDomainOptions) SetContainerID(containerID string) *GetDataProductByDomainOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductByDomainOptions) SetHeaders(param map[string]string) *GetDataProductByDomainOptions {
	options.Headers = param
	return options
}

// GetDataProductDraftContractTermsOptions : The GetDataProductDraftContractTerms options.
type GetDataProductDraftContractTermsOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// The type of the response: application/odcs+yaml or application/json.
	Accept *string `json:"Accept,omitempty"`

	// Set to false to exclude external contract documents (e.g., Terms and Conditions URLs) from the response. By default,
	// these are included.
	IncludeContractDocuments *bool `json:"include_contract_documents,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDataProductDraftContractTermsOptions : Instantiate GetDataProductDraftContractTermsOptions
func (*DataProductHubAPIServiceV1) NewGetDataProductDraftContractTermsOptions(dataProductID string, draftID string, contractTermsID string) *GetDataProductDraftContractTermsOptions {
	return &GetDataProductDraftContractTermsOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetDataProductDraftContractTermsOptions) SetDataProductID(dataProductID string) *GetDataProductDraftContractTermsOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *GetDataProductDraftContractTermsOptions) SetDraftID(draftID string) *GetDataProductDraftContractTermsOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *GetDataProductDraftContractTermsOptions) SetContractTermsID(contractTermsID string) *GetDataProductDraftContractTermsOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *GetDataProductDraftContractTermsOptions) SetAccept(accept string) *GetDataProductDraftContractTermsOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetIncludeContractDocuments : Allow user to set IncludeContractDocuments
func (_options *GetDataProductDraftContractTermsOptions) SetIncludeContractDocuments(includeContractDocuments bool) *GetDataProductDraftContractTermsOptions {
	_options.IncludeContractDocuments = core.BoolPtr(includeContractDocuments)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductDraftContractTermsOptions) SetHeaders(param map[string]string) *GetDataProductDraftContractTermsOptions {
	options.Headers = param
	return options
}

// GetDataProductDraftOptions : The GetDataProductDraft options.
type GetDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDataProductDraftOptions : Instantiate GetDataProductDraftOptions
func (*DataProductHubAPIServiceV1) NewGetDataProductDraftOptions(dataProductID string, draftID string) *GetDataProductDraftOptions {
	return &GetDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDataProductOptions : Instantiate GetDataProductOptions
func (*DataProductHubAPIServiceV1) NewGetDataProductOptions(dataProductID string) *GetDataProductOptions {
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

	// If the value is true, then it will be verfied whether the caller is present in the data access request pre-approved
	// user list.
	CheckCallerApproval *bool `json:"check_caller_approval,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDataProductReleaseOptions : Instantiate GetDataProductReleaseOptions
func (*DataProductHubAPIServiceV1) NewGetDataProductReleaseOptions(dataProductID string, releaseID string) *GetDataProductReleaseOptions {
	return &GetDataProductReleaseOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID: core.StringPtr(releaseID),
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

// SetCheckCallerApproval : Allow user to set CheckCallerApproval
func (_options *GetDataProductReleaseOptions) SetCheckCallerApproval(checkCallerApproval bool) *GetDataProductReleaseOptions {
	_options.CheckCallerApproval = core.BoolPtr(checkCallerApproval)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDataProductReleaseOptions) SetHeaders(param map[string]string) *GetDataProductReleaseOptions {
	options.Headers = param
	return options
}

// GetDomainOptions : The GetDomain options.
type GetDomainOptions struct {
	// Domain id.
	DomainID *string `json:"domain_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDomainOptions : Instantiate GetDomainOptions
func (*DataProductHubAPIServiceV1) NewGetDomainOptions(domainID string) *GetDomainOptions {
	return &GetDomainOptions{
		DomainID: core.StringPtr(domainID),
	}
}

// SetDomainID : Allow user to set DomainID
func (_options *GetDomainOptions) SetDomainID(domainID string) *GetDomainOptions {
	_options.DomainID = core.StringPtr(domainID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDomainOptions) SetHeaders(param map[string]string) *GetDomainOptions {
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDraftContractTermsDocumentOptions : Instantiate GetDraftContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewGetDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string) *GetDraftContractTermsDocumentOptions {
	return &GetDraftContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInitializeStatusOptions : Instantiate GetInitializeStatusOptions
func (*DataProductHubAPIServiceV1) NewGetInitializeStatusOptions() *GetInitializeStatusOptions {
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

// GetPublishedDataProductDraftContractTermsOptions : The GetPublishedDataProductDraftContractTerms options.
type GetPublishedDataProductDraftContractTermsOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// The type of the response: application/odcs+yaml or application/json.
	Accept *string `json:"Accept,omitempty"`

	// Set to false to exclude external contract documents (e.g., Terms and Conditions URLs) from the response. By default,
	// these are included.
	IncludeContractDocuments *bool `json:"include_contract_documents,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPublishedDataProductDraftContractTermsOptions : Instantiate GetPublishedDataProductDraftContractTermsOptions
func (*DataProductHubAPIServiceV1) NewGetPublishedDataProductDraftContractTermsOptions(dataProductID string, releaseID string, contractTermsID string) *GetPublishedDataProductDraftContractTermsOptions {
	return &GetPublishedDataProductDraftContractTermsOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID: core.StringPtr(releaseID),
		ContractTermsID: core.StringPtr(contractTermsID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *GetPublishedDataProductDraftContractTermsOptions) SetDataProductID(dataProductID string) *GetPublishedDataProductDraftContractTermsOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetReleaseID : Allow user to set ReleaseID
func (_options *GetPublishedDataProductDraftContractTermsOptions) SetReleaseID(releaseID string) *GetPublishedDataProductDraftContractTermsOptions {
	_options.ReleaseID = core.StringPtr(releaseID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *GetPublishedDataProductDraftContractTermsOptions) SetContractTermsID(contractTermsID string) *GetPublishedDataProductDraftContractTermsOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *GetPublishedDataProductDraftContractTermsOptions) SetAccept(accept string) *GetPublishedDataProductDraftContractTermsOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetIncludeContractDocuments : Allow user to set IncludeContractDocuments
func (_options *GetPublishedDataProductDraftContractTermsOptions) SetIncludeContractDocuments(includeContractDocuments bool) *GetPublishedDataProductDraftContractTermsOptions {
	_options.IncludeContractDocuments = core.BoolPtr(includeContractDocuments)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPublishedDataProductDraftContractTermsOptions) SetHeaders(param map[string]string) *GetPublishedDataProductDraftContractTermsOptions {
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetReleaseContractTermsDocumentOptions : Instantiate GetReleaseContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewGetReleaseContractTermsDocumentOptions(dataProductID string, releaseID string, contractTermsID string, documentID string) *GetReleaseContractTermsDocumentOptions {
	return &GetReleaseContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID: core.StringPtr(releaseID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
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

// GetS3BucketValidationOptions : The GetS3BucketValidation options.
type GetS3BucketValidationOptions struct {
	// Name of the bucket to validate.
	BucketName *string `json:"bucket_name" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetS3BucketValidationOptions : Instantiate GetS3BucketValidationOptions
func (*DataProductHubAPIServiceV1) NewGetS3BucketValidationOptions(bucketName string) *GetS3BucketValidationOptions {
	return &GetS3BucketValidationOptions{
		BucketName: core.StringPtr(bucketName),
	}
}

// SetBucketName : Allow user to set BucketName
func (_options *GetS3BucketValidationOptions) SetBucketName(bucketName string) *GetS3BucketValidationOptions {
	_options.BucketName = core.StringPtr(bucketName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetS3BucketValidationOptions) SetHeaders(param map[string]string) *GetS3BucketValidationOptions {
	options.Headers = param
	return options
}

// GetServiceIDCredentialsOptions : The GetServiceIDCredentials options.
type GetServiceIDCredentialsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetServiceIDCredentialsOptions : Instantiate GetServiceIDCredentialsOptions
func (*DataProductHubAPIServiceV1) NewGetServiceIDCredentialsOptions() *GetServiceIDCredentialsOptions {
	return &GetServiceIDCredentialsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceIDCredentialsOptions) SetHeaders(param map[string]string) *GetServiceIDCredentialsOptions {
	options.Headers = param
	return options
}

// InitializeOptions : The Initialize options.
type InitializeOptions struct {
	// Container reference.
	Container *ContainerReference `json:"container,omitempty"`

	// List of configuration options to (re-)initialize.
	Include []string `json:"include,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the InitializeOptions.Include property.
const (
	InitializeOptionsIncludeCatalogConfigurationsConst = "catalog_configurations"
	InitializeOptionsIncludeDataProductSamplesConst = "data_product_samples"
	InitializeOptionsIncludeDeliveryMethodsConst = "delivery_methods"
	InitializeOptionsIncludeDomainsMultiIndustryConst = "domains_multi_industry"
	InitializeOptionsIncludeFunctionalAdminUserGroupConst = "functional_admin_user_group"
	InitializeOptionsIncludeProjectConst = "project"
	InitializeOptionsIncludeWorkflowsConst = "workflows"
)

// NewInitializeOptions : Instantiate InitializeOptions
func (*DataProductHubAPIServiceV1) NewInitializeOptions() *InitializeOptions {
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
	Status *string `json:"status" validate:"required"`

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

	// Resource defining provided workflow definitions.
	Workflows *ProvidedCatalogWorkflows `json:"workflows,omitempty"`
}

// Constants associated with the InitializeResource.Status property.
// Status of the initialize operation.
const (
	InitializeResourceStatusFailedConst = "failed"
	InitializeResourceStatusInProgressConst = "in_progress"
	InitializeResourceStatusNotStartedConst = "not_started"
	InitializeResourceStatusSucceededConst = "succeeded"
)

// UnmarshalInitializeResource unmarshals an instance of InitializeResource from the specified map of raw messages.
func UnmarshalInitializeResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializeResource)
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		err = core.SDKErrorf(err, "", "trace-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModelResource)
	if err != nil {
		err = core.SDKErrorf(err, "", "errors-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_started_at", &obj.LastStartedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_started_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_finished_at", &obj.LastFinishedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_finished_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "initialized_options", &obj.InitializedOptions, UnmarshalInitializedOption)
	if err != nil {
		err = core.SDKErrorf(err, "", "initialized_options-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "workflows", &obj.Workflows, UnmarshalProvidedCatalogWorkflows)
	if err != nil {
		err = core.SDKErrorf(err, "", "workflows-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InitializeSubDomain : The subdomain for a data product domain.
type InitializeSubDomain struct {
	// The name of the data product subdomain.
	Name *string `json:"name,omitempty"`

	// The identifier of the data product subdomain.
	ID *string `json:"id,omitempty"`

	// The description of the data product subdomain.
	Description *string `json:"description,omitempty"`
}

// UnmarshalInitializeSubDomain unmarshals an instance of InitializeSubDomain from the specified map of raw messages.
func UnmarshalInitializeSubDomain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InitializeSubDomain)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		err = core.SDKErrorf(err, "", "description-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
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
	JSONPatchOperationOpAddConst = "add"
	JSONPatchOperationOpCopyConst = "copy"
	JSONPatchOperationOpMoveConst = "move"
	JSONPatchOperationOpRemoveConst = "remove"
	JSONPatchOperationOpReplaceConst = "replace"
	JSONPatchOperationOpTestConst = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		err = core.SDKErrorf(err, "", "op-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		err = core.SDKErrorf(err, "", "path-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		err = core.SDKErrorf(err, "", "from-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDataProductContractTemplateOptions : The ListDataProductContractTemplate options.
type ListDataProductContractTemplateOptions struct {
	// Container ID of the data product catalog. If not supplied, the data product catalog is looked up by using the uid of
	// the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Name of the data product contract template. If not supplied, the data product templates within the catalog will
	// returned.
	ContractTemplateName *string `json:"contract_template.name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDataProductContractTemplateOptions : Instantiate ListDataProductContractTemplateOptions
func (*DataProductHubAPIServiceV1) NewListDataProductContractTemplateOptions() *ListDataProductContractTemplateOptions {
	return &ListDataProductContractTemplateOptions{}
}

// SetContainerID : Allow user to set ContainerID
func (_options *ListDataProductContractTemplateOptions) SetContainerID(containerID string) *ListDataProductContractTemplateOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetContractTemplateName : Allow user to set ContractTemplateName
func (_options *ListDataProductContractTemplateOptions) SetContractTemplateName(contractTemplateName string) *ListDataProductContractTemplateOptions {
	_options.ContractTemplateName = core.StringPtr(contractTemplateName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductContractTemplateOptions) SetHeaders(param map[string]string) *ListDataProductContractTemplateOptions {
	options.Headers = param
	return options
}

// ListDataProductDomainsOptions : The ListDataProductDomains options.
type ListDataProductDomainsOptions struct {
	// Container ID of the data product catalog. If not supplied, the data product catalog is looked up by using the uid of
	// the default data product catalog.
	ContainerID *string `json:"container.id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDataProductDomainsOptions : Instantiate ListDataProductDomainsOptions
func (*DataProductHubAPIServiceV1) NewListDataProductDomainsOptions() *ListDataProductDomainsOptions {
	return &ListDataProductDomainsOptions{}
}

// SetContainerID : Allow user to set ContainerID
func (_options *ListDataProductDomainsOptions) SetContainerID(containerID string) *ListDataProductDomainsOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDataProductDomainsOptions) SetHeaders(param map[string]string) *ListDataProductDomainsOptions {
	options.Headers = param
	return options
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDataProductDraftsOptions : Instantiate ListDataProductDraftsOptions
func (*DataProductHubAPIServiceV1) NewListDataProductDraftsOptions(dataProductID string) *ListDataProductDraftsOptions {
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the ListDataProductReleasesOptions.State property.
const (
	ListDataProductReleasesOptionsStateAvailableConst = "available"
	ListDataProductReleasesOptionsStateRetiredConst = "retired"
)

// NewListDataProductReleasesOptions : Instantiate ListDataProductReleasesOptions
func (*DataProductHubAPIServiceV1) NewListDataProductReleasesOptions(dataProductID string) *ListDataProductReleasesOptions {
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDataProductsOptions : Instantiate ListDataProductsOptions
func (*DataProductHubAPIServiceV1) NewListDataProductsOptions() *ListDataProductsOptions {
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

// ManageAPIKeysOptions : The ManageAPIKeys options.
type ManageAPIKeysOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewManageAPIKeysOptions : Instantiate ManageAPIKeysOptions
func (*DataProductHubAPIServiceV1) NewManageAPIKeysOptions() *ManageAPIKeysOptions {
	return &ManageAPIKeysOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ManageAPIKeysOptions) SetHeaders(param map[string]string) *ManageAPIKeysOptions {
	options.Headers = param
	return options
}

// MemberRolesSchema : Member roles of a corresponding asset.
type MemberRolesSchema struct {
	// User id.
	UserIamID *string `json:"user_iam_id,omitempty"`

	// Roles of the given user.
	Roles []string `json:"roles,omitempty"`
}

// UnmarshalMemberRolesSchema unmarshals an instance of MemberRolesSchema from the specified map of raw messages.
func UnmarshalMemberRolesSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberRolesSchema)
	err = core.UnmarshalPrimitive(m, "user_iam_id", &obj.UserIamID)
	if err != nil {
		err = core.SDKErrorf(err, "", "user_iam_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "roles", &obj.Roles)
	if err != nil {
		err = core.SDKErrorf(err, "", "roles-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		err = core.SDKErrorf(err, "", "start-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Overview : Overview details of a data contract.
type Overview struct {
	// The API version of the contract.
	APIVersion *string `json:"api_version,omitempty"`

	// The kind of contract.
	Kind *string `json:"kind,omitempty"`

	// The name of the contract.
	Name *string `json:"name,omitempty"`

	// The version of the contract.
	Version *string `json:"version" validate:"required"`

	// Domain that the data product version belongs to. If this is the first version of a data product, this field is
	// required. If this is a new version of an existing data product, the domain will default to the domain of the
	// previous version of the data product.
	Domain *Domain `json:"domain" validate:"required"`

	// Additional information links about the contract.
	MoreInfo *string `json:"more_info,omitempty"`
}

// NewOverview : Instantiate Overview (Generic Model Constructor)
func (*DataProductHubAPIServiceV1) NewOverview(version string, domain *Domain) (_model *Overview, err error) {
	_model = &Overview{
		Version: core.StringPtr(version),
		Domain: domain,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalOverview unmarshals an instance of Overview from the specified map of raw messages.
func UnmarshalOverview(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Overview)
	err = core.UnmarshalPrimitive(m, "api_version", &obj.APIVersion)
	if err != nil {
		err = core.SDKErrorf(err, "", "api_version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		err = core.SDKErrorf(err, "", "kind-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "domain", &obj.Domain, UnmarshalDomain)
	if err != nil {
		err = core.SDKErrorf(err, "", "domain-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		err = core.SDKErrorf(err, "", "more_info-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Pricing : Represents the pricing details of the contract.
type Pricing struct {
	// The amount for the contract pricing.
	Amount *string `json:"amount,omitempty"`

	// The currency for the pricing amount.
	Currency *string `json:"currency,omitempty"`

	// The unit associated with the pricing.
	Unit *string `json:"unit,omitempty"`
}

// UnmarshalPricing unmarshals an instance of Pricing from the specified map of raw messages.
func UnmarshalPricing(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pricing)
	err = core.UnmarshalPrimitive(m, "amount", &obj.Amount)
	if err != nil {
		err = core.SDKErrorf(err, "", "amount-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		err = core.SDKErrorf(err, "", "currency-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		err = core.SDKErrorf(err, "", "unit-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProducerInputModel : Parameters for delivery that are set by a data product producer.
type ProducerInputModel struct {
	// Engine details as defined by the data product producer.
	EngineDetails *EngineDetailsModel `json:"engine_details,omitempty"`
}

// UnmarshalProducerInputModel unmarshals an instance of ProducerInputModel from the specified map of raw messages.
func UnmarshalProducerInputModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProducerInputModel)
	err = core.UnmarshalModel(m, "engine_details", &obj.EngineDetails, UnmarshalEngineDetailsModel)
	if err != nil {
		err = core.SDKErrorf(err, "", "engine_details-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PropertiesSchema : Properties of the corresponding asset.
type PropertiesSchema struct {
	// Value of the property object.
	Value *string `json:"value,omitempty"`
}

// UnmarshalPropertiesSchema unmarshals an instance of PropertiesSchema from the specified map of raw messages.
func UnmarshalPropertiesSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PropertiesSchema)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProvidedCatalogWorkflows : Resource defining provided workflow definitions.
type ProvidedCatalogWorkflows struct {
	// A reference to a workflow definition.
	DataAccess *ProvidedWorkflowResource `json:"data_access,omitempty"`

	// A reference to a workflow definition.
	RequestNewProduct *ProvidedWorkflowResource `json:"request_new_product,omitempty"`
}

// UnmarshalProvidedCatalogWorkflows unmarshals an instance of ProvidedCatalogWorkflows from the specified map of raw messages.
func UnmarshalProvidedCatalogWorkflows(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProvidedCatalogWorkflows)
	err = core.UnmarshalModel(m, "data_access", &obj.DataAccess, UnmarshalProvidedWorkflowResource)
	if err != nil {
		err = core.SDKErrorf(err, "", "data_access-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "request_new_product", &obj.RequestNewProduct, UnmarshalProvidedWorkflowResource)
	if err != nil {
		err = core.SDKErrorf(err, "", "request_new_product-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProvidedWorkflowResource : A reference to a workflow definition.
type ProvidedWorkflowResource struct {
	// Reference to a workflow definition.
	Definition *WorkflowDefinitionReference `json:"definition,omitempty"`
}

// UnmarshalProvidedWorkflowResource unmarshals an instance of ProvidedWorkflowResource from the specified map of raw messages.
func UnmarshalProvidedWorkflowResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProvidedWorkflowResource)
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalWorkflowDefinitionReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "definition-error", common.GetComponentInfo())
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewPublishDataProductDraftOptions : Instantiate PublishDataProductDraftOptions
func (*DataProductHubAPIServiceV1) NewPublishDataProductDraftOptions(dataProductID string, draftID string) *PublishDataProductDraftOptions {
	return &PublishDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
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

// ReinitiateDataAssetVisualizationOptions : The ReinitiateDataAssetVisualization options.
type ReinitiateDataAssetVisualizationOptions struct {
	// Data product hub asset and it's related part asset.
	Assets []DataAssetRelationship `json:"assets,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReinitiateDataAssetVisualizationOptions : Instantiate ReinitiateDataAssetVisualizationOptions
func (*DataProductHubAPIServiceV1) NewReinitiateDataAssetVisualizationOptions() *ReinitiateDataAssetVisualizationOptions {
	return &ReinitiateDataAssetVisualizationOptions{}
}

// SetAssets : Allow user to set Assets
func (_options *ReinitiateDataAssetVisualizationOptions) SetAssets(assets []DataAssetRelationship) *ReinitiateDataAssetVisualizationOptions {
	_options.Assets = assets
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReinitiateDataAssetVisualizationOptions) SetHeaders(param map[string]string) *ReinitiateDataAssetVisualizationOptions {
	options.Headers = param
	return options
}

// ReplaceDataProductDraftContractTermsOptions : The ReplaceDataProductDraftContractTerms options.
type ReplaceDataProductDraftContractTermsOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// The reference schema for a asset in a container.
	Asset *AssetReference `json:"asset,omitempty"`

	// ID of the contract terms.
	ID *string `json:"id,omitempty"`

	// Collection of contract terms documents.
	Documents []ContractTermsDocument `json:"documents,omitempty"`

	// An error message, if existing, relating to the contract terms.
	ErrorMsg *string `json:"error_msg,omitempty"`

	// Overview details of a data contract.
	Overview *Overview `json:"overview,omitempty"`

	// Description details of a data contract.
	Description *Description `json:"description,omitempty"`

	// List of sub domains to be added within a domain.
	Organization []ContractTemplateOrganization `json:"organization,omitempty"`

	// List of roles associated with the contract.
	Roles []Roles `json:"roles,omitempty"`

	// Represents the pricing details of the contract.
	Price *Pricing `json:"price,omitempty"`

	// Service Level Agreement details.
	SLA []ContractTemplateSLA `json:"sla,omitempty"`

	// Support and communication details for the contract.
	SupportAndCommunication []ContractTemplateSupportAndCommunication `json:"support_and_communication,omitempty"`

	// Custom properties that are not part of the standard contract.
	CustomProperties []ContractTemplateCustomProperty `json:"custom_properties,omitempty"`

	// Contains the contract test status and related metadata.
	ContractTest *ContractTest `json:"contract_test,omitempty"`

	// Schema details of the data asset.
	Schema []ContractSchema `json:"schema,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceDataProductDraftContractTermsOptions : Instantiate ReplaceDataProductDraftContractTermsOptions
func (*DataProductHubAPIServiceV1) NewReplaceDataProductDraftContractTermsOptions(dataProductID string, draftID string, contractTermsID string) *ReplaceDataProductDraftContractTermsOptions {
	return &ReplaceDataProductDraftContractTermsOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *ReplaceDataProductDraftContractTermsOptions) SetDataProductID(dataProductID string) *ReplaceDataProductDraftContractTermsOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *ReplaceDataProductDraftContractTermsOptions) SetDraftID(draftID string) *ReplaceDataProductDraftContractTermsOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *ReplaceDataProductDraftContractTermsOptions) SetContractTermsID(contractTermsID string) *ReplaceDataProductDraftContractTermsOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetAsset : Allow user to set Asset
func (_options *ReplaceDataProductDraftContractTermsOptions) SetAsset(asset *AssetReference) *ReplaceDataProductDraftContractTermsOptions {
	_options.Asset = asset
	return _options
}

// SetID : Allow user to set ID
func (_options *ReplaceDataProductDraftContractTermsOptions) SetID(id string) *ReplaceDataProductDraftContractTermsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDocuments : Allow user to set Documents
func (_options *ReplaceDataProductDraftContractTermsOptions) SetDocuments(documents []ContractTermsDocument) *ReplaceDataProductDraftContractTermsOptions {
	_options.Documents = documents
	return _options
}

// SetErrorMsg : Allow user to set ErrorMsg
func (_options *ReplaceDataProductDraftContractTermsOptions) SetErrorMsg(errorMsg string) *ReplaceDataProductDraftContractTermsOptions {
	_options.ErrorMsg = core.StringPtr(errorMsg)
	return _options
}

// SetOverview : Allow user to set Overview
func (_options *ReplaceDataProductDraftContractTermsOptions) SetOverview(overview *Overview) *ReplaceDataProductDraftContractTermsOptions {
	_options.Overview = overview
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ReplaceDataProductDraftContractTermsOptions) SetDescription(description *Description) *ReplaceDataProductDraftContractTermsOptions {
	_options.Description = description
	return _options
}

// SetOrganization : Allow user to set Organization
func (_options *ReplaceDataProductDraftContractTermsOptions) SetOrganization(organization []ContractTemplateOrganization) *ReplaceDataProductDraftContractTermsOptions {
	_options.Organization = organization
	return _options
}

// SetRoles : Allow user to set Roles
func (_options *ReplaceDataProductDraftContractTermsOptions) SetRoles(roles []Roles) *ReplaceDataProductDraftContractTermsOptions {
	_options.Roles = roles
	return _options
}

// SetPrice : Allow user to set Price
func (_options *ReplaceDataProductDraftContractTermsOptions) SetPrice(price *Pricing) *ReplaceDataProductDraftContractTermsOptions {
	_options.Price = price
	return _options
}

// SetSLA : Allow user to set SLA
func (_options *ReplaceDataProductDraftContractTermsOptions) SetSLA(sla []ContractTemplateSLA) *ReplaceDataProductDraftContractTermsOptions {
	_options.SLA = sla
	return _options
}

// SetSupportAndCommunication : Allow user to set SupportAndCommunication
func (_options *ReplaceDataProductDraftContractTermsOptions) SetSupportAndCommunication(supportAndCommunication []ContractTemplateSupportAndCommunication) *ReplaceDataProductDraftContractTermsOptions {
	_options.SupportAndCommunication = supportAndCommunication
	return _options
}

// SetCustomProperties : Allow user to set CustomProperties
func (_options *ReplaceDataProductDraftContractTermsOptions) SetCustomProperties(customProperties []ContractTemplateCustomProperty) *ReplaceDataProductDraftContractTermsOptions {
	_options.CustomProperties = customProperties
	return _options
}

// SetContractTest : Allow user to set ContractTest
func (_options *ReplaceDataProductDraftContractTermsOptions) SetContractTest(contractTest *ContractTest) *ReplaceDataProductDraftContractTermsOptions {
	_options.ContractTest = contractTest
	return _options
}

// SetSchema : Allow user to set Schema
func (_options *ReplaceDataProductDraftContractTermsOptions) SetSchema(schema []ContractSchema) *ReplaceDataProductDraftContractTermsOptions {
	_options.Schema = schema
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceDataProductDraftContractTermsOptions) SetHeaders(param map[string]string) *ReplaceDataProductDraftContractTermsOptions {
	options.Headers = param
	return options
}

// RetireDataProductReleaseOptions : The RetireDataProductRelease options.
type RetireDataProductReleaseOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product release id.
	ReleaseID *string `json:"release_id" validate:"required,ne="`

	// Revoke's Access from all the Subscriptions of the Data Product. No user's can able to see the subscribed assets
	// anymore.
	RevokeAccess *bool `json:"revoke_access,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewRetireDataProductReleaseOptions : Instantiate RetireDataProductReleaseOptions
func (*DataProductHubAPIServiceV1) NewRetireDataProductReleaseOptions(dataProductID string, releaseID string) *RetireDataProductReleaseOptions {
	return &RetireDataProductReleaseOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID: core.StringPtr(releaseID),
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

// SetRevokeAccess : Allow user to set RevokeAccess
func (_options *RetireDataProductReleaseOptions) SetRevokeAccess(revokeAccess bool) *RetireDataProductReleaseOptions {
	_options.RevokeAccess = core.BoolPtr(revokeAccess)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RetireDataProductReleaseOptions) SetHeaders(param map[string]string) *RetireDataProductReleaseOptions {
	options.Headers = param
	return options
}

// Roles : Represents a role associated with the contract.
type Roles struct {
	// The role associated with the contract.
	Role *string `json:"role,omitempty"`
}

// UnmarshalRoles unmarshals an instance of Roles from the specified map of raw messages.
func UnmarshalRoles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Roles)
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		err = core.SDKErrorf(err, "", "role-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceIDCredentials : Service id credentials.
type ServiceIDCredentials struct {
	// Name of the api key of the service id.
	Name *string `json:"name,omitempty"`

	// Created date of the api key of the service id.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// UnmarshalServiceIDCredentials unmarshals an instance of ServiceIDCredentials from the specified map of raw messages.
func UnmarshalServiceIDCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceIDCredentials)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateDataProductContractTemplateOptions : The UpdateDataProductContractTemplate options.
type UpdateDataProductContractTemplateOptions struct {
	// Data Product Contract Template id.
	ContractTemplateID *string `json:"contract_template_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDataProductContractTemplateOptions : Instantiate UpdateDataProductContractTemplateOptions
func (*DataProductHubAPIServiceV1) NewUpdateDataProductContractTemplateOptions(contractTemplateID string, containerID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductContractTemplateOptions {
	return &UpdateDataProductContractTemplateOptions{
		ContractTemplateID: core.StringPtr(contractTemplateID),
		ContainerID: core.StringPtr(containerID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetContractTemplateID : Allow user to set ContractTemplateID
func (_options *UpdateDataProductContractTemplateOptions) SetContractTemplateID(contractTemplateID string) *UpdateDataProductContractTemplateOptions {
	_options.ContractTemplateID = core.StringPtr(contractTemplateID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *UpdateDataProductContractTemplateOptions) SetContainerID(containerID string) *UpdateDataProductContractTemplateOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductContractTemplateOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductContractTemplateOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductContractTemplateOptions) SetHeaders(param map[string]string) *UpdateDataProductContractTemplateOptions {
	options.Headers = param
	return options
}

// UpdateDataProductDomainOptions : The UpdateDataProductDomain options.
type UpdateDataProductDomainOptions struct {
	// Domain id.
	DomainID *string `json:"domain_id" validate:"required,ne="`

	// Container ID of the data product catalog.
	ContainerID *string `json:"container.id" validate:"required"`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDataProductDomainOptions : Instantiate UpdateDataProductDomainOptions
func (*DataProductHubAPIServiceV1) NewUpdateDataProductDomainOptions(domainID string, containerID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDomainOptions {
	return &UpdateDataProductDomainOptions{
		DomainID: core.StringPtr(domainID),
		ContainerID: core.StringPtr(containerID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDomainID : Allow user to set DomainID
func (_options *UpdateDataProductDomainOptions) SetDomainID(domainID string) *UpdateDataProductDomainOptions {
	_options.DomainID = core.StringPtr(domainID)
	return _options
}

// SetContainerID : Allow user to set ContainerID
func (_options *UpdateDataProductDomainOptions) SetContainerID(containerID string) *UpdateDataProductDomainOptions {
	_options.ContainerID = core.StringPtr(containerID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductDomainOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDomainOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductDomainOptions) SetHeaders(param map[string]string) *UpdateDataProductDomainOptions {
	options.Headers = param
	return options
}

// UpdateDataProductDraftContractTermsOptions : The UpdateDataProductDraftContractTerms options.
type UpdateDataProductDraftContractTermsOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// Contract terms id.
	ContractTermsID *string `json:"contract_terms_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDataProductDraftContractTermsOptions : Instantiate UpdateDataProductDraftContractTermsOptions
func (*DataProductHubAPIServiceV1) NewUpdateDataProductDraftContractTermsOptions(dataProductID string, draftID string, contractTermsID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDraftContractTermsOptions {
	return &UpdateDataProductDraftContractTermsOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		JSONPatchInstructions: jsonPatchInstructions,
	}
}

// SetDataProductID : Allow user to set DataProductID
func (_options *UpdateDataProductDraftContractTermsOptions) SetDataProductID(dataProductID string) *UpdateDataProductDraftContractTermsOptions {
	_options.DataProductID = core.StringPtr(dataProductID)
	return _options
}

// SetDraftID : Allow user to set DraftID
func (_options *UpdateDataProductDraftContractTermsOptions) SetDraftID(draftID string) *UpdateDataProductDraftContractTermsOptions {
	_options.DraftID = core.StringPtr(draftID)
	return _options
}

// SetContractTermsID : Allow user to set ContractTermsID
func (_options *UpdateDataProductDraftContractTermsOptions) SetContractTermsID(contractTermsID string) *UpdateDataProductDraftContractTermsOptions {
	_options.ContractTermsID = core.StringPtr(contractTermsID)
	return _options
}

// SetJSONPatchInstructions : Allow user to set JSONPatchInstructions
func (_options *UpdateDataProductDraftContractTermsOptions) SetJSONPatchInstructions(jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDraftContractTermsOptions {
	_options.JSONPatchInstructions = jsonPatchInstructions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDataProductDraftContractTermsOptions) SetHeaders(param map[string]string) *UpdateDataProductDraftContractTermsOptions {
	options.Headers = param
	return options
}

// UpdateDataProductDraftOptions : The UpdateDataProductDraft options.
type UpdateDataProductDraftOptions struct {
	// Data product ID. Use '-' to skip specifying the data product ID explicitly.
	DataProductID *string `json:"data_product_id" validate:"required,ne="`

	// Data product draft id.
	DraftID *string `json:"draft_id" validate:"required,ne="`

	// A set of patch operations as defined in RFC 6902. See http://jsonpatch.com/ for more information.
	JSONPatchInstructions []JSONPatchOperation `json:"jsonPatchInstructions" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDataProductDraftOptions : Instantiate UpdateDataProductDraftOptions
func (*DataProductHubAPIServiceV1) NewUpdateDataProductDraftOptions(dataProductID string, draftID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductDraftOptions {
	return &UpdateDataProductDraftOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDataProductReleaseOptions : Instantiate UpdateDataProductReleaseOptions
func (*DataProductHubAPIServiceV1) NewUpdateDataProductReleaseOptions(dataProductID string, releaseID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDataProductReleaseOptions {
	return &UpdateDataProductReleaseOptions{
		DataProductID: core.StringPtr(dataProductID),
		ReleaseID: core.StringPtr(releaseID),
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

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateDraftContractTermsDocumentOptions : Instantiate UpdateDraftContractTermsDocumentOptions
func (*DataProductHubAPIServiceV1) NewUpdateDraftContractTermsDocumentOptions(dataProductID string, draftID string, contractTermsID string, documentID string, jsonPatchInstructions []JSONPatchOperation) *UpdateDraftContractTermsDocumentOptions {
	return &UpdateDraftContractTermsDocumentOptions{
		DataProductID: core.StringPtr(dataProductID),
		DraftID: core.StringPtr(draftID),
		ContractTermsID: core.StringPtr(contractTermsID),
		DocumentID: core.StringPtr(documentID),
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
func (*DataProductHubAPIServiceV1) NewUseCase(id string) (_model *UseCase, err error) {
	_model = &UseCase{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalUseCase unmarshals an instance of UseCase from the specified map of raw messages.
func UnmarshalUseCase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UseCase)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "container", &obj.Container, UnmarshalContainerReference)
	if err != nil {
		err = core.SDKErrorf(err, "", "container-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Visualization : Data members for visualization.
type Visualization struct {
	// Visualization identifier.
	ID *string `json:"id,omitempty"`

	// Visualization name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalVisualization unmarshals an instance of Visualization from the specified map of raw messages.
func UnmarshalVisualization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Visualization)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WorkflowDefinitionReference : Reference to a workflow definition.
type WorkflowDefinitionReference struct {
	// ID of a workflow definition.
	ID *string `json:"id,omitempty"`
}

// UnmarshalWorkflowDefinitionReference unmarshals an instance of WorkflowDefinitionReference from the specified map of raw messages.
func UnmarshalWorkflowDefinitionReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WorkflowDefinitionReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// DataProductsPager can be used to simplify the use of the "ListDataProducts" method.
//
type DataProductsPager struct {
	hasNext bool
	options *ListDataProductsOptions
	client  *DataProductHubAPIServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductsPager returns a new DataProductsPager instance.
func (dataProductHubApiService *DataProductHubAPIServiceV1) NewDataProductsPager(options *ListDataProductsOptions) (pager *DataProductsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
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

//
// DataProductDraftsPager can be used to simplify the use of the "ListDataProductDrafts" method.
//
type DataProductDraftsPager struct {
	hasNext bool
	options *ListDataProductDraftsOptions
	client  *DataProductHubAPIServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductDraftsPager returns a new DataProductDraftsPager instance.
func (dataProductHubApiService *DataProductHubAPIServiceV1) NewDataProductDraftsPager(options *ListDataProductDraftsOptions) (pager *DataProductDraftsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
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
func (pager *DataProductDraftsPager) GetNextWithContext(ctx context.Context) (page []DataProductDraftSummary, err error) {
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
func (pager *DataProductDraftsPager) GetAllWithContext(ctx context.Context) (allItems []DataProductDraftSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductDraftSummary
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
func (pager *DataProductDraftsPager) GetNext() (page []DataProductDraftSummary, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductDraftsPager) GetAll() (allItems []DataProductDraftSummary, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

//
// DataProductReleasesPager can be used to simplify the use of the "ListDataProductReleases" method.
//
type DataProductReleasesPager struct {
	hasNext bool
	options *ListDataProductReleasesOptions
	client  *DataProductHubAPIServiceV1
	pageContext struct {
		next *string
	}
}

// NewDataProductReleasesPager returns a new DataProductReleasesPager instance.
func (dataProductHubApiService *DataProductHubAPIServiceV1) NewDataProductReleasesPager(options *ListDataProductReleasesOptions) (pager *DataProductReleasesPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
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
func (pager *DataProductReleasesPager) GetNextWithContext(ctx context.Context) (page []DataProductReleaseSummary, err error) {
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
func (pager *DataProductReleasesPager) GetAllWithContext(ctx context.Context) (allItems []DataProductReleaseSummary, err error) {
	for pager.HasNext() {
		var nextPage []DataProductReleaseSummary
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
func (pager *DataProductReleasesPager) GetNext() (page []DataProductReleaseSummary, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DataProductReleasesPager) GetAll() (allItems []DataProductReleaseSummary, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
