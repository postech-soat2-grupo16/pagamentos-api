package pagamento

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"gorm.io/gorm"
)

type UseCase struct {
	pagamentoGateway    interfaces.PagamentoGatewayI
	mercadoPagoGateway  interfaces.MercadoPagoGatewayI
	queueGateway        interfaces.QueueGatewayI
	notificationGateway interfaces.NotificationGatewayI
	pedidoGateway       interfaces.PedidoGatewayI
	clienteGateway      interfaces.ClienteGatewayI
}

func NewUseCase(pagamentoGateway interfaces.PagamentoGatewayI,
	mercadoPagoGateway interfaces.MercadoPagoGatewayI,
	queueGateway interfaces.QueueGatewayI,
	notificationGateway interfaces.NotificationGatewayI,
	pedidoGateway interfaces.PedidoGatewayI,
	clienteGateway interfaces.ClienteGatewayI,
) UseCase {
	return UseCase{pagamentoGateway: pagamentoGateway,
		mercadoPagoGateway:  mercadoPagoGateway,
		queueGateway:        queueGateway,
		notificationGateway: notificationGateway,
		pedidoGateway:       pedidoGateway,
		clienteGateway:      clienteGateway,
	}
}

func (p UseCase) CreateQRCode(pedidoID string) (*string, error) {
	_, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	//qrCode, err := p.mercadoPagoGateway.CreateQRCodeForPedido(*pedidoResponse)
	qrCode := "00020101021226940014BR.GOV.BCB.PIX2572pix-qr.mercadopago.com/instore/o/v2/6c46ee45-795a-4a2c-a594-69e7ae531cdb5204000053039865802BR5910Teste FIAP6009SAO PAULO62070503***6304DB48"

	return &qrCode, nil
}

func (p UseCase) UpdatePaymentStatusByPaymentID(pagamentoID uint32) (*entities.Pagamento, error) {
	var statusPagamento = entities.StatusPagamentoQRCodeCriado

	if pagamentoID%2 == 0 {
		fmt.Printf("Erro ao criar QR CODE para o pagamento %d", pagamentoID)
		statusPagamento = entities.StatusPagamentoQRCodeErro
	}

	var payment, err = p.pagamentoGateway.UpdatePaymentStatusByPaymentID(pagamentoID, statusPagamento)
	if err != nil {
		fmt.Printf("Error updating payment status: %s\n", err)
		return nil, err
	}
	fmt.Printf("Callback do QR-Code %d, status do pagamento atualizado para %s\n", payment.ID, payment.Status)

	//TODO envio da notificação SNS

	return payment, nil
}

func (p UseCase) CreatePayment(pedidoID string) (*entities.Pagamento, error) {
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	clienteID, _ := strconv.ParseUint(pedido.ClientID, 10, 32)
	newPayment := entities.Pagamento{
		PedidoID:  pedido.ID,
		ClienteID: uint32(clienteID),
		Amount:    pedido.GetAmount(),
		Status:    entities.StatusPagamentoIniciado,
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

func (p UseCase) ProcessPaymentStatus(pagamentoID uint32, statusPagamento string) (*entities.Pagamento, error) {
	pagamento, err := p.GetByID(pagamentoID)
	if err != nil {
		return nil, err
	}

	cliente, err := p.clienteGateway.GetByID(pagamento.ClienteID)
	if err != nil {
		return nil, err
	}

	updatedPayment, err := p.pagamentoGateway.UpdatePaymentStatusByPaymentID(pagamentoID, statusPagamento)
	if err != nil {
		fmt.Printf("Error updating payment status: %s\n", err)
		return nil, err
	}

	fmt.Printf("Atualização do status do pagamento %d, status do pagamento atualizado para %s\n", updatedPayment.ID, updatedPayment.Status)

	err = p.notificationGateway.SendNotification(pagamento, cliente.Email)
	if err != nil {
		fmt.Printf("Error sending payment notification: %s\n", err)
		return nil, err
	}

	if updatedPayment.IsPaymentApproved() {
		err = p.queueGateway.SendMessage(pagamento)
		if err != nil {
			fmt.Printf("Error sending payment message: %s\n", err)
			return nil, err
		}
	}

	return pagamento, nil
}
