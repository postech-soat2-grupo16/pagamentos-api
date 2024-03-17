package tests

import (
	"errors"
	"github.com/joaocampari/postech-soat2-grupo16/controllers"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces/mocks"
	"github.com/joaocampari/postech-soat2-grupo16/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var (
	errUsecaseFailure = errors.New("ErrUsecaseFailed")
)

func TestPagamentoController_GetQRCodeByPedidoID_Error_400(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos//qr-code", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Bad request response is expected")
}

func TestPagamentoController_GetQRCodeByPedidoID_Error_500_by_CreatePayment(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("CreatePayment").Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/65619d06-f3fb-4726-b9fa-be597efa0417/qr-code", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPagamentoController_GetQRCodeByPedidoID_Error_500_by_CreateQRCode(t *testing.T) {
	payment := entities.Pagamento{
		ID:        1321,
		PedidoID:  "abc",
		Amount:    100,
		Status:    entities.StatusPagamentoIniciado,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("CreatePayment").Return(&payment, nil)
	useCase.On("CreateQRCode").Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/65619d06-f3fb-4726-b9fa-be597efa0417/qr-code", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPagamentoController_GetQRCodeByPedidoID_Error_404_by_CreateQRCode(t *testing.T) {
	payment := entities.Pagamento{
		ID:        1321,
		PedidoID:  "abc",
		Amount:    100,
		Status:    entities.StatusPagamentoIniciado,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("CreatePayment").Return(&payment, nil)
	useCase.On("CreateQRCode").Return(nil, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/65619d06-f3fb-4726-b9fa-be597efa0417/qr-code", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code, "Not found response is expected")
}

func TestPagamentoController_PaymentWebhookCreate_Error_400(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	badJSON := `{"invalid json`
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/pagamentos/mp-webhook", strings.NewReader(badJSON))

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Bad request response is expected")
}

func TestPagamentoController_PaymentWebhookCreate_Error_422(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	useCase.On("UpdatePaymentStatusByPaymentID").Return(nil, util.NewErrorDomain("domain"))

	body := `{"id":1,"live_mode":true,"type":"payment","date_created":"2015-03-25T10:04:58.396-04:00","user_id":44444,"api_version":"v1","action":"payment.created","data":{"id":"1"}}`
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/pagamentos/mp-webhook", strings.NewReader(body))

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code, "Unprocessable Entity response is expected")
}

func TestPagamentoController_PaymentWebhookCreate_Error_500(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	useCase.On("UpdatePaymentStatusByPaymentID").Return(nil, errUsecaseFailure)

	body := `{"id":1,"live_mode":true,"type":"payment","date_created":"2015-03-25T10:04:58.396-04:00","user_id":44444,"api_version":"v1","action":"payment.created","data":{"id":"1"}}`
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/pagamentos/mp-webhook", strings.NewReader(body))

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPagamentoController_PaymentWebhookCreate_Error_404(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	useCase.On("UpdatePaymentStatusByPaymentID").Return(nil, nil)

	body := `{"id":1,"live_mode":true,"type":"payment","date_created":"2015-03-25T10:04:58.396-04:00","user_id":44444,"api_version":"v1","action":"payment.created","data":{"id":"1"}}`
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/pagamentos/mp-webhook", strings.NewReader(body))

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code, "Not found response is expected")
}

func TestPagamentoController_GetPaymentById_Error_400(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/erroid", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Bad request response is expected")
}

func TestPagamentoController_GetPaymentById_Error_500(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("GetByID").Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/1", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestPagamentoController_GetPaymentById_Error_404(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("GetByID").Return(nil, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/1", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code, "Not Found response is expected")
}
