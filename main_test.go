package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/skpr/certificate-expiry-notification-lambda/internal/acm"
	"github.com/skpr/certificate-expiry-notification-lambda/internal/slack"
	util "github.com/skpr/certificate-expiry-notification-lambda/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestHandleCert(t *testing.T) {
	data := []byte(`{
		"version": "0",
		"id": "9c95e8e4-96a4-ef3f-b739-b6aa5b193afb",
		"detail-type": "ACM Certificate Approaching Expiration",
		"source": "aws.acm",
		"account": "123456789012",
		"time": "2020-09-30T06:51:08Z",
		"region": "us-east-1",
		"resources": ["arn:aws:acm:us-east-1:123456789012:certificate/61f50cd4-45b9-4259-b049-d0a53682fa4b"],
		"detail": {
		  "DaysToExpiry": 31,
		  "CommonName": "example.com"
		}
	  }`)
	var event acm.Event
	err := json.Unmarshal(data, &event)

	if err != nil {
		fmt.Println(err)
	}

	// os.Setenv("SLACK_WEBHOOK_URL", "")
	os.Setenv("SLACK_WEBHOOK_URL", "https://hooks.slack.com/")

	fmt.Printf("Event: %v\n", len(event.Resources))

	assert.Equal(t, "ACM Certificate Approaching Expiration", event.DetailType)

	config, err := util.LoadConfig(".")
	assert.NoError(t, err)

	errs := config.Validate()
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Println(e)
		}
	}
	// slackClient, err := slack.NewClient(config.SlackWebhookURL)
	// assert.NoError(t, err)

	slackMock := &slack.MockClient{}

	err = handleCert(event, slackMock)

	assert.NoError(t, err)
}
