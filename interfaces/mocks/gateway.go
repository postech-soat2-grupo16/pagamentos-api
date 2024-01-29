package mocks

import (
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/stretchr/testify/mock"
)

type MockPagamentoGateway struct {
	mock.Mock
}

type MockPedidoGateway struct {
	mock.Mock
}

type MockQueueGateway struct {
	mock.Mock
}

func (m *MockPagamentoGateway) UpdatePaymentStatusByPaymentID(pagamentoID uint32, status string) (*entities.Pagamento, error) {
	args := m.Called(pagamentoID, status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPagamentoGateway) CreatePayment(pagamento entities.Pagamento) (*entities.Pagamento, error) {
	args := m.Called(pagamento)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPagamentoGateway) GetByID(paymentID uint32) (*entities.Pagamento, error) {
	args := m.Called(paymentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPedidoGateway) GetByID(pedidoID string) (*pagamento.Pedido, error) {
	args := m.Called(pedidoID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pagamento.Pedido), args.Error(1)
}

func (m *MockQueueGateway) SendMessage(pagamento *entities.Pagamento) error {
	args := m.Called(pagamento)
	return args.Error(0)
}
