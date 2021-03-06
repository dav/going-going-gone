package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

const DEFAULT_HTML_RESPONSE = "<head><meta http-equiv=\"refresh\" content=\"0; URL=\"https://www.youtube.com/watch?v=CPNK0VspQ0M\" /></head>"

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       DEFAULT_HTML_RESPONSE,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
