/*
Fatture in Cloud API v2 - API Reference

Connect your software with Fatture in Cloud, the invoicing platform chosen by more than 400.000 businesses in Italy.   The Fatture in Cloud API is based on REST, and makes possible to interact with the user related data prior authorization via OAuth2 protocol.

API version: 2.0.19
Contact: info@fattureincloud.it
*/

package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	fattureincloud "github.com/fattureincloud/fattureincloud-go-sdk/v2/api"
	. "github.com/fattureincloud/fattureincloud-go-sdk/v2/model"
	"github.com/stretchr/testify/assert"
)

func TestGetCompanyInfo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"id":12346,"name":"Studio Commercialista","email":"mario.rossi@example.com","type":"accountant","access_info":{"role":"master","through_accountant":false},"plan_info":{"limits":{"clients":5000,"suppliers":5000,"products":5000,"documents":3000},"functions":{"archive":true,"cerved":true,"document_attachments":true,"e_invoice":true,"genius":true,"mail_tracking":true,"payment_notifications":true,"paypal":true,"receipts":true,"recurring":true,"smtp":true,"sofort":false,"stock":true,"subaccounts":true,"tessera_sanitaria":true,"ts_digital":true,"ts_invoice_trading":true,"ts_pay":true},"functions_status":{"ts_digital":{"active":true},"ts_pay":{"active":false}}},"accountant_id":12345,"is_accountant":true}}`))
	}))
	defer s.Close()

	serverURL, _ := url.Parse(s.URL)
	configuration := fattureincloud.NewConfiguration()
	configuration.HTTPClient = s.Client()
	configuration.Host = serverURL.Host
	configuration.Scheme = "http"
	apiClient := fattureincloud.NewAPIClient(configuration)

	actual, _, err := apiClient.CompaniesApi.GetCompanyInfo(context.Background(), 2).Execute()
	assert.NoError(t, err, "errore in chiamata api")

	expected := NewCompanyInfo().
		SetId(12346).
		SetName("Studio Commercialista").
		SetEmail("mario.rossi@example.com").
		SetType(CompanyTypes.ACCOUNTANT).
		SetIsAccountant(true).
		SetAccountantId(12345).
		SetAccessInfo(*NewCompanyInfoAccessInfo().
			SetRole(UserCompanyRoles.MASTER).
			SetThroughAccountant(false)).
		SetPlanInfo(*NewCompanyInfoPlanInfo().
			SetLimits(*NewCompanyInfoPlanInfoLimits().
				SetClients(5000).
				SetSuppliers(5000).
				SetProducts(5000).
				SetDocuments(3000)).
			SetFunctions(*NewCompanyInfoPlanInfoFunctions().
				SetDocumentAttachments(true).
				SetArchive(true).
				SetPaymentNotifications(true).
				SetPaypal(true).
				SetReceipts(true).
				SetEInvoice(true).
				SetGenius(true).
				SetStock(true).
				SetSmtp(true).
				SetMailTracking(true).
				SetSubaccounts(true).
				SetTesseraSanitaria(true).
				SetRecurring(true).
				SetSofort(false).
				SetCerved(true).
				SetTsDigital(true).
				SetTsPay(true).
				SetTsInvoiceTrading(true)).
			SetFunctionsStatus(*NewCompanyInfoPlanInfoFunctionsStatus().
				SetTsDigital(*NewFunctionStatus().
					SetActive(true)).
				SetTsPay(*NewFunctionStatus().
					SetActive(false))))

	assert.True(t, reflect.DeepEqual(expected, actual.Data))
}
