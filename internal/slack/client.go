// Package slack is used to interact with Slack.
package slack

// ClientInterface is used to interact with Slack.
type ClientInterface interface {
	PostMessage(params PostMessageParams) error
}

// Client is used to interact with Slack.
type Client struct {
	webhooks []string
}

// NewClient for interacting with Slack.
func NewClient(webhooks []string) (*Client, error) {
	client := &Client{
		webhooks: webhooks,
	}

	return client, nil
}

// MockClient for testing purposes.
type MockClient struct {
	PostMessageParams PostMessageParams
}

// PostMessage mocks the Slack API.
func (m *MockClient) PostMessage(params PostMessageParams) error {
	// Store the parameters so we can check them in our tests.
	m.PostMessageParams = params
	return nil
}
