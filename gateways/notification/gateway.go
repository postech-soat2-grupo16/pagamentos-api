package notification

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"os"
)

type GatewayInterface interface {
	SendNotification(pagamento *entities.Pagamento, email string) error
}

type Gateway struct {
	notificationTopic string
	notification      *sns.SNS
}

type GatewayMock struct {
}

func (g GatewayMock) SendNotification(pagamento *entities.Pagamento, email string) error {
	return nil
}

func NewGateway(notificationClient *sns.SNS) GatewayInterface {
	if notificationClient == nil {
		return NewGatewayMock()
	}
	return &Gateway{
		notificationTopic: os.Getenv("NOTIFICATION_TOPIC"),
		notification:      notificationClient,
	}
}

func NewGatewayMock() *GatewayMock {
	return &GatewayMock{}
}

func (g *Gateway) SendNotification(pagamento *entities.Pagamento, email string) error {

	notificationMessage := fmt.Sprintf("Status do Pagamento %d atualizado para %s", pagamento.ID, pagamento.Status)
	fmt.Printf("Sending message: %s\n", notificationMessage)

	//Build message
	message := &sns.PublishInput{
		TopicArn: &g.notificationTopic,
		Subject:  aws.String(fmt.Sprintf("Status do Pagamento %d", pagamento.ID)),
		Message:  &notificationMessage,
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"target": {
				DataType:    aws.String("String"),
				StringValue: aws.String(email),
			},
		}}

	fmt.Printf("Enviando Notificação de pagamento ID %d\n", pagamento.ID)
	messageResult, err := g.notification.Publish(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}
	fmt.Printf("Notificação de pagamento ID %d enviada com sucesso: %v\n", pagamento.ID, messageResult)

	return nil
}
