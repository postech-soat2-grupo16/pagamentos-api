package message

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type GatewayInterface interface {
	SendMessage(pagamento *entities.Pagamento) error
}

type Gateway struct {
	queueURL string
	queue    *sqs.SQS
}

type GatewayMock struct {
}

func (g GatewayMock) SendMessage(pagamento *entities.Pagamento) error {
	return nil
}

func NewGateway(queueClient *sqs.SQS) GatewayInterface {
	if queueClient == nil {
		return NewGatewayMock()
	}
	return &Gateway{
		queueURL: os.Getenv("QUEUE_URL"),
		queue:    queueClient,
	}
}

func NewGatewayMock() *GatewayMock {
	return &GatewayMock{}
}

func (g *Gateway) SendMessage(pagamento *entities.Pagamento) error {
	jsonBody, err := json.Marshal(pagamento)
	if err != nil {
		fmt.Println("Erro ao serializar o objeto para JSON:", err)
		return nil
	}

	stringMessage := string(jsonBody)
	fmt.Printf("Sending message: %s\n", jsonBody)

	//Build message
	message := &sqs.SendMessageInput{
		QueueUrl:    &g.queueURL,
		MessageBody: &stringMessage,
	}

	messageResult, err := g.queue.SendMessage(message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}
	fmt.Printf("Message result: %s\n", messageResult)

	return nil
}
