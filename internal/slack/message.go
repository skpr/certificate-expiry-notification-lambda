package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// PostMessageParams are the parameters required to post a message to Slack.
type PostMessageParams struct {
	// Domain this event is related to.
	Domain string
	// Certificate ARN this event is related to.
	CertificateArn string
	// Project this event is related to.
	Expiry string
	// Message to be displayed.
	Description string
}

// Validate the parameters.
func (p PostMessageParams) Validate() error {
	var errs []error

	if p.Domain == "" {
		errs = append(errs, fmt.Errorf("domain is required"))
	}

	if p.CertificateArn == "" {
		errs = append(errs, fmt.Errorf("certificate ARN is required"))
	}

	if p.Expiry == "" {
		errs = append(errs, fmt.Errorf("expiry is required"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// PostMessage to Slack channel.
func (c *Client) PostMessage(params PostMessageParams) error {
	if err := params.Validate(); err != nil {
		return fmt.Errorf("invalid parameters: %w", err)
	}

	var message Message

	context := BlockContext{
		Type: BlockTypeContext,
	}

	context.Elements = []BlockContextElement{
		{
			Type: BlockElementTypeMarkdown,
			Text: fmt.Sprintf("*Domain* = %s", params.Domain),
		},
		{
			Type: BlockElementTypeMarkdown,
			Text: fmt.Sprintf("*CertificateArn* = %s", params.CertificateArn),
		},
		{
			Type: BlockElementTypeMarkdown,
			Text: fmt.Sprintf("*Expiry Within* = %s", params.Expiry),
		},
	}

	message.Blocks = append(message.Blocks, context)

	// Separate the context from the content.
	message.Blocks = append(message.Blocks, BlockDivider{
		Type: BlockTypeDivider,
	})

	// Details of the alarm.
	details := BlockSection{
		Type: BlockTypeSection,
		Text: BlockSectionText{
			Type: BlockTextTypeMarkdown,
			Text: fmt.Sprintf("*%s*", params.Description),
		},
	}

	message.Blocks = append(message.Blocks, details)

	request, err := json.Marshal(message)
	if err != nil {
		return err
	}

	fmt.Println(string(request))

	for _, webhook := range c.webhooks {
		req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(request))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)

		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("returned status code: %d", resp.StatusCode)
		}
	}

	return nil
}
