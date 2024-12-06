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

package dataproducthubapiservicev1_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/data-product-exchange-go-sdk/dataproducthubapiservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`DataProductHubAPIServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(dataProductHubAPIService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(dataProductHubAPIService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
				URL: "https://dataproducthubapiservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(dataProductHubAPIService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_HUB_API_SERVICE_URL":       "https://dataproducthubapiservicev1/api",
				"DATA_PRODUCT_HUB_API_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{})
				Expect(dataProductHubAPIService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := dataProductHubAPIService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductHubAPIService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductHubAPIService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductHubAPIService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(dataProductHubAPIService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dataProductHubAPIService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductHubAPIService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductHubAPIService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductHubAPIService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{})
				err := dataProductHubAPIService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := dataProductHubAPIService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != dataProductHubAPIService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(dataProductHubAPIService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(dataProductHubAPIService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_HUB_API_SERVICE_URL":       "https://dataproducthubapiservicev1/api",
				"DATA_PRODUCT_HUB_API_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(dataProductHubAPIService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"DATA_PRODUCT_HUB_API_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(dataProductHubAPIService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = dataproducthubapiservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions) - Operation response error`, func() {
		getInitializeStatusPath := "/data_product_exchange/v1/configuration/initialize/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetInitializeStatus with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproducthubapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
		getInitializeStatusPath := "/data_product_exchange/v1/configuration/initialize/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}], "workflows": {"data_access": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}, "request_new_product": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproducthubapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetInitializeStatusWithContext(ctx, getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getInitializeStatusPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["container.id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}], "workflows": {"data_access": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}, "request_new_product": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}}`)
				}))
			})
			It(`Invoke GetInitializeStatus successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetInitializeStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproducthubapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetInitializeStatus with error: Operation request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproducthubapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetInitializeStatus successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := new(dataproducthubapiservicev1.GetInitializeStatusOptions)
				getInitializeStatusOptionsModel.ContainerID = core.StringPtr("testString")
				getInitializeStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceIDCredentials(getServiceIDCredentialsOptions *GetServiceIDCredentialsOptions) - Operation response error`, func() {
		getServiceIDCredentialsPath := "/data_product_exchange/v1/configuration/credentials"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDCredentialsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceIDCredentials with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := new(dataproducthubapiservicev1.GetServiceIDCredentialsOptions)
				getServiceIDCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServiceIDCredentials(getServiceIDCredentialsOptions *GetServiceIDCredentialsOptions)`, func() {
		getServiceIDCredentialsPath := "/data_product_exchange/v1/configuration/credentials"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "data-product-admin-service-id-API-key", "created_at": "2024-03-15T04:07+0000"}`)
				}))
			})
			It(`Invoke GetServiceIDCredentials successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := new(dataproducthubapiservicev1.GetServiceIDCredentialsOptions)
				getServiceIDCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetServiceIDCredentialsWithContext(ctx, getServiceIDCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetServiceIDCredentialsWithContext(ctx, getServiceIDCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServiceIDCredentialsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "data-product-admin-service-id-API-key", "created_at": "2024-03-15T04:07+0000"}`)
				}))
			})
			It(`Invoke GetServiceIDCredentials successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetServiceIDCredentials(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := new(dataproducthubapiservicev1.GetServiceIDCredentialsOptions)
				getServiceIDCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetServiceIDCredentials with error: Operation request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := new(dataproducthubapiservicev1.GetServiceIDCredentialsOptions)
				getServiceIDCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetServiceIDCredentials successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := new(dataproducthubapiservicev1.GetServiceIDCredentialsOptions)
				getServiceIDCredentialsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Initialize(initializeOptions *InitializeOptions) - Operation response error`, func() {
		initializePath := "/data_product_exchange/v1/configuration/initialize"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Initialize with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproducthubapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Initialize(initializeOptions *InitializeOptions)`, func() {
		initializePath := "/data_product_exchange/v1/configuration/initialize"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}], "workflows": {"data_access": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}, "request_new_product": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}}`)
				}))
			})
			It(`Invoke Initialize successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproducthubapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.InitializeWithContext(ctx, initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.InitializeWithContext(ctx, initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initializePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "href": "https://api.example.com/configuration/initialize/status?catalog_id=d29c42eb-7100-4b7a-8257-c196dbcca1cd", "status": "not_started", "trace": "Trace", "errors": [{"code": "request_body_error", "message": "Message", "extra": {"anyKey": "anyValue"}, "more_info": "MoreInfo"}], "last_started_at": "2023-08-21T15:24:06.021Z", "last_finished_at": "2023-08-21T20:24:34.450Z", "initialized_options": [{"name": "Name", "version": 1}], "workflows": {"data_access": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}, "request_new_product": {"definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}}`)
				}))
			})
			It(`Invoke Initialize successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.Initialize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproducthubapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Initialize with error: Operation request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproducthubapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke Initialize successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := new(dataproducthubapiservicev1.InitializeOptions)
				initializeOptionsModel.Container = containerReferenceModel
				initializeOptionsModel.Include = []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}
				initializeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.Initialize(initializeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ManageAPIKeys(manageAPIKeysOptions *ManageAPIKeysOptions)`, func() {
		manageAPIKeysPath := "/data_product_exchange/v1/configuration/rotate_credentials"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(manageAPIKeysPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke ManageAPIKeys successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dataProductHubAPIService.ManageAPIKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ManageAPIKeysOptions model
				manageAPIKeysOptionsModel := new(dataproducthubapiservicev1.ManageAPIKeysOptions)
				manageAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dataProductHubAPIService.ManageAPIKeys(manageAPIKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke ManageAPIKeys with error: Operation request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ManageAPIKeysOptions model
				manageAPIKeysOptionsModel := new(dataproducthubapiservicev1.ManageAPIKeysOptions)
				manageAPIKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dataProductHubAPIService.ManageAPIKeys(manageAPIKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) - Operation response error`, func() {
		listDataProductsPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProducts with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproducthubapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions)`, func() {
		listDataProductsPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproducthubapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.ListDataProductsWithContext(ctx, listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "data_products": [{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}`)
				}))
			})
			It(`Invoke ListDataProducts successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.ListDataProducts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproducthubapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProducts with error: Operation request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproducthubapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataProducts successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := new(dataproducthubapiservicev1.ListDataProductsOptions)
				listDataProductsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductsOptionsModel.Start = core.StringPtr("testString")
				listDataProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.ListDataProducts(listDataProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductSummaryCollection)
				nextObject := new(dataproducthubapiservicev1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductSummaryCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"data_products":[{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductsPager.GetNext successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductsOptionsModel := &dataproducthubapiservicev1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dataproducthubapiservicev1.DataProductSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductsPager.GetAll successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductsOptionsModel := &dataproducthubapiservicev1.ListDataProductsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductsPager(listDataProductsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions) - Operation response error`, func() {
		createDataProductPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataProduct with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dataproducthubapiservicev1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {
		createDataProductPath := "/data_product_exchange/v1/data_products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke CreateDataProduct successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dataproducthubapiservicev1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.CreateDataProductWithContext(ctx, createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.CreateDataProductWithContext(ctx, createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke CreateDataProduct successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.CreateDataProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dataproducthubapiservicev1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProduct with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dataproducthubapiservicev1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductOptions model with no property values
				createDataProductOptionsModelNew := new(dataproducthubapiservicev1.CreateDataProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDataProduct successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsModel := new(dataproducthubapiservicev1.CreateDataProductOptions)
				createDataProductOptionsModel.Drafts = []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}
				createDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.CreateDataProduct(createDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProduct(getDataProductOptions *GetDataProductOptions) - Operation response error`, func() {
		getDataProductPath := "/data_product_exchange/v1/data_products/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProduct with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproducthubapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
		getDataProductPath := "/data_product_exchange/v1/data_products/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke GetDataProduct successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproducthubapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetDataProductWithContext(ctx, getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "latest_release": {"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke GetDataProduct successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetDataProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproducthubapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProduct with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproducthubapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductOptions model with no property values
				getDataProductOptionsModelNew := new(dataproducthubapiservicev1.GetDataProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.GetDataProduct(getDataProductOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataProduct successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductOptions model
				getDataProductOptionsModel := new(dataproducthubapiservicev1.GetDataProductOptions)
				getDataProductOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetDataProduct(getDataProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions) - Operation response error`, func() {
		completeDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString/complete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions)`, func() {
		completeDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString/complete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocumentWithContext(ctx, completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.CompleteDraftContractTermsDocumentWithContext(ctx, completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteDraftContractTermsDocumentOptions model with no property values
				completeDraftContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CompleteDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				completeDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions)
				completeDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				completeDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) - Operation response error`, func() {
		listDataProductDraftsPath := "/data_product_exchange/v1/data_products/testString/drafts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProductDrafts with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions)`, func() {
		listDataProductDraftsPath := "/data_product_exchange/v1/data_products/testString/drafts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductDrafts successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.ListDataProductDraftsWithContext(ctx, listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.ListDataProductDraftsWithContext(ctx, listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "drafts": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductDrafts successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.ListDataProductDrafts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductDrafts with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDataProductDraftsOptions model with no property values
				listDataProductDraftsOptionsModelNew := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataProductDrafts successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductDraftsOptions model
				listDataProductDraftsOptionsModel := new(dataproducthubapiservicev1.ListDataProductDraftsOptions)
				listDataProductDraftsOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Version = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductDraftsOptionsModel.Start = core.StringPtr("testString")
				listDataProductDraftsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductDraftCollection)
				nextObject := new(dataproducthubapiservicev1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductDraftCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductDraftsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"drafts":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"name":"My Data Product","description":"This is a description of My Data Product.","tags":["Tags"],"use_cases":[{"id":"ID","name":"Name","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}],"types":["data"],"contract_terms":[{"asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"id":"ID","documents":[{"url":"URL","type":"terms_and_conditions","name":"Name","id":"2b0bf220-079c-11ee-be56-0242ac120002","attachment":{"id":"ID"},"upload_url":"UploadURL"}],"error_msg":"ErrorMsg"}],"is_restricted":true,"id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"drafts":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"name":"My Data Product","description":"This is a description of My Data Product.","tags":["Tags"],"use_cases":[{"id":"ID","name":"Name","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}],"types":["data"],"contract_terms":[{"asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"id":"ID","documents":[{"url":"URL","type":"terms_and_conditions","name":"Name","id":"2b0bf220-079c-11ee-be56-0242ac120002","attachment":{"id":"ID"},"upload_url":"UploadURL"}],"error_msg":"ErrorMsg"}],"is_restricted":true,"id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductDraftsPager.GetNext successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductDraftsOptionsModel := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
					DataProductID:    core.StringPtr("testString"),
					AssetContainerID: core.StringPtr("testString"),
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dataproducthubapiservicev1.DataProductVersionSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductDraftsPager.GetAll successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductDraftsOptionsModel := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
					DataProductID:    core.StringPtr("testString"),
					AssetContainerID: core.StringPtr("testString"),
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions) - Operation response error`, func() {
		createDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDataProductDraft with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Asset = assetPrototypeModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.Workflows = dataProductWorkflowsModel
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions)`, func() {
		createDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CreateDataProductDraft successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Asset = assetPrototypeModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.Workflows = dataProductWorkflowsModel
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.CreateDataProductDraftWithContext(ctx, createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.CreateDataProductDraftWithContext(ctx, createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke CreateDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.CreateDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Asset = assetPrototypeModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.Workflows = dataProductWorkflowsModel
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDataProductDraft with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Asset = assetPrototypeModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.Workflows = dataProductWorkflowsModel
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDataProductDraftOptions model with no property values
				createDataProductDraftOptionsModelNew := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the CreateDataProductDraftOptions model
				createDataProductDraftOptionsModel := new(dataproducthubapiservicev1.CreateDataProductDraftOptions)
				createDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Asset = assetPrototypeModel
				createDataProductDraftOptionsModel.Version = core.StringPtr("1.2.0")
				createDataProductDraftOptionsModel.State = core.StringPtr("draft")
				createDataProductDraftOptionsModel.DataProduct = dataProductIdentityModel
				createDataProductDraftOptionsModel.Name = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Description = core.StringPtr("testString")
				createDataProductDraftOptionsModel.Tags = []string{"testString"}
				createDataProductDraftOptionsModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				createDataProductDraftOptionsModel.Types = []string{"data"}
				createDataProductDraftOptionsModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				createDataProductDraftOptionsModel.IsRestricted = core.BoolPtr(true)
				createDataProductDraftOptionsModel.Domain = domainModel
				createDataProductDraftOptionsModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				createDataProductDraftOptionsModel.Workflows = dataProductWorkflowsModel
				createDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions) - Operation response error`, func() {
		createDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
		createDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocumentWithContext(ctx, createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.CreateDraftContractTermsDocumentWithContext(ctx, createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDraftContractTermsDocumentOptions model with no property values
				createDraftContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				createDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions)
				createDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Type = core.StringPtr("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.Name = core.StringPtr("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.URL = core.StringPtr("testString")
				createDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions) - Operation response error`, func() {
		getDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProductDraft with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions)`, func() {
		getDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDataProductDraft successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetDataProductDraftWithContext(ctx, getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetDataProductDraftWithContext(ctx, getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductDraftPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductDraft with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductDraftOptions model with no property values
				getDataProductDraftOptionsModelNew := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductDraftOptions model
				getDataProductDraftOptionsModel := new(dataproducthubapiservicev1.GetDataProductDraftOptions)
				getDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				getDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions)`, func() {
		deleteDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDataProductDraftPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dataProductHubAPIService.DeleteDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDataProductDraftOptions model
				deleteDataProductDraftOptionsModel := new(dataproducthubapiservicev1.DeleteDataProductDraftOptions)
				deleteDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				deleteDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				deleteDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDataProductDraft with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the DeleteDataProductDraftOptions model
				deleteDataProductDraftOptionsModel := new(dataproducthubapiservicev1.DeleteDataProductDraftOptions)
				deleteDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				deleteDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				deleteDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDataProductDraftOptions model with no property values
				deleteDataProductDraftOptionsModelNew := new(dataproducthubapiservicev1.DeleteDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions) - Operation response error`, func() {
		updateDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDataProductDraft with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions)`, func() {
		updateDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateDataProductDraft successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.UpdateDataProductDraftWithContext(ctx, updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.UpdateDataProductDraftWithContext(ctx, updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductDraftPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductDraft with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductDraftOptions model with no property values
				updateDataProductDraftOptionsModelNew := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductDraftOptions model
				updateDataProductDraftOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductDraftOptions)
				updateDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				updateDataProductDraftOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions) - Operation response error`, func() {
		getDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions)`, func() {
		getDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetDraftContractTermsDocumentWithContext(ctx, getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetDraftContractTermsDocumentWithContext(ctx, getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDraftContractTermsDocumentOptions model with no property values
				getDraftContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				getDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions)
				getDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
		deleteDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := dataProductHubAPIService.DeleteDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				deleteDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions)
				deleteDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = dataProductHubAPIService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				deleteDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions)
				deleteDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				deleteDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := dataProductHubAPIService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDraftContractTermsDocumentOptions model with no property values
				deleteDraftContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = dataProductHubAPIService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions) - Operation response error`, func() {
		updateDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions)`, func() {
		updateDraftContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocumentWithContext(ctx, updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.UpdateDraftContractTermsDocumentWithContext(ctx, updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDraftContractTermsDocumentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDraftContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDraftContractTermsDocumentOptions model with no property values
				updateDraftContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDraftContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				updateDraftContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions)
				updateDraftContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DraftID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDraftContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions) - Operation response error`, func() {
		publishDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/publish"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PublishDataProductDraft with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
		publishDataProductDraftPath := "/data_product_exchange/v1/data_products/testString/drafts/testString/publish"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PublishDataProductDraft successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.PublishDataProductDraftWithContext(ctx, publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.PublishDataProductDraftWithContext(ctx, publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishDataProductDraftPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke PublishDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.PublishDataProductDraft(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PublishDataProductDraft with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PublishDataProductDraftOptions model with no property values
				publishDataProductDraftOptionsModelNew := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PublishDataProductDraft successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the PublishDataProductDraftOptions model
				publishDataProductDraftOptionsModel := new(dataproducthubapiservicev1.PublishDataProductDraftOptions)
				publishDataProductDraftOptionsModel.DataProductID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.DraftID = core.StringPtr("testString")
				publishDataProductDraftOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions) - Operation response error`, func() {
		getDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for check_caller_approval query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDataProductRelease with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.CheckCallerApproval = core.BoolPtr(false)
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
		getDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for check_caller_approval query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDataProductRelease successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.CheckCallerApproval = core.BoolPtr(false)
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetDataProductReleaseWithContext(ctx, getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetDataProductReleaseWithContext(ctx, getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDataProductReleasePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for check_caller_approval query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.CheckCallerApproval = core.BoolPtr(false)
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDataProductRelease with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.CheckCallerApproval = core.BoolPtr(false)
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDataProductReleaseOptions model with no property values
				getDataProductReleaseOptionsModelNew := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetDataProductReleaseOptions model
				getDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.GetDataProductReleaseOptions)
				getDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				getDataProductReleaseOptionsModel.CheckCallerApproval = core.BoolPtr(false)
				getDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions) - Operation response error`, func() {
		updateDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDataProductRelease with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions)`, func() {
		updateDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateDataProductRelease successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.UpdateDataProductReleaseWithContext(ctx, updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.UpdateDataProductReleaseWithContext(ctx, updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDataProductReleasePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke UpdateDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDataProductRelease with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDataProductReleaseOptions model with no property values
				updateDataProductReleaseOptionsModelNew := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"

				// Construct an instance of the UpdateDataProductReleaseOptions model
				updateDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.UpdateDataProductReleaseOptions)
				updateDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				updateDataProductReleaseOptionsModel.JSONPatchInstructions = []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions) - Operation response error`, func() {
		getReleaseContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/releases/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions)`, func() {
		getReleaseContractTermsDocumentPath := "/data_product_exchange/v1/data_products/testString/releases/testString/contract_terms/testString/documents/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocumentWithContext(ctx, getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.GetReleaseContractTermsDocumentWithContext(ctx, getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getReleaseContractTermsDocumentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}`)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetReleaseContractTermsDocument with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetReleaseContractTermsDocumentOptions model with no property values
				getReleaseContractTermsDocumentOptionsModelNew := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetReleaseContractTermsDocument successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				getReleaseContractTermsDocumentOptionsModel := new(dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions)
				getReleaseContractTermsDocumentOptionsModel.DataProductID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ReleaseID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.ContractTermsID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.DocumentID = core.StringPtr("testString")
				getReleaseContractTermsDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) - Operation response error`, func() {
		listDataProductReleasesPath := "/data_product_exchange/v1/data_products/testString/releases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDataProductReleases with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions)`, func() {
		listDataProductReleasesPath := "/data_product_exchange/v1/data_products/testString/releases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "releases": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductReleases successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.ListDataProductReleasesWithContext(ctx, listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.ListDataProductReleasesWithContext(ctx, listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["asset.container.id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 200, "first": {"href": "https://api.example.com/collection"}, "next": {"href": "https://api.example.com/collection?start=eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9", "start": "eyJvZmZzZXQiOjAsImRvbmUiOnRydWV9"}, "releases": [{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}}]}`)
				}))
			})
			It(`Invoke ListDataProductReleases successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.ListDataProductReleases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDataProductReleases with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDataProductReleasesOptions model with no property values
				listDataProductReleasesOptionsModelNew := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDataProductReleases successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the ListDataProductReleasesOptions model
				listDataProductReleasesOptionsModel := new(dataproducthubapiservicev1.ListDataProductReleasesOptions)
				listDataProductReleasesOptionsModel.DataProductID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.AssetContainerID = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.State = []string{"available"}
				listDataProductReleasesOptionsModel.Version = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listDataProductReleasesOptionsModel.Start = core.StringPtr("testString")
				listDataProductReleasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductReleaseCollection)
				nextObject := new(dataproducthubapiservicev1.NextPage)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(dataproducthubapiservicev1.DataProductReleaseCollection)

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDataProductReleasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"releases":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"name":"My Data Product","description":"This is a description of My Data Product.","tags":["Tags"],"use_cases":[{"id":"ID","name":"Name","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}],"types":["data"],"contract_terms":[{"asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"id":"ID","documents":[{"url":"URL","type":"terms_and_conditions","name":"Name","id":"2b0bf220-079c-11ee-be56-0242ac120002","attachment":{"id":"ID"},"upload_url":"UploadURL"}],"error_msg":"ErrorMsg"}],"is_restricted":true,"id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"releases":[{"version":"1.0.0","state":"draft","data_product":{"id":"b38df608-d34b-4d58-8136-ed25e6c6684e","release":{"id":"18bdbde1-918e-4ecf-aa23-6727bf319e14"},"container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"name":"My Data Product","description":"This is a description of My Data Product.","tags":["Tags"],"use_cases":[{"id":"ID","name":"Name","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}],"types":["data"],"contract_terms":[{"asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}},"id":"ID","documents":[{"url":"URL","type":"terms_and_conditions","name":"Name","id":"2b0bf220-079c-11ee-be56-0242ac120002","attachment":{"id":"ID"},"upload_url":"UploadURL"}],"error_msg":"ErrorMsg"}],"is_restricted":true,"id":"2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd","asset":{"id":"2b0bf220-079c-11ee-be56-0242ac120002","container":{"id":"d29c42eb-7100-4b7a-8257-c196dbcca1cd","type":"catalog"}}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DataProductReleasesPager.GetNext successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductReleasesOptionsModel := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
					DataProductID:    core.StringPtr("testString"),
					AssetContainerID: core.StringPtr("testString"),
					State:            []string{"available"},
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []dataproducthubapiservicev1.DataProductVersionSummary
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DataProductReleasesPager.GetAll successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				listDataProductReleasesOptionsModel := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
					DataProductID:    core.StringPtr("testString"),
					AssetContainerID: core.StringPtr("testString"),
					State:            []string{"available"},
					Version:          core.StringPtr("testString"),
					Limit:            core.Int64Ptr(int64(10)),
				}

				pager, err := dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions) - Operation response error`, func() {
		retireDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString/retire"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RetireDataProductRelease with error: Operation response processing error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				dataProductHubAPIService.EnableRetries(0, 0)
				result, response, operationErr = dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions)`, func() {
		retireDataProductReleasePath := "/data_product_exchange/v1/data_products/testString/releases/testString/retire"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke RetireDataProductRelease successfully with retries`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())
				dataProductHubAPIService.EnableRetries(0, 0)

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := dataProductHubAPIService.RetireDataProductReleaseWithContext(ctx, retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				dataProductHubAPIService.DisableRetries()
				result, response, operationErr := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = dataProductHubAPIService.RetireDataProductReleaseWithContext(ctx, retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(retireDataProductReleasePath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": "1.0.0", "state": "draft", "data_product": {"id": "b38df608-d34b-4d58-8136-ed25e6c6684e", "release": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}, "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "name": "My Data Product", "description": "This is a description of My Data Product.", "tags": ["Tags"], "use_cases": [{"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}], "types": ["data"], "contract_terms": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "id": "ID", "documents": [{"url": "URL", "type": "terms_and_conditions", "name": "Name", "id": "2b0bf220-079c-11ee-be56-0242ac120002", "attachment": {"id": "ID"}, "upload_url": "UploadURL"}], "error_msg": "ErrorMsg"}], "is_restricted": true, "id": "2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd", "asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "domain": {"id": "ID", "name": "Name", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}, "parts_out": [{"asset": {"id": "2b0bf220-079c-11ee-be56-0242ac120002", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}, "type": "data_asset"}, "delivery_methods": [{"id": "09cf5fcc-cb9d-4995-a8e4-16517b25229f", "container": {"id": "d29c42eb-7100-4b7a-8257-c196dbcca1cd", "type": "catalog"}}]}], "published_by": "PublishedBy", "published_at": "2019-01-01T12:00:00.000Z", "created_by": "CreatedBy", "created_at": "2019-01-01T12:00:00.000Z", "workflows": {"order_access_request": {"task_assignee_users": ["TaskAssigneeUsers"], "pre_approved_users": ["PreApprovedUsers"], "custom_workflow_definition": {"id": "18bdbde1-918e-4ecf-aa23-6727bf319e14"}}}, "properties": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke RetireDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := dataProductHubAPIService.RetireDataProductRelease(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RetireDataProductRelease with error: Operation validation and request error`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := dataProductHubAPIService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RetireDataProductReleaseOptions model with no property values
				retireDataProductReleaseOptionsModelNew := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RetireDataProductRelease successfully`, func() {
				dataProductHubAPIService, serviceErr := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(dataProductHubAPIService).ToNot(BeNil())

				// Construct an instance of the RetireDataProductReleaseOptions model
				retireDataProductReleaseOptionsModel := new(dataproducthubapiservicev1.RetireDataProductReleaseOptions)
				retireDataProductReleaseOptionsModel.DataProductID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.ReleaseID = core.StringPtr("testString")
				retireDataProductReleaseOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			dataProductHubAPIService, _ := dataproducthubapiservicev1.NewDataProductHubAPIServiceV1(&dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{
				URL:           "http://dataproducthubapiservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAssetPartReference successfully`, func() {
				var container *dataproducthubapiservicev1.ContainerReference = nil
				_, err := dataProductHubAPIService.NewAssetPartReference(container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAssetPrototype successfully`, func() {
				var container *dataproducthubapiservicev1.ContainerIdentity = nil
				_, err := dataProductHubAPIService.NewAssetPrototype(container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewAssetReference successfully`, func() {
				var container *dataproducthubapiservicev1.ContainerReference = nil
				_, err := dataProductHubAPIService.NewAssetReference(container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCompleteDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the CompleteDraftContractTermsDocumentOptions model
				dataProductID := "testString"
				draftID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				completeDraftContractTermsDocumentOptionsModel := dataProductHubAPIService.NewCompleteDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				completeDraftContractTermsDocumentOptionsModel.SetDataProductID("testString")
				completeDraftContractTermsDocumentOptionsModel.SetDraftID("testString")
				completeDraftContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				completeDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				completeDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(completeDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(completeDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(completeDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(completeDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(completeDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewContainerIdentity successfully`, func() {
				id := "d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				_model, err := dataProductHubAPIService.NewContainerIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContainerReference successfully`, func() {
				id := "d29c42eb-7100-4b7a-8257-c196dbcca1cd"
				typeVar := "catalog"
				_model, err := dataProductHubAPIService.NewContainerReference(id, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContractTermsDocument successfully`, func() {
				typeVar := "terms_and_conditions"
				name := "testString"
				id := "2b0bf220-079c-11ee-be56-0242ac120002"
				_model, err := dataProductHubAPIService.NewContractTermsDocument(typeVar, name, id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewContractTermsDocumentPatch successfully`, func() {
				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocument := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocument.URL = core.StringPtr("testString")
				contractTermsDocument.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocument.Name = core.StringPtr("testString")
				contractTermsDocument.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocument.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocument.UploadURL = core.StringPtr("testString")

				contractTermsDocumentPatch := dataProductHubAPIService.NewContractTermsDocumentPatch(contractTermsDocument)
				Expect(contractTermsDocumentPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dataproducthubapiservicev1.JSONPatchOperation).Path
				}
				Expect(contractTermsDocumentPatch).To(MatchAllElements(_path, Elements{
					"/url": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/url")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.URL),
					}),
					"/type": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/type")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Type),
					}),
					"/name": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/name")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Name),
					}),
					"/id": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/id")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.ID),
					}),
					"/attachment": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/attachment")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.Attachment),
					}),
					"/upload_url": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/upload_url")),
						"From":  BeNil(),
						"Value": Equal(contractTermsDocument.UploadURL),
					}),
				}))
			})
			It(`Invoke NewCreateDataProductDraftOptions successfully`, func() {
				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				Expect(containerIdentityModel).ToNot(BeNil())
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				Expect(containerIdentityModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				Expect(assetPrototypeModel).ToNot(BeNil())
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel
				Expect(assetPrototypeModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPrototypeModel.Container).To(Equal(containerIdentityModel))

				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				Expect(dataProductDraftVersionReleaseModel).ToNot(BeNil())
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")
				Expect(dataProductDraftVersionReleaseModel.ID).To(Equal(core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3")))

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				Expect(dataProductIdentityModel).ToNot(BeNil())
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel
				Expect(dataProductIdentityModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(dataProductIdentityModel.Release).To(Equal(dataProductDraftVersionReleaseModel))

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				Expect(useCaseModel).ToNot(BeNil())
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel
				Expect(useCaseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				Expect(assetReferenceModel).ToNot(BeNil())
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel
				Expect(assetReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetReferenceModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				Expect(contractTermsDocumentModel).ToNot(BeNil())
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")
				Expect(contractTermsDocumentModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(contractTermsDocumentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(contractTermsDocumentModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(contractTermsDocumentModel.UploadURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				Expect(dataProductContractTermsModel).ToNot(BeNil())
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")
				Expect(dataProductContractTermsModel.Asset).To(Equal(assetReferenceModel))
				Expect(dataProductContractTermsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(dataProductContractTermsModel.Documents).To(Equal([]dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}))
				Expect(dataProductContractTermsModel.ErrorMsg).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				Expect(domainModel).ToNot(BeNil())
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel
				Expect(domainModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				Expect(assetPartReferenceModel).ToNot(BeNil())
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")
				Expect(assetPartReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPartReferenceModel.Container).To(Equal(containerReferenceModel))
				Expect(assetPartReferenceModel.Type).To(Equal(core.StringPtr("data_asset")))

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				Expect(deliveryMethodModel).ToNot(BeNil())
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel
				Expect(deliveryMethodModel.ID).To(Equal(core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")))
				Expect(deliveryMethodModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				Expect(dataProductPartModel).ToNot(BeNil())
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}
				Expect(dataProductPartModel.Asset).To(Equal(assetPartReferenceModel))
				Expect(dataProductPartModel.DeliveryMethods).To(Equal([]dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}))

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				Expect(dataProductCustomWorkflowDefinitionModel).ToNot(BeNil())
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")
				Expect(dataProductCustomWorkflowDefinitionModel.ID).To(Equal(core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")))

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				Expect(dataProductOrderAccessRequestModel).ToNot(BeNil())
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel
				Expect(dataProductOrderAccessRequestModel.TaskAssigneeUsers).To(Equal([]string{"testString"}))
				Expect(dataProductOrderAccessRequestModel.PreApprovedUsers).To(Equal([]string{"testString"}))
				Expect(dataProductOrderAccessRequestModel.CustomWorkflowDefinition).To(Equal(dataProductCustomWorkflowDefinitionModel))

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				Expect(dataProductWorkflowsModel).ToNot(BeNil())
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel
				Expect(dataProductWorkflowsModel.OrderAccessRequest).To(Equal(dataProductOrderAccessRequestModel))

				// Construct an instance of the CreateDataProductDraftOptions model
				dataProductID := "testString"
				var createDataProductDraftOptionsAsset *dataproducthubapiservicev1.AssetPrototype = nil
				createDataProductDraftOptionsModel := dataProductHubAPIService.NewCreateDataProductDraftOptions(dataProductID, createDataProductDraftOptionsAsset)
				createDataProductDraftOptionsModel.SetDataProductID("testString")
				createDataProductDraftOptionsModel.SetAsset(assetPrototypeModel)
				createDataProductDraftOptionsModel.SetVersion("1.2.0")
				createDataProductDraftOptionsModel.SetState("draft")
				createDataProductDraftOptionsModel.SetDataProduct(dataProductIdentityModel)
				createDataProductDraftOptionsModel.SetName("testString")
				createDataProductDraftOptionsModel.SetDescription("testString")
				createDataProductDraftOptionsModel.SetTags([]string{"testString"})
				createDataProductDraftOptionsModel.SetUseCases([]dataproducthubapiservicev1.UseCase{*useCaseModel})
				createDataProductDraftOptionsModel.SetTypes([]string{"data"})
				createDataProductDraftOptionsModel.SetContractTerms([]dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel})
				createDataProductDraftOptionsModel.SetIsRestricted(true)
				createDataProductDraftOptionsModel.SetDomain(domainModel)
				createDataProductDraftOptionsModel.SetPartsOut([]dataproducthubapiservicev1.DataProductPart{*dataProductPartModel})
				createDataProductDraftOptionsModel.SetWorkflows(dataProductWorkflowsModel)
				createDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(createDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductDraftOptionsModel.Asset).To(Equal(assetPrototypeModel))
				Expect(createDataProductDraftOptionsModel.Version).To(Equal(core.StringPtr("1.2.0")))
				Expect(createDataProductDraftOptionsModel.State).To(Equal(core.StringPtr("draft")))
				Expect(createDataProductDraftOptionsModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(createDataProductDraftOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductDraftOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDataProductDraftOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createDataProductDraftOptionsModel.UseCases).To(Equal([]dataproducthubapiservicev1.UseCase{*useCaseModel}))
				Expect(createDataProductDraftOptionsModel.Types).To(Equal([]string{"data"}))
				Expect(createDataProductDraftOptionsModel.ContractTerms).To(Equal([]dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}))
				Expect(createDataProductDraftOptionsModel.IsRestricted).To(Equal(core.BoolPtr(true)))
				Expect(createDataProductDraftOptionsModel.Domain).To(Equal(domainModel))
				Expect(createDataProductDraftOptionsModel.PartsOut).To(Equal([]dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}))
				Expect(createDataProductDraftOptionsModel.Workflows).To(Equal(dataProductWorkflowsModel))
				Expect(createDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDataProductOptions successfully`, func() {
				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				Expect(dataProductDraftVersionReleaseModel).ToNot(BeNil())
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")
				Expect(dataProductDraftVersionReleaseModel.ID).To(Equal(core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")))

				// Construct an instance of the DataProductIdentity model
				dataProductIdentityModel := new(dataproducthubapiservicev1.DataProductIdentity)
				Expect(dataProductIdentityModel).ToNot(BeNil())
				dataProductIdentityModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductIdentityModel.Release = dataProductDraftVersionReleaseModel
				Expect(dataProductIdentityModel.ID).To(Equal(core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")))
				Expect(dataProductIdentityModel.Release).To(Equal(dataProductDraftVersionReleaseModel))

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				Expect(useCaseModel).ToNot(BeNil())
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel
				Expect(useCaseModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(useCaseModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				Expect(assetReferenceModel).ToNot(BeNil())
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel
				Expect(assetReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetReferenceModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				Expect(contractTermsDocumentAttachmentModel).ToNot(BeNil())
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")
				Expect(contractTermsDocumentAttachmentModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				Expect(contractTermsDocumentModel).ToNot(BeNil())
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")
				Expect(contractTermsDocumentModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(contractTermsDocumentModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(contractTermsDocumentModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(contractTermsDocumentModel.Attachment).To(Equal(contractTermsDocumentAttachmentModel))
				Expect(contractTermsDocumentModel.UploadURL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				Expect(dataProductContractTermsModel).ToNot(BeNil())
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")
				Expect(dataProductContractTermsModel.Asset).To(Equal(assetReferenceModel))
				Expect(dataProductContractTermsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(dataProductContractTermsModel.Documents).To(Equal([]dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}))
				Expect(dataProductContractTermsModel.ErrorMsg).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ContainerIdentity model
				containerIdentityModel := new(dataproducthubapiservicev1.ContainerIdentity)
				Expect(containerIdentityModel).ToNot(BeNil())
				containerIdentityModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				Expect(containerIdentityModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))

				// Construct an instance of the AssetPrototype model
				assetPrototypeModel := new(dataproducthubapiservicev1.AssetPrototype)
				Expect(assetPrototypeModel).ToNot(BeNil())
				assetPrototypeModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPrototypeModel.Container = containerIdentityModel
				Expect(assetPrototypeModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPrototypeModel.Container).To(Equal(containerIdentityModel))

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				Expect(domainModel).ToNot(BeNil())
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel
				Expect(domainModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(domainModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				Expect(assetPartReferenceModel).ToNot(BeNil())
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")
				Expect(assetPartReferenceModel.ID).To(Equal(core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")))
				Expect(assetPartReferenceModel.Container).To(Equal(containerReferenceModel))
				Expect(assetPartReferenceModel.Type).To(Equal(core.StringPtr("data_asset")))

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				Expect(deliveryMethodModel).ToNot(BeNil())
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel
				Expect(deliveryMethodModel.ID).To(Equal(core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")))
				Expect(deliveryMethodModel.Container).To(Equal(containerReferenceModel))

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				Expect(dataProductPartModel).ToNot(BeNil())
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}
				Expect(dataProductPartModel.Asset).To(Equal(assetPartReferenceModel))
				Expect(dataProductPartModel.DeliveryMethods).To(Equal([]dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}))

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				Expect(dataProductCustomWorkflowDefinitionModel).ToNot(BeNil())
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")
				Expect(dataProductCustomWorkflowDefinitionModel.ID).To(Equal(core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")))

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				Expect(dataProductOrderAccessRequestModel).ToNot(BeNil())
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel
				Expect(dataProductOrderAccessRequestModel.TaskAssigneeUsers).To(Equal([]string{"testString"}))
				Expect(dataProductOrderAccessRequestModel.PreApprovedUsers).To(Equal([]string{"testString"}))
				Expect(dataProductOrderAccessRequestModel.CustomWorkflowDefinition).To(Equal(dataProductCustomWorkflowDefinitionModel))

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				Expect(dataProductWorkflowsModel).ToNot(BeNil())
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel
				Expect(dataProductWorkflowsModel.OrderAccessRequest).To(Equal(dataProductOrderAccessRequestModel))

				// Construct an instance of the DataProductVersionPrototype model
				dataProductVersionPrototypeModel := new(dataproducthubapiservicev1.DataProductVersionPrototype)
				Expect(dataProductVersionPrototypeModel).ToNot(BeNil())
				dataProductVersionPrototypeModel.Version = core.StringPtr("1.0.0")
				dataProductVersionPrototypeModel.State = core.StringPtr("draft")
				dataProductVersionPrototypeModel.DataProduct = dataProductIdentityModel
				dataProductVersionPrototypeModel.Name = core.StringPtr("My New Data Product")
				dataProductVersionPrototypeModel.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersionPrototypeModel.Tags = []string{"testString"}
				dataProductVersionPrototypeModel.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersionPrototypeModel.Types = []string{"data"}
				dataProductVersionPrototypeModel.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersionPrototypeModel.IsRestricted = core.BoolPtr(true)
				dataProductVersionPrototypeModel.Asset = assetPrototypeModel
				dataProductVersionPrototypeModel.Domain = domainModel
				dataProductVersionPrototypeModel.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersionPrototypeModel.Workflows = dataProductWorkflowsModel
				Expect(dataProductVersionPrototypeModel.Version).To(Equal(core.StringPtr("1.0.0")))
				Expect(dataProductVersionPrototypeModel.State).To(Equal(core.StringPtr("draft")))
				Expect(dataProductVersionPrototypeModel.DataProduct).To(Equal(dataProductIdentityModel))
				Expect(dataProductVersionPrototypeModel.Name).To(Equal(core.StringPtr("My New Data Product")))
				Expect(dataProductVersionPrototypeModel.Description).To(Equal(core.StringPtr("This is a description of My Data Product.")))
				Expect(dataProductVersionPrototypeModel.Tags).To(Equal([]string{"testString"}))
				Expect(dataProductVersionPrototypeModel.UseCases).To(Equal([]dataproducthubapiservicev1.UseCase{*useCaseModel}))
				Expect(dataProductVersionPrototypeModel.Types).To(Equal([]string{"data"}))
				Expect(dataProductVersionPrototypeModel.ContractTerms).To(Equal([]dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}))
				Expect(dataProductVersionPrototypeModel.IsRestricted).To(Equal(core.BoolPtr(true)))
				Expect(dataProductVersionPrototypeModel.Asset).To(Equal(assetPrototypeModel))
				Expect(dataProductVersionPrototypeModel.Domain).To(Equal(domainModel))
				Expect(dataProductVersionPrototypeModel.PartsOut).To(Equal([]dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}))
				Expect(dataProductVersionPrototypeModel.Workflows).To(Equal(dataProductWorkflowsModel))

				// Construct an instance of the CreateDataProductOptions model
				createDataProductOptionsDrafts := []dataproducthubapiservicev1.DataProductVersionPrototype{}
				createDataProductOptionsModel := dataProductHubAPIService.NewCreateDataProductOptions(createDataProductOptionsDrafts)
				createDataProductOptionsModel.SetDrafts([]dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel})
				createDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDataProductOptionsModel).ToNot(BeNil())
				Expect(createDataProductOptionsModel.Drafts).To(Equal([]dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel}))
				Expect(createDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the CreateDraftContractTermsDocumentOptions model
				dataProductID := "testString"
				draftID := "testString"
				contractTermsID := "testString"
				createDraftContractTermsDocumentOptionsType := "terms_and_conditions"
				createDraftContractTermsDocumentOptionsName := "Terms and conditions document"
				createDraftContractTermsDocumentOptionsModel := dataProductHubAPIService.NewCreateDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, createDraftContractTermsDocumentOptionsType, createDraftContractTermsDocumentOptionsName)
				createDraftContractTermsDocumentOptionsModel.SetDataProductID("testString")
				createDraftContractTermsDocumentOptionsModel.SetDraftID("testString")
				createDraftContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				createDraftContractTermsDocumentOptionsModel.SetType("terms_and_conditions")
				createDraftContractTermsDocumentOptionsModel.SetName("Terms and conditions document")
				createDraftContractTermsDocumentOptionsModel.SetURL("testString")
				createDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(createDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.Type).To(Equal(core.StringPtr("terms_and_conditions")))
				Expect(createDraftContractTermsDocumentOptionsModel.Name).To(Equal(core.StringPtr("Terms and conditions document")))
				Expect(createDraftContractTermsDocumentOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDataProductIdentity successfully`, func() {
				id := "b38df608-d34b-4d58-8136-ed25e6c6684e"
				_model, err := dataProductHubAPIService.NewDataProductIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDataProductPart successfully`, func() {
				var asset *dataproducthubapiservicev1.AssetPartReference = nil
				_, err := dataProductHubAPIService.NewDataProductPart(asset)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDataProductVersionPatch successfully`, func() {
				// Construct an instance of the DataProductDraftVersionRelease model
				dataProductDraftVersionReleaseModel := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
				dataProductDraftVersionReleaseModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")

				// Construct an instance of the DataProductVersionDataProduct model
				dataProductVersionDataProductModel := new(dataproducthubapiservicev1.DataProductVersionDataProduct)
				dataProductVersionDataProductModel.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
				dataProductVersionDataProductModel.Release = dataProductDraftVersionReleaseModel
				dataProductVersionDataProductModel.Container = containerReferenceModel

				// Construct an instance of the UseCase model
				useCaseModel := new(dataproducthubapiservicev1.UseCase)
				useCaseModel.ID = core.StringPtr("testString")
				useCaseModel.Name = core.StringPtr("testString")
				useCaseModel.Container = containerReferenceModel

				// Construct an instance of the AssetReference model
				assetReferenceModel := new(dataproducthubapiservicev1.AssetReference)
				assetReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetReferenceModel.Container = containerReferenceModel

				// Construct an instance of the ContractTermsDocumentAttachment model
				contractTermsDocumentAttachmentModel := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
				contractTermsDocumentAttachmentModel.ID = core.StringPtr("testString")

				// Construct an instance of the ContractTermsDocument model
				contractTermsDocumentModel := new(dataproducthubapiservicev1.ContractTermsDocument)
				contractTermsDocumentModel.URL = core.StringPtr("testString")
				contractTermsDocumentModel.Type = core.StringPtr("terms_and_conditions")
				contractTermsDocumentModel.Name = core.StringPtr("testString")
				contractTermsDocumentModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				contractTermsDocumentModel.Attachment = contractTermsDocumentAttachmentModel
				contractTermsDocumentModel.UploadURL = core.StringPtr("testString")

				// Construct an instance of the DataProductContractTerms model
				dataProductContractTermsModel := new(dataproducthubapiservicev1.DataProductContractTerms)
				dataProductContractTermsModel.Asset = assetReferenceModel
				dataProductContractTermsModel.ID = core.StringPtr("testString")
				dataProductContractTermsModel.Documents = []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel}
				dataProductContractTermsModel.ErrorMsg = core.StringPtr("testString")

				// Construct an instance of the Domain model
				domainModel := new(dataproducthubapiservicev1.Domain)
				domainModel.ID = core.StringPtr("testString")
				domainModel.Name = core.StringPtr("testString")
				domainModel.Container = containerReferenceModel

				// Construct an instance of the AssetPartReference model
				assetPartReferenceModel := new(dataproducthubapiservicev1.AssetPartReference)
				assetPartReferenceModel.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
				assetPartReferenceModel.Container = containerReferenceModel
				assetPartReferenceModel.Type = core.StringPtr("data_asset")

				// Construct an instance of the DeliveryMethod model
				deliveryMethodModel := new(dataproducthubapiservicev1.DeliveryMethod)
				deliveryMethodModel.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
				deliveryMethodModel.Container = containerReferenceModel

				// Construct an instance of the DataProductPart model
				dataProductPartModel := new(dataproducthubapiservicev1.DataProductPart)
				dataProductPartModel.Asset = assetPartReferenceModel
				dataProductPartModel.DeliveryMethods = []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel}

				// Construct an instance of the DataProductCustomWorkflowDefinition model
				dataProductCustomWorkflowDefinitionModel := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
				dataProductCustomWorkflowDefinitionModel.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

				// Construct an instance of the DataProductOrderAccessRequest model
				dataProductOrderAccessRequestModel := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
				dataProductOrderAccessRequestModel.TaskAssigneeUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.PreApprovedUsers = []string{"testString"}
				dataProductOrderAccessRequestModel.CustomWorkflowDefinition = dataProductCustomWorkflowDefinitionModel

				// Construct an instance of the DataProductWorkflows model
				dataProductWorkflowsModel := new(dataproducthubapiservicev1.DataProductWorkflows)
				dataProductWorkflowsModel.OrderAccessRequest = dataProductOrderAccessRequestModel

				// Construct an instance of the DataProductVersion model
				dataProductVersion := new(dataproducthubapiservicev1.DataProductVersion)
				dataProductVersion.Version = core.StringPtr("1.0.0")
				dataProductVersion.State = core.StringPtr("draft")
				dataProductVersion.DataProduct = dataProductVersionDataProductModel
				dataProductVersion.Name = core.StringPtr("My Data Product")
				dataProductVersion.Description = core.StringPtr("This is a description of My Data Product.")
				dataProductVersion.Tags = []string{"testString"}
				dataProductVersion.UseCases = []dataproducthubapiservicev1.UseCase{*useCaseModel}
				dataProductVersion.Types = []string{"data"}
				dataProductVersion.ContractTerms = []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel}
				dataProductVersion.IsRestricted = core.BoolPtr(true)
				dataProductVersion.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002@d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				dataProductVersion.Asset = assetReferenceModel
				dataProductVersion.Domain = domainModel
				dataProductVersion.PartsOut = []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel}
				dataProductVersion.PublishedBy = core.StringPtr("testString")
				dataProductVersion.PublishedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.CreatedBy = core.StringPtr("testString")
				dataProductVersion.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				dataProductVersion.Workflows = dataProductWorkflowsModel
				dataProductVersion.Properties = map[string]interface{}{"anyKey": "anyValue"}

				dataProductVersionPatch := dataProductHubAPIService.NewDataProductVersionPatch(dataProductVersion)
				Expect(dataProductVersionPatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(dataproducthubapiservicev1.JSONPatchOperation).Path
				}
				Expect(dataProductVersionPatch).To(MatchAllElements(_path, Elements{
					"/version": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/version")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Version),
					}),
					"/state": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/state")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.State),
					}),
					"/data_product": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/data_product")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.DataProduct),
					}),
					"/name": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/name")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Name),
					}),
					"/description": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/description")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Description),
					}),
					"/tags": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/tags")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Tags),
					}),
					"/use_cases": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/use_cases")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.UseCases),
					}),
					"/types": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/types")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Types),
					}),
					"/contract_terms": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/contract_terms")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.ContractTerms),
					}),
					"/is_restricted": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/is_restricted")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.IsRestricted),
					}),
					"/id": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/id")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.ID),
					}),
					"/asset": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/asset")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Asset),
					}),
					"/domain": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/domain")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Domain),
					}),
					"/parts_out": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/parts_out")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PartsOut),
					}),
					"/published_by": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/published_by")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PublishedBy),
					}),
					"/published_at": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/published_at")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.PublishedAt),
					}),
					"/created_by": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/created_by")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.CreatedBy),
					}),
					"/created_at": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/created_at")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.CreatedAt),
					}),
					"/workflows": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/workflows")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Workflows),
					}),
					"/properties": MatchAllFields(Fields{
						"Op":    PointTo(Equal(dataproducthubapiservicev1.JSONPatchOperationOpAddConst)),
						"Path":  PointTo(Equal("/properties")),
						"From":  BeNil(),
						"Value": Equal(dataProductVersion.Properties),
					}),
				}))
			})
			It(`Invoke NewDataProductVersionPrototype successfully`, func() {
				var asset *dataproducthubapiservicev1.AssetPrototype = nil
				_, err := dataProductHubAPIService.NewDataProductVersionPrototype(asset)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDeleteDataProductDraftOptions successfully`, func() {
				// Construct an instance of the DeleteDataProductDraftOptions model
				dataProductID := "testString"
				draftID := "testString"
				deleteDataProductDraftOptionsModel := dataProductHubAPIService.NewDeleteDataProductDraftOptions(dataProductID, draftID)
				deleteDataProductDraftOptionsModel.SetDataProductID("testString")
				deleteDataProductDraftOptionsModel.SetDraftID("testString")
				deleteDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(deleteDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the DeleteDraftContractTermsDocumentOptions model
				dataProductID := "testString"
				draftID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				deleteDraftContractTermsDocumentOptionsModel := dataProductHubAPIService.NewDeleteDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				deleteDraftContractTermsDocumentOptionsModel.SetDataProductID("testString")
				deleteDraftContractTermsDocumentOptionsModel.SetDraftID("testString")
				deleteDraftContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				deleteDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				deleteDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(deleteDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeliveryMethod successfully`, func() {
				id := "09cf5fcc-cb9d-4995-a8e4-16517b25229f"
				var container *dataproducthubapiservicev1.ContainerReference = nil
				_, err := dataProductHubAPIService.NewDeliveryMethod(id, container)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewDomain successfully`, func() {
				id := "testString"
				_model, err := dataProductHubAPIService.NewDomain(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetDataProductDraftOptions successfully`, func() {
				// Construct an instance of the GetDataProductDraftOptions model
				dataProductID := "testString"
				draftID := "testString"
				getDataProductDraftOptionsModel := dataProductHubAPIService.NewGetDataProductDraftOptions(dataProductID, draftID)
				getDataProductDraftOptionsModel.SetDataProductID("testString")
				getDataProductDraftOptionsModel.SetDraftID("testString")
				getDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(getDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductOptions successfully`, func() {
				// Construct an instance of the GetDataProductOptions model
				dataProductID := "testString"
				getDataProductOptionsModel := dataProductHubAPIService.NewGetDataProductOptions(dataProductID)
				getDataProductOptionsModel.SetDataProductID("testString")
				getDataProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductOptionsModel).ToNot(BeNil())
				Expect(getDataProductOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the GetDataProductReleaseOptions model
				dataProductID := "testString"
				releaseID := "testString"
				getDataProductReleaseOptionsModel := dataProductHubAPIService.NewGetDataProductReleaseOptions(dataProductID, releaseID)
				getDataProductReleaseOptionsModel.SetDataProductID("testString")
				getDataProductReleaseOptionsModel.SetReleaseID("testString")
				getDataProductReleaseOptionsModel.SetCheckCallerApproval(false)
				getDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(getDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("testString")))
				Expect(getDataProductReleaseOptionsModel.CheckCallerApproval).To(Equal(core.BoolPtr(false)))
				Expect(getDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the GetDraftContractTermsDocumentOptions model
				dataProductID := "testString"
				draftID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				getDraftContractTermsDocumentOptionsModel := dataProductHubAPIService.NewGetDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID)
				getDraftContractTermsDocumentOptionsModel.SetDataProductID("testString")
				getDraftContractTermsDocumentOptionsModel.SetDraftID("testString")
				getDraftContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				getDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				getDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(getDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(getDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(getDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(getDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetInitializeStatusOptions successfully`, func() {
				// Construct an instance of the GetInitializeStatusOptions model
				getInitializeStatusOptionsModel := dataProductHubAPIService.NewGetInitializeStatusOptions()
				getInitializeStatusOptionsModel.SetContainerID("testString")
				getInitializeStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getInitializeStatusOptionsModel).ToNot(BeNil())
				Expect(getInitializeStatusOptionsModel.ContainerID).To(Equal(core.StringPtr("testString")))
				Expect(getInitializeStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetReleaseContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the GetReleaseContractTermsDocumentOptions model
				dataProductID := "testString"
				releaseID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				getReleaseContractTermsDocumentOptionsModel := dataProductHubAPIService.NewGetReleaseContractTermsDocumentOptions(dataProductID, releaseID, contractTermsID, documentID)
				getReleaseContractTermsDocumentOptionsModel.SetDataProductID("testString")
				getReleaseContractTermsDocumentOptionsModel.SetReleaseID("testString")
				getReleaseContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				getReleaseContractTermsDocumentOptionsModel.SetDocumentID("testString")
				getReleaseContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getReleaseContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(getReleaseContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(getReleaseContractTermsDocumentOptionsModel.ReleaseID).To(Equal(core.StringPtr("testString")))
				Expect(getReleaseContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(getReleaseContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getReleaseContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceIDCredentialsOptions successfully`, func() {
				// Construct an instance of the GetServiceIDCredentialsOptions model
				getServiceIDCredentialsOptionsModel := dataProductHubAPIService.NewGetServiceIDCredentialsOptions()
				getServiceIDCredentialsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceIDCredentialsOptionsModel).ToNot(BeNil())
				Expect(getServiceIDCredentialsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInitializeOptions successfully`, func() {
				// Construct an instance of the ContainerReference model
				containerReferenceModel := new(dataproducthubapiservicev1.ContainerReference)
				Expect(containerReferenceModel).ToNot(BeNil())
				containerReferenceModel.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
				containerReferenceModel.Type = core.StringPtr("catalog")
				Expect(containerReferenceModel.ID).To(Equal(core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")))
				Expect(containerReferenceModel.Type).To(Equal(core.StringPtr("catalog")))

				// Construct an instance of the InitializeOptions model
				initializeOptionsModel := dataProductHubAPIService.NewInitializeOptions()
				initializeOptionsModel.SetContainer(containerReferenceModel)
				initializeOptionsModel.SetInclude([]string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"})
				initializeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(initializeOptionsModel).ToNot(BeNil())
				Expect(initializeOptionsModel.Container).To(Equal(containerReferenceModel))
				Expect(initializeOptionsModel.Include).To(Equal([]string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"}))
				Expect(initializeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := dataProductHubAPIService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListDataProductDraftsOptions successfully`, func() {
				// Construct an instance of the ListDataProductDraftsOptions model
				dataProductID := "testString"
				listDataProductDraftsOptionsModel := dataProductHubAPIService.NewListDataProductDraftsOptions(dataProductID)
				listDataProductDraftsOptionsModel.SetDataProductID("testString")
				listDataProductDraftsOptionsModel.SetAssetContainerID("testString")
				listDataProductDraftsOptionsModel.SetVersion("testString")
				listDataProductDraftsOptionsModel.SetLimit(int64(10))
				listDataProductDraftsOptionsModel.SetStart("testString")
				listDataProductDraftsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductDraftsOptionsModel).ToNot(BeNil())
				Expect(listDataProductDraftsOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.AssetContainerID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductDraftsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductDraftsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataProductReleasesOptions successfully`, func() {
				// Construct an instance of the ListDataProductReleasesOptions model
				dataProductID := "testString"
				listDataProductReleasesOptionsModel := dataProductHubAPIService.NewListDataProductReleasesOptions(dataProductID)
				listDataProductReleasesOptionsModel.SetDataProductID("testString")
				listDataProductReleasesOptionsModel.SetAssetContainerID("testString")
				listDataProductReleasesOptionsModel.SetState([]string{"available"})
				listDataProductReleasesOptionsModel.SetVersion("testString")
				listDataProductReleasesOptionsModel.SetLimit(int64(10))
				listDataProductReleasesOptionsModel.SetStart("testString")
				listDataProductReleasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductReleasesOptionsModel).ToNot(BeNil())
				Expect(listDataProductReleasesOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.AssetContainerID).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.State).To(Equal([]string{"available"}))
				Expect(listDataProductReleasesOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductReleasesOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductReleasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDataProductsOptions successfully`, func() {
				// Construct an instance of the ListDataProductsOptions model
				listDataProductsOptionsModel := dataProductHubAPIService.NewListDataProductsOptions()
				listDataProductsOptionsModel.SetLimit(int64(10))
				listDataProductsOptionsModel.SetStart("testString")
				listDataProductsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDataProductsOptionsModel).ToNot(BeNil())
				Expect(listDataProductsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDataProductsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listDataProductsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewManageAPIKeysOptions successfully`, func() {
				// Construct an instance of the ManageAPIKeysOptions model
				manageAPIKeysOptionsModel := dataProductHubAPIService.NewManageAPIKeysOptions()
				manageAPIKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(manageAPIKeysOptionsModel).ToNot(BeNil())
				Expect(manageAPIKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPublishDataProductDraftOptions successfully`, func() {
				// Construct an instance of the PublishDataProductDraftOptions model
				dataProductID := "testString"
				draftID := "testString"
				publishDataProductDraftOptionsModel := dataProductHubAPIService.NewPublishDataProductDraftOptions(dataProductID, draftID)
				publishDataProductDraftOptionsModel.SetDataProductID("testString")
				publishDataProductDraftOptionsModel.SetDraftID("testString")
				publishDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(publishDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(publishDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(publishDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(publishDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRetireDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the RetireDataProductReleaseOptions model
				dataProductID := "testString"
				releaseID := "testString"
				retireDataProductReleaseOptionsModel := dataProductHubAPIService.NewRetireDataProductReleaseOptions(dataProductID, releaseID)
				retireDataProductReleaseOptionsModel.SetDataProductID("testString")
				retireDataProductReleaseOptionsModel.SetReleaseID("testString")
				retireDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(retireDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(retireDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(retireDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("testString")))
				Expect(retireDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDataProductDraftOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDataProductDraftOptions model
				dataProductID := "testString"
				draftID := "testString"
				jsonPatchInstructions := []dataproducthubapiservicev1.JSONPatchOperation{}
				updateDataProductDraftOptionsModel := dataProductHubAPIService.NewUpdateDataProductDraftOptions(dataProductID, draftID, jsonPatchInstructions)
				updateDataProductDraftOptionsModel.SetDataProductID("testString")
				updateDataProductDraftOptionsModel.SetDraftID("testString")
				updateDataProductDraftOptionsModel.SetJSONPatchInstructions([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductDraftOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductDraftOptionsModel).ToNot(BeNil())
				Expect(updateDataProductDraftOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductDraftOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductDraftOptionsModel.JSONPatchInstructions).To(Equal([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductDraftOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDataProductReleaseOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDataProductReleaseOptions model
				dataProductID := "testString"
				releaseID := "testString"
				jsonPatchInstructions := []dataproducthubapiservicev1.JSONPatchOperation{}
				updateDataProductReleaseOptionsModel := dataProductHubAPIService.NewUpdateDataProductReleaseOptions(dataProductID, releaseID, jsonPatchInstructions)
				updateDataProductReleaseOptionsModel.SetDataProductID("testString")
				updateDataProductReleaseOptionsModel.SetReleaseID("testString")
				updateDataProductReleaseOptionsModel.SetJSONPatchInstructions([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDataProductReleaseOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDataProductReleaseOptionsModel).ToNot(BeNil())
				Expect(updateDataProductReleaseOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductReleaseOptionsModel.ReleaseID).To(Equal(core.StringPtr("testString")))
				Expect(updateDataProductReleaseOptionsModel.JSONPatchInstructions).To(Equal([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDataProductReleaseOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateDraftContractTermsDocumentOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(dataproducthubapiservicev1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = "testString"
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal("testString"))

				// Construct an instance of the UpdateDraftContractTermsDocumentOptions model
				dataProductID := "testString"
				draftID := "testString"
				contractTermsID := "testString"
				documentID := "testString"
				jsonPatchInstructions := []dataproducthubapiservicev1.JSONPatchOperation{}
				updateDraftContractTermsDocumentOptionsModel := dataProductHubAPIService.NewUpdateDraftContractTermsDocumentOptions(dataProductID, draftID, contractTermsID, documentID, jsonPatchInstructions)
				updateDraftContractTermsDocumentOptionsModel.SetDataProductID("testString")
				updateDraftContractTermsDocumentOptionsModel.SetDraftID("testString")
				updateDraftContractTermsDocumentOptionsModel.SetContractTermsID("testString")
				updateDraftContractTermsDocumentOptionsModel.SetDocumentID("testString")
				updateDraftContractTermsDocumentOptionsModel.SetJSONPatchInstructions([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDraftContractTermsDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDraftContractTermsDocumentOptionsModel).ToNot(BeNil())
				Expect(updateDraftContractTermsDocumentOptionsModel.DataProductID).To(Equal(core.StringPtr("testString")))
				Expect(updateDraftContractTermsDocumentOptionsModel.DraftID).To(Equal(core.StringPtr("testString")))
				Expect(updateDraftContractTermsDocumentOptionsModel.ContractTermsID).To(Equal(core.StringPtr("testString")))
				Expect(updateDraftContractTermsDocumentOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(updateDraftContractTermsDocumentOptionsModel.JSONPatchInstructions).To(Equal([]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDraftContractTermsDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUseCase successfully`, func() {
				id := "testString"
				_model, err := dataProductHubAPIService.NewUseCase(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAssetPartReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.AssetPartReference)
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Container = nil
			model.Type = core.StringPtr("data_asset")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.AssetPartReference
			err = dataproducthubapiservicev1.UnmarshalAssetPartReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssetPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.AssetPrototype)
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.AssetPrototype
			err = dataproducthubapiservicev1.UnmarshalAssetPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalAssetReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.AssetReference)
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.AssetReference
			err = dataproducthubapiservicev1.UnmarshalAssetReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContainerIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.ContainerIdentity)
			model.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.ContainerIdentity
			err = dataproducthubapiservicev1.UnmarshalContainerIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContainerReference successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.ContainerReference)
			model.ID = core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd")
			model.Type = core.StringPtr("catalog")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.ContainerReference
			err = dataproducthubapiservicev1.UnmarshalContainerReference(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContractTermsDocument successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.ContractTermsDocument)
			model.URL = core.StringPtr("testString")
			model.Type = core.StringPtr("terms_and_conditions")
			model.Name = core.StringPtr("testString")
			model.ID = core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002")
			model.Attachment = nil
			model.UploadURL = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.ContractTermsDocument
			err = dataproducthubapiservicev1.UnmarshalContractTermsDocument(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalContractTermsDocumentAttachment successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.ContractTermsDocumentAttachment)
			model.ID = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.ContractTermsDocumentAttachment
			err = dataproducthubapiservicev1.UnmarshalContractTermsDocumentAttachment(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductContractTerms successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductContractTerms)
			model.Asset = nil
			model.ID = core.StringPtr("testString")
			model.Documents = nil
			model.ErrorMsg = core.StringPtr("testString")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductContractTerms
			err = dataproducthubapiservicev1.UnmarshalDataProductContractTerms(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductCustomWorkflowDefinition successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductCustomWorkflowDefinition)
			model.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductCustomWorkflowDefinition
			err = dataproducthubapiservicev1.UnmarshalDataProductCustomWorkflowDefinition(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductDraftVersionRelease successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductDraftVersionRelease)
			model.ID = core.StringPtr("18bdbde1-918e-4ecf-aa23-6727bf319e14")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductDraftVersionRelease
			err = dataproducthubapiservicev1.UnmarshalDataProductDraftVersionRelease(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductIdentity)
			model.ID = core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e")
			model.Release = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductIdentity
			err = dataproducthubapiservicev1.UnmarshalDataProductIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductOrderAccessRequest successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductOrderAccessRequest)
			model.TaskAssigneeUsers = []string{"testString"}
			model.PreApprovedUsers = []string{"testString"}
			model.CustomWorkflowDefinition = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductOrderAccessRequest
			err = dataproducthubapiservicev1.UnmarshalDataProductOrderAccessRequest(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductPart successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductPart)
			model.Asset = nil
			model.DeliveryMethods = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductPart
			err = dataproducthubapiservicev1.UnmarshalDataProductPart(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductVersionPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductVersionPrototype)
			model.Version = core.StringPtr("1.0.0")
			model.State = core.StringPtr("draft")
			model.DataProduct = nil
			model.Name = core.StringPtr("My Data Product")
			model.Description = core.StringPtr("This is a description of My Data Product.")
			model.Tags = []string{"testString"}
			model.UseCases = nil
			model.Types = []string{"data"}
			model.ContractTerms = nil
			model.IsRestricted = core.BoolPtr(true)
			model.Asset = nil
			model.Domain = nil
			model.PartsOut = nil
			model.Workflows = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductVersionPrototype
			err = dataproducthubapiservicev1.UnmarshalDataProductVersionPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDataProductWorkflows successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DataProductWorkflows)
			model.OrderAccessRequest = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DataProductWorkflows
			err = dataproducthubapiservicev1.UnmarshalDataProductWorkflows(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDeliveryMethod successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.DeliveryMethod)
			model.ID = core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.DeliveryMethod
			err = dataproducthubapiservicev1.UnmarshalDeliveryMethod(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalDomain successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.Domain)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.Domain
			err = dataproducthubapiservicev1.UnmarshalDomain(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalJSONPatchOperation successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.JSONPatchOperation)
			model.Op = core.StringPtr("add")
			model.Path = core.StringPtr("testString")
			model.From = core.StringPtr("testString")
			model.Value = "testString"

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.JSONPatchOperation
			err = dataproducthubapiservicev1.UnmarshalJSONPatchOperation(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalUseCase successfully`, func() {
			// Construct an instance of the model.
			model := new(dataproducthubapiservicev1.UseCase)
			model.ID = core.StringPtr("testString")
			model.Name = core.StringPtr("testString")
			model.Container = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *dataproducthubapiservicev1.UseCase
			err = dataproducthubapiservicev1.UnmarshalUseCase(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
