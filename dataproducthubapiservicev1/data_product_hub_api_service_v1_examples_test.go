//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/data-product-exchange-go-sdk/v4/dataproducthubapiservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
)

//
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
//
var _ = Describe(`DataProductHubAPIServiceV1 Examples Tests`, func() {

	const externalConfigFile = "../dph_v1.env"

	var (
		dataProductHubAPIServiceService *dataproducthubapiservicev1.DataProductHubAPIServiceV1
		config       map[string]string

		// Variables to hold link values
		completeADraftByContractTermsIDLink string
		completeADraftByDraftIDLink string
		completeContractTermsDocumentByDocumentIDLink string
		completeDraftContractTermsByDataProductIDLink string
		createAContractTermsDocByContractTermsIDLink string
		createAContractTermsDocByDraftIDLink string
		createDataProductByCatalogIDLink string
		createDraftByContainerIDLink string
		createNewDraftByDataProductIDLink string
		deleteAContractDocumentByDraftIDLink string
		deleteADraftByContractTermsIDLink string
		deleteADraftByDraftIDLink string
		deleteContractDocumentByDataProductIDLink string
		deleteContractTermsDocumentByDocumentIDLink string
		deleteDraftOfDataProductByDataProductIDLink string
		getADraftByContractTermsIDLink string
		getADraftContractDocumentByDraftIDLink string
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
		publishADraftByDraftIDLink string
		publishADraftOfDataProductByDataProductIDLink string
		retireAReleaseContractTermsByReleaseIDLink string
		retireAReleasesOfDataProductByDataProductIDLink string
		updateADraftByContractTermsIDLink string
		updateADraftByDraftIDLink string
		updateAReleaseByReleaseIDLink string
		updateContractDocumentByDataProductIDLink string
		updateContractDocumentByDraftIDLink string
		updateContractTermsDocumentByDocumentIDLink string
		updateDraftOfDataProductByDataProductIDLink string
		updateReleaseOfDataProductByDataProductIDLink string
		uploadContractTermsDocByDataProductIDLink string
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

			dataProductHubAPIServiceService, err = dataproducthubapiservicev1.NewDataProductHubAPIServiceV1UsingExternalConfig(dataProductHubAPIServiceServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(dataProductHubAPIServiceService).ToNot(BeNil())
		})
	})

	Describe(`DataProductHubAPIServiceV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Initialize request example`, func() {
			fmt.Println("\nInitialize() result:")
			// begin-initialize

			initializeOptions := dataProductHubAPIServiceService.NewInitializeOptions()
			initializeOptions.SetInclude([]string{"delivery_methods", "domains_multi_industry", "data_product_samples", "workflows", "project", "catalog_configurations"})

			initializeResource, response, err := dataProductHubAPIServiceService.Initialize(initializeOptions)
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

			dataProductDraftPrototypeModel := &dataproducthubapiservicev1.DataProductDraftPrototype{
				Name: core.StringPtr("My New Data Product"),
				Asset: assetPrototypeModel,
			}

			createDataProductOptions := dataProductHubAPIServiceService.NewCreateDataProductOptions(
				[]dataproducthubapiservicev1.DataProductDraftPrototype{*dataProductDraftPrototypeModel},
			)

			dataProduct, response, err := dataProductHubAPIServiceService.CreateDataProduct(createDataProductOptions)
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
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Release: dataProductDraftVersionReleaseModel,
			}

			createDataProductDraftOptions := dataProductHubAPIServiceService.NewCreateDataProductDraftOptions(
				createNewDraftByDataProductIDLink,
				assetPrototypeModel,
			)
			createDataProductDraftOptions.SetVersion("1.2.0")
			createDataProductDraftOptions.SetDataProduct(dataProductIdentityModel)

			dataProductDraft, response, err := dataProductHubAPIServiceService.CreateDataProductDraft(createDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDraft, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductDraft).ToNot(BeNil())

			getADraftContractDocumentByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftContractDocumentByDraftIDLink value: %v\n", getADraftContractDocumentByDraftIDLink)
			updateADraftByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByContractTermsIDLink value: %v\n", updateADraftByContractTermsIDLink)
			createAContractTermsDocByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByContractTermsIDLink value: %v\n", createAContractTermsDocByContractTermsIDLink)
			updateContractDocumentByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateContractDocumentByDraftIDLink value: %v\n", updateContractDocumentByDraftIDLink)
			getAReleaseContractTermsByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getAReleaseContractTermsByContractTermsIDLink value: %v\n", getAReleaseContractTermsByContractTermsIDLink)
			completeADraftByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByContractTermsIDLink value: %v\n", completeADraftByContractTermsIDLink)
			getDraftByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved getDraftByDraftIDLink value: %v\n", getDraftByDraftIDLink)
			publishADraftByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved publishADraftByDraftIDLink value: %v\n", publishADraftByDraftIDLink)
			updateADraftByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved updateADraftByDraftIDLink value: %v\n", updateADraftByDraftIDLink)
			createAContractTermsDocByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved createAContractTermsDocByDraftIDLink value: %v\n", createAContractTermsDocByDraftIDLink)
			deleteAContractDocumentByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteAContractDocumentByDraftIDLink value: %v\n", deleteAContractDocumentByDraftIDLink)
			deleteADraftByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByContractTermsIDLink value: %v\n", deleteADraftByContractTermsIDLink)
			deleteADraftByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved deleteADraftByDraftIDLink value: %v\n", deleteADraftByDraftIDLink)
			completeADraftByDraftIDLink = *dataProductDraft.ID
			fmt.Fprintf(GinkgoWriter, "Saved completeADraftByDraftIDLink value: %v\n", completeADraftByDraftIDLink)
			getADraftByContractTermsIDLink = *dataProductDraft.ContractTerms[0].ID
			fmt.Fprintf(GinkgoWriter, "Saved getADraftByContractTermsIDLink value: %v\n", getADraftByContractTermsIDLink)
		})
		It(`CreateDraftContractTermsDocument request example`, func() {
			fmt.Println("\nCreateDraftContractTermsDocument() result:")
			// begin-create_draft_contract_terms_document

			createDraftContractTermsDocumentOptions := dataProductHubAPIServiceService.NewCreateDraftContractTermsDocumentOptions(
				uploadContractTermsDocByDataProductIDLink,
				createAContractTermsDocByDraftIDLink,
				createAContractTermsDocByContractTermsIDLink,
				"terms_and_conditions",
				"Terms and conditions document",
			)

			contractTermsDocument, response, err := dataProductHubAPIServiceService.CreateDraftContractTermsDocument(createDraftContractTermsDocumentOptions)
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

			publishDataProductDraftOptions := dataProductHubAPIServiceService.NewPublishDataProductDraftOptions(
				publishADraftOfDataProductByDataProductIDLink,
				publishADraftByDraftIDLink,
			)

			dataProductRelease, response, err := dataProductHubAPIServiceService.PublishDataProductDraft(publishDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductRelease, "", "  ")
			fmt.Println(string(b))

			// end-publish_data_product_draft

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
		It(`GetInitializeStatus request example`, func() {
			fmt.Println("\nGetInitializeStatus() result:")
			// begin-get_initialize_status

			getInitializeStatusOptions := dataProductHubAPIServiceService.NewGetInitializeStatusOptions()

			initializeResource, response, err := dataProductHubAPIServiceService.GetInitializeStatus(getInitializeStatusOptions)
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

			getServiceIDCredentialsOptions := dataProductHubAPIServiceService.NewGetServiceIDCredentialsOptions()

			serviceIDCredentials, response, err := dataProductHubAPIServiceService.GetServiceIDCredentials(getServiceIDCredentialsOptions)
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

			manageAPIKeysOptions := dataProductHubAPIServiceService.NewManageAPIKeysOptions()

			response, err := dataProductHubAPIServiceService.ManageAPIKeys(manageAPIKeysOptions)
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
		It(`CreateDataAssetVisualization request example`, func() {
			fmt.Println("\nCreateDataAssetVisualization() result:")
			// begin-create_data_asset_visualization

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("2be8f727-c5d2-4cb0-9216-f9888f428048"),
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dataproducthubapiservicev1.AssetReference{
				ID: core.StringPtr("caeee3f3-756e-47d5-846d-da4600809e22"),
				Container: containerReferenceModel,
			}

			dataAssetRelationshipModel := &dataproducthubapiservicev1.DataAssetRelationship{
				Asset: assetReferenceModel,
				RelatedAsset: assetReferenceModel,
			}

			createDataAssetVisualizationOptions := dataProductHubAPIServiceService.NewCreateDataAssetVisualizationOptions()
			createDataAssetVisualizationOptions.SetAssets([]dataproducthubapiservicev1.DataAssetRelationship{*dataAssetRelationshipModel})

			dataAssetVisualizationRes, response, err := dataProductHubAPIServiceService.CreateDataAssetVisualization(createDataAssetVisualizationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataAssetVisualizationRes, "", "  ")
			fmt.Println(string(b))

			// end-create_data_asset_visualization

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataAssetVisualizationRes).ToNot(BeNil())
		})
		It(`ReinitiateDataAssetVisualization request example`, func() {
			fmt.Println("\nReinitiateDataAssetVisualization() result:")
			// begin-reinitiate_data_asset_visualization

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("2be8f727-c5d2-4cb0-9216-f9888f428048"),
				Type: core.StringPtr("catalog"),
			}

			assetReferenceModel := &dataproducthubapiservicev1.AssetReference{
				ID: core.StringPtr("caeee3f3-756e-47d5-846d-da4600809e22"),
				Container: containerReferenceModel,
			}

			dataAssetRelationshipModel := &dataproducthubapiservicev1.DataAssetRelationship{
				Asset: assetReferenceModel,
				RelatedAsset: assetReferenceModel,
			}

			reinitiateDataAssetVisualizationOptions := dataProductHubAPIServiceService.NewReinitiateDataAssetVisualizationOptions()
			reinitiateDataAssetVisualizationOptions.SetAssets([]dataproducthubapiservicev1.DataAssetRelationship{*dataAssetRelationshipModel})

			dataAssetVisualizationRes, response, err := dataProductHubAPIServiceService.ReinitiateDataAssetVisualization(reinitiateDataAssetVisualizationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataAssetVisualizationRes, "", "  ")
			fmt.Println(string(b))

			// end-reinitiate_data_asset_visualization

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataAssetVisualizationRes).ToNot(BeNil())
		})
		It(`ListDataProducts request example`, func() {
			fmt.Println("\nListDataProducts() result:")
			// begin-list_data_products
			listDataProductsOptions := &dataproducthubapiservicev1.ListDataProductsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubAPIServiceService.NewDataProductsPager(listDataProductsOptions)
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

			getDataProductOptions := dataProductHubAPIServiceService.NewGetDataProductOptions(
				getDataProductByDataProductIDLink,
			)

			dataProduct, response, err := dataProductHubAPIServiceService.GetDataProduct(getDataProductOptions)
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

			completeDraftContractTermsDocumentOptions := dataProductHubAPIServiceService.NewCompleteDraftContractTermsDocumentOptions(
				completeDraftContractTermsByDataProductIDLink,
				completeADraftByDraftIDLink,
				completeADraftByContractTermsIDLink,
				completeContractTermsDocumentByDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIServiceService.CompleteDraftContractTermsDocument(completeDraftContractTermsDocumentOptions)
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
				DataProductID: &getListOfDataProductDraftsByDataProductIDLink,
				AssetContainerID: core.StringPtr("testString"),
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubAPIServiceService.NewDataProductDraftsPager(listDataProductDraftsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproducthubapiservicev1.DataProductDraftSummary
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

			getDataProductDraftOptions := dataProductHubAPIServiceService.NewGetDataProductDraftOptions(
				getADraftOfDataProductByDataProductIDLink,
				getDraftByDraftIDLink,
			)

			dataProductDraft, response, err := dataProductHubAPIServiceService.GetDataProductDraft(getDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDraft, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDraft).ToNot(BeNil())
		})
		It(`UpdateDataProductDraft request example`, func() {
			fmt.Println("\nUpdateDataProductDraft() result:")
			// begin-update_data_product_draft

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductDraftOptions := dataProductHubAPIServiceService.NewUpdateDataProductDraftOptions(
				updateDraftOfDataProductByDataProductIDLink,
				updateADraftByDraftIDLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductDraft, response, err := dataProductHubAPIServiceService.UpdateDataProductDraft(updateDataProductDraftOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDraft, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_draft

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDraft).ToNot(BeNil())
		})
		It(`GetDraftContractTermsDocument request example`, func() {
			fmt.Println("\nGetDraftContractTermsDocument() result:")
			// begin-get_draft_contract_terms_document

			getDraftContractTermsDocumentOptions := dataProductHubAPIServiceService.NewGetDraftContractTermsDocumentOptions(
				getContractDocumentByDataProductIDLink,
				getADraftContractDocumentByDraftIDLink,
				getADraftByContractTermsIDLink,
				getContractTermsDocumentByIDDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIServiceService.GetDraftContractTermsDocument(getDraftContractTermsDocumentOptions)
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
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDraftContractTermsDocumentOptions := dataProductHubAPIServiceService.NewUpdateDraftContractTermsDocumentOptions(
				updateContractDocumentByDataProductIDLink,
				updateContractDocumentByDraftIDLink,
				updateADraftByContractTermsIDLink,
				updateContractTermsDocumentByDocumentIDLink,
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTermsDocument, response, err := dataProductHubAPIServiceService.UpdateDraftContractTermsDocument(updateDraftContractTermsDocumentOptions)
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
		It(`GetDataProductDraftContractTerms request example`, func() {
			fmt.Println("\nGetDataProductDraftContractTerms() result:")
			// begin-get_data_product_draft_contract_terms

			getDataProductDraftContractTermsOptions := dataProductHubAPIServiceService.NewGetDataProductDraftContractTermsOptions(
				"testString",
				"testString",
				"testString",
			)

			result, response, err := dataProductHubAPIServiceService.GetDataProductDraftContractTerms(getDataProductDraftContractTermsOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil { panic(err) }
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil { panic(err) }
			}

			// end-get_data_product_draft_contract_terms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ReplaceDataProductDraftContractTerms request example`, func() {
			fmt.Println("\nReplaceDataProductDraftContractTerms() result:")
			// begin-replace_data_product_draft_contract_terms

			contractTermsDocumentModel := &dataproducthubapiservicev1.ContractTermsDocument{
				URL: core.StringPtr("https://ibm.com/document"),
				Type: core.StringPtr("terms_and_conditions"),
				Name: core.StringPtr("Terms and Conditions"),
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
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

			replaceDataProductDraftContractTermsOptions := dataProductHubAPIServiceService.NewReplaceDataProductDraftContractTermsOptions(
				"testString",
				"testString",
				"testString",
			)
			replaceDataProductDraftContractTermsOptions.SetDocuments([]dataproducthubapiservicev1.ContractTermsDocument{*contractTermsDocumentModel})
			replaceDataProductDraftContractTermsOptions.SetOverview(overviewModel)
			replaceDataProductDraftContractTermsOptions.SetDescription(descriptionModel)
			replaceDataProductDraftContractTermsOptions.SetOrganization([]dataproducthubapiservicev1.ContractTemplateOrganization{*contractTemplateOrganizationModel})
			replaceDataProductDraftContractTermsOptions.SetRoles([]dataproducthubapiservicev1.Roles{*rolesModel})
			replaceDataProductDraftContractTermsOptions.SetPrice(pricingModel)
			replaceDataProductDraftContractTermsOptions.SetSLA([]dataproducthubapiservicev1.ContractTemplateSLA{*contractTemplateSLAModel})
			replaceDataProductDraftContractTermsOptions.SetSupportAndCommunication([]dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{*contractTemplateSupportAndCommunicationModel})
			replaceDataProductDraftContractTermsOptions.SetCustomProperties([]dataproducthubapiservicev1.ContractTemplateCustomProperty{*contractTemplateCustomPropertyModel})

			contractTerms, response, err := dataProductHubAPIServiceService.ReplaceDataProductDraftContractTerms(replaceDataProductDraftContractTermsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTerms, "", "  ")
			fmt.Println(string(b))

			// end-replace_data_product_draft_contract_terms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTerms).ToNot(BeNil())
		})
		It(`UpdateDataProductDraftContractTerms request example`, func() {
			fmt.Println("\nUpdateDataProductDraftContractTerms() result:")
			// begin-update_data_product_draft_contract_terms

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductDraftContractTermsOptions := dataProductHubAPIServiceService.NewUpdateDataProductDraftContractTermsOptions(
				"testString",
				"testString",
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			contractTerms, response, err := dataProductHubAPIServiceService.UpdateDataProductDraftContractTerms(updateDataProductDraftContractTermsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(contractTerms, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_draft_contract_terms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(contractTerms).ToNot(BeNil())
		})
		It(`GetDataProductRelease request example`, func() {
			fmt.Println("\nGetDataProductRelease() result:")
			// begin-get_data_product_release

			getDataProductReleaseOptions := dataProductHubAPIServiceService.NewGetDataProductReleaseOptions(
				getAReleaseOfDataProductByDataProductIDLink,
				getAReleaseByReleaseIDLink,
			)

			dataProductRelease, response, err := dataProductHubAPIServiceService.GetDataProductRelease(getDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductRelease, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
		It(`UpdateDataProductRelease request example`, func() {
			fmt.Println("\nUpdateDataProductRelease() result:")
			// begin-update_data_product_release

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductReleaseOptions := dataProductHubAPIServiceService.NewUpdateDataProductReleaseOptions(
				updateReleaseOfDataProductByDataProductIDLink,
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductRelease, response, err := dataProductHubAPIServiceService.UpdateDataProductRelease(updateDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductRelease, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
		It(`GetReleaseContractTermsDocument request example`, func() {
			fmt.Println("\nGetReleaseContractTermsDocument() result:")
			// begin-get_release_contract_terms_document

			getReleaseContractTermsDocumentOptions := dataProductHubAPIServiceService.NewGetReleaseContractTermsDocumentOptions(
				getReleaseContractDocumentByDataProductIDLink,
				getAReleaseContractTermsByReleaseIDLink,
				getAReleaseContractTermsByContractTermsIDLink,
				getReleaseContractDocumentByDocumentIDLink,
			)

			contractTermsDocument, response, err := dataProductHubAPIServiceService.GetReleaseContractTermsDocument(getReleaseContractTermsDocumentOptions)
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
		It(`GetPublishedDataProductDraftContractTerms request example`, func() {
			fmt.Println("\nGetPublishedDataProductDraftContractTerms() result:")
			// begin-get_published_data_product_draft_contract_terms

			getPublishedDataProductDraftContractTermsOptions := dataProductHubAPIServiceService.NewGetPublishedDataProductDraftContractTermsOptions(
				"testString",
				"testString",
				"testString",
			)

			result, response, err := dataProductHubAPIServiceService.GetPublishedDataProductDraftContractTerms(getPublishedDataProductDraftContractTermsOptions)
			if err != nil {
				panic(err)
			}
			if result != nil {
				defer result.Close()
				outFile, err := os.Create("result.out")
				if err != nil { panic(err) }
				defer outFile.Close()
				_, err = io.Copy(outFile, result)
				if err != nil { panic(err) }
			}

			// end-get_published_data_product_draft_contract_terms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
		It(`ListDataProductReleases request example`, func() {
			fmt.Println("\nListDataProductReleases() result:")
			// begin-list_data_product_releases
			listDataProductReleasesOptions := &dataproducthubapiservicev1.ListDataProductReleasesOptions{
				DataProductID: &getListOfReleasesOfDataProductByDataProductIDLink,
				AssetContainerID: core.StringPtr("testString"),
				State: []string{"available"},
				Version: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := dataProductHubAPIServiceService.NewDataProductReleasesPager(listDataProductReleasesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []dataproducthubapiservicev1.DataProductReleaseSummary
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

			retireDataProductReleaseOptions := dataProductHubAPIServiceService.NewRetireDataProductReleaseOptions(
				retireAReleasesOfDataProductByDataProductIDLink,
				retireAReleaseContractTermsByReleaseIDLink,
			)

			dataProductRelease, response, err := dataProductHubAPIServiceService.RetireDataProductRelease(retireDataProductReleaseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductRelease, "", "  ")
			fmt.Println(string(b))

			// end-retire_data_product_release

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductRelease).ToNot(BeNil())
		})
		It(`ListDataProductContractTemplate request example`, func() {
			fmt.Println("\nListDataProductContractTemplate() result:")
			// begin-list_data_product_contract_template

			listDataProductContractTemplateOptions := dataProductHubAPIServiceService.NewListDataProductContractTemplateOptions()

			dataProductContractTemplateCollection, response, err := dataProductHubAPIServiceService.ListDataProductContractTemplate(listDataProductContractTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductContractTemplateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_data_product_contract_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplateCollection).ToNot(BeNil())
		})
		It(`CreateContractTemplate request example`, func() {
			fmt.Println("\nCreateContractTemplate() result:")
			// begin-create_contract_template

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("f531f74a-01c8-4e91-8e29-b018db683c86"),
				Type: core.StringPtr("catalog"),
			}

			domainModel := &dataproducthubapiservicev1.Domain{
				ID: core.StringPtr("b38df608-d34b-4d58-8136-ed25e6c6684e"),
				Name: core.StringPtr("domain_name"),
			}

			overviewModel := &dataproducthubapiservicev1.Overview{
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

			contractTermsModel := &dataproducthubapiservicev1.ContractTerms{
				Overview: overviewModel,
				Description: descriptionModel,
				Organization: []dataproducthubapiservicev1.ContractTemplateOrganization{*contractTemplateOrganizationModel},
				Roles: []dataproducthubapiservicev1.Roles{*rolesModel},
				Price: pricingModel,
				SLA: []dataproducthubapiservicev1.ContractTemplateSLA{*contractTemplateSLAModel},
				SupportAndCommunication: []dataproducthubapiservicev1.ContractTemplateSupportAndCommunication{*contractTemplateSupportAndCommunicationModel},
				CustomProperties: []dataproducthubapiservicev1.ContractTemplateCustomProperty{*contractTemplateCustomPropertyModel},
			}

			createContractTemplateOptions := dataProductHubAPIServiceService.NewCreateContractTemplateOptions(
				containerReferenceModel,
			)
			createContractTemplateOptions.SetName("Sample Data Contract Template")
			createContractTemplateOptions.SetContractTerms(contractTermsModel)

			dataProductContractTemplate, response, err := dataProductHubAPIServiceService.CreateContractTemplate(createContractTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductContractTemplate, "", "  ")
			fmt.Println(string(b))

			// end-create_contract_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
		It(`GetContractTemplate request example`, func() {
			fmt.Println("\nGetContractTemplate() result:")
			// begin-get_contract_template

			getContractTemplateOptions := dataProductHubAPIServiceService.NewGetContractTemplateOptions(
				"testString",
				"testString",
			)

			dataProductContractTemplate, response, err := dataProductHubAPIServiceService.GetContractTemplate(getContractTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductContractTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_contract_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
		It(`UpdateDataProductContractTemplate request example`, func() {
			fmt.Println("\nUpdateDataProductContractTemplate() result:")
			// begin-update_data_product_contract_template

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductContractTemplateOptions := dataProductHubAPIServiceService.NewUpdateDataProductContractTemplateOptions(
				"testString",
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductContractTemplate, response, err := dataProductHubAPIServiceService.UpdateDataProductContractTemplate(updateDataProductContractTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductContractTemplate, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_contract_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductContractTemplate).ToNot(BeNil())
		})
		It(`ListDataProductDomains request example`, func() {
			fmt.Println("\nListDataProductDomains() result:")
			// begin-list_data_product_domains

			listDataProductDomainsOptions := dataProductHubAPIServiceService.NewListDataProductDomainsOptions()

			dataProductDomainCollection, response, err := dataProductHubAPIServiceService.ListDataProductDomains(listDataProductDomainsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDomainCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_data_product_domains

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomainCollection).ToNot(BeNil())
		})
		It(`CreateDataProductDomain request example`, func() {
			fmt.Println("\nCreateDataProductDomain() result:")
			// begin-create_data_product_domain

			containerReferenceModel := &dataproducthubapiservicev1.ContainerReference{
				ID: core.StringPtr("ed580171-a6e4-4b93-973f-ae2f2f62991b"),
				Type: core.StringPtr("catalog"),
			}

			initializeSubDomainModel := &dataproducthubapiservicev1.InitializeSubDomain{
				Name: core.StringPtr("Sub domain 1"),
				Description: core.StringPtr("New sub domain 1"),
			}

			createDataProductDomainOptions := dataProductHubAPIServiceService.NewCreateDataProductDomainOptions(
				containerReferenceModel,
			)
			createDataProductDomainOptions.SetName("Test domain")
			createDataProductDomainOptions.SetDescription("The sample description for new domain")
			createDataProductDomainOptions.SetSubDomains([]dataproducthubapiservicev1.InitializeSubDomain{*initializeSubDomainModel})

			dataProductDomain, response, err := dataProductHubAPIServiceService.CreateDataProductDomain(createDataProductDomainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDomain, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_domain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dataProductDomain).ToNot(BeNil())
		})
		It(`CreateDataProductSubdomain request example`, func() {
			fmt.Println("\nCreateDataProductSubdomain() result:")
			// begin-create_data_product_subdomain

			createDataProductSubdomainOptions := dataProductHubAPIServiceService.NewCreateDataProductSubdomainOptions(
				"testString",
				"testString",
			)
			createDataProductSubdomainOptions.SetName("Sub domain 1")
			createDataProductSubdomainOptions.SetDescription("New sub domain 1")

			initializeSubDomain, response, err := dataProductHubAPIServiceService.CreateDataProductSubdomain(createDataProductSubdomainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(initializeSubDomain, "", "  ")
			fmt.Println(string(b))

			// end-create_data_product_subdomain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(initializeSubDomain).ToNot(BeNil())
		})
		It(`GetDomain request example`, func() {
			fmt.Println("\nGetDomain() result:")
			// begin-get_domain

			getDomainOptions := dataProductHubAPIServiceService.NewGetDomainOptions(
				"testString",
			)

			dataProductDomain, response, err := dataProductHubAPIServiceService.GetDomain(getDomainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDomain, "", "  ")
			fmt.Println(string(b))

			// end-get_domain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomain).ToNot(BeNil())
		})
		It(`UpdateDataProductDomain request example`, func() {
			fmt.Println("\nUpdateDataProductDomain() result:")
			// begin-update_data_product_domain

			jsonPatchOperationModel := &dataproducthubapiservicev1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDataProductDomainOptions := dataProductHubAPIServiceService.NewUpdateDataProductDomainOptions(
				"testString",
				"testString",
				[]dataproducthubapiservicev1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			dataProductDomain, response, err := dataProductHubAPIServiceService.UpdateDataProductDomain(updateDataProductDomainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductDomain, "", "  ")
			fmt.Println(string(b))

			// end-update_data_product_domain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductDomain).ToNot(BeNil())
		})
		It(`GetDataProductByDomain request example`, func() {
			fmt.Println("\nGetDataProductByDomain() result:")
			// begin-get_data_product_by_domain

			getDataProductByDomainOptions := dataProductHubAPIServiceService.NewGetDataProductByDomainOptions(
				"testString",
				"testString",
			)

			dataProductVersionCollection, response, err := dataProductHubAPIServiceService.GetDataProductByDomain(getDataProductByDomainOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataProductVersionCollection, "", "  ")
			fmt.Println(string(b))

			// end-get_data_product_by_domain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataProductVersionCollection).ToNot(BeNil())
		})
		It(`CreateS3Bucket request example`, func() {
			fmt.Println("\nCreateS3Bucket() result:")
			// begin-create_s3_bucket

			createS3BucketOptions := dataProductHubAPIServiceService.NewCreateS3BucketOptions(
				true,
			)

			bucketResponse, response, err := dataProductHubAPIServiceService.CreateS3Bucket(createS3BucketOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(bucketResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_s3_bucket

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bucketResponse).ToNot(BeNil())
		})
		It(`GetS3BucketValidation request example`, func() {
			fmt.Println("\nGetS3BucketValidation() result:")
			// begin-get_s3_bucket_validation

			getS3BucketValidationOptions := dataProductHubAPIServiceService.NewGetS3BucketValidationOptions(
				"testString",
			)

			bucketValidationResponse, response, err := dataProductHubAPIServiceService.GetS3BucketValidation(getS3BucketValidationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(bucketValidationResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_s3_bucket_validation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bucketValidationResponse).ToNot(BeNil())
		})
		It(`DeleteDraftContractTermsDocument request example`, func() {
			// begin-delete_draft_contract_terms_document

			deleteDraftContractTermsDocumentOptions := dataProductHubAPIServiceService.NewDeleteDraftContractTermsDocumentOptions(
				deleteContractDocumentByDataProductIDLink,
				deleteAContractDocumentByDraftIDLink,
				deleteADraftByContractTermsIDLink,
				deleteContractTermsDocumentByDocumentIDLink,
			)

			response, err := dataProductHubAPIServiceService.DeleteDraftContractTermsDocument(deleteDraftContractTermsDocumentOptions)
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

			deleteDataProductDraftOptions := dataProductHubAPIServiceService.NewDeleteDataProductDraftOptions(
				deleteDraftOfDataProductByDataProductIDLink,
				deleteADraftByDraftIDLink,
			)

			response, err := dataProductHubAPIServiceService.DeleteDataProductDraft(deleteDataProductDraftOptions)
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
		It(`DeleteDataProductContractTemplate request example`, func() {
			// begin-delete_data_product_contract_template

			deleteDataProductContractTemplateOptions := dataProductHubAPIServiceService.NewDeleteDataProductContractTemplateOptions(
				"testString",
				"testString",
			)

			response, err := dataProductHubAPIServiceService.DeleteDataProductContractTemplate(deleteDataProductContractTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDataProductContractTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_data_product_contract_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteDomain request example`, func() {
			// begin-delete_domain

			deleteDomainOptions := dataProductHubAPIServiceService.NewDeleteDomainOptions(
				"testString",
			)

			response, err := dataProductHubAPIServiceService.DeleteDomain(deleteDomainOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteDomain(): %d\n", response.StatusCode)
			}

			// end-delete_domain

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
