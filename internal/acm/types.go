package acm

// Event used to parse to the lambda.
type Event struct {
	Version    string      `json:"version"`
	ID         string      `json:"id"`
	DetailType string      `json:"detail-type"`
	Source     string      `json:"source"`
	Account    string      `json:"account"`
	Time       string      `json:"time"`
	Region     string      `json:"region"`
	Resources  []string    `json:"resources"`
	Detail     EventDetail `json:"detail"`
}

// EventDetail used to understand the event.
type EventDetail struct {
	DaysToExpiry int    `json:"DaysToExpiry"`
	CommonName   string `json:"CommonName"`
}
