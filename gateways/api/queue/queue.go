package queue

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type APIRepository struct {
	AWSAccessKey string
	AWSSecretKey string
	AWSRegion    string
}

func NewGateway(awsAccessKey, awsSecretKey, awsRegion string) *APIRepository {
	return &APIRepository{
		awsAccessKey,
		awsSecretKey,
		awsRegion,
	}
}

func (a *APIRepository) Publish(pagamento entities.Pagamento) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(a.AWSRegion),
		Credentials: credentials.NewStaticCredentials(a.AWSAccessKey, a.AWSSecretKey, ""),
	})
	if err != nil {
		fmt.Println("Erro ao criar sessão da AWS:", err)
		return nil
	}

	sqsClient := sqs.New(sess)

	queueURL := "http://localhost/payment-queue"

	jsonBody, err := json.Marshal(pagamento)
	if err != nil {
		fmt.Println("Erro ao serializar o objeto para JSON:", err)
		return nil
	}

	sendMessageOutput, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(string(jsonBody)),
		QueueUrl:     &queueURL,
		DelaySeconds: aws.Int64(0),
	})
	if err != nil {
		fmt.Println("Erro ao enviar mensagem para a fila:", err)
		return nil
	}

	fmt.Println("Mensagem enviada com sucesso. ID:", *sendMessageOutput.MessageId)

	return nil
}
