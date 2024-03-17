package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type PagamentoUseCase interface {
	CreateQRCode(pedidoID string) (*string, error)
	UpdatePaymentStatusByPaymentID(pagamentoID uint32) (*entities.Pagamento, error)
	CreatePayment(pedidoID string) (*entities.Pagamento, error)
	GetByID(paymentID uint32) (*entities.Pagamento, error)
	ProcessPaymentStatus(pagamentoID uint32, statusPagamento string) (*entities.Pagamento, error)
}
