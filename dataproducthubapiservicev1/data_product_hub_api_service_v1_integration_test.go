//go:build integration

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

package dataproducthubapiservicev1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/data-product-exchange-go-sdk/v4/dataproducthubapiservicev1"
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
		err          error
		dataProductHubAPIService *dataproducthubapiservicev1.DataProductHubAPIServiceV1
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		completeContractTermsDocumentByDocumentIDLink string
		completeDraftContractTermsByDataProductIDLink string
		createAContractTermsDocByContractTermsIDLink string
		createAContractTermsDocByDraftIDLink string
		createDataProductByCatalogIDLink string
		createDraftByContainerIDLink string
		createNewDraftByDataProductIDLink string
		createContractTemplateID string
		deleteAContractDocumentByDraftIDLink string
		deleteADraftByContractTermsIDLink string
		deleteADraftByDraftIDLink string
		deleteContractDocumentByDataProductIDLink string
		deleteContractTermsDocumentByDocumentIDLink string
		deleteDraftOfDataProductByDataProductIDLink string
		getADraftOfDataProductByDataProductIDLink string
		getAReleaseByReleaseIDLink string
		getAReleaseContractTermsByContractTermsIDLink string
		getAReleaseContractTermsByReleaseIDLink string
		getAReleaseOfDataProductByDataProductIDLink string
		getContractDocumentByDataProductIDLink string
		getContractTermsDocumentByIDDocumentIDLink string
		getDataProductByDataProductIDLink string
		getDraftByDraftIDLink string
		getListOfDataProductDraftsByDataProductIDLink string
		getListOfReleasesOfDataProductByDataProductIDLink string
		getReleaseContractDocumentByDataProductIDLink string
		getReleaseContractDocumentByDocumentIDLink string
		getStatusByCatalogIDLink string
		publishADraftOfDataProductByDataProductIDLink string
		retireAReleaseContractTermsByReleaseIDLink string
		retireAReleasesOfDataProductByDataProductIDLink string
		updateAReleaseByReleaseIDLink string
		updateContractDocumentByDataProductIDLink string
		updateContractTermsDocumentByDocumentIDLink string
		updateDraftOfDataProductByDataProductIDLink string
		updateReleaseOfDataProductByDataProductIDLink string
		uploadContractTermsDocByDataProductIDLink string
		createDataProductDomainID string
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
			dataProductHubAPIServiceOptions := &dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{}

			dataProductHubAPIService, err = dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(dataProductHubAPIServiceOptions)
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
				ID: core.StringPtr("8aa199e8-71e0-4710-8263-e64084e8d30c"),
				Type: core.StringPtr("catalog"),
			}

			initializeOptions := &dataproducthubapiservicev1.InitializeOptions{
				Container: containerReferenceModel,
				Include: []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"},
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
			getServiceIDCredentialsOptions := &dataproducthubapiservicev1.GetServiceIDCredentialsOptions{
			}

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
			manageAPIKeysOptions := &dataproducthubapiservicev1.ManageAPIKeysOptions{
			}

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
				ID:        core.StringPtr("ac7c75cb-e188-4b06-b1b5-1b8dcfa41cf6"),
				Container: containerReferenceModel,
				Type:      core.StringPtr("data_asset"),
			}

			deliveryMethodModel := &dataproducthubapiservicev1.DeliveryMethod{
				ID:        core.StringPtr("ff700594-f4cb-44b6-9735-11c085df8eef"),
				Container: containerReferenceModel,
			}

			dataProductPartModel := &dataproducthubapiservicev1.DataProductPart{
				Asset:           assetPartReferenceModel,
				DeliveryMethods: []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel},
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID:        core.StringPtr("e26d713d-c3a6-4c74-a235-bd5f4cfa8a70"),
				Name:      core.StringPtr("Sustainability"),
				Container: containerReferenceModel,
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductDraftPrototype{
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
				Drafts: []dataproducthubapiservicev1.DataProductDraftPrototype{*dataProductVersionPrototypeModel},
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
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) with pagination`, func(){
			listDataProductsOptions := &dataproducthubapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			var allResults []dataproducthubapiservicev1.DataProductSummary
			for {
				dataProductCollection, response, err := dataProductHubAPIService.ListDataProducts(listDataProductsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(dataProductCollection).ToNot(BeNil())
				allResults = append(allResults, dataProductCollection.DataProducts...)

				listDataProductsOptions.Start, err = dataProductCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listDataProductsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListDataProducts(listDataProductsOptions *ListDataProductsOptions) using DataProductsPager`, func(){
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
				DataProductID: &getADraftOfDataProductByDataProductIDLink,
				DraftID: &getDraftByDraftIDLink,
			}

			dataProductDraft, response, err := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDraft).ToNot(BeNil())
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
				DataProductID:core.StringPtr("-"),
				DraftID: &createAContractTermsDocByDraftIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductDraft, response, err := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDraft).ToNot(BeNil())
		})
	})

	Describe(`CreateDraftContractTermsDocument - Upload a contract document to the data product draft contract terms`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions *CreateDraftContractTermsDocumentOptions)`, func() {
			createDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.CreateDraftContractTermsDocumentOptions{
				DataProductID: &uploadContractTermsDocByDataProductIDLink,
				DraftID: &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("Terms and conditions document"),
				URL: core.StringPtr("https://www.ibm.com/contract_document"),
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
				Op: core.StringPtr("add"),
				Path: core.StringPtr("/name"),
				Value: "updated terms and condition",
			}

			updateDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions{
				DataProductID: &getContractDocumentByDataProductIDLink,
				DraftID: &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				DocumentID: &getContractTermsDocumentByIDDocumentIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTermsDocument, response, err := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})
	

	Describe(`GetDataProductDraftContractTerms - Retrieve a data product contract terms identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductDraftContractTerms(getDataProductDraftContractTermsOptions *GetDataProductDraftContractTermsOptions)`, func() {
			getDataProductDraftContractTermsOptions := &dataproducthubapiservicev1.GetDataProductDraftContractTermsOptions{
				DataProductID: &getContractDocumentByDataProductIDLink,
				DraftID: &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				IncludeContractDocuments: core.BoolPtr(true),
			}

			result, response, err := dataProductHubAPIService.GetDataProductDraftContractTerms(getDataProductDraftContractTermsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`PublishDataProductDraft - Publish a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
			publishDataProductDraftOptions := &dataproducthubapiservicev1.PublishDataProductDraftOptions{
				DataProductID: &publishADraftOfDataProductByDataProductIDLink,
				DraftID: &getDraftByDraftIDLink,
			}

			dataProductRelease, response, err := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())

			updateAReleaseByReleaseIDLink = *dataProductRelease.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateAReleaseByReleaseIDLink value: %v\n", updateAReleaseByReleaseIDLink)
			getAReleaseContractTermsByReleaseIDLink = *dataProductRelease.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByReleaseIDLink value: %v\n", getAReleaseContractTermsByReleaseIDLink)
			retireAReleaseContractTermsByReleaseIDLink = *dataProductRelease.ID
			fmt.Fprintf(GinkgoWriter, "Saved retireAReleaseContractTermsByReleaseIDLink value: %v\n", retireAReleaseContractTermsByReleaseIDLink)
			getAReleaseByReleaseIDLink = *dataProductRelease.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseByReleaseIDLink value: %v\n", getAReleaseByReleaseIDLink)
		})
	})


	Describe(`GetDataProductRelease - Get a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
			getDataProductReleaseOptions := &dataproducthubapiservicev1.GetDataProductReleaseOptions{
				DataProductID: &getAReleaseOfDataProductByDataProductIDLink,
				ReleaseID: &getAReleaseByReleaseIDLink,
				CheckCallerApproval: core.BoolPtr(false),
			}

			dataProductRelease, response, err := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductRelease - Update the data product release identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductRelease(updateDataProductReleaseOptions *UpdateDataProductReleaseOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("replace"),
				Path: core.StringPtr("/description"),
				
				Value: "New description for my data product",
			}

			updateDataProductReleaseOptions := &dataproducthubapiservicev1.UpdateDataProductReleaseOptions{
				DataProductID: &updateReleaseOfDataProductByDataProductIDLink,
				ReleaseID: &getAReleaseByReleaseIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductRelease, response, err := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
	})


	Describe(`GetReleaseContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions *GetReleaseContractTermsDocumentOptions)`, func() {
			getReleaseContractTermsDocumentOptions := &dataproducthubapiservicev1.GetReleaseContractTermsDocumentOptions{
				DataProductID: &getReleaseContractDocumentByDataProductIDLink,
				ReleaseID: &getAReleaseContractTermsByReleaseIDLink,
				ContractTermsID: &getAReleaseContractTermsByContractTermsIDLink,
				DocumentID: &getReleaseContractDocumentByDocumentIDLink,
			}

			contractTermsDocument, response, err := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`ReplaceDataProductDraftContractTerms - Update a data product contract terms identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceDataProductDraftContractTerms(replaceDataProductDraftContractTermsOptions *ReplaceDataProductDraftContractTermsOptions)`, func() {

			contractTermsDocumentAttachmentModel := &dataproducthubapiservicev1.ContractTermsDocumentAttachment{
				ID: &getContractTermsDocumentByIDDocumentIDLink,
			}

			contractTermsDocumentModel := &dataproducthubapiservicev1.ContractTermsDocument{
				URL: core.StringPtr("https://ibm.com/document"),
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("Terms and Conditions"),
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Attachment: contractTermsDocumentAttachmentModel,
				
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Name: core.StringPtr("domain_name"),
				
			}

			overviewModel := &dataproducthubapiservicev1.Overview{
				APIVersion: core.StringPtr("v3.0.1"),
				Kind: core.StringPtr("DataContract"),
				Name: core.StringPtr("Sample Data Contract"),
				Version: core.StringPtr("v0.0"),
				Domain: domainModel,
				MoreInfo: core.StringPtr("List of links to sources that provide more details on the data contract."),
			}

			contractTermsMoreInfoModel := &dataproducthubapiservicev1.ContractTermsMoreInfo{
				Type: core.StringPtr("privacy-statement"),
				URL: core.StringPtr("https://www.moreinfo.example.coms"),
			}

			descriptionModel := &dataproducthubapiservicev1.Description{
				Purpose: core.StringPtr("Intended purpose for the provided data."),
				Limitations: core.StringPtr("Technical, compliance, and legal limitations for data use."),
				Usage: core.StringPtr("Recommended usage of the data."),
				MoreInfo: []dataproducthubapiservicev1.ContractTermsMoreInfo{*contractTermsMoreInfoModel},
				CustomProperties: core.StringPtr("Custom properties that are not part of the standard."),
			}

			contractTemplateOrganizationModel := &dataproducthubapiservicev1.ContractTemplateOrganization{
				UserID: core.StringPtr("IBMid-691000IN4G"),
				Role: core.StringPtr("owner"),
			}

			rolesModel := &dataproducthubapiservicev1.Roles{
				Role: core.StringPtr("IAM Role"),
			}

			pricingModel := &dataproducthubapiservicev1.Pricing{
				Amount: core.StringPtr("Amount"),
				Currency: core.StringPtr("Currency"),
				Unit: core.StringPtr("Unit"),
			}

			contractTemplateSLAPropertyModel := &dataproducthubapiservicev1.ContractTemplateSLAProperty{
				Property: core.StringPtr("slaproperty"),
				Value: core.StringPtr("slavalue"),
			}

			contractTemplateSLAModel := &dataproducthubapiservicev1.ContractTemplateSLA{
				DefaultElement: core.StringPtr("sladefaultelement"),
				Properties: []dataproducthubapiservicev1.ContractTemplateSLAProperty{*contractTemplateSLAPropertyModel},
			}

			contractTemplateSupportAndCommunicationModel := &dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{
				Channel: core.StringPtr("channel"),
				URL: core.StringPtr("https://www.example.coms"),
			}

			contractTemplateCustomPropertyModel := &dataproducthubapiservicev1.ContractTemplateCustomProperty{
				Key: core.StringPtr("The name of the key."),
				Value: core.StringPtr("The value of the key."),
			}

			contractTestModel := &dataproducthubapiservicev1.ContractTest{
				Status: core.StringPtr("pass"),
				LastTestedTime: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
			}

			replaceDataProductDraftContractTermsOptions := &dataproducthubapiservicev1.ReplaceDataProductDraftContractTermsOptions{
				DataProductID: &getContractDocumentByDataProductIDLink,
				DraftID: &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				
				ID: core.StringPtr("testString"),
				Documents: []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel},
				ErrorMsg: core.StringPtr("testString"),
				Overview: overviewModel,
				Description: descriptionModel,
				Organization: []dataproducthubapiservicev1.ContractTemplateOrganization{*contractTemplateOrganizationModel},
				Roles: []dataproducthubapiservicev1.Roles{*rolesModel},
				Price: pricingModel,
				SLA: []dataproducthubapiservicev1.ContractTemplateSLA{*contractTemplateSLAModel},
				SupportAndCommunication: []dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{*contractTemplateSupportAndCommunicationModel},
				CustomProperties: []dataproducthubapiservicev1.ContractTemplateCustomProperty{*contractTemplateCustomPropertyModel},
				ContractTest: contractTestModel,
				
			}

			contractTerms, response, err := dataProductHubAPIService.ReplaceDataProductDraftContractTerms(replaceDataProductDraftContractTermsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTerms).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductDraftContractTerms - Update a contract terms property`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductDraftContractTerms(updateDataProductDraftContractTermsOptions *UpdateDataProductDraftContractTermsOptions)`, func() {

			overviewMap := map[string]interface{}{
				"name":        "Sample Data Contract",
				"version":     "v0.0",
				"more_info":   "List of links to sources that provide more details on the data contract.",
				"domain": map[string]interface{}{
				    "id":   "b38df608-d34b-4d58-8136-ed25e6c6684e",
					"name": "domain_name",
				    },
			}

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/overview"),
				Value: overviewMap,
			}

			updateDataProductDraftContractTermsOptions := &dataproducthubapiservicev1.UpdateDataProductDraftContractTermsOptions{
				DataProductID: &getContractDocumentByDataProductIDLink,
				DraftID: &createAContractTermsDocByDraftIDLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTerms, response, err := dataProductHubAPIService.UpdateDataProductDraftContractTerms(updateDataProductDraftContractTermsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTerms).ToNot(BeNil())
		})
	})
	Describe(`ListDataProductReleases - Retrieve a list of data product releases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) with pagination`, func(){
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID: &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: &createDataProductByCatalogIDLink,
				State: []string{"available"},
				
				Limit: core.Int64Ptr(int64(10)),
				
			}

			listDataProductReleasesOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductReleaseSummary
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
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) using DataProductReleasesPager`, func(){
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID: &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: &createDataProductByCatalogIDLink,
				State: []string{"available"},
				
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproducthubapiservicev1.DataProductReleaseSummary
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
				ReleaseID: &retireAReleaseContractTermsByReleaseIDLink,
				RevokeAccess: core.BoolPtr(false),
			}

			dataProductRelease, response, err := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
	})

	Describe(`CreateDataAssetVisualization - Create visualization asset and initialize profiling for the provided data assets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataAssetVisualization(createDataAssetVisualizationOptions *CreateDataAssetVisualizationOptions)`, func() {
			visualizationModel := &dataproducthubapiservicev1.Visualization{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("2be8f727-c5d2-4cb0-9216-f9888f428048"),
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dataproducthubapiservicev1.AssetReference{
				ID: core.StringPtr("caeee3f3-756e-47d5-846d-da4600809e22"),
				Name: core.StringPtr("testString"),
				Container: containerReferenceModel,
			}

			errorMessageModel := &dataproducthubapiservicev1.ErrorMessage{
				Code: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
			}

			dataAssetRelationshipModel := &dataproducthubapiservicev1.DataAssetRelationship{
				Visualization: visualizationModel,
				Asset: assetReferenceModel,
				RelatedAsset: assetReferenceModel,
				Error: errorMessageModel,
			}

			createDataAssetVisualizationOptions := &dataproducthubapiservicev1.CreateDataAssetVisualizationOptions{
				Assets: []dataproducthubapiservicev1.DataAssetRelationship{*dataAssetRelationshipModel},
			}

			dataAssetVisualizationRes, response, err := dataProductHubAPIService.CreateDataAssetVisualization(createDataAssetVisualizationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataAssetVisualizationRes).ToNot(BeNil())
		})
	})

	Describe(`ReinitiateDataAssetVisualization - Reinitiate visualization for an asset`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReinitiateDataAssetVisualization(reinitiateDataAssetVisualizationOptions *ReinitiateDataAssetVisualizationOptions)`, func() {
			visualizationModel := &dataproducthubapiservicev1.Visualization{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
			}

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("2be8f727-c5d2-4cb0-9216-f9888f428048"),
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dataproducthubapiservicev1.AssetReference{
				ID: core.StringPtr("caeee3f3-756e-47d5-846d-da4600809e22"),
				Name: core.StringPtr("testString"),
				Container: containerReferenceModel,
			}

			errorMessageModel := &dataproducthubapiservicev1.ErrorMessage{
				Code: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
			}

			dataAssetRelationshipModel := &dataproducthubapiservicev1.DataAssetRelationship{
				Visualization: visualizationModel,
				Asset: assetReferenceModel,
				RelatedAsset: assetReferenceModel,
				Error: errorMessageModel,
			}

			reinitiateDataAssetVisualizationOptions := &dataproducthubapiservicev1.ReinitiateDataAssetVisualizationOptions{
				Assets: []dataproducthubapiservicev1.DataAssetRelationship{*dataAssetRelationshipModel},
			}

			dataAssetVisualizationRes, response, err := dataProductHubAPIService.ReinitiateDataAssetVisualization(reinitiateDataAssetVisualizationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataAssetVisualizationRes).ToNot(BeNil())
		})
	})

	Describe(`CreateContractTemplate - Create new data product contract template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateContractTemplate(createContractTemplateOptions *CreateContractTemplateOptions)`, func() {
			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: &getStatusByCatalogIDLink,
				Type: core.StringPtr("catalog"),
			}

			errorMessageModel := &dataproducthubapiservicev1.ErrorMessage{
				Code: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
			}

			assetReferenceModel := &dataproducthubapiservicev1.AssetReference{
				ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Name: core.StringPtr("testString"),
				Container: containerReferenceModel,
			}

			contractTermsDocumentAttachmentModel := &dataproducthubapiservicev1.ContractTermsDocumentAttachment{
				ID: core.StringPtr("testString"),
			}

			contractTermsDocumentModel := &dataproducthubapiservicev1.ContractTermsDocument{
				URL: core.StringPtr("testString"),
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("testString"),
				ID: core.StringPtr("2b0bf220-079c-11ee-be56-0242ac120002"),
				Attachment: contractTermsDocumentAttachmentModel,
				UploadURL: core.StringPtr("testString"),
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Name: core.StringPtr("domain_name"),
				Container: containerReferenceModel,
			}

			overviewModel := &dataproducthubapiservicev1.Overview{
				APIVersion: core.StringPtr("v3.0.1"),
				Kind: core.StringPtr("DataContract"),
				Name: core.StringPtr("Sample Data Contract"),
				Version: core.StringPtr("0.0.0"),
				Domain: domainModel,
				MoreInfo: core.StringPtr("List of links to sources that provide more details on the data contract."),
			}

			contractTermsMoreInfoModel := &dataproducthubapiservicev1.ContractTermsMoreInfo{
				Type: core.StringPtr("privacy-statement"),
				URL: core.StringPtr("https://www.moreinfo.example.coms"),
			}

			descriptionModel := &dataproducthubapiservicev1.Description{
				Purpose: core.StringPtr("Intended purpose for the provided data."),
				Limitations: core.StringPtr("Technical, compliance, and legal limitations for data use."),
				Usage: core.StringPtr("Recommended usage of the data."),
				MoreInfo: []dataproducthubapiservicev1.ContractTermsMoreInfo{*contractTermsMoreInfoModel},
				CustomProperties: core.StringPtr("Custom properties that are not part of the standard."),
			}

			contractTemplateOrganizationModel := &dataproducthubapiservicev1.ContractTemplateOrganization{
				UserID: core.StringPtr("IBMid-691000IN4G"),
				Role: core.StringPtr("owner"),
			}

			rolesModel := &dataproducthubapiservicev1.Roles{
				Role: core.StringPtr("IAM Role"),
			}

			pricingModel := &dataproducthubapiservicev1.Pricing{
				Amount: core.StringPtr("100.00"),
				Currency: core.StringPtr("USD"),
				Unit: core.StringPtr("megabyte"),
			}

			contractTemplateSLAPropertyModel := &dataproducthubapiservicev1.ContractTemplateSLAProperty{
				Property: core.StringPtr("slaproperty"),
				Value: core.StringPtr("slavalue"),
			}

			contractTemplateSLAModel := &dataproducthubapiservicev1.ContractTemplateSLA{
				DefaultElement: core.StringPtr("sladefaultelement"),
				Properties: []dataproducthubapiservicev1.ContractTemplateSLAProperty{*contractTemplateSLAPropertyModel},
			}

			contractTemplateSupportAndCommunicationModel := &dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{
				Channel: core.StringPtr("channel"),
				URL: core.StringPtr("https://www.example.coms"),
			}

			contractTemplateCustomPropertyModel := &dataproducthubapiservicev1.ContractTemplateCustomProperty{
				Key: core.StringPtr("propertykey"),
				Value: core.StringPtr("propertyvalue"),
			}

			contractTestModel := &dataproducthubapiservicev1.ContractTest{
				Status: core.StringPtr("pass"),
				LastTestedTime: core.StringPtr("testString"),
				Message: core.StringPtr("testString"),
			}

			contractSchemaPropertyTypeModel := &dataproducthubapiservicev1.ContractSchemaPropertyType{
				Type: core.StringPtr("testString"),
				Length: core.StringPtr("testString"),
				Scale: core.StringPtr("testString"),
				Nullable: core.StringPtr("testString"),
				Signed: core.StringPtr("testString"),
				NativeType: core.StringPtr("testString"),
			}

			contractSchemaPropertyModel := &dataproducthubapiservicev1.ContractSchemaProperty{
				Name: core.StringPtr("testString"),
				Type: contractSchemaPropertyTypeModel,
			}

			contractSchemaModel := &dataproducthubapiservicev1.ContractSchema{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				PhysicalType: core.StringPtr("testString"),
				Properties: []dataproducthubapiservicev1.ContractSchemaProperty{*contractSchemaPropertyModel},
			}

			contractTermsModel := &dataproducthubapiservicev1.ContractTerms{
				Asset: assetReferenceModel,
				ID: core.StringPtr("testString"),
				Documents: []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel},
				ErrorMsg: core.StringPtr("testString"),
				Overview: overviewModel,
				Description: descriptionModel,
				Organization: []dataproducthubapiservicev1.ContractTemplateOrganization{*contractTemplateOrganizationModel},
				Roles: []dataproducthubapiservicev1.Roles{*rolesModel},
				Price: pricingModel,
				SLA: []dataproducthubapiservicev1.ContractTemplateSLA{*contractTemplateSLAModel},
				SupportAndCommunication: []dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{*contractTemplateSupportAndCommunicationModel},
				CustomProperties: []dataproducthubapiservicev1.ContractTemplateCustomProperty{*contractTemplateCustomPropertyModel},
				ContractTest: contractTestModel,
				Schema: []dataproducthubapiservicev1.ContractSchema{*contractSchemaModel},
			}

			createContractTemplateOptions := &dataproducthubapiservicev1.CreateContractTemplateOptions{
				Container: containerReferenceModel,
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("Sample Data Contract Template"),
				Error: errorMessageModel,
				ContractTerms: contractTermsModel,
				ContainerID: &getStatusByCatalogIDLink,
				ContractTemplateName: core.StringPtr("testString"),
			}

			dataProductContractTemplate, response, err := dataProductHubAPIService.CreateContractTemplate(createContractTemplateOptions)
			createContractTemplateID = *dataProductContractTemplate.ID


			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
	})


	Describe(`GetContractTemplate - Retrieve a data product contract template identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetContractTemplate(getContractTemplateOptions *GetContractTemplateOptions)`, func() {
			getContractTemplateOptions := &dataproducthubapiservicev1.GetContractTemplateOptions{
				ContractTemplateID: &createContractTemplateID,
				ContainerID: &getStatusByCatalogIDLink,
			}

			dataProductContractTemplate, response, err := dataProductHubAPIService.GetContractTemplate(getContractTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductContractTemplate - Update the data product contract template identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductContractTemplate(updateDataProductContractTemplateOptions *UpdateDataProductContractTemplateOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("replace"),
				Path: core.StringPtr("/name"),
				
				Value: "contract template name ",
			}

			updateDataProductContractTemplateOptions := &dataproducthubapiservicev1.UpdateDataProductContractTemplateOptions{
				ContractTemplateID: &createContractTemplateID,
				ContainerID: &getStatusByCatalogIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductContractTemplate, response, err := dataProductHubAPIService.UpdateDataProductContractTemplate(updateDataProductContractTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
	})

	Describe(`ListDataProductContractTemplate - Retrieve a list of data product contract templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductContractTemplate(listDataProductContractTemplateOptions *ListDataProductContractTemplateOptions)`, func() {
			listDataProductContractTemplateOptions := &dataproducthubapiservicev1.ListDataProductContractTemplateOptions{
				ContainerID: &getStatusByCatalogIDLink,
				ContractTemplateName: core.StringPtr("testString"),
			}

			dataProductContractTemplateCollection, response, err := dataProductHubAPIService.ListDataProductContractTemplate(listDataProductContractTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplateCollection).ToNot(BeNil())
		})
	})

	Describe(`DeleteDataProductContractTemplate - Delete a data product contract template identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductContractTemplate(deleteDataProductContractTemplateOptions *DeleteDataProductContractTemplateOptions)`, func() {
			deleteDataProductContractTemplateOptions := &dataproducthubapiservicev1.DeleteDataProductContractTemplateOptions{
				ContractTemplateID: &createContractTemplateID,
				ContainerID: &getStatusByCatalogIDLink,
			}

			response, err := dataProductHubAPIService.DeleteDataProductContractTemplate(deleteDataProductContractTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
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

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductDraftPrototype{
				Version:     core.StringPtr("1.0.0"),
				State:       core.StringPtr("draft"),
				Name:        core.StringPtr("New Delete Draft DP created from IntegrationTest Go SDK"),
				Description: core.StringPtr("This is a description of My Data Product which will get deleted."),
				Types:       []string{"data"},
				Asset:       assetPrototypeModel,
			}

			createDataProductOptions := &dataproducthubapiservicev1.CreateDataProductOptions{
				Drafts: []dataproducthubapiservicev1.DataProductDraftPrototype{*dataProductVersionPrototypeModel},
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

	Describe(`ListDataProductDrafts - Retrieve a list of data product drafts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) with pagination`, func(){
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID: core.StringPtr("-"),
				AssetContainerID: &createDataProductByCatalogIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			listDataProductDraftsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductDraftSummary
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
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) using DataProductDraftsPager`, func(){
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID: core.StringPtr("-"),
				AssetContainerID: &createDataProductByCatalogIDLink,
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []dataproducthubapiservicev1.DataProductDraftSummary
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

	Describe(`ListDataProductDomains - Retrieve a list of data product domains`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductDomains(listDataProductDomainsOptions *ListDataProductDomainsOptions)`, func() {
			listDataProductDomainsOptions := &dataproducthubapiservicev1.ListDataProductDomainsOptions{
				ContainerID: &createDataProductByCatalogIDLink,
			}

			dataProductDomainCollection, response, err := dataProductHubAPIService.ListDataProductDomains(listDataProductDomainsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomainCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDataProductDomain - Create new data product domain`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductDomain(createDataProductDomainOptions *CreateDataProductDomainOptions)`, func() {
			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: &getStatusByCatalogIDLink,
				Type: core.StringPtr("catalog"),
			}

			initializeSubDomainModel := &dataproducthubapiservicev1.InitializeSubDomain{
				Name: core.StringPtr("Sub domain 1"),
				ID: core.StringPtr("testString"),
				Description: core.StringPtr("New sub domain 1"),
			}

			createDataProductDomainOptions := &dataproducthubapiservicev1.CreateDataProductDomainOptions{
				Container: containerReferenceModel,
				Trace: core.StringPtr("testString"),
				Name: core.StringPtr("Test domain"),
				Description: core.StringPtr("The sample description for new domain"),
				ID: &getStatusByCatalogIDLink,
				SubDomains: []dataproducthubapiservicev1.InitializeSubDomain{*initializeSubDomainModel},
				ContainerID: core.StringPtr("testString"),
			}

			dataProductDomain, response, err := dataProductHubAPIService.CreateDataProductDomain(createDataProductDomainOptions)
			createDataProductDomainID = *dataProductDomain.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductDomain).ToNot(BeNil())
		})
	})

	Describe(`CreateDataProductSubdomain - Create data product subdomains for a specific domain identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductSubdomain(createDataProductSubdomainOptions *CreateDataProductSubdomainOptions)`, func() {
			createDataProductSubdomainOptions := &dataproducthubapiservicev1.CreateDataProductSubdomainOptions{
				DomainID: &createDataProductDomainID,
				ContainerID: &getStatusByCatalogIDLink,
				Name: core.StringPtr("Sub domain 2"),
				
				Description: core.StringPtr("New sub domain 2"),
			}

			initializeSubDomain, response, err := dataProductHubAPIService.CreateDataProductSubdomain(createDataProductSubdomainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(initializeSubDomain).ToNot(BeNil())
		})
	})

	Describe(`GetDomain - Retrieve a data product domain or subdomain identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDomain(getDomainOptions *GetDomainOptions)`, func() {
			getDomainOptions := &dataproducthubapiservicev1.GetDomainOptions{
			DomainID: &createDataProductDomainID,
			}

			dataProductDomain, response, err := dataProductHubAPIService.GetDomain(getDomainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomain).ToNot(BeNil())
		})
	})

	Describe(`UpdateDataProductDomain - Update the data product domain identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDataProductDomain(updateDataProductDomainOptions *UpdateDataProductDomainOptions)`, func() {
			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("replace"),
				Path: core.StringPtr("/name"),
				Value: "update domain name ",
			}

			updateDataProductDomainOptions := &dataproducthubapiservicev1.UpdateDataProductDomainOptions{
				DomainID: &createDataProductDomainID,
				ContainerID: &getStatusByCatalogIDLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductDomain, response, err := dataProductHubAPIService.UpdateDataProductDomain(updateDataProductDomainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomain).ToNot(BeNil())
		})
	})

	Describe(`GetDataProductByDomain - Retrieve all data products in a domain specified by id or any of it's subdomains`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductByDomain(getDataProductByDomainOptions *GetDataProductByDomainOptions)`, func() {
			getDataProductByDomainOptions := &dataproducthubapiservicev1.GetDataProductByDomainOptions{
				DomainID: &createDataProductDomainID,
				ContainerID: &getStatusByCatalogIDLink,
			}

			dataProductVersionCollection, response, err := dataProductHubAPIService.GetDataProductByDomain(getDataProductByDomainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersionCollection).ToNot(BeNil())
		})
	})

	Describe(`DeleteDomain - Delete a data product domain identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDomain(deleteDomainOptions *DeleteDomainOptions)`, func() {
			deleteDomainOptions := &dataproducthubapiservicev1.DeleteDomainOptions{
				DomainID: &createDataProductDomainID,
			}

			response, err := dataProductHubAPIService.DeleteDomain(deleteDomainOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteDraftContractTermsDocument - Delete a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
			deleteDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions{
				DataProductID: core.StringPtr("-"),
				DraftID: &deleteAContractDocumentByDraftIDLink,
				ContractTermsID: &deleteADraftByContractTermsIDLink,
				DocumentID: &deleteContractTermsDocumentByDocumentIDLink,
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
				DraftID: &deleteADraftByDraftIDLink,
			}

			response, err := dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateS3Bucket - Create a new Bucket`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateS3Bucket(createS3BucketOptions *CreateS3BucketOptions)`, func() {
			Skip("Skipping CreateS3Bucket test")
			createS3BucketOptions := &dataproducthubapiservicev1.CreateS3BucketOptions{
				IsShared: core.BoolPtr(true),
			}

			bucketResponse, response, err := dataProductHubAPIService.CreateS3Bucket(createS3BucketOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bucketResponse).ToNot(BeNil())
		})
	})

	Describe(`GetS3BucketValidation - Validate the Bucket Existence`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetS3BucketValidation(getS3BucketValidationOptions *GetS3BucketValidationOptions)`, func() {
			Skip("Skipping bucket validation test")

			getS3BucketValidationOptions := &dataproducthubapiservicev1.GetS3BucketValidationOptions{
				BucketName: core.StringPtr("testString"),
			}

			bucketValidationResponse, response, err := dataProductHubAPIService.GetS3BucketValidation(getS3BucketValidationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketValidationResponse).ToNot(BeNil())
		})
	})
})
