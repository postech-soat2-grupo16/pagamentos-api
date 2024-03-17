package mocks

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/stretchr/testify/mock"
)

type MockPagamentosUseCase struct {
	mock.Mock
}

func (m *MockPagamentosUseCase) CreateQRCode(pedidoID string) (*string, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockPagamentosUseCase) UpdatePaymentStatusByPaymentID(pagamentoID uint32) (*entities.Pagamento, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPagamentosUseCase) CreatePayment(pedidoID string) (*entities.Pagamento, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPagamentosUseCase) GetByID(paymentID uint32) (*entities.Pagamento, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}

func (m *MockPagamentosUseCase) ProcessPaymentStatus(pagamentoID uint32, statusPagamento string) (*entities.Pagamento, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Pagamento), args.Error(1)
}
