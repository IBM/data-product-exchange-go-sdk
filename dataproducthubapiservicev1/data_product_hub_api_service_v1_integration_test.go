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

var _ = Describe(`DataProductHubApiServiceV1 Integration Tests`, func() {
	const externalConfigFile = "../data_product_hub_api_service_v1.env"

	var (
		err                      error
		dataProductHubApiService *dataproducthubapiservicev1.DataProductHubApiServiceV1
		serviceURL               string
		config                   map[string]string

		// Variables to hold link values
		completeADraftByContractTermsIdLink               string
		completeADraftByDraftIdLink                       string
		completeContractTermsDocumentByDocumentIdLink     string
		completeDraftContractTermsByDataProductIdLink     string
		createAContractTermsDocByContractTermsIdLink      string
		createAContractTermsDocByDraftIdLink              string
		createDataProductByCatalogIdLink                  string
		createDraftByContainerIdLink                      string
		createNewDraftByDataProductIdLink                 string
		deleteAContractDocumentByDraftIdLink              string
		deleteADraftByContractTermsIdLink                 string
		deleteADraftByDraftIdLink                         string
		deleteContractDocumentByDataProductIdLink         string
		deleteContractTermsDocumentByDocumentIdLink       string
		deleteDraftOfDataProductByDataProductIdLink       string
		getADraftByContractTermsIdLink                    string
		getADraftContractDocumentByDraftIdLink            string
		getADraftOfDataProductByDataProductIdLink         string
		getAReleaseByReleaseIdLink                        string
		getAReleaseContractTermsByContractTermsIdLink     string
		getAReleaseContractTermsByReleaseIdLink           string
		getAReleaseOfDataProductByDataProductIdLink       string
		getContractDocumentByDataProductIdLink            string
		getContractTermsDocumentByIdDocumentIdLink        string
		getDataProductByDataProductIdLink                 string
		getDraftByDraftIdLink                             string
		getListOfDataProductDraftsByDataProductIdLink     string
		getListOfReleasesOfDataProductByDataProductIdLink string
		getReleaseContractDocumentByDataProductIdLink     string
		getReleaseContractDocumentByDocumentIdLink        string
		getStatusByCatalogIdLink                          string
		publishADraftByDraftIdLink                        string
		publishADraftOfDataProductByDataProductIdLink     string
		retireAReleaseContractTermsByReleaseIdLink        string
		retireAReleasesOfDataProductByDataProductIdLink   string
		updateADraftByContractTermsIdLink                 string
		updateADraftByDraftIdLink                         string
		updateAReleaseByReleaseIdLink                     string
		updateContractDocumentByDataProductIdLink         string
		updateContractDocumentByDraftIdLink               string
		updateContractTermsDocumentByDocumentIdLink       string
		updateDraftOfDataProductByDataProductIdLink       string
		updateReleaseOfDataProductByDataProductIdLink     string
		uploadContractTermsDocByDataProductIdLink         string
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
			dataProductHubApiServiceServiceOptions := &dataproducthubapiservicev1.DataProductHubApiServiceV1Options{}

			dataProductHubApiService, err = dataproducthubapiservicev1.NewDataProductHubApiServiceV1UsingExternalConfig(dataProductHubApiServiceServiceOptions)
			Expect(err).To(BeNil())
			Expect(dataProductHubApiService).ToNot(BeNil())
			Expect(dataProductHubApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			dataProductHubApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`Initialize - Initialize resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize(initializeOptions *InitializeOptions)`, func() {
			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID:   core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Type: core.StringPtr("catalog"),
			}

			initializeOptions := &dataproducthubapiservicev1.InitializeOptions{
				Container: containerReferenceModel,
				Include:   []string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project"},
			}

			initializeResource, response, err := dataProductHubApiService.Initialize(initializeOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(initializeResource).ToNot(BeNil())

			createDraftByContainerIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved createDraftByContainerIdLink value: %v\n", createDraftByContainerIdLink)
			createDataProductByCatalogIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved createDataProductByCatalogIdLink value: %v\n", createDataProductByCatalogIdLink)
			getStatusByCatalogIdLink = *initializeResource.Container.ID
			fmt.Fprintf(GinkgoWriter, "Saved getStatusByCatalogIdLink value: %v\n", getStatusByCatalogIdLink)
		})
	})

	Describe(`CreateDataProduct - Create a new data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProduct(createDataProductOptions *CreateDataProductOptions)`, func() {

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID:   core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Type: core.StringPtr("catalog"),
			}

			contractTermsDocumentModel := &dataproducthubapiservicev1.ContractTermsDocument{
				URL:        core.StringPtr("https://data.un.org/Host.aspx?Content=UNdataUse"),
				Type:       core.StringPtr("terms_and_conditions"),
				Name:       core.StringPtr("Contract and Terms"),
				ID:         core.StringPtr("d343dbcc-db69-4fdd-98f9-e2df58c8b690"),
				Attachment: nil,
				UploadURL:  nil,
			}

			dataProductContractTermsModel := &dataproducthubapiservicev1.DataProductContractTerms{
				ID:        core.StringPtr("17840b80-78f5-4928-9631-925668c3d530@78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Documents: []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel},
			}

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				ID:        core.StringPtr("e69d51ba-e08c-4665-bc89-c24064de08dd"),
				Container: containerIdentityModel,
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID:        core.StringPtr("352c8ccc-950d-4881-8139-9dc641fa87cd"),
				Name:      core.StringPtr("Human Resources"),
				Container: containerReferenceModel,
			}

			assetPartReferenceModel := &dataproducthubapiservicev1.AssetPartReference{
				ID:        core.StringPtr("e69d51ba-e08c-4665-bc89-c24064de08dd"),
				Container: containerReferenceModel,
				Type:      core.StringPtr("data_asset"),
			}

			deliveryMethodModel := &dataproducthubapiservicev1.DeliveryMethod{
				ID:        core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
				Container: containerReferenceModel,
			}

			dataProductPartModel := &dataproducthubapiservicev1.DataProductPart{
				Asset:           assetPartReferenceModel,
				DeliveryMethods: []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel},
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductVersionPrototype{
				Version:       core.StringPtr("1.0.0"),
				State:         core.StringPtr("draft"),
				DataProduct:   nil,
				Name:          core.StringPtr("GoSDK Integration Test Data Product"),
				Description:   core.StringPtr("GoSDK Integration Test Data Product"),
				Types:         []string{"data"},
				ContractTerms: []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel},
				IsRestricted:  core.BoolPtr(true),
				Asset:         assetPrototypeModel,
				Domain:        domainModel,
				PartsOut:      []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel},
			}

			createDataProductOptions := &dataproducthubapiservicev1.CreateDataProductOptions{
				Drafts: []dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			}

			dataProduct, response, err := dataProductHubApiService.CreateDataProduct(createDataProductOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProduct).ToNot(BeNil())

			createNewDraftByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved createNewDraftByDataProductIdLink value: %v\n", createNewDraftByDataProductIdLink)
			getContractDocumentByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractDocumentByDataProductIdLink value: %v\n", getContractDocumentByDataProductIdLink)
			retireAReleasesOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved retireAReleasesOfDataProductByDataProductIdLink value: %v\n", retireAReleasesOfDataProductByDataProductIdLink)
			getDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDataProductByDataProductIdLink value: %v\n", getDataProductByDataProductIdLink)
			updateDraftOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateDraftOfDataProductByDataProductIdLink value: %v\n", updateDraftOfDataProductByDataProductIdLink)
			updateContractDocumentByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractDocumentByDataProductIdLink value: %v\n", updateContractDocumentByDataProductIdLink)
			deleteDraftOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteDraftOfDataProductByDataProductIdLink value: %v\n", deleteDraftOfDataProductByDataProductIdLink)
			getAReleaseOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseOfDataProductByDataProductIdLink value: %v\n", getAReleaseOfDataProductByDataProductIdLink)
			completeDraftContractTermsByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeDraftContractTermsByDataProductIdLink value: %v\n", completeDraftContractTermsByDataProductIdLink)
			deleteContractDocumentByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractDocumentByDataProductIdLink value: %v\n", deleteContractDocumentByDataProductIdLink)
			getListOfDataProductDraftsByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getListOfDataProductDraftsByDataProductIdLink value: %v\n", getListOfDataProductDraftsByDataProductIdLink)
			getADraftOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftOfDataProductByDataProductIdLink value: %v\n", getADraftOfDataProductByDataProductIdLink)
			getReleaseContractDocumentByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getReleaseContractDocumentByDataProductIdLink value: %v\n", getReleaseContractDocumentByDataProductIdLink)
			publishADraftOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved publishADraftOfDataProductByDataProductIdLink value: %v\n", publishADraftOfDataProductByDataProductIdLink)
			getListOfReleasesOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved getListOfReleasesOfDataProductByDataProductIdLink value: %v\n", getListOfReleasesOfDataProductByDataProductIdLink)
			updateReleaseOfDataProductByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateReleaseOfDataProductByDataProductIdLink value: %v\n", updateReleaseOfDataProductByDataProductIdLink)
			uploadContractTermsDocByDataProductIdLink = *dataProduct.ID
			fmt.Fprintf(GinkgoWriter, "Saved uploadContractTermsDocByDataProductIdLink value: %v\n", uploadContractTermsDocByDataProductIdLink)
		})
	})

	Describe(`GetDataProduct - Retrieve a data product identified by id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProduct(getDataProductOptions *GetDataProductOptions)`, func() {
			getDataProductOptions := &dataproducthubapiservicev1.GetDataProductOptions{
				DataProductID: &getDataProductByDataProductIdLink,
			}

			dataProduct, response, err := dataProductHubApiService.GetDataProduct(getDataProductOptions)
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

			listDataProductsOptions.Start = nil
			listDataProductsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductSummary
			for {
				dataProductSummaryCollection, response, err := dataProductHubApiService.ListDataProducts(listDataProductsOptions)
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
			pager, err := dataProductHubApiService.NewDataProductsPager(listDataProductsOptions)
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
			pager, err = dataProductHubApiService.NewDataProductsPager(listDataProductsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProducts() returned a total of %d item(s) using DataProductsPager.\n", len(allResults))
		})
	})

	Describe(`CreateDataProductDraft - Create a new draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDataProductDraft(createDataProductDraftOptions *CreateDataProductDraftOptions)`, func() {
			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				ID:        core.StringPtr("e69d51ba-e08c-4665-bc89-c24064de08dd"),
				Container: containerIdentityModel,
			}

			dataProductIdentityModel := &dataproducthubapiservicev1.DataProductIdentity{
				ID: &getDataProductByDataProductIdLink,
			}

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID:   core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Type: core.StringPtr("catalog"),
			}

			contractTermsDocumentModel := &dataproducthubapiservicev1.ContractTermsDocument{
				URL:       core.StringPtr("https://data.un.org/Host.aspx?Content=UNdataUse"),
				Type:      core.StringPtr("terms_and_conditions"),
				Name:      core.StringPtr("Contract Terms Documents"),
				ID:        core.StringPtr("d343dbcc-db69-4fdd-98f9-e2df58c8b690"),
				UploadURL: nil,
			}

			dataProductContractTermsModel := &dataproducthubapiservicev1.DataProductContractTerms{
				ID:        core.StringPtr("5f608558-4976-4a5b-9308-f651bd24725a@78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Documents: []dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel},
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID:        core.StringPtr("352c8ccc-950d-4881-8139-9dc641fa87cd"),
				Name:      core.StringPtr("Human Resources"),
				Container: containerReferenceModel,
			}

			assetPartReferenceModel := &dataproducthubapiservicev1.AssetPartReference{
				ID:        core.StringPtr("e69d51ba-e08c-4665-bc89-c24064de08dd"),
				Container: containerReferenceModel,
				Type:      core.StringPtr("data_asset"),
			}

			deliveryMethodModel := &dataproducthubapiservicev1.DeliveryMethod{
				ID:        core.StringPtr("09cf5fcc-cb9d-4995-a8e4-16517b25229f"),
				Container: containerReferenceModel,
			}

			dataProductPartModel := &dataproducthubapiservicev1.DataProductPart{
				Asset:           assetPartReferenceModel,
				DeliveryMethods: []dataproducthubapiservicev1.DeliveryMethod{*deliveryMethodModel},
			}

			createDataProductDraftOptions := &dataproducthubapiservicev1.CreateDataProductDraftOptions{
				DataProductID: &getDataProductByDataProductIdLink,
				Asset:         assetPrototypeModel,
				Version:       core.StringPtr("1.2.0"),
				State:         core.StringPtr("draft"),
				DataProduct:   dataProductIdentityModel,
				Name:          core.StringPtr("GoSDK Integration Test "),
				Description:   core.StringPtr("Integration test Description"),
				Types:         []string{"data"},
				ContractTerms: []dataproducthubapiservicev1.DataProductContractTerms{*dataProductContractTermsModel},
				IsRestricted:  core.BoolPtr(true),
				Domain:        domainModel,
				PartsOut:      []dataproducthubapiservicev1.DataProductPart{*dataProductPartModel},
			}

			dataProductVersion, response, err := dataProductHubApiService.CreateDataProductDraft(createDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			getADraftContractDocumentByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftContractDocumentByDraftIdLink value: %v\n", getADraftContractDocumentByDraftIdLink)
			updateADraftByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByContractTermsIdLink value: %v\n", updateADraftByContractTermsIdLink)
			createAContractTermsDocByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByContractTermsIdLink value: %v\n", createAContractTermsDocByContractTermsIdLink)
			updateContractDocumentByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractDocumentByDraftIdLink value: %v\n", updateContractDocumentByDraftIdLink)
			getAReleaseContractTermsByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByContractTermsIdLink value: %v\n", getAReleaseContractTermsByContractTermsIdLink)
			completeADraftByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByContractTermsIdLink value: %v\n", completeADraftByContractTermsIdLink)
			getDraftByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDraftByDraftIdLink value: %v\n", getDraftByDraftIdLink)
			publishADraftByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved publishADraftByDraftIdLink value: %v\n", publishADraftByDraftIdLink)
			updateADraftByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByDraftIdLink value: %v\n", updateADraftByDraftIdLink)
			createAContractTermsDocByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByDraftIdLink value: %v\n", createAContractTermsDocByDraftIdLink)
			deleteAContractDocumentByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteAContractDocumentByDraftIdLink value: %v\n", deleteAContractDocumentByDraftIdLink)
			deleteADraftByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByContractTermsIdLink value: %v\n", deleteADraftByContractTermsIdLink)
			deleteADraftByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByDraftIdLink value: %v\n", deleteADraftByDraftIdLink)
			completeADraftByDraftIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByDraftIdLink value: %v\n", completeADraftByDraftIdLink)
			getADraftByContractTermsIdLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftByContractTermsIdLink value: %v\n", getADraftByContractTermsIdLink)
		})
	})

	Describe(`GetDataProductDraft - Get a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductDraft(getDataProductDraftOptions *GetDataProductDraftOptions)`, func() {
			getDataProductDraftOptions := &dataproducthubapiservicev1.GetDataProductDraftOptions{
				DataProductID: &getADraftOfDataProductByDataProductIdLink,
				DraftID:       &getDraftByDraftIdLink,
			}

			dataProductVersion, response, err := dataProductHubApiService.GetDataProductDraft(getDataProductDraftOptions)
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
				Value: "Updated Data Product Draft",
			}

			updateDataProductDraftOptions := &dataproducthubapiservicev1.UpdateDataProductDraftOptions{
				DataProductID:         &updateDraftOfDataProductByDataProductIdLink,
				DraftID:               &updateADraftByDraftIdLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dataProductHubApiService.UpdateDataProductDraft(updateDataProductDraftOptions)
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
				DataProductID:   &uploadContractTermsDocByDataProductIdLink,
				DraftID:         &createAContractTermsDocByDraftIdLink,
				ContractTermsID: &createAContractTermsDocByContractTermsIdLink,
				Type:            core.StringPtr("terms_and_conditions"),
				Name:            core.StringPtr("Terms and conditions document"),
				ID:              core.StringPtr("cf54ad0c-89d3-4926-8d2d-bb8bca9ee2bf"),
				URL:             core.StringPtr("https://data.un.org/Host.aspx?Content=UNdataUse"),
				UploadURL:       core.StringPtr("https://data.un.org/Host.aspx?Content=UNdataUse"),
			}

			contractTermsDocument, response, err := dataProductHubApiService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(contractTermsDocument).ToNot(BeNil())

			getReleaseContractDocumentByDocumentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved getReleaseContractDocumentByDocumentIdLink value: %v\n", getReleaseContractDocumentByDocumentIdLink)
			deleteContractTermsDocumentByDocumentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteContractTermsDocumentByDocumentIdLink value: %v\n", deleteContractTermsDocumentByDocumentIdLink)
			getContractTermsDocumentByIdDocumentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved getContractTermsDocumentByIdDocumentIdLink value: %v\n", getContractTermsDocumentByIdDocumentIdLink)
			updateContractTermsDocumentByDocumentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractTermsDocumentByDocumentIdLink value: %v\n", updateContractTermsDocumentByDocumentIdLink)
			completeContractTermsDocumentByDocumentIdLink = *contractTermsDocument.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeContractTermsDocumentByDocumentIdLink value: %v\n", completeContractTermsDocumentByDocumentIdLink)
		})
	})

	Describe(`GetDraftContractTermsDocument - Get a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions *GetDraftContractTermsDocumentOptions)`, func() {
			getDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.GetDraftContractTermsDocumentOptions{
				DataProductID:   &getContractDocumentByDataProductIdLink,
				DraftID:         &getADraftContractDocumentByDraftIdLink,
				ContractTermsID: &getADraftByContractTermsIdLink,
				DocumentID:      &getContractTermsDocumentByIdDocumentIdLink,
			}

			contractTermsDocument, response, err := dataProductHubApiService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
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
				Op:    core.StringPtr("replace"),
				Path:  core.StringPtr("/type"),
				Value: "terms_and_conditions",
			}

			updateDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.UpdateDraftContractTermsDocumentOptions{
				DataProductID:         &updateContractDocumentByDataProductIdLink,
				DraftID:               &updateContractDocumentByDraftIdLink,
				ContractTermsID:       &updateADraftByContractTermsIdLink,
				DocumentID:            &updateContractTermsDocumentByDocumentIdLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			contractTermsDocument, response, err := dataProductHubApiService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`CompleteDraftContractTermsDocument - Complete a contract document upload operation`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions *CompleteDraftContractTermsDocumentOptions)`, func() {
			completeDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.CompleteDraftContractTermsDocumentOptions{
				DataProductID:   &completeDraftContractTermsByDataProductIdLink,
				DraftID:         &completeADraftByDraftIdLink,
				ContractTermsID: &completeADraftByContractTermsIdLink,
				DocumentID:      &completeContractTermsDocumentByDocumentIdLink,
			}

			contractTermsDocument, response, err := dataProductHubApiService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`DeleteDataProductDraft - Delete a data product draft identified by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDataProductDraft(deleteDataProductDraftOptions *DeleteDataProductDraftOptions)`, func() {
			deleteDataProductDraftOptions := &dataproducthubapiservicev1.DeleteDataProductDraftOptions{
				DataProductID: &deleteDraftOfDataProductByDataProductIdLink,
				DraftID:       &deleteADraftByDraftIdLink,
			}

			response, err := dataProductHubApiService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`PublishDataProductDraft - Publish a draft of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PublishDataProductDraft(publishDataProductDraftOptions *PublishDataProductDraftOptions)`, func() {
			publishDataProductDraftOptions := &dataproducthubapiservicev1.PublishDataProductDraftOptions{
				DataProductID: &publishADraftOfDataProductByDataProductIdLink,
				DraftID:       &publishADraftByDraftIdLink,
			}

			dataProductVersion, response, err := dataProductHubApiService.PublishDataProductDraft(publishDataProductDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())

			updateAReleaseByReleaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateAReleaseByReleaseIdLink value: %v\n", updateAReleaseByReleaseIdLink)
			getAReleaseContractTermsByReleaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByReleaseIdLink value: %v\n", getAReleaseContractTermsByReleaseIdLink)
			retireAReleaseContractTermsByReleaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved retireAReleaseContractTermsByReleaseIdLink value: %v\n", retireAReleaseContractTermsByReleaseIdLink)
			getAReleaseByReleaseIdLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseByReleaseIdLink value: %v\n", getAReleaseByReleaseIdLink)
		})
	})

	Describe(`ListDataProductDrafts - Retrieve a list of data product drafts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductDrafts(listDataProductDraftsOptions *ListDataProductDraftsOptions) with pagination`, func() {
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID:    &getListOfDataProductDraftsByDataProductIdLink,
				AssetContainerID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Version:          core.StringPtr("1.3.0"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			listDataProductDraftsOptions.Start = nil
			listDataProductDraftsOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for {
				dataProductDraftCollection, response, err := dataProductHubApiService.ListDataProductDrafts(listDataProductDraftsOptions)
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
				DataProductID:    &getListOfDataProductDraftsByDataProductIdLink,
				AssetContainerID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				Version:          core.StringPtr("1.0.0"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubApiService.NewDataProductDraftsPager(listDataProductDraftsOptions)
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
			pager, err = dataProductHubApiService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListDataProductDrafts() returned a total of %d item(s) using DataProductDraftsPager.\n", len(allResults))
		})
	})

	Describe(`GetInitializeStatus - Get resource initialization status`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInitializeStatus(getInitializeStatusOptions *GetInitializeStatusOptions)`, func() {
			getInitializeStatusOptions := &dataproducthubapiservicev1.GetInitializeStatusOptions{
				ContainerID: &getStatusByCatalogIdLink,
			}

			initializeResource, response, err := dataProductHubApiService.GetInitializeStatus(getInitializeStatusOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
	})

	Describe(`GetServiceIdCredentials - Get service id credentials`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceIdCredentials(getServiceIdCredentialsOptions *GetServiceIdCredentialsOptions)`, func() {
			getServiceIdCredentialsOptions := &dataproducthubapiservicev1.GetServiceIdCredentialsOptions{}

			serviceIdCredentials, response, err := dataProductHubApiService.GetServiceIdCredentials(getServiceIdCredentialsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdCredentials).ToNot(BeNil())
		})
	})

	Describe(`ManageApiKeys - Rotate credentials for a Data Product Hub instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ManageApiKeys(manageApiKeysOptions *ManageApiKeysOptions)`, func() {
			manageApiKeysOptions := &dataproducthubapiservicev1.ManageApiKeysOptions{}

			response, err := dataProductHubApiService.ManageApiKeys(manageApiKeysOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetDataProductRelease - Get a release of an existing data product`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDataProductRelease(getDataProductReleaseOptions *GetDataProductReleaseOptions)`, func() {
			getDataProductReleaseOptions := &dataproducthubapiservicev1.GetDataProductReleaseOptions{
				DataProductID: &getAReleaseOfDataProductByDataProductIdLink,
				ReleaseID:     &getAReleaseByReleaseIdLink,
			}

			dataProductVersion, response, err := dataProductHubApiService.GetDataProductRelease(getDataProductReleaseOptions)
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
				Value: "Updated Data Product Release",
			}

			updateDataProductReleaseOptions := &dataproducthubapiservicev1.UpdateDataProductReleaseOptions{
				DataProductID:         &updateReleaseOfDataProductByDataProductIdLink,
				ReleaseID:             &getAReleaseByReleaseIdLink,
				JSONPatchInstructions: []dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			}

			dataProductVersion, response, err := dataProductHubApiService.UpdateDataProductRelease(updateDataProductReleaseOptions)
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
				DataProductID:   &getReleaseContractDocumentByDataProductIdLink,
				ReleaseID:       &getAReleaseContractTermsByReleaseIdLink,
				ContractTermsID: &getAReleaseContractTermsByContractTermsIdLink,
				DocumentID:      &getReleaseContractDocumentByDocumentIdLink,
			}

			contractTermsDocument, response, err := dataProductHubApiService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
	})

	Describe(`DeleteDraftContractTermsDocument - Delete a contract document`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions *DeleteDraftContractTermsDocumentOptions)`, func() {
			deleteDraftContractTermsDocumentOptions := &dataproducthubapiservicev1.DeleteDraftContractTermsDocumentOptions{
				DataProductID:   &deleteContractDocumentByDataProductIdLink,
				DraftID:         &deleteAContractDocumentByDraftIdLink,
				ContractTermsID: &deleteADraftByContractTermsIdLink,
				DocumentID:      &deleteContractTermsDocumentByDocumentIdLink,
			}

			response, err := dataProductHubApiService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListDataProductReleases - Retrieve a list of data product releases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDataProductReleases(listDataProductReleasesOptions *ListDataProductReleasesOptions) with pagination`, func() {
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIdLink,
				AssetContainerID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				State:            []string{"available"},
				Version:          core.StringPtr("1.3.0"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			listDataProductReleasesOptions.Start = nil
			listDataProductReleasesOptions.Limit = core.Int64Ptr(1)

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for {
				dataProductReleaseCollection, response, err := dataProductHubApiService.ListDataProductReleases(listDataProductReleasesOptions)
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
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIdLink,
				AssetContainerID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
				State:            []string{"available"},
				Version:          core.StringPtr("1.3.0"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := dataProductHubApiService.NewDataProductReleasesPager(listDataProductReleasesOptions)
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
			pager, err = dataProductHubApiService.NewDataProductReleasesPager(listDataProductReleasesOptions)
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
				DataProductID: &retireAReleasesOfDataProductByDataProductIdLink,
				ReleaseID:     &retireAReleaseContractTermsByReleaseIdLink,
			}

			dataProductVersion, response, err := dataProductHubApiService.RetireDataProductRelease(retireDataProductReleaseOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
	})

})

//
// Utility functions are declared in the unit test file
//
