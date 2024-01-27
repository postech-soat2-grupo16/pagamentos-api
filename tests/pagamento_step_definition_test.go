package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"io"
	"net/http"
	"time"
)

func (i *Input) getURL(endpoint string) string {
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}

func (i *Input) sendRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, i.getURL(endpoint), body)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (i *Input) theAPIShouldRespondWith(statusCode int) error {
	if i.statusCode != statusCode {
		return fmt.Errorf("Expected status code %d, but got %d", statusCode, i.statusCode)
	}
	return nil
}

func (i *Input) theAPIShouldRespondWithAQRCode() error {
	if i.statusCode != http.StatusOK {
		return fmt.Errorf("Expected status code %d, but got %d", http.StatusOK, i.statusCode)
	}

	var qrCode pagamento.QRCode
	err := json.NewDecoder(i.body).Decode(&qrCode)
	if err != nil {
		return fmt.Errorf("Error decoding QR Code response: %v", err)
	}

	return nil
}

func aMercadoPagoPaymentCallback() error {
	callbackData := pagamento.PaymentCallback{
		Id:          123,
		LiveMode:    false,
		Type:        "",
		DateCreated: time.Time{},
		UserId:      0,
		ApiVersion:  "",
		Action:      "",
	}

	callbackBody, err := json.Marshal(callbackData)
	if err != nil {
		return fmt.Errorf("Error encoding callback data: %v", err)
	}

	response, err := inputs.sendRequest("POST", "/pagamentos/mp-webhook", bytes.NewReader(callbackBody))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	inputs.statusCode = response.StatusCode
	inputs.body = response.Body

	return inputs.theAPIShouldRespondWith(http.StatusOK)
}

func aPaymentID() error {
	inputs.paymentID = "1"
	return nil
}

func theAPIShouldRespondWith(arg1 string) error {
	return godog.ErrPending
}

func theAPIShouldRespondWithThePaymentDetails() error {
	return godog.ErrPending
}

func (i *Input) theAPIShouldRespondWithThePaymentDetails() error {
	if i.statusCode != http.StatusOK {
		return fmt.Errorf("Expected status code %d, but got %d", http.StatusOK, i.statusCode)
	}

	var payment pagamento.Pagamento
	err := json.NewDecoder(i.body).Decode(&payment)
	if err != nil {
		return fmt.Errorf("Error decoding Payment response: %v", err)
	}

	return nil
}

func theHealthEndpointIsAccessed() error {
	response, err := inputs.sendRequest(http.MethodGet, "/pagamentos/health", nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	inputs.statusCode = response.StatusCode
	inputs.body = response.Body

	return inputs.theAPIShouldRespondWith(http.StatusOK)
}

func thePaymentStatusShouldBeUpdated() error {
	return godog.ErrPending
}

func theUserRequestsThePaymentByID() error {
	response, err := inputs.sendRequest(http.MethodGet, fmt.Sprintf("/pagamentos/%s", inputs.paymentID), nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	inputs.statusCode = response.StatusCode
	inputs.body = response.Body

	return inputs.theAPIShouldRespondWithThePaymentDetails()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a MercadoPago payment callback$`, aMercadoPagoPaymentCallback)
	ctx.Step(`^a payment ID$`, aPaymentID)
	ctx.Step(`^the API should respond with "([^"]*)"$`, theAPIShouldRespondWith)
	ctx.Step(`^the API should respond with the payment details$`, theAPIShouldRespondWithThePaymentDetails)
	ctx.Step(`^the health endpoint is accessed$`, theHealthEndpointIsAccessed)
	ctx.Step(`^the payment status should be updated$`, thePaymentStatusShouldBeUpdated)
	ctx.Step(`^the user requests the payment by ID$`, theUserRequestsThePaymentByID)
}

var inputs Input

type Input struct {
	pedidoID   string
	paymentID  string
	statusCode int
	status     string
	body       io.ReadCloser
}
