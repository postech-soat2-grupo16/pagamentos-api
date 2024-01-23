package message

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"os"
)

type Gateway struct {
	queueURL string
	queue    *sqs.SQS
}

func NewGateway(queueClient *sqs.SQS) *Gateway {
	return &Gateway{
		queueURL: os.Getenv("QUEUE_URL"),
		queue:    queueClient,
	}
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
