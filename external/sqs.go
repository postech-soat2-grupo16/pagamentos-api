package external

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

func GetSqsClient() *sqs.SQS {
	if os.Getenv("IS_LOCAL") == "true" {

	}

	region := "us-east-1"
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	sqsClient := sqs.New(session.Must(session.NewSession(awsConfig)))
	fmt.Printf("sqs client connected: %v\n", *sqsClient)

	return sqsClient
}
