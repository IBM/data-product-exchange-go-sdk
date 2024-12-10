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
var _ = Describe(`DataProductHubAPIServiceV1 Examples Tests`, func() {

	const externalConfigFile = "../data_product_hub_api_service_v1.env"

	var (
		dataProductHubAPIService *dataproducthubapiservicev1.DataProductHubAPIServiceV1
		config                   map[string]string

		// Variables to hold link values
		completeADraftByContractTermsIDLink               string
		completeADraftByDraftIDLink                       string
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
		getADraftByContractTermsIDLink                    string
		getADraftContractDocumentByDraftIDLink            string
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
		publishADraftByDraftIDLink                        string
		publishADraftOfDataProductByDataProductIDLink     string
		retireAReleaseContractTermsByReleaseIDLink        string
		retireAReleasesOfDataProductByDataProductIDLink   string
		updateADraftByContractTermsIDLink                 string
		updateADraftByDraftIDLink                         string
		updateAReleaseByReleaseIDLink                     string
		updateContractDocumentByDataProductIDLink         string
		updateContractDocumentByDraftIDLink               string
		updateContractTermsDocumentByDocumentIDLink       string
		updateDraftOfDataProductByDataProductIDLink       string
		updateReleaseOfDataProductByDataProductIDLink     string
		uploadContractTermsDocByDataProductIDLink         string
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

			dataProductHubAPIServiceServiceOptions := &dataproducthubapiservicev1.DataProductHubAPIServiceV1Options{}

			dataProductHubAPIService, err = dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(dataProductHubAPIServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(dataProductHubAPIService).ToNot(BeNil())
		})
	})

	Describe(`DataProductHubAPIServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize request example`, func() {
			fmt.Println("\nInitialize() result:")
			// begin-initialize

			initializeOptions := dataProductHubAPIService.NewInitializeOptions()
			initializeOptions.SetInclude([]string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"})

			initializeResource, response, err := dataProductHubAPIService.Initialize(initializeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeResource, "", "  ")
			fmt.Println(string(b))

			// end-initialize

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
		It(`CreateDataProduct request example`, func() {
			fmt.Println("\nCreateDataProduct() result:")
			// begin-create_data_product

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			dataProductVersionPrototypeModel := &dataproducthubapiservicev1.DataProductVersionPrototype{
				Name:  core.StringPtr("My New Data Product"),
				Asset: assetPrototypeModel,
			}

			createDataProductOptions := dataProductHubAPIService.NewCreateDataProductOptions(
				[]dataproducthubapiservicev1.DataProductVersionPrototype{*dataProductVersionPrototypeModel},
			)

			dataProduct, response, err := dataProductHubAPIService.CreateDataProduct(createDataProductOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProduct, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product

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
		})
		It(`CreateDataProductDraft request example`, func() {
			fmt.Println("\nCreateDataProductDraft() result:")
			// begin-create_data_product_draft

			containerIdentityModel := &dataproducthubapiservicev1.ContainerIdentity{
				ID: core.StringPtr("d29c42eb-7100-4b7a-8257-c196dbcca1cd"),
			}

			assetPrototypeModel := &dataproducthubapiservicev1.AssetPrototype{
				Container: containerIdentityModel,
			}

			dataProductDraftVersionReleaseModel := &dataproducthubapiservicev1.DataProductDraftVersionRelease{
				ID: core.StringPtr("8bf83660-11fe-4427-a72a-8d8359af24e3"),
			}

			dataProductIdentityModel := &dataproducthubapiservicev1.DataProductIdentity{
				ID:      core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Release: dataProductDraftVersionReleaseModel,
			}

			createDataProductDraftOptions := dataProductHubAPIService.NewCreateDataProductDraftOptions(
				createNewDraftByDataProductIDLink,
				assetPrototypeModel,
			)
			createDataProductDraftOptions.SetVersion("1.2.0")
			createDataProductDraftOptions.SetDataProduct(dataProductIdentityModel)

			dataProductVersion, response, err := dataProductHubAPIService.CreateDataProductDraft(createDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductVersion).ToNot(BeNil())

			getADraftContractDocumentByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftContractDocumentByDraftIDLink value: %v\n", getADraftContractDocumentByDraftIDLink)
			updateADraftByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByContractTermsIDLink value: %v\n", updateADraftByContractTermsIDLink)
			createAContractTermsDocByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByContractTermsIDLink value: %v\n", createAContractTermsDocByContractTermsIDLink)
			updateContractDocumentByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractDocumentByDraftIDLink value: %v\n", updateContractDocumentByDraftIDLink)
			getAReleaseContractTermsByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByContractTermsIDLink value: %v\n", getAReleaseContractTermsByContractTermsIDLink)
			completeADraftByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByContractTermsIDLink value: %v\n", completeADraftByContractTermsIDLink)
			getDraftByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDraftByDraftIDLink value: %v\n", getDraftByDraftIDLink)
			publishADraftByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved publishADraftByDraftIDLink value: %v\n", publishADraftByDraftIDLink)
			updateADraftByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByDraftIDLink value: %v\n", updateADraftByDraftIDLink)
			createAContractTermsDocByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByDraftIDLink value: %v\n", createAContractTermsDocByDraftIDLink)
			deleteAContractDocumentByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteAContractDocumentByDraftIDLink value: %v\n", deleteAContractDocumentByDraftIDLink)
			deleteADraftByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByContractTermsIDLink value: %v\n", deleteADraftByContractTermsIDLink)
			deleteADraftByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByDraftIDLink value: %v\n", deleteADraftByDraftIDLink)
			completeADraftByDraftIDLink = *dataProductVersion.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByDraftIDLink value: %v\n", completeADraftByDraftIDLink)
			getADraftByContractTermsIDLink = *dataProductVersion.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftByContractTermsIDLink value: %v\n", getADraftByContractTermsIDLink)
		})
		It(`CreateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCreateDraftContractTermsDocument() result:")
			// begin-create_draft_contract_terms_document

			createDraftContractTermsDocumentOptions := dataProductHubAPIService.NewCreateDraftContractTermsDocumentOptions(
				uploadContractTermsDocByDataProductIDLink,
				createAContractTermsDocByDraftIDLink,
				createAContractTermsDocByContractTermsIDLink,
				"terms_and_conditions",
				"Terms and conditions document",
			)

			contractTermsDocument, response, err := dataProductHubAPIService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTermsDocument, "", "  ")
			fmt.Println(string(b))

			// end-create_draft_contract_terms_document

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
		It(`PublishDataProductDraft request example`, func() {
			fmt.Println("\nPublishDataProductDraft() result:")
			// begin-publish_data_product_draft

			publishDataProductDraftOptions := dataProductHubAPIService.NewPublishDataProductDraftOptions(
				publishADraftOfDataProductByDataProductIDLink,
				publishADraftByDraftIDLink,
			)

			dataProductVersion, response, err := dataProductHubAPIService.PublishDataProductDraft(publishDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersion, "", "  ")
			fmt.Println(string(b))

			// end-publish_data_product_draft

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
		It(`GetInitializeStatus request example`, func() {
			fmt.Println("\nGetInitializeStatus() result:")
			// begin-get_initialize_status

			getInitializeStatusOptions := dataProductHubAPIService.NewGetInitializeStatusOptions()

			initializeResource, response, err := dataProductHubAPIService.GetInitializeStatus(getInitializeStatusOptions)
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
		It(`GetServiceIDCredentials request example`, func() {
			fmt.Println("\nGetServiceIDCredentials() result:")
			// begin-get_service_id_credentials

			getServiceIDCredentialsOptions := dataProductHubAPIService.NewGetServiceIDCredentialsOptions()

			serviceIDCredentials, response, err := dataProductHubAPIService.GetServiceIDCredentials(getServiceIDCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIDCredentials, "", "  ")
			fmt.Println(string(b))

			// end-get_service_id_credentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIDCredentials).ToNot(BeNil())
		})
		It(`ManageAPIKeys request example`, func() {
			// begin-manage_api_keys

			manageAPIKeysOptions := dataProductHubAPIService.NewManageAPIKeysOptions()

			response, err := dataProductHubAPIService.ManageAPIKeys(manageAPIKeysOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from ManageAPIKeys(): %d\n", response.StatusCode)
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

			pager, err := dataProductHubAPIService.NewDataProductsPager(listDataProductsOptions)
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

			getDataProductOptions := dataProductHubAPIService.NewGetDataProductOptions(
				getDataProductByDataProductIDLink,
			)

			dataProduct, response, err := dataProductHubAPIService.GetDataProduct(getDataProductOptions)
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

			completeDraftContractTermsDocumentOptions := dataProductHubAPIService.NewCompleteDraftContractTermsDocumentOptions(
				completeDraftContractTermsByDataProductIDLink,
				completeADraftByDraftIDLink,
				completeADraftByContractTermsIDLink,
				completeContractTermsDocumentByDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
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
				DataProductID:    &getListOfDataProductDraftsByDataProductIDLink,
				AssetContainerID: core.StringPtr("testString"),
				Version:          core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubAPIService.NewDataProductDraftsPager(listDataProductDraftsOptions)
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

			getDataProductDraftOptions := dataProductHubAPIService.NewGetDataProductDraftOptions(
				getADraftOfDataProductByDataProductIDLink,
				getDraftByDraftIDLink,
			)

			dataProductVersion, response, err := dataProductHubAPIService.GetDataProductDraft(getDataProductDraftOptions)
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

			updateDataProductDraftOptions := dataProductHubAPIService.NewUpdateDataProductDraftOptions(
				updateDraftOfDataProductByDataProductIDLink,
				updateADraftByDraftIDLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dataProductHubAPIService.UpdateDataProductDraft(updateDataProductDraftOptions)
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

			getDraftContractTermsDocumentOptions := dataProductHubAPIService.NewGetDraftContractTermsDocumentOptions(
				getContractDocumentByDataProductIDLink,
				getADraftContractDocumentByDraftIDLink,
				getADraftByContractTermsIDLink,
				getContractTermsDocumentByIDDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
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

			updateDraftContractTermsDocumentOptions := dataProductHubAPIService.NewUpdateDraftContractTermsDocumentOptions(
				updateContractDocumentByDataProductIDLink,
				updateContractDocumentByDraftIDLink,
				updateADraftByContractTermsIDLink,
				updateContractTermsDocumentByDocumentIDLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTermsDocument, response, err := dataProductHubAPIService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
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

			getDataProductReleaseOptions := dataProductHubAPIService.NewGetDataProductReleaseOptions(
				getAReleaseOfDataProductByDataProductIDLink,
				getAReleaseByReleaseIDLink,
			)

			dataProductVersion, response, err := dataProductHubAPIService.GetDataProductRelease(getDataProductReleaseOptions)
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

			updateDataProductReleaseOptions := dataProductHubAPIService.NewUpdateDataProductReleaseOptions(
				updateReleaseOfDataProductByDataProductIDLink,
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductVersion, response, err := dataProductHubAPIService.UpdateDataProductRelease(updateDataProductReleaseOptions)
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

			getReleaseContractTermsDocumentOptions := dataProductHubAPIService.NewGetReleaseContractTermsDocumentOptions(
				getReleaseContractDocumentByDataProductIDLink,
				getAReleaseContractTermsByReleaseIDLink,
				getAReleaseContractTermsByContractTermsIDLink,
				getReleaseContractDocumentByDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
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
				DataProductID:    &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: core.StringPtr("testString"),
				State:            []string{"available"},
				Version:          core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubAPIService.NewDataProductReleasesPager(listDataProductReleasesOptions)
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

			retireDataProductReleaseOptions := dataProductHubAPIService.NewRetireDataProductReleaseOptions(
				retireAReleasesOfDataProductByDataProductIDLink,
				retireAReleaseContractTermsByReleaseIDLink,
			)

			dataProductVersion, response, err := dataProductHubAPIService.RetireDataProductRelease(retireDataProductReleaseOptions)
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

			deleteDraftContractTermsDocumentOptions := dataProductHubAPIService.NewDeleteDraftContractTermsDocumentOptions(
				deleteContractDocumentByDataProductIDLink,
				deleteAContractDocumentByDraftIDLink,
				deleteADraftByContractTermsIDLink,
				deleteContractTermsDocumentByDocumentIDLink,
			)

			response, err := dataProductHubAPIService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
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

			deleteDataProductDraftOptions := dataProductHubAPIService.NewDeleteDataProductDraftOptions(
				deleteDraftOfDataProductByDataProductIDLink,
				deleteADraftByDraftIDLink,
			)

			response, err := dataProductHubAPIService.DeleteDataProductDraft(deleteDataProductDraftOptions)
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
