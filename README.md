# FattureInCloud Go SDK

![Release](https://img.shields.io/github/v/release/fattureincloud/fattureincloud-go-sdk) ![unit tests](https://github.com/fattureincloud/fattureincloud-go-sdk/actions/workflows/validate.yaml/badge.svg)

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 500.000 businesses in Italy. 

The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.24
- Package version: 2.0.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.fattureincloud.it](https://www.fattureincloud.it)

## Installation

Install the following dependencies:

```shell
go get github.com/fattureincloud/fattureincloud-go-sdk/v2
```

## Documentation for API Endpoints

All URIs are relative to *https://api-v2.fattureincloud.it*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ArchiveApi* | [**CreateArchiveDocument**](docs/ArchiveApi.md#createarchivedocument) | **Post** /c/{company_id}/archive | Create Archive Document
*ArchiveApi* | [**DeleteArchiveDocument**](docs/ArchiveApi.md#deletearchivedocument) | **Delete** /c/{company_id}/archive/{document_id} | Delete Archive Document
*ArchiveApi* | [**GetArchiveDocument**](docs/ArchiveApi.md#getarchivedocument) | **Get** /c/{company_id}/archive/{document_id} | Get Archive Document
*ArchiveApi* | [**ListArchiveDocuments**](docs/ArchiveApi.md#listarchivedocuments) | **Get** /c/{company_id}/archive | List Archive Documents
*ArchiveApi* | [**ModifyArchiveDocument**](docs/ArchiveApi.md#modifyarchivedocument) | **Put** /c/{company_id}/archive/{document_id} | Modify Archive Document
*ArchiveApi* | [**UploadArchiveDocumentAttachment**](docs/ArchiveApi.md#uploadarchivedocumentattachment) | **Post** /c/{company_id}/archive/attachment | Upload Archive Document Attachment
*CashbookApi* | [**CreateCashbookEntry**](docs/CashbookApi.md#createcashbookentry) | **Post** /c/{company_id}/cashbook | Create Cashbook Entry
*CashbookApi* | [**DeleteCashbookEntry**](docs/CashbookApi.md#deletecashbookentry) | **Delete** /c/{company_id}/cashbook/{document_id} | Delete Cashbook Entry
*CashbookApi* | [**GetCashbookEntry**](docs/CashbookApi.md#getcashbookentry) | **Get** /c/{company_id}/cashbook/{document_id} | Get Cashbook Entry
*CashbookApi* | [**ListCashbookEntries**](docs/CashbookApi.md#listcashbookentries) | **Get** /c/{company_id}/cashbook | List Cashbook Entries
*CashbookApi* | [**ModifyCashbookEntry**](docs/CashbookApi.md#modifycashbookentry) | **Put** /c/{company_id}/cashbook/{document_id} | Modify Cashbook Entry
*ClientsApi* | [**CreateClient**](docs/ClientsApi.md#createclient) | **Post** /c/{company_id}/entities/clients | Create Client
*ClientsApi* | [**DeleteClient**](docs/ClientsApi.md#deleteclient) | **Delete** /c/{company_id}/entities/clients/{client_id} | Delete Client
*ClientsApi* | [**GetClient**](docs/ClientsApi.md#getclient) | **Get** /c/{company_id}/entities/clients/{client_id} | Get Client
*ClientsApi* | [**ListClients**](docs/ClientsApi.md#listclients) | **Get** /c/{company_id}/entities/clients | List Clients
*ClientsApi* | [**ModifyClient**](docs/ClientsApi.md#modifyclient) | **Put** /c/{company_id}/entities/clients/{client_id} | Modify Client
*CompaniesApi* | [**GetCompanyInfo**](docs/CompaniesApi.md#getcompanyinfo) | **Get** /c/{company_id}/company/info | Get Company Info
*EmailsApi* | [**ListEmails**](docs/EmailsApi.md#listemails) | **Get** /c/{company_id}/emails | List emails
*InfoApi* | [**ListArchiveCategories**](docs/InfoApi.md#listarchivecategories) | **Get** /c/{company_id}/info/archive_categories | List Archive Categories
*InfoApi* | [**ListCities**](docs/InfoApi.md#listcities) | **Get** /info/cities | List Cities
*InfoApi* | [**ListCostCenters**](docs/InfoApi.md#listcostcenters) | **Get** /c/{company_id}/info/cost_centers | List Cost Centers
*InfoApi* | [**ListCountries**](docs/InfoApi.md#listcountries) | **Get** /info/countries | List Countries
*InfoApi* | [**ListCurrencies**](docs/InfoApi.md#listcurrencies) | **Get** /info/currencies | List Currencies
*InfoApi* | [**ListDeliveryNotesDefaultCausals**](docs/InfoApi.md#listdeliverynotesdefaultcausals) | **Get** /info/dn_causals | List Delivery Notes Default Causals
*InfoApi* | [**ListDetailedCountries**](docs/InfoApi.md#listdetailedcountries) | **Get** /info/detailed_countries | List Detailed Countries
*InfoApi* | [**ListLanguages**](docs/InfoApi.md#listlanguages) | **Get** /info/languages | List Languages
*InfoApi* | [**ListPaymentAccounts**](docs/InfoApi.md#listpaymentaccounts) | **Get** /c/{company_id}/info/payment_accounts | List Payment Accounts
*InfoApi* | [**ListPaymentMethods**](docs/InfoApi.md#listpaymentmethods) | **Get** /c/{company_id}/info/payment_methods | List Payment Methods
*InfoApi* | [**ListProductCategories**](docs/InfoApi.md#listproductcategories) | **Get** /c/{company_id}/info/product_categories | List Product Categories
*InfoApi* | [**ListReceivedDocumentCategories**](docs/InfoApi.md#listreceiveddocumentcategories) | **Get** /c/{company_id}/info/received_document_categories | List Received Document Categories
*InfoApi* | [**ListRevenueCenters**](docs/InfoApi.md#listrevenuecenters) | **Get** /c/{company_id}/info/revenue_centers | List Revenue Centers
*InfoApi* | [**ListTemplates**](docs/InfoApi.md#listtemplates) | **Get** /info/templates | List Templates
*InfoApi* | [**ListUnitsOfMeasure**](docs/InfoApi.md#listunitsofmeasure) | **Get** /info/measures | List Units of Measure
*InfoApi* | [**ListVatTypes**](docs/InfoApi.md#listvattypes) | **Get** /c/{company_id}/info/vat_types | List Vat Types
*IssuedDocumentsApi* | [**CreateIssuedDocument**](docs/IssuedDocumentsApi.md#createissueddocument) | **Post** /c/{company_id}/issued_documents | Create Issued Document
*IssuedDocumentsApi* | [**DeleteIssuedDocument**](docs/IssuedDocumentsApi.md#deleteissueddocument) | **Delete** /c/{company_id}/issued_documents/{document_id} | Delete Issued Document
*IssuedDocumentsApi* | [**DeleteIssuedDocumentAttachment**](docs/IssuedDocumentsApi.md#deleteissueddocumentattachment) | **Delete** /c/{company_id}/issued_documents/{document_id}/attachment | Delete Issued Document Attachment
*IssuedDocumentsApi* | [**GetEmailData**](docs/IssuedDocumentsApi.md#getemaildata) | **Get** /c/{company_id}/issued_documents/{document_id}/email | Get Email Data
*IssuedDocumentsApi* | [**GetExistingIssuedDocumentTotals**](docs/IssuedDocumentsApi.md#getexistingissueddocumenttotals) | **Post** /c/{company_id}/issued_documents/{document_id}/totals | Get Existing Issued Document Totals
*IssuedDocumentsApi* | [**GetIssuedDocument**](docs/IssuedDocumentsApi.md#getissueddocument) | **Get** /c/{company_id}/issued_documents/{document_id} | Get Issued Document
*IssuedDocumentsApi* | [**GetIssuedDocumentPreCreateInfo**](docs/IssuedDocumentsApi.md#getissueddocumentprecreateinfo) | **Get** /c/{company_id}/issued_documents/info | Get Issued Document Pre-create info
*IssuedDocumentsApi* | [**GetNewIssuedDocumentTotals**](docs/IssuedDocumentsApi.md#getnewissueddocumenttotals) | **Post** /c/{company_id}/issued_documents/totals | Get New Issued Document Totals
*IssuedDocumentsApi* | [**JoinIssuedDocuments**](docs/IssuedDocumentsApi.md#joinissueddocuments) | **Get** /c/{company_id}/issued_documents/join | Join issued documents
*IssuedDocumentsApi* | [**ListIssuedDocuments**](docs/IssuedDocumentsApi.md#listissueddocuments) | **Get** /c/{company_id}/issued_documents | List Issued Documents
*IssuedDocumentsApi* | [**ModifyIssuedDocument**](docs/IssuedDocumentsApi.md#modifyissueddocument) | **Put** /c/{company_id}/issued_documents/{document_id} | Modify Issued Document
*IssuedDocumentsApi* | [**ScheduleEmail**](docs/IssuedDocumentsApi.md#scheduleemail) | **Post** /c/{company_id}/issued_documents/{document_id}/email | Schedule Email
*IssuedDocumentsApi* | [**TransformIssuedDocument**](docs/IssuedDocumentsApi.md#transformissueddocument) | **Get** /c/{company_id}/issued_documents/transform | Transform issued document
*IssuedDocumentsApi* | [**UploadIssuedDocumentAttachment**](docs/IssuedDocumentsApi.md#uploadissueddocumentattachment) | **Post** /c/{company_id}/issued_documents/attachment | Upload Issued Document Attachment
*IssuedEInvoicesApi* | [**GetEInvoiceRejectionReason**](docs/IssuedEInvoicesApi.md#geteinvoicerejectionreason) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/error_reason | Get e-invoice rejection reason
*IssuedEInvoicesApi* | [**GetEInvoiceXml**](docs/IssuedEInvoicesApi.md#geteinvoicexml) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/xml | Get e-invoice XML
*IssuedEInvoicesApi* | [**SendEInvoice**](docs/IssuedEInvoicesApi.md#sendeinvoice) | **Post** /c/{company_id}/issued_documents/{document_id}/e_invoice/send | Send the e-invoice
*IssuedEInvoicesApi* | [**VerifyEInvoiceXml**](docs/IssuedEInvoicesApi.md#verifyeinvoicexml) | **Get** /c/{company_id}/issued_documents/{document_id}/e_invoice/xml_verify | Verify e-invoice XML
*ProductsApi* | [**CreateProduct**](docs/ProductsApi.md#createproduct) | **Post** /c/{company_id}/products | Create Product
*ProductsApi* | [**DeleteProduct**](docs/ProductsApi.md#deleteproduct) | **Delete** /c/{company_id}/products/{product_id} | Delete Product
*ProductsApi* | [**GetProduct**](docs/ProductsApi.md#getproduct) | **Get** /c/{company_id}/products/{product_id} | Get Product
*ProductsApi* | [**ListProducts**](docs/ProductsApi.md#listproducts) | **Get** /c/{company_id}/products | List Products
*ProductsApi* | [**ModifyProduct**](docs/ProductsApi.md#modifyproduct) | **Put** /c/{company_id}/products/{product_id} | Modify Product
*ReceiptsApi* | [**CreateReceipt**](docs/ReceiptsApi.md#createreceipt) | **Post** /c/{company_id}/receipts | Create Receipt
*ReceiptsApi* | [**DeleteReceipt**](docs/ReceiptsApi.md#deletereceipt) | **Delete** /c/{company_id}/receipts/{document_id} | Delete Receipt
*ReceiptsApi* | [**GetReceipt**](docs/ReceiptsApi.md#getreceipt) | **Get** /c/{company_id}/receipts/{document_id} | Get Receipt
*ReceiptsApi* | [**GetReceiptPreCreateInfo**](docs/ReceiptsApi.md#getreceiptprecreateinfo) | **Get** /c/{company_id}/receipts/info | Get Receipt Pre-Create Info
*ReceiptsApi* | [**GetReceiptsMonthlyTotals**](docs/ReceiptsApi.md#getreceiptsmonthlytotals) | **Get** /c/{company_id}/receipts/monthly_totals | Get Receipts Monthly Totals
*ReceiptsApi* | [**ListReceipts**](docs/ReceiptsApi.md#listreceipts) | **Get** /c/{company_id}/receipts | List Receipts
*ReceiptsApi* | [**ModifyReceipt**](docs/ReceiptsApi.md#modifyreceipt) | **Put** /c/{company_id}/receipts/{document_id} | Modify Receipt
*ReceivedDocumentsApi* | [**CreateReceivedDocument**](docs/ReceivedDocumentsApi.md#createreceiveddocument) | **Post** /c/{company_id}/received_documents | Create Received Document
*ReceivedDocumentsApi* | [**DeleteReceivedDocument**](docs/ReceivedDocumentsApi.md#deletereceiveddocument) | **Delete** /c/{company_id}/received_documents/{document_id} | Delete Received Document
*ReceivedDocumentsApi* | [**DeleteReceivedDocumentAttachment**](docs/ReceivedDocumentsApi.md#deletereceiveddocumentattachment) | **Delete** /c/{company_id}/received_documents/{document_id}/attachment | Delete Received Document Attachment
*ReceivedDocumentsApi* | [**GetExistingReceivedDocumentTotals**](docs/ReceivedDocumentsApi.md#getexistingreceiveddocumenttotals) | **Post** /c/{company_id}/received_documents/{document_id}/totals | Get Existing Received Document Totals
*ReceivedDocumentsApi* | [**GetNewReceivedDocumentTotals**](docs/ReceivedDocumentsApi.md#getnewreceiveddocumenttotals) | **Post** /c/{company_id}/received_documents/totals | Get New Received Document Totals
*ReceivedDocumentsApi* | [**GetReceivedDocument**](docs/ReceivedDocumentsApi.md#getreceiveddocument) | **Get** /c/{company_id}/received_documents/{document_id} | Get Received Document
*ReceivedDocumentsApi* | [**GetReceivedDocumentPreCreateInfo**](docs/ReceivedDocumentsApi.md#getreceiveddocumentprecreateinfo) | **Get** /c/{company_id}/received_documents/info | Get Received Document Pre-Create Info
*ReceivedDocumentsApi* | [**ListReceivedDocuments**](docs/ReceivedDocumentsApi.md#listreceiveddocuments) | **Get** /c/{company_id}/received_documents | List Received Documents
*ReceivedDocumentsApi* | [**ModifyReceivedDocument**](docs/ReceivedDocumentsApi.md#modifyreceiveddocument) | **Put** /c/{company_id}/received_documents/{document_id} | Modify Received Document
*ReceivedDocumentsApi* | [**UploadReceivedDocumentAttachment**](docs/ReceivedDocumentsApi.md#uploadreceiveddocumentattachment) | **Post** /c/{company_id}/received_documents/attachment | Upload Received Document Attachment
*SettingsApi* | [**CreatePaymentAccount**](docs/SettingsApi.md#createpaymentaccount) | **Post** /c/{company_id}/settings/payment_accounts | Create Payment Account
*SettingsApi* | [**CreatePaymentMethod**](docs/SettingsApi.md#createpaymentmethod) | **Post** /c/{company_id}/settings/payment_methods | Create Payment Method
*SettingsApi* | [**CreateVatType**](docs/SettingsApi.md#createvattype) | **Post** /c/{company_id}/settings/vat_types | Create Vat Type
*SettingsApi* | [**DeletePaymentAccount**](docs/SettingsApi.md#deletepaymentaccount) | **Delete** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Delete Payment Account
*SettingsApi* | [**DeletePaymentMethod**](docs/SettingsApi.md#deletepaymentmethod) | **Delete** /c/{company_id}/settings/payment_methods/{payment_method_id} | Delete Payment Method
*SettingsApi* | [**DeleteVatType**](docs/SettingsApi.md#deletevattype) | **Delete** /c/{company_id}/settings/vat_types/{vat_type_id} | Delete Vat Type
*SettingsApi* | [**GetPaymentAccount**](docs/SettingsApi.md#getpaymentaccount) | **Get** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Get Payment Account
*SettingsApi* | [**GetPaymentMethod**](docs/SettingsApi.md#getpaymentmethod) | **Get** /c/{company_id}/settings/payment_methods/{payment_method_id} | Get Payment Method
*SettingsApi* | [**GetVatType**](docs/SettingsApi.md#getvattype) | **Get** /c/{company_id}/settings/vat_types/{vat_type_id} | Get Vat Type
*SettingsApi* | [**ModifyPaymentAccount**](docs/SettingsApi.md#modifypaymentaccount) | **Put** /c/{company_id}/settings/payment_accounts/{payment_account_id} | Modify Payment Account
*SettingsApi* | [**ModifyPaymentMethod**](docs/SettingsApi.md#modifypaymentmethod) | **Put** /c/{company_id}/settings/payment_methods/{payment_method_id} | Modify Payment Method
*SettingsApi* | [**ModifyVatType**](docs/SettingsApi.md#modifyvattype) | **Put** /c/{company_id}/settings/vat_types/{vat_type_id} | Modify Vat Type
*SuppliersApi* | [**CreateSupplier**](docs/SuppliersApi.md#createsupplier) | **Post** /c/{company_id}/entities/suppliers | Create Supplier
*SuppliersApi* | [**DeleteSupplier**](docs/SuppliersApi.md#deletesupplier) | **Delete** /c/{company_id}/entities/suppliers/{supplier_id} | Delete Supplier
*SuppliersApi* | [**GetSupplier**](docs/SuppliersApi.md#getsupplier) | **Get** /c/{company_id}/entities/suppliers/{supplier_id} | Get Supplier
*SuppliersApi* | [**ListSuppliers**](docs/SuppliersApi.md#listsuppliers) | **Get** /c/{company_id}/entities/suppliers | List Suppliers
*SuppliersApi* | [**ModifySupplier**](docs/SuppliersApi.md#modifysupplier) | **Put** /c/{company_id}/entities/suppliers/{supplier_id} | Modify Supplier
*TaxesApi* | [**CreateF24**](docs/TaxesApi.md#createf24) | **Post** /c/{company_id}/taxes | Create F24
*TaxesApi* | [**DeleteF24**](docs/TaxesApi.md#deletef24) | **Delete** /c/{company_id}/taxes/{document_id} | Delete F24
*TaxesApi* | [**DeleteF24Attachment**](docs/TaxesApi.md#deletef24attachment) | **Delete** /c/{company_id}/taxes/{document_id}/attachment | Delete F24 Attachment
*TaxesApi* | [**GetF24**](docs/TaxesApi.md#getf24) | **Get** /c/{company_id}/taxes/{document_id} | Get F24
*TaxesApi* | [**ListF24**](docs/TaxesApi.md#listf24) | **Get** /c/{company_id}/taxes | List F24
*TaxesApi* | [**ModifyF24**](docs/TaxesApi.md#modifyf24) | **Put** /c/{company_id}/taxes/{document_id} | Modify F24
*TaxesApi* | [**UploadF24Attachment**](docs/TaxesApi.md#uploadf24attachment) | **Post** /c/{company_id}/taxes/attachment | Upload F24 Attachment
*UserApi* | [**GetUserInfo**](docs/UserApi.md#getuserinfo) | **Get** /user/info | Get User Info
*UserApi* | [**ListUserCompanies**](docs/UserApi.md#listusercompanies) | **Get** /user/companies | List User Companies


## Documentation For Models

 - [ArchiveDocument](docs/ArchiveDocument.md)
 - [AttachmentData](docs/AttachmentData.md)
 - [CashbookEntry](docs/CashbookEntry.md)
 - [CashbookEntryDocument](docs/CashbookEntryDocument.md)
 - [CashbookEntryKind](docs/CashbookEntryKind.md)
 - [CashbookEntryType](docs/CashbookEntryType.md)
 - [City](docs/City.md)
 - [Client](docs/Client.md)
 - [ClientType](docs/ClientType.md)
 - [Company](docs/Company.md)
 - [CompanyInfo](docs/CompanyInfo.md)
 - [CompanyInfoAccessInfo](docs/CompanyInfoAccessInfo.md)
 - [CompanyInfoPlanInfo](docs/CompanyInfoPlanInfo.md)
 - [CompanyInfoPlanInfoFunctions](docs/CompanyInfoPlanInfoFunctions.md)
 - [CompanyInfoPlanInfoFunctionsStatus](docs/CompanyInfoPlanInfoFunctionsStatus.md)
 - [CompanyInfoPlanInfoLimits](docs/CompanyInfoPlanInfoLimits.md)
 - [CompanyType](docs/CompanyType.md)
 - [ControlledCompany](docs/ControlledCompany.md)
 - [CreateArchiveDocumentRequest](docs/CreateArchiveDocumentRequest.md)
 - [CreateArchiveDocumentResponse](docs/CreateArchiveDocumentResponse.md)
 - [CreateCashbookEntryRequest](docs/CreateCashbookEntryRequest.md)
 - [CreateCashbookEntryResponse](docs/CreateCashbookEntryResponse.md)
 - [CreateClientRequest](docs/CreateClientRequest.md)
 - [CreateClientResponse](docs/CreateClientResponse.md)
 - [CreateF24Request](docs/CreateF24Request.md)
 - [CreateF24Response](docs/CreateF24Response.md)
 - [CreateIssuedDocumentRequest](docs/CreateIssuedDocumentRequest.md)
 - [CreateIssuedDocumentResponse](docs/CreateIssuedDocumentResponse.md)
 - [CreatePaymentAccountRequest](docs/CreatePaymentAccountRequest.md)
 - [CreatePaymentAccountResponse](docs/CreatePaymentAccountResponse.md)
 - [CreatePaymentMethodRequest](docs/CreatePaymentMethodRequest.md)
 - [CreatePaymentMethodResponse](docs/CreatePaymentMethodResponse.md)
 - [CreateProductRequest](docs/CreateProductRequest.md)
 - [CreateProductResponse](docs/CreateProductResponse.md)
 - [CreateReceiptRequest](docs/CreateReceiptRequest.md)
 - [CreateReceiptResponse](docs/CreateReceiptResponse.md)
 - [CreateReceivedDocumentRequest](docs/CreateReceivedDocumentRequest.md)
 - [CreateReceivedDocumentResponse](docs/CreateReceivedDocumentResponse.md)
 - [CreateSupplierRequest](docs/CreateSupplierRequest.md)
 - [CreateSupplierResponse](docs/CreateSupplierResponse.md)
 - [CreateVatTypeRequest](docs/CreateVatTypeRequest.md)
 - [CreateVatTypeResponse](docs/CreateVatTypeResponse.md)
 - [Currency](docs/Currency.md)
 - [DetailedCountry](docs/DetailedCountry.md)
 - [DocumentTemplate](docs/DocumentTemplate.md)
 - [EInvoiceRejectionReason](docs/EInvoiceRejectionReason.md)
 - [Email](docs/Email.md)
 - [EmailAttachment](docs/EmailAttachment.md)
 - [EmailData](docs/EmailData.md)
 - [EmailDataDefaultSenderEmail](docs/EmailDataDefaultSenderEmail.md)
 - [EmailRecipientStatus](docs/EmailRecipientStatus.md)
 - [EmailSchedule](docs/EmailSchedule.md)
 - [EmailScheduleInclude](docs/EmailScheduleInclude.md)
 - [EmailStatus](docs/EmailStatus.md)
 - [Entity](docs/Entity.md)
 - [EntityType](docs/EntityType.md)
 - [F24](docs/F24.md)
 - [F24Status](docs/F24Status.md)
 - [FunctionStatus](docs/FunctionStatus.md)
 - [GetArchiveDocumentResponse](docs/GetArchiveDocumentResponse.md)
 - [GetCashbookEntryResponse](docs/GetCashbookEntryResponse.md)
 - [GetClientResponse](docs/GetClientResponse.md)
 - [GetCompanyInfoResponse](docs/GetCompanyInfoResponse.md)
 - [GetEInvoiceRejectionReasonResponse](docs/GetEInvoiceRejectionReasonResponse.md)
 - [GetEmailDataResponse](docs/GetEmailDataResponse.md)
 - [GetExistingIssuedDocumentTotalsRequest](docs/GetExistingIssuedDocumentTotalsRequest.md)
 - [GetExistingIssuedDocumentTotalsResponse](docs/GetExistingIssuedDocumentTotalsResponse.md)
 - [GetExistingReceivedDocumentTotalsRequest](docs/GetExistingReceivedDocumentTotalsRequest.md)
 - [GetExistingReceivedDocumentTotalsResponse](docs/GetExistingReceivedDocumentTotalsResponse.md)
 - [GetF24Response](docs/GetF24Response.md)
 - [GetIssuedDocumentPreCreateInfoResponse](docs/GetIssuedDocumentPreCreateInfoResponse.md)
 - [GetIssuedDocumentResponse](docs/GetIssuedDocumentResponse.md)
 - [GetNewIssuedDocumentTotalsRequest](docs/GetNewIssuedDocumentTotalsRequest.md)
 - [GetNewIssuedDocumentTotalsResponse](docs/GetNewIssuedDocumentTotalsResponse.md)
 - [GetNewReceivedDocumentTotalsRequest](docs/GetNewReceivedDocumentTotalsRequest.md)
 - [GetNewReceivedDocumentTotalsResponse](docs/GetNewReceivedDocumentTotalsResponse.md)
 - [GetPaymentAccountResponse](docs/GetPaymentAccountResponse.md)
 - [GetPaymentMethodResponse](docs/GetPaymentMethodResponse.md)
 - [GetProductResponse](docs/GetProductResponse.md)
 - [GetReceiptPreCreateInfoResponse](docs/GetReceiptPreCreateInfoResponse.md)
 - [GetReceiptResponse](docs/GetReceiptResponse.md)
 - [GetReceiptsMonthlyTotalsResponse](docs/GetReceiptsMonthlyTotalsResponse.md)
 - [GetReceivedDocumentPreCreateInfoResponse](docs/GetReceivedDocumentPreCreateInfoResponse.md)
 - [GetReceivedDocumentResponse](docs/GetReceivedDocumentResponse.md)
 - [GetSupplierResponse](docs/GetSupplierResponse.md)
 - [GetUserInfoResponse](docs/GetUserInfoResponse.md)
 - [GetUserInfoResponseEmailConfirmationState](docs/GetUserInfoResponseEmailConfirmationState.md)
 - [GetUserInfoResponseInfo](docs/GetUserInfoResponseInfo.md)
 - [GetVatTypeResponse](docs/GetVatTypeResponse.md)
 - [IssuedDocument](docs/IssuedDocument.md)
 - [IssuedDocumentEiData](docs/IssuedDocumentEiData.md)
 - [IssuedDocumentExtraData](docs/IssuedDocumentExtraData.md)
 - [IssuedDocumentItemsListItem](docs/IssuedDocumentItemsListItem.md)
 - [IssuedDocumentOptions](docs/IssuedDocumentOptions.md)
 - [IssuedDocumentPaymentsListItem](docs/IssuedDocumentPaymentsListItem.md)
 - [IssuedDocumentPaymentsListItemPaymentTerms](docs/IssuedDocumentPaymentsListItemPaymentTerms.md)
 - [IssuedDocumentPreCreateInfo](docs/IssuedDocumentPreCreateInfo.md)
 - [IssuedDocumentPreCreateInfoDefaultValues](docs/IssuedDocumentPreCreateInfoDefaultValues.md)
 - [IssuedDocumentPreCreateInfoExtraDataDefaultValues](docs/IssuedDocumentPreCreateInfoExtraDataDefaultValues.md)
 - [IssuedDocumentPreCreateInfoItemsDefaultValues](docs/IssuedDocumentPreCreateInfoItemsDefaultValues.md)
 - [IssuedDocumentStatus](docs/IssuedDocumentStatus.md)
 - [IssuedDocumentTotals](docs/IssuedDocumentTotals.md)
 - [IssuedDocumentType](docs/IssuedDocumentType.md)
 - [JoinIssuedDocumentsResponse](docs/JoinIssuedDocumentsResponse.md)
 - [Language](docs/Language.md)
 - [ListArchiveCategoriesResponse](docs/ListArchiveCategoriesResponse.md)
 - [ListArchiveDocumentsResponse](docs/ListArchiveDocumentsResponse.md)
 - [ListArchiveDocumentsResponsePage](docs/ListArchiveDocumentsResponsePage.md)
 - [ListCashbookEntriesResponse](docs/ListCashbookEntriesResponse.md)
 - [ListCitiesResponse](docs/ListCitiesResponse.md)
 - [ListClientsResponse](docs/ListClientsResponse.md)
 - [ListClientsResponsePage](docs/ListClientsResponsePage.md)
 - [ListCostCentersResponse](docs/ListCostCentersResponse.md)
 - [ListCountriesResponse](docs/ListCountriesResponse.md)
 - [ListCurrenciesResponse](docs/ListCurrenciesResponse.md)
 - [ListDeliveryNotesDefaultCausalsResponse](docs/ListDeliveryNotesDefaultCausalsResponse.md)
 - [ListDetailedCountriesResponse](docs/ListDetailedCountriesResponse.md)
 - [ListEmailsResponse](docs/ListEmailsResponse.md)
 - [ListEmailsResponsePage](docs/ListEmailsResponsePage.md)
 - [ListF24Response](docs/ListF24Response.md)
 - [ListF24ResponseAggregatedData](docs/ListF24ResponseAggregatedData.md)
 - [ListF24ResponseAggregation](docs/ListF24ResponseAggregation.md)
 - [ListF24ResponsePage](docs/ListF24ResponsePage.md)
 - [ListIssuedDocumentsResponse](docs/ListIssuedDocumentsResponse.md)
 - [ListIssuedDocumentsResponsePage](docs/ListIssuedDocumentsResponsePage.md)
 - [ListLanguagesResponse](docs/ListLanguagesResponse.md)
 - [ListPaymentAccountsResponse](docs/ListPaymentAccountsResponse.md)
 - [ListPaymentMethodsResponse](docs/ListPaymentMethodsResponse.md)
 - [ListProductCategoriesResponse](docs/ListProductCategoriesResponse.md)
 - [ListProductsResponse](docs/ListProductsResponse.md)
 - [ListProductsResponsePage](docs/ListProductsResponsePage.md)
 - [ListReceiptsResponse](docs/ListReceiptsResponse.md)
 - [ListReceiptsResponsePage](docs/ListReceiptsResponsePage.md)
 - [ListReceivedDocumentCategoriesResponse](docs/ListReceivedDocumentCategoriesResponse.md)
 - [ListReceivedDocumentsResponse](docs/ListReceivedDocumentsResponse.md)
 - [ListReceivedDocumentsResponsePage](docs/ListReceivedDocumentsResponsePage.md)
 - [ListRevenueCentersResponse](docs/ListRevenueCentersResponse.md)
 - [ListSuppliersResponse](docs/ListSuppliersResponse.md)
 - [ListSuppliersResponsePage](docs/ListSuppliersResponsePage.md)
 - [ListTemplatesResponse](docs/ListTemplatesResponse.md)
 - [ListUnitsOfMeasureResponse](docs/ListUnitsOfMeasureResponse.md)
 - [ListUserCompaniesResponse](docs/ListUserCompaniesResponse.md)
 - [ListUserCompaniesResponseData](docs/ListUserCompaniesResponseData.md)
 - [ListVatTypesResponse](docs/ListVatTypesResponse.md)
 - [ModifyArchiveDocumentRequest](docs/ModifyArchiveDocumentRequest.md)
 - [ModifyArchiveDocumentResponse](docs/ModifyArchiveDocumentResponse.md)
 - [ModifyCashbookEntryRequest](docs/ModifyCashbookEntryRequest.md)
 - [ModifyCashbookEntryResponse](docs/ModifyCashbookEntryResponse.md)
 - [ModifyClientRequest](docs/ModifyClientRequest.md)
 - [ModifyClientResponse](docs/ModifyClientResponse.md)
 - [ModifyF24Request](docs/ModifyF24Request.md)
 - [ModifyF24Response](docs/ModifyF24Response.md)
 - [ModifyIssuedDocumentRequest](docs/ModifyIssuedDocumentRequest.md)
 - [ModifyIssuedDocumentResponse](docs/ModifyIssuedDocumentResponse.md)
 - [ModifyPaymentAccountRequest](docs/ModifyPaymentAccountRequest.md)
 - [ModifyPaymentAccountResponse](docs/ModifyPaymentAccountResponse.md)
 - [ModifyPaymentMethodRequest](docs/ModifyPaymentMethodRequest.md)
 - [ModifyPaymentMethodResponse](docs/ModifyPaymentMethodResponse.md)
 - [ModifyProductRequest](docs/ModifyProductRequest.md)
 - [ModifyProductResponse](docs/ModifyProductResponse.md)
 - [ModifyReceiptRequest](docs/ModifyReceiptRequest.md)
 - [ModifyReceiptResponse](docs/ModifyReceiptResponse.md)
 - [ModifyReceivedDocumentRequest](docs/ModifyReceivedDocumentRequest.md)
 - [ModifyReceivedDocumentResponse](docs/ModifyReceivedDocumentResponse.md)
 - [ModifySupplierRequest](docs/ModifySupplierRequest.md)
 - [ModifySupplierResponse](docs/ModifySupplierResponse.md)
 - [ModifyVatTypeRequest](docs/ModifyVatTypeRequest.md)
 - [ModifyVatTypeResponse](docs/ModifyVatTypeResponse.md)
 - [MonthlyTotal](docs/MonthlyTotal.md)
 - [OriginalDocumentType](docs/OriginalDocumentType.md)
 - [Pagination](docs/Pagination.md)
 - [PaymentAccount](docs/PaymentAccount.md)
 - [PaymentAccountType](docs/PaymentAccountType.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentMethodDetails](docs/PaymentMethodDetails.md)
 - [PaymentMethodType](docs/PaymentMethodType.md)
 - [PaymentTermsType](docs/PaymentTermsType.md)
 - [PermissionLevel](docs/PermissionLevel.md)
 - [Permissions](docs/Permissions.md)
 - [PermissionsFicIssuedDocumentsDetailed](docs/PermissionsFicIssuedDocumentsDetailed.md)
 - [Product](docs/Product.md)
 - [Receipt](docs/Receipt.md)
 - [ReceiptItemsListItem](docs/ReceiptItemsListItem.md)
 - [ReceiptPreCreateInfo](docs/ReceiptPreCreateInfo.md)
 - [ReceiptType](docs/ReceiptType.md)
 - [ReceivedDocument](docs/ReceivedDocument.md)
 - [ReceivedDocumentEntity](docs/ReceivedDocumentEntity.md)
 - [ReceivedDocumentInfo](docs/ReceivedDocumentInfo.md)
 - [ReceivedDocumentInfoDefaultValues](docs/ReceivedDocumentInfoDefaultValues.md)
 - [ReceivedDocumentInfoItemsDefaultValues](docs/ReceivedDocumentInfoItemsDefaultValues.md)
 - [ReceivedDocumentItemsListItem](docs/ReceivedDocumentItemsListItem.md)
 - [ReceivedDocumentPaymentsListItem](docs/ReceivedDocumentPaymentsListItem.md)
 - [ReceivedDocumentPaymentsListItemPaymentTerms](docs/ReceivedDocumentPaymentsListItemPaymentTerms.md)
 - [ReceivedDocumentTotals](docs/ReceivedDocumentTotals.md)
 - [ReceivedDocumentType](docs/ReceivedDocumentType.md)
 - [ScheduleEmailRequest](docs/ScheduleEmailRequest.md)
 - [SendEInvoiceRequest](docs/SendEInvoiceRequest.md)
 - [SendEInvoiceRequestData](docs/SendEInvoiceRequestData.md)
 - [SendEInvoiceResponse](docs/SendEInvoiceResponse.md)
 - [SendEInvoiceResponseData](docs/SendEInvoiceResponseData.md)
 - [SenderEmail](docs/SenderEmail.md)
 - [ShowTotalsMode](docs/ShowTotalsMode.md)
 - [Supplier](docs/Supplier.md)
 - [SupplierType](docs/SupplierType.md)
 - [TransformIssuedDocumentResponse](docs/TransformIssuedDocumentResponse.md)
 - [UploadArchiveAttachmentResponse](docs/UploadArchiveAttachmentResponse.md)
 - [UploadF24AttachmentResponse](docs/UploadF24AttachmentResponse.md)
 - [UploadIssuedDocumentAttachmentResponse](docs/UploadIssuedDocumentAttachmentResponse.md)
 - [UploadReceivedDocumentAttachmentResponse](docs/UploadReceivedDocumentAttachmentResponse.md)
 - [User](docs/User.md)
 - [UserCompanyRole](docs/UserCompanyRole.md)
 - [VatItem](docs/VatItem.md)
 - [VatKind](docs/VatKind.md)
 - [VatType](docs/VatType.md)
 - [VerifyEInvoiceXmlErrorResponse](docs/VerifyEInvoiceXmlErrorResponse.md)
 - [VerifyEInvoiceXmlErrorResponseError](docs/VerifyEInvoiceXmlErrorResponseError.md)
 - [VerifyEInvoiceXmlErrorResponseErrorValidationResult](docs/VerifyEInvoiceXmlErrorResponseErrorValidationResult.md)
 - [VerifyEInvoiceXmlErrorResponseExtra](docs/VerifyEInvoiceXmlErrorResponseExtra.md)
 - [VerifyEInvoiceXmlResponse](docs/VerifyEInvoiceXmlResponse.md)
 - [VerifyEInvoiceXmlResponseData](docs/VerifyEInvoiceXmlResponseData.md)


## Documentation For Authorization



### OAuth2AuthenticationCodeFlow


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://api-v2.fattureincloud.it/oauth/authorize
- **Scopes**: 
 - **entity.clients:r**: Read permission to the Clients registry
 - **entity.clients:a**: Write permission to the Clients registry
 - **entity.suppliers:r**: Read permission to the Suppliers registry
 - **entity.suppliers:a**: Write permission to the Suppliers registry
 - **products:r**: Read permission to the Products
 - **products:a**: Write permission to the Products
 - **issued_documents.invoices:r**: Read permission to the issued Invoices
 - **issued_documents.credit_notes:r**: Read permission to the issued Credit Notes
 - **issued_documents.receipts:r**: Read permission to the issued Receipts
 - **issued_documents.orders:r**: Read permission to the issued Orders
 - **issued_documents.quotes:r**: Read permission to the issued Quotes
 - **issued_documents.proformas:r**: Read permission to the issued Proformas
 - **issued_documents.delivery_notes:r**: Read permission to the issued Delivery Notes
 - **issued_documents.work_reports:r**: Read permission to the issued Work Reports
 - **issued_documents.supplier_orders:r**: Read permission to the issued Supplier Orders
 - **issued_documents.self_invoices:r**: Read permission to the issued Self Invoices
 - **issued_documents.invoices:a**: Write permission to the issued Invoices
 - **issued_documents.credit_notes:a**: Write permission to the issued Credit Notes
 - **issued_documents.receipts:a**: Write permission to the issued issued Receipts
 - **issued_documents.orders:a**: Write permission to the issued Orders
 - **issued_documents.quotes:a**: Write permission to the issued Quotes
 - **issued_documents.proformas:a**: Write permission to the issued Proformas
 - **issued_documents.delivery_notes:a**: Write permission to the issued Delivery Notes
 - **issued_documents.work_reports:a**: Write permission to the issued Work Reports
 - **issued_documents.supplier_orders:a**: Write permission to the issued Supplier Orders
 - **issued_documents.self_invoices:a**: Write permission to the issued Self Invoices
 - **received_documents:r**: Read permission to the Received Documents
 - **received_documents:a**: Write permission to the Received Documents
 - **stock:r**: Read permission to the Stock movements
 - **stock:a**: Write permission to the Stock movements
 - **receipts:r**: Read permission to the Receipts
 - **receipts:a**: Write permission to the Receipts
 - **taxes:r**: Read permission to the Taxes
 - **taxes:a**: Write permission to the Taxes
 - **archive:r**: Read permission to the Archive Documents
 - **archive:a**: Read permission to the Archive Documents
 - **cashbook:r**: Read permission to the Cashbook
 - **cashbook:a**: Write permission to the Cashbook
 - **settings:r**: Read permission to the Settings
 - **settings:a**: Write permission to the Settings
 - **situation:r**: Read permission to the company Situation

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@fattureincloud.it

