package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type PagamentoUseCase interface {
	CreateQRCode(pedidoID uint32) (*string, error)
	UpdatePaymentStatusByPaymentID(pagamentoID uint32) (*entities.Pagamento, error)
	CreatePayment(pedidoID uint32) (*entities.Pagamento, error)
	GetByID(paymentID uint32) (*entities.Pagamento, error)
}
