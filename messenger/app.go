/* This package handles receiving messages from SQS and forward it to slack */
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	// This is built on lambda so lambda context
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// defines the slack request body type
type SlackRequestBody struct {
	Text string `json: "text"`
}

/*
this is the handler that execute the main loop that sends message
to slack based on the webhook stored in environment
*/
func SendMessageToSlack(ctx context.Context) {
	// get the webhook url
	WebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	// get lambda context to extract the body content
	
	// send it to the sendslacknotification function
	
}

/*
	This function does the actual work of doing a post request to the webhook URL
	takes in a WebhookUrl of type string and a msg of type string
	returns a error of type error
*/
func SendSlackNotification(WebhookUrl string, msg string) error {
	DataBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	// send the actual request
	req, err := http.NewRequest(http.MethodPost, WebhookUrl, bytes.NewBuffer(DataBody))
	if err != nil {
		return nil
	}
	// add json header
	req.Header.Add("Content-Type", "application/json")

	// actually send out the request and give it a 5second timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	// read through the result to check status
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("something happened from the request to slack")
	}
	return nil
}

/*
Start the lambda handler with the main function
*/
func main() {
	// start the function sendmessagetoslack with the context.Context{} interface
	lambda.Start(SendMessageToSlack)
}