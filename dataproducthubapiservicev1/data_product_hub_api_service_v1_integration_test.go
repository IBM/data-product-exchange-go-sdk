//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/data-product-exchange-go-sdk/dataproducthubapiservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the dataproducthubapiservicev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`DataProductHubAPIServiceV1 Integration Tests`, func() {
	const externalConfigFile = "../dph_v1.env"

	var (
		err                      error
		dataProductHubAPIService *dataproducthubapiservicev1.DataProductHubAPIServiceV1
		serviceURL               string
		config                   map[string]string

		// Variables to hold link values
		completeContractTermsDocumentByDocumentIDLink     string
		completeDraftContractTermsByDataProductIDLink     string
		createAContractTermsDocByContractTermsIDLink      string
		createAContractTermsDocByDraftIDLink              string
		createDataProductByCatalogIDLink                  string
		createDraftByContainerIDLink                      string
		createNewDraftByDataProductIDLink                 string
		deleteAContractDocumentByDraftIDLink              string
		deleteADraftByContractTermsIDLink                 string
		deleteADraftByDraftIDLink                         string
		deleteContractDocumentByDataProductIDLink         string
		deleteContractTermsDocumentByDocumentIDLink       string
		deleteDraftOfDataProductByDataProductIDLink       string
		getADraftOfDataProductByDataProductIDLink         string
		getAReleaseByReleaseIDLink                        string
		getAReleaseContractTermsByContractTermsIDLink     string
		getAReleaseContractTermsByReleaseIDLink           string
		getAReleaseOfDataProductByDataProductIDLink       string
		getContractDocumentByDataProductIDLink            string
		getContractTermsDocumentByIDDocumentIDLink        string
		getDataProductByDataProductIDLink                 string
		getDraftByDraftIDLink                             string
		getListOfDataProductDraftsByDataProductIDLink     string
		getListOfReleasesOfDataProductByDataProductIDLink string
		getReleaseContractDocumentByDataProductIDLink     string
		getReleaseContractDocumentByDocumentIDLink        string
		getStatusByCatalogIDLink                          string
		publishADraftOfDataProductByDataProductIDLink     string
		retireAReleaseContractTermsByReleaseIDLink        string
		retireAReleasesOfDataProductByDataProductIDLink   string
		updateAReleaseByReleaseIDLink                     string
		updateContractDocumentByDataProductIDLink         string
		updateContractTermsDocumentByDocumentIDLink       string
		updateDraftOfDataProductByDataProductIDLink       string
		updateReleaseOfDataProductByDataProductIDLink     string
		uploadContractTermsDocByDataProductIDLink         string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(dataproducthubapiservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			dataProductHubAPIServiceServiceOptions := &dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{}

			dataProductHubAPIService, err = dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(dataProductHubAPIServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(dataProductHubAPIService).ToNot(BeNil())
			Expect(dataProductHubAPIService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			dataProductHubAPIService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`Initialize - Initialize resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize(initializeOptions *InitializeOptions)`, func() {
			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID:   core.StringPtr("a7ca67e8-1fac-4061-ae9b-7604e15c4ab3"),
				Type: core.StringPtr("catalog"),
			}

			initializeOptions := &dataproducthubapiservicev1.InitializeOptions{
				Container: containerReferenceModel,
				Include:   []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"},
			}

			initializeResource, response, err := dataProductHubAPIService.Initialize(initializeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(initializeResource).ToNot(BeNil())

			createDraftByContainerIDLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved createDraftByContainerIDLink value: %v\n", createDraftByContainerIDLink)
			createDataProductByCatalogIDLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved createDataProductByCatalogIDLink value: %v\n", createDataProductByCatalogIDLink)
			getStatusByCatalogIDLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved getStatusByCatalogIDLink value: %v\n", getStatusByCatalogIDLink)
		})
	})

	Describe(`GetInitializeStatus - Get resource initialization status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
			getInitializeStatusOptions := &dataproducthubapiservicev1.GetInitializeStatusOptions{
				ContainerID: &getStatusByCatalogIDLink,
			}

			initializeResource, response, err := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
	})

	Describe(`GetServiceIDCredentials - Get service id credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceIDCredentials(getServiceIDCredentialsOptions *GetServiceIDCredentialsOptions)`, func() {
			getServiceIDCredentialsOptions := &dataproducthubapiservicev1.GetServiceIDCredentialsOptions{}

			serviceIDCredentials, response, err := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIDCredentials).ToNot(BeNil())
		})
	})

	Describe(`ManageAPIKeys - Rotate credentials for a Data Product Hub instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ManageAPIKeys(manageAPIKeysOptions *ManageAPIKeysOptions)`, func() {
			manageAPIKeysOptions := &dataproducthubapiservicev1.ManageAPIKeysOptions{}

			response, err := dataProductHubAPIService.ManageAPIKeys(manageAPIKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateDataProduct - Create a new data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID:   core.StringPtr(createDataProductByCatalogIDLink),
				Type: core.StringPtr("catalog"),
			}

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr(createDataProductByCatalogIDLink),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			assetPartReferenceModel := &dataproducthubapiservicev1.AssetPartReference{
				ID:        core.StringPtr("16a8f683-f947-48d9-a92c-b81758b1a5f5"),
				Container: containerReferenceModel,
				Type:      core.StringPtr("data_asset"),
			}

			deliveryMethodModel := &dataproducthubapiservicev1.DeliveryMethod{
				ID:        core.StringPtr("8848fd43-7384-4435-aff3-6a9f113768c4"),
				Container: containerReferenceModel,
			}

			dataProductPartModel := &dataproducthubapiservicev1.DataProductPart{
				Asset:           assetPartReferenceModel,
				DeliveryMethods: []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel},
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID:        core.StringPtr("3f0688f0-69c3-441e-b49b-7c223daa1804"),
				Name:      core.StringPtr("Risk Management"),
				Container: containerReferenceModel,
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductVersionPrototype{
				Version:      core.StringPtr("1.0.0"),
				State:        core.StringPtr("draft"),
				Name:         core.StringPtr("My New Data Product created using Go SDK"),
				Description:  core.StringPtr("This is a description of My Data Product."),
				Types:        []string{"data"},
				IsRestricted: core.BoolPtr(false),
				Asset:        assetPrototypeModel,
				Domain:       domainModel,
				PartsOut:     []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel},
			}

			createDataProductOptions := &dataproducthubapiservicev1.CreateDataProductOptions{
				Drafts: []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			}

			dataProduct, response, err := dataProductHubAPIService.CreateDataProduct(createDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProduct).ToNot(BeNil())

			createNewDraftByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved createNewDraftByDataProductIDLink value: %v\n", createNewDraftByDataProductIDLink)
			getContractDocumentByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractDocumentByDataProductIDLink value: %v\n", getContractDocumentByDataProductIDLink)
			retireAReleasesOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved retireAReleasesOfDataProductByDataProductIDLink value: %v\n", retireAReleasesOfDataProductByDataProductIDLink)
			getDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDataProductByDataProductIDLink value: %v\n", getDataProductByDataProductIDLink)
			updateDraftOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateDraftOfDataProductByDataProductIDLink value: %v\n", updateDraftOfDataProductByDataProductIDLink)
			updateContractDocumentByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractDocumentByDataProductIDLink value: %v\n", updateContractDocumentByDataProductIDLink)
			deleteDraftOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteDraftOfDataProductByDataProductIDLink value: %v\n", deleteDraftOfDataProductByDataProductIDLink)
			getAReleaseOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseOfDataProductByDataProductIDLink value: %v\n", getAReleaseOfDataProductByDataProductIDLink)
			completeDraftContractTermsByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeDraftContractTermsByDataProductIDLink value: %v\n", completeDraftContractTermsByDataProductIDLink)
			deleteContractDocumentByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractDocumentByDataProductIDLink value: %v\n", deleteContractDocumentByDataProductIDLink)
			getListOfDataProductDraftsByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getListOfDataProductDraftsByDataProductIDLink value: %v\n", getListOfDataProductDraftsByDataProductIDLink)
			getADraftOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftOfDataProductByDataProductIDLink value: %v\n", getADraftOfDataProductByDataProductIDLink)
			getReleaseContractDocumentByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getReleaseContractDocumentByDataProductIDLink value: %v\n", getReleaseContractDocumentByDataProductIDLink)
			publishADraftOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved publishADraftOfDataProductByDataProductIDLink value: %v\n", publishADraftOfDataProductByDataProductIDLink)
			getListOfReleasesOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getListOfReleasesOfDataProductByDataProductIDLink value: %v\n", getListOfReleasesOfDataProductByDataProductIDLink)
			updateReleaseOfDataProductByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateReleaseOfDataProductByDataProductIDLink value: %v\n", updateReleaseOfDataProductByDataProductIDLink)
			uploadContractTermsDocByDataProductIDLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved uploadContractTermsDocByDataProductIDLink value: %v\n", uploadContractTermsDocByDataProductIDLink)
			createAContractTermsDocByContractTermsIDLink = *dataProduct.Drafts[0].ContractTerms[0].ID
			getAReleaseContractTermsByContractTermsIDLink = *dataProduct.Drafts[0].ContractTerms[0].ID
			createAContractTermsDocByDraftIDLink = *dataProduct.Drafts[0].ID
			getDraftByDraftIDLink = *dataProduct.Drafts[0].ID
		})
	})

	Describe(`GetDataProduct - Retrieve a data product identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
			getDataProductOptions := &dataproducthubapiservicev1.GetDataProductOptions{
				DataProductID: &getDataProductByDataProductIDLink,
			}

			dataProduct, response, err := dataProductHubAPIService.GetDataProduct(getDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())
		})
	})

	Describe(`ListDataProducts - Retrieve a list of data products`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) with pagination`, func() {
			listDataProductsOptions := &dataproducthubapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			var allResults []dataproducthubapiservicev1.DataProductSummary
			for {
				dataProductSummaryCollection, response, err := dataProductHubAPIService.ListDataProducts(listDataProductsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductSummaryCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductSummaryCollection.DataProducts...)

				listDataProductsOptions.Start, err = dataProductSummaryCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) using DataProductsPager`, func() {
			listDataProductsOptions := &dataproducthubapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubAPIService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproducthubapiservicev1.DataProductSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dataProductHubAPIService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProducts() returned a total of %d item(s) using DataProductsPager.\n", len(allResults))
		})
	})

	Describe(`GetDataProductDraft - Get a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions)`, func() {
			getDataProductDraftOptions := &dataproducthubapiservicev1.GetDataProductDraftOptions{
				DataProductID: core.StringPtr("-"),
				DraftID:       &getDraftByDraftIDLink,
			}

			dataProductVersion, response, err := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductDraft - Update the data product draft identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductDraft(updateDataProductDraftOptions *UpdateDataProductDraftOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/description"),
				Value: "Updated the description.",
			}

			updateDataProductDraftOptions := &dataproducthubapiservicev1.UpdateDataProductDraftOptions{
				DataProductID:         core.StringPtr("-"),
				DraftID:               &getDraftByDraftIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`CreateDraftContractTermsDocument - Upload a contract document to the data product draft contract terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
			createDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions{
				DataProductID:   &uploadContractTermsDocByDataProductIDLink,
				DraftID:         &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				Type:            core.StringPtr("terms_and_conditions"),
				Name:            core.StringPtr("Terms and conditions document"),
				URL:             core.StringPtr("https://www.ibm.com/contract_document"),
			}

			contractTermsDocument, response, err := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			getReleaseContractDocumentByDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved getReleaseContractDocumentByDocumentIDLink value: %v\n", getReleaseContractDocumentByDocumentIDLink)
			deleteContractTermsDocumentByDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentByDocumentIDLink value: %v\n", deleteContractTermsDocumentByDocumentIDLink)
			getContractTermsDocumentByIDDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractTermsDocumentByIDDocumentIDLink value: %v\n", getContractTermsDocumentByIDDocumentIDLink)
			updateContractTermsDocumentByDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractTermsDocumentByDocumentIDLink value: %v\n", updateContractTermsDocumentByDocumentIDLink)
			completeContractTermsDocumentByDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeContractTermsDocumentByDocumentIDLink value: %v\n", completeContractTermsDocumentByDocumentIDLink)
		})
	})

	Describe(`GetDraftContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions)`, func() {
			getDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions{
				DataProductID:   &getContractDocumentByDataProductIDLink,
				DraftID:         &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				DocumentID:      &getContractTermsDocumentByIDDocumentIDLink,
			}

			contractTermsDocument, response, err := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`UpdateDraftContractTermsDocument - Update a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions *UpdateDraftContractTermsDocumentOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:    core.StringPtr("add"),
				Path:  core.StringPtr("/name"),
				Value: "updated Terms and Conditions",
			}

			updateDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions{
				DataProductID:         &getContractDocumentByDataProductIDLink,
				DraftID:               &createAContractTermsDocByDraftIDLink,
				ContractTermsID:       &createAContractTermsDocByContractTermsIDLink,
				DocumentID:            &getContractTermsDocumentByIDDocumentIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTermsDocument, response, err := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`PublishDataProductDraft - Publish a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
			publishDataProductDraftOptions := &dataproducthubapiservicev1.PublishDataProductDraftOptions{
				DataProductID: &publishADraftOfDataProductByDataProductIDLink,
				DraftID:       &getDraftByDraftIDLink,
			}

			dataProductVersion, response, err := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())

			updateAReleaseByReleaseIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateAReleaseByReleaseIDLink value: %v\n", updateAReleaseByReleaseIDLink)
			getAReleaseContractTermsByReleaseIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByReleaseIDLink value: %v\n", getAReleaseContractTermsByReleaseIDLink)
			retireAReleaseContractTermsByReleaseIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved retireAReleaseContractTermsByReleaseIDLink value: %v\n", retireAReleaseContractTermsByReleaseIDLink)
			getAReleaseByReleaseIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseByReleaseIDLink value: %v\n", getAReleaseByReleaseIDLink)
		})
	})

	Describe(`GetDataProductRelease - Get a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
			getDataProductReleaseOptions := &dataproducthubapiservicev1.GetDataProductReleaseOptions{
				DataProductID:       &getAReleaseOfDataProductByDataProductIDLink,
				ReleaseID:           &getAReleaseByReleaseIDLink,
				CheckCallerApproval: core.BoolPtr(false),
			}

			dataProductVersion, response, err := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductRelease - Update the data product release identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/description"),
				Value: "New description for my data product",
			}

			updateDataProductReleaseOptions := &dataproducthubapiservicev1.UpdateDataProductReleaseOptions{
				DataProductID:         &updateReleaseOfDataProductByDataProductIDLink,
				ReleaseID:             &getAReleaseByReleaseIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`GetReleaseContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions)`, func() {
			getReleaseContractTermsDocumentOptions := &dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions{
				DataProductID:   &getReleaseContractDocumentByDataProductIDLink,
				ReleaseID:       &getAReleaseContractTermsByReleaseIDLink,
				ContractTermsID: &getAReleaseContractTermsByContractTermsIDLink,
				DocumentID:      &getReleaseContractDocumentByDocumentIDLink,
			}

			contractTermsDocument, response, err := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`ListDataProductReleases - Retrieve a list of data product releases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) with pagination`, func() {
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: &createDataProductByCatalogIDLink,
				State:            []string{"available"},
				Limit:            core.Int64Ptr(int64(10)),
			}

			listDataProductReleasesOptions.Start = nil
			listDataProductReleasesOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for {
				dataProductReleaseCollection, response, err := dataProductHubAPIService.ListDataProductReleases(listDataProductReleasesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductReleaseCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductReleaseCollection.Releases...)

				listDataProductReleasesOptions.Start, err = dataProductReleaseCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductReleasesOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) using DataProductReleasesPager`, func() {
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: &createDataProductByCatalogIDLink,
				State:            []string{"available"},
				Limit:            core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductReleases() returned a total of %d item(s) using DataProductReleasesPager.\n", len(allResults))
		})
	})

	Describe(`RetireDataProductRelease - Retire a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RetireDataProductRelease(retireDataProductReleaseOptions *RetireDataProductReleaseOptions)`, func() {
			retireDataProductReleaseOptions := &dataproducthubapiservicev1.RetireDataProductReleaseOptions{
				DataProductID: &retireAReleasesOfDataProductByDataProductIDLink,
				ReleaseID:     &retireAReleaseContractTermsByReleaseIDLink,
			}

			dataProductVersion, response, err := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

	Describe(`CreateDataProductDraftForDeleteOp - Create a new data product draft`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr(createDataProductByCatalogIDLink),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductVersionPrototype{
				Version:     core.StringPtr("1.0.0"),
				State:       core.StringPtr("draft"),
				Name:        core.StringPtr("New Delete Draft DP created from IntegrationTest Go SDK"),
				Description: core.StringPtr("This is a description of My Data Product which will get deleted."),
				Types:       []string{"data"},
				Asset:       assetPrototypeModel,
			}

			createDataProductOptions := &dataproducthubapiservicev1.CreateDataProductOptions{
				Drafts: []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			}

			dataProduct, response, err := dataProductHubAPIService.CreateDataProduct(createDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProduct).ToNot(BeNil())

			deleteAContractDocumentByDraftIDLink = *dataProduct.Drafts[0].ID
			deleteADraftByContractTermsIDLink = *dataProduct.Drafts[0].ContractTerms[0].ID
			createAContractTermsDocByContractTermsIDLink = *dataProduct.Drafts[0].ContractTerms[0].ID
			deleteADraftByDraftIDLink = *dataProduct.Drafts[0].ID
			createAContractTermsDocByDraftIDLink = *dataProduct.Drafts[0].ID
		})
	})

	Describe(`ListDataProductDrafts - Retrieve a list of data product drafts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) with pagination`, func() {
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID:    core.StringPtr("-"),
				AssetContainerID: &createDataProductByCatalogIDLink,
				Limit:            core.Int64Ptr(int64(10)),
			}

			listDataProductDraftsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for {
				dataProductDraftCollection, response, err := dataProductHubAPIService.ListDataProductDrafts(listDataProductDraftsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductDraftCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductDraftCollection.Drafts...)

				listDataProductDraftsOptions.Start, err = dataProductDraftCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductDraftsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) using DataProductDraftsPager`, func() {
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID:    core.StringPtr("-"),
				AssetContainerID: &createDataProductByCatalogIDLink,
				Limit:            core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductDrafts() returned a total of %d item(s) using DataProductDraftsPager.\n", len(allResults))
		})
	})

	Describe(`CreateDraftContractTermsDocumentForDeleteOp - Create a contract document for Delete Option`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
			createDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions{
				DataProductID:   core.StringPtr("-"),
				DraftID:         &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				Type:            core.StringPtr("terms_and_conditions"),
				Name:            core.StringPtr("Terms and conditions document"),
				URL:             core.StringPtr("https://data.un.org/Host.aspx?Content=UNdataUse"),
			}

			contractTermsDocument, response, err := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			deleteContractTermsDocumentByDocumentIDLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentByDocumentIDLink value: %v\n", deleteContractTermsDocumentByDocumentIDLink)
		})
	})

	Describe(`DeleteDraftContractTermsDocument - Delete a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
			deleteDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions{
				DataProductID:   core.StringPtr("-"),
				DraftID:         &deleteAContractDocumentByDraftIDLink,
				ContractTermsID: &deleteADraftByContractTermsIDLink,
				DocumentID:      &deleteContractTermsDocumentByDocumentIDLink,
			}

			response, err := dataProductHubAPIService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDataProductDraft - Delete a data product draft identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions)`, func() {
			deleteDataProductDraftOptions := &dataproducthubapiservicev1.DeleteDataProductDraftOptions{
				DataProductID: core.StringPtr("-"),
				DraftID:       &deleteADraftByDraftIDLink,
			}

			response, err := dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

})

//
// Utility functions are declared in the unit test file
//
