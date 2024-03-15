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

	notificationMessage := fmt.Sprintf("Status do Pagamento %d atualizado PARA %s", pagamento.ID, pagamento.Status)
	fmt.Printf("Sending message: %s\n", notificationMessage)

	//Build message
	//TODO: Adicionar o envio a apenas 1 email (destino)
	message := &sns.PublishInput{
		TopicArn:  &g.notificationTopic,
		Message:   &notificationMessage,
		TargetArn: aws.String(fmt.Sprintf("%s:endpoint/email/%s", os.Getenv("NOTIFICATION_TOPIC"), email)),
		Subject:   aws.String("Subject of the message"),
	}

	fmt.Printf("Enviando mensagem de Notificação de pagamento ID %d\n", pagamento.ID)
	messageResult, err := g.notification.Publish(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}
	fmt.Printf("Mensagem de Notificação de pagamento ID %d enviada com sucesso: %v\n", pagamento.ID, messageResult)

	return nil
}
