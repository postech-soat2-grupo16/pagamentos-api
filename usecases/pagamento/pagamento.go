package pagamento

import (
	"errors"
	"fmt"
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"time"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"gorm.io/gorm"
)

type UseCase struct {
	pagamentoGateway   interfaces.PagamentoGatewayI
	mercadoPagoGateway interfaces.MercadoPagoGatewayI
	queueGateway       interfaces.QueueGatewayI
	pedidoGateway      interfaces.PedidoGatewayI
}

func NewUseCase(pagamentoGateway interfaces.PagamentoGatewayI,
	mercadoPagoGateway interfaces.MercadoPagoGatewayI,
	queueGateway interfaces.QueueGatewayI,
	pedidoGateway interfaces.PedidoGatewayI) UseCase {
	return UseCase{pagamentoGateway: pagamentoGateway,
		mercadoPagoGateway: mercadoPagoGateway,
		queueGateway:       queueGateway,
		pedidoGateway:      pedidoGateway}
}

func (p UseCase) CreateQRCode(pedidoID uint32) (*string, error) {
	pedidoResponse, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	qrCode, err := p.mercadoPagoGateway.CreateQRCodeForPedido(*pedidoResponse)
	if err != nil {
		return nil, err
	}

	return &qrCode, nil
}

func (p UseCase) UpdatePaymentStatusByPaymentID(pagamentoID uint32) (*entities.Pagamento, error) {
	var payment, err = p.pagamentoGateway.UpdatePaymentStatusByPaymentID(pagamentoID, pagamento.StatusPagamentoAprovado)
	if err != nil {
		fmt.Printf("Error updating payment status: %s\n", err)
		return nil, err
	}

	err = p.queueGateway.SendMessage(payment)
	if err != nil {
		fmt.Printf("Error sending payment message: %s\n", err)
		return nil, err
	}

	return payment, nil
}

func (p UseCase) CreatePayment(pedidoID uint32) (*entities.Pagamento, error) {
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	newPayment := entities.Pagamento{
		PedidoID:  pedido.ID,
		Amount:    pedido.GetAmount(),
		Status:    pagamento.StatusPagamentoIniciado,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return p.pagamentoGateway.CreatePayment(newPayment)
}

func (p UseCase) GetByID(paymentID uint32) (*entities.Pagamento, error) {
	payment, err := p.pagamentoGateway.GetByID(paymentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return payment, nil
}
