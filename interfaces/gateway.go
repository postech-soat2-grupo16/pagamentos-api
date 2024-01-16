package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type PedidoGatewayI interface {
	GetByID(pedidoID uint32) (*pedido.Pedido, error)
}

type MercadoPagoGatewayI interface {
	GetPedidoIDByPaymentID(paymentID string) (uint32, error)
	CreateQRCodeForPedido(pedido pedido.Pedido) (string, error)
}

type PagamentoGatewayI interface {
	UpdatePaymentStatusByPaymentID(pagamentoID uint32, status string) (*entities.Pagamento, error)
}

type QueueGatewayI interface {
	Publish(pagamento entities.Pagamento) error
}
