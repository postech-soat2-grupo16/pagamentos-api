package tests

import (
	"errors"
	"github.com/joaocampari/postech-soat2-grupo16/controllers"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var (
	errUsecaseFailure  = errors.New("ErrUsecaseFailed")
	errUsecaseNotFound = errors.New("ErrUsecaseNotFound")
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

func TestPagamentoController_GetQRCodeByPedidoID_Error_500(t *testing.T) {
	useCase := new(mocks.MockPagamentosUseCase)
	useCase.On("CreatePayment").Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/pagamentos/1/qr-code", nil)

	c := chi.NewRouter()
	controllers.NewPagamentoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}
