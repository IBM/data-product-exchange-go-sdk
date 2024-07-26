//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/data-product-exchange-go-sdk/dataproducthubapiservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Data Product Hub API Service service.
//
// The following configuration properties are assumed to be defined:
// DATA_PRODUCT_HUB_API_SERVICE_URL=<service base url>
// DATA_PRODUCT_HUB_API_SERVICE_AUTH_TYPE=iam
// DATA_PRODUCT_HUB_API_SERVICE_APIKEY=<IAM apikey>
// DATA_PRODUCT_HUB_API_SERVICE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`DataProductHubApiServiceV1 Examples Tests`, func() {

	const externalConfigFile = "data_product_hub_api_service_v1.env"

	var (
		dataProductHubApiServiceService *dataproducthubapiservicev1.DataProductHubApiServiceV1
		config                          map[string]string

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
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(dataproducthubapiservicev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			dataProductHubApiServiceServiceOptions := &dataproducthubapiservicev1.DataProductHubApiServiceV1Options{}

			dataProductHubApiServiceService, err = dataproducthubapiservicev1.NewDataProductHubApiServiceV1UsingExternalConfig(dataProductHubApiServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(dataProductHubApiServiceService).ToNot(BeNil())
		})
	})

	Describe(`DataProductHubApiServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize request example`, func() {
			fmt.Println("\nInitialize() result:")
			// begin-initialize

			initializeOptions := dataProductHubApiServiceService.NewInitializeOptions()
			initializeOptions.SetInclude([]string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project"})

			initializeResource, response, err := dataProductHubApiServiceService.Initialize(initializeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-initialize

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
		It(`CreateDataProduct request example`, func() {
			fmt.Println("\nCreateDataProduct() result:")
			// begin-create_data_product

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductVersionPrototype{
				Name:  core.StringPtr("My New Data Product"),
				Asset: assetPrototypeModel,
			}

			createDataProductOptions := dataProductHubApiServiceService.NewCreateDataProductOptions(
				[]dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			)

			dataProduct, response, err := dataProductHubApiServiceService.CreateDataProduct(createDataProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product

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
		It(`CreateDataProductDraft request example`, func() {
			fmt.Println("\nCreateDataProductDraft() result:")
			// begin-create_data_product_draft

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("78c6e2f0-d6ed-46c0-8dee-2c2bd578d58b"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			dataProductIdentityModel := &dataproducthubapiservicev1.DataProductIdentity{
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
			}

			createDataProductDraftOptions := dataProductHubApiServiceService.NewCreateDataProductDraftOptions(
				createNewDraftByDataProductIdLink,
				assetPrototypeModel,
			)
			createDataProductDraftOptions.SetVersion("1.2.0")
			createDataProductDraftOptions.SetDataProduct(dataProductIdentityModel)

			dataProductVersion, response, err := dataProductHubApiServiceService.CreateDataProductDraft(createDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_draft

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
		It(`CreateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCreateDraftContractTermsDocument() result:")
			// begin-create_draft_contract_terms_document

			createDraftContractTermsDocumentOptions := dataProductHubApiServiceService.NewCreateDraftContractTermsDocumentOptions(
				uploadContractTermsDocByDataProductIdLink,
				createAContractTermsDocByDraftIdLink,
				createAContractTermsDocByContractTermsIdLink,
				"terms_and_conditions",
				"Terms and conditions document",
				"b38df608-d34b-4d58-8136-ed25e6c6684e",
			)

			contractTermsDocument, response, err := dataProductHubApiServiceService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-create_draft_contract_terms_document

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
		It(`PublishDataProductDraft request example`, func() {
			fmt.Println("\nPublishDataProductDraft() result:")
			// begin-publish_data_product_draft

			publishDataProductDraftOptions := dataProductHubApiServiceService.NewPublishDataProductDraftOptions(
				publishADraftOfDataProductByDataProductIdLink,
				publishADraftByDraftIdLink,
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.PublishDataProductDraft(publishDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-publish_data_product_draft

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
		It(`GetInitializeStatus request example`, func() {
			fmt.Println("\nGetInitializeStatus() result:")
			// begin-get_initialize_status

			getInitializeStatusOptions := dataProductHubApiServiceService.NewGetInitializeStatusOptions()

			initializeResource, response, err := dataProductHubApiServiceService.GetInitializeStatus(getInitializeStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-get_initialize_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initializeResource).ToNot(BeNil())
		})
		It(`GetServiceIdCredentials request example`, func() {
			fmt.Println("\nGetServiceIdCredentials() result:")
			// begin-get_service_id_credentials

			getServiceIdCredentialsOptions := dataProductHubApiServiceService.NewGetServiceIdCredentialsOptions()

			serviceIdCredentials, response, err := dataProductHubApiServiceService.GetServiceIdCredentials(getServiceIdCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIdCredentials, "", "  ")
			fmt.Println(string(b))

			// end-get_service_id_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdCredentials).ToNot(BeNil())
		})
		It(`ManageApiKeys request example`, func() {
			// begin-manage_api_keys

			manageApiKeysOptions := dataProductHubApiServiceService.NewManageApiKeysOptions()

			response, err := dataProductHubApiServiceService.ManageApiKeys(manageApiKeysOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from ManageApiKeys(): %d\n", response.StatusCode)
			}

			// end-manage_api_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ListDataProducts request example`, func() {
			fmt.Println("\nListDataProducts() result:")
			// begin-list_data_products
			listDataProductsOptions := &dataproducthubapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubApiServiceService.NewDataProductsPager(listDataProductsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproducthubapiservicev1.DataProductSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_products
		})
		It(`GetDataProduct request example`, func() {
			fmt.Println("\nGetDataProduct() result:")
			// begin-get_data_product

			getDataProductOptions := dataProductHubApiServiceService.NewGetDataProductOptions(
				getDataProductByDataProductIdLink,
			)

			dataProduct, response, err := dataProductHubApiServiceService.GetDataProduct(getDataProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProduct, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProduct).ToNot(BeNil())
		})
		It(`CompleteDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCompleteDraftContractTermsDocument() result:")
			// begin-complete_draft_contract_terms_document

			completeDraftContractTermsDocumentOptions := dataProductHubApiServiceService.NewCompleteDraftContractTermsDocumentOptions(
				completeDraftContractTermsByDataProductIdLink,
				completeADraftByDraftIdLink,
				completeADraftByContractTermsIdLink,
				completeContractTermsDocumentByDocumentIdLink,
			)

			contractTermsDocument, response, err := dataProductHubApiServiceService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-complete_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`ListDataProductDrafts request example`, func() {
			fmt.Println("\nListDataProductDrafts() result:")
			// begin-list_data_product_drafts
			listDataProductDraftsOptions := &dataproducthubapiservicev1.ListDataProductDraftsOptions{
				DataProductID:    &getListOfDataProductDraftsByDataProductIdLink,
				AssetContainerID: core.StringPtr("testString"),
				Version:          core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubApiServiceService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_product_drafts
		})
		It(`GetDataProductDraft request example`, func() {
			fmt.Println("\nGetDataProductDraft() result:")
			// begin-get_data_product_draft

			getDataProductDraftOptions := dataProductHubApiServiceService.NewGetDataProductDraftOptions(
				getADraftOfDataProductByDataProductIdLink,
				getDraftByDraftIdLink,
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.GetDataProductDraft(getDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`UpdateDataProductDraft request example`, func() {
			fmt.Println("\nUpdateDataProductDraft() result:")
			// begin-update_data_product_draft

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductDraftOptions := dataProductHubApiServiceService.NewUpdateDataProductDraftOptions(
				updateDraftOfDataProductByDataProductIdLink,
				updateADraftByDraftIdLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.UpdateDataProductDraft(updateDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`GetDraftContractTermsDocument request example`, func() {
			fmt.Println("\nGetDraftContractTermsDocument() result:")
			// begin-get_draft_contract_terms_document

			getDraftContractTermsDocumentOptions := dataProductHubApiServiceService.NewGetDraftContractTermsDocumentOptions(
				getContractDocumentByDataProductIdLink,
				getADraftContractDocumentByDraftIdLink,
				getADraftByContractTermsIdLink,
				getContractTermsDocumentByIdDocumentIdLink,
			)

			contractTermsDocument, response, err := dataProductHubApiServiceService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-get_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`UpdateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nUpdateDraftContractTermsDocument() result:")
			// begin-update_draft_contract_terms_document

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDraftContractTermsDocumentOptions := dataProductHubApiServiceService.NewUpdateDraftContractTermsDocumentOptions(
				updateContractDocumentByDataProductIdLink,
				updateContractDocumentByDraftIdLink,
				updateADraftByContractTermsIdLink,
				updateContractTermsDocumentByDocumentIdLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTermsDocument, response, err := dataProductHubApiServiceService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-update_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`GetDataProductRelease request example`, func() {
			fmt.Println("\nGetDataProductRelease() result:")
			// begin-get_data_product_release

			getDataProductReleaseOptions := dataProductHubApiServiceService.NewGetDataProductReleaseOptions(
				getAReleaseOfDataProductByDataProductIdLink,
				getAReleaseByReleaseIdLink,
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.GetDataProductRelease(getDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`UpdateDataProductRelease request example`, func() {
			fmt.Println("\nUpdateDataProductRelease() result:")
			// begin-update_data_product_release

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op:   core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductReleaseOptions := dataProductHubApiServiceService.NewUpdateDataProductReleaseOptions(
				updateReleaseOfDataProductByDataProductIdLink,
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`GetReleaseContractTermsDocument request example`, func() {
			fmt.Println("\nGetReleaseContractTermsDocument() result:")
			// begin-get_release_contract_terms_document

			getReleaseContractTermsDocumentOptions := dataProductHubApiServiceService.NewGetReleaseContractTermsDocumentOptions(
				getReleaseContractDocumentByDataProductIdLink,
				getAReleaseContractTermsByReleaseIdLink,
				getAReleaseContractTermsByContractTermsIdLink,
				getReleaseContractDocumentByDocumentIdLink,
			)

			contractTermsDocument, response, err := dataProductHubApiServiceService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-get_release_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTermsDocument).ToNot(BeNil())
		})
		It(`ListDataProductReleases request example`, func() {
			fmt.Println("\nListDataProductReleases() result:")
			// begin-list_data_product_releases
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIdLink,
				AssetContainerID: core.StringPtr("testString"),
				State:            []string{"available"},
				Version:          core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubApiServiceService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproducthubapiservicev1.DataProductVersionSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_data_product_releases
		})
		It(`RetireDataProductRelease request example`, func() {
			fmt.Println("\nRetireDataProductRelease() result:")
			// begin-retire_data_product_release

			retireDataProductReleaseOptions := dataProductHubApiServiceService.NewRetireDataProductReleaseOptions(
				retireAReleasesOfDataProductByDataProductIdLink,
				retireAReleaseContractTermsByReleaseIdLink,
			)

			dataProductVersion, response, err := dataProductHubApiServiceService.RetireDataProductRelease(retireDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-retire_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersion).ToNot(BeNil())
		})
		It(`DeleteDraftContractTermsDocument request example`, func() {
			// begin-delete_draft_contract_terms_document

			deleteDraftContractTermsDocumentOptions := dataProductHubApiServiceService.NewDeleteDraftContractTermsDocumentOptions(
				deleteContractDocumentByDataProductIdLink,
				deleteAContractDocumentByDraftIdLink,
				deleteADraftByContractTermsIdLink,
				deleteContractTermsDocumentByDocumentIdLink,
			)

			response, err := dataProductHubApiServiceService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDraftContractTermsDocument(): %d\n", response.StatusCode)
			}

			// end-delete_draft_contract_terms_document

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteDataProductDraft request example`, func() {
			// begin-delete_data_product_draft

			deleteDataProductDraftOptions := dataProductHubApiServiceService.NewDeleteDataProductDraftOptions(
				deleteDraftOfDataProductByDataProductIdLink,
				deleteADraftByDraftIdLink,
			)

			response, err := dataProductHubApiServiceService.DeleteDataProductDraft(deleteDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDataProductDraft(): %d\n", response.StatusCode)
			}

			// end-delete_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
