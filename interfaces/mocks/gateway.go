package mocks

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/stretchr/testify/mock"
)

type MockPagamentoGateway struct {
	mock.Mock
}

func (m *MockPagamentoGateway) Update(pagamento entities.Pagamento) (*entities.Pagamento, error) {
	args := m.Called(pagamento)
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
