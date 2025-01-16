// Package main is the entry point for the lambda.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/skpr/certificate-expiry-notification-lambda/internal/acm"
	"github.com/skpr/certificate-expiry-notification-lambda/internal/slack"
	util "github.com/skpr/certificate-expiry-notification-lambda/internal/utils"
)

const (
	// EventDetailType is the type of event we are looking for in the payload.
	EventDetailType = "ACM Certificate Approaching Expiration"
)

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(_ context.Context, rawEvent json.RawMessage) error {
	var event acm.Event
	err := json.Unmarshal([]byte(rawEvent), &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	if event.DetailType != EventDetailType || len(event.Resources) > 0 {
		return fmt.Errorf("event type not supported")
	}

	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	errs := config.Validate()
	if len(errs) > 0 {
		return fmt.Errorf("configuration error: %s", strings.Join(errs, "\n"))
	}
	slackClient, err := slack.NewClient(config.SlackWebhookURL)
	if err != nil {
		return fmt.Errorf("failed to create Slack client: %w", err)
	}

	err = handleCert(event, slackClient)
	if err != nil {
		return err
	}
	return nil
}

func handleCert(event acm.Event, slackClient slack.ClientInterface) error {
	err := postNotification(slackClient, event)
	if err != nil {
		return err
	}

	return nil
}

func postNotification(slackClient slack.ClientInterface, event acm.Event) error {
	err := slackClient.PostMessage(slack.PostMessageParams{
		Domain:         event.Detail.CommonName,
		CertificateArn: event.Resources[0],
		Expiry:         strconv.Itoa(event.Detail.DaysToExpiry),
		Description:    "The above certificate is expiring within 45 days.",
	})
	if err != nil {
		return fmt.Errorf("failed to post Slack message: %w", err)
	}
	return nil
}
