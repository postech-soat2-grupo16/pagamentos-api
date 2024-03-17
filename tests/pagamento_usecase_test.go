package tests

import (
	"errors"
	pagamento2 "github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces/mocks"
	"github.com/joaocampari/postech-soat2-grupo16/usecases/pagamento"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateQRCode_Fail(t *testing.T) {
	mockPedidoGateway := new(mocks.MockPedidoGateway)
	mockPedidoGateway.On("GetByID", mock.Anything).Return(nil, errors.New("request error"))

	useCase := pagamento.NewUseCase(nil, nil, nil, nil,
		mockPedidoGateway, nil)

	_, err := useCase.CreateQRCode("abc-ab31323")
	assert.Error(t, err)
}

func TestUpdatePaymentStatusByPaymentID_DBFail(t *testing.T) {
	mockPagamentoGateway := new(mocks.MockPagamentoGateway)
	mockPagamentoGateway.On("UpdatePaymentStatusByPaymentID", mock.Anything, mock.Anything).
		Return(nil, errors.New("db error"))

	useCase := pagamento.NewUseCase(mockPagamentoGateway, nil, nil, nil,
		nil, nil)

	_, err := useCase.UpdatePaymentStatusByPaymentID(90)
	assert.Error(t, err)
}

//func TestUpdatePaymentStatusByPaymentID_QueueFail(t *testing.T) {
//	payment := &entities.Pagamento{
//		ID:        93,
//		PedidoID:  "abc0432-445vdsfv",
//		Amount:    100,
//		Status:    "APROVADO",
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//
//	mockPagamentoGateway := new(mocks.MockPagamentoGateway)
//	mockQueueGateway := new(mocks.MockQueueGateway)
//
//	mockPagamentoGateway.On("UpdatePaymentStatusByPaymentID", mock.Anything, mock.Anything).
//		Return(payment, nil)
//	mockQueueGateway.On("SendMessage", mock.Anything).
//		Return(errors.New("queue error"))
//
//	useCase := pagamento.NewUseCase(mockPagamentoGateway, nil, mockQueueGateway, nil,
//		nil, nil)
//
//	_, err := useCase.UpdatePaymentStatusByPaymentID(94)
//	assert.Error(t, err)
//}

func TestCreatePayment_RequestFail(t *testing.T) {
	mockPedidoGateway := new(mocks.MockPedidoGateway)
	mockPedidoGateway.On("GetByID", mock.Anything).Return(nil, errors.New("request error"))

	useCase := pagamento.NewUseCase(nil, nil, nil, nil,
		mockPedidoGateway, nil)

	_, err := useCase.CreatePayment("abc-ab31323")
	assert.Error(t, err)
}

func TestCreatePayment_DBFail(t *testing.T) {
	var items []pagamento2.Item
	item := pagamento2.Item{
		ItemID:      "2",
		Price:       10,
		Quantity:    2,
		Name:        "abc",
		Category:    "aa",
		Description: "aaa",
	}
	items = append(items, item)

	pedido := &pagamento2.Pedido{
		ID:        "ac-432452",
		ClientID:  "11234",
		Status:    "CRIADO",
		Items:     items,
		Notes:     "abc",
		CreatedAt: "",
		UpdatedAt: "",
	}

	mockPedidoGateway := new(mocks.MockPedidoGateway)
	mockPedidoGateway.On("GetByID", mock.Anything).Return(pedido, nil)
	mockPagamentoGateway := new(mocks.MockPagamentoGateway)
	mockPagamentoGateway.On("CreatePayment", mock.Anything).Return(nil, errors.New("db error"))

	useCase := pagamento.NewUseCase(mockPagamentoGateway, nil, nil, nil,
		mockPedidoGateway, nil)

	_, err := useCase.CreatePayment("abc-ab31323")
	assert.Error(t, err)
}
