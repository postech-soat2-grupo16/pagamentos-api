package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type PedidoGatewayI interface {
	GetByID(pedidoID string) (*pagamento.Pedido, error)
}

type MercadoPagoGatewayI interface {
	GetPedidoIDByPaymentID(paymentID string) (uint32, error)
	CreateQRCodeForPedido(pedido pagamento.Pedido) (string, error)
}

type PagamentoGatewayI interface {
	UpdatePaymentStatusByPaymentID(pagamentoID uint32, status string) (*entities.Pagamento, error)
	CreatePayment(pagamento entities.Pagamento) (*entities.Pagamento, error)
	GetByID(paymentID uint32) (*entities.Pagamento, error)
}

type QueueGatewayI interface {
	SendMessage(pagamento *entities.Pagamento) error
}

type NotificationGatewayI interface {
	SendNotification(pagamento *entities.Pagamento, email string) error
}
