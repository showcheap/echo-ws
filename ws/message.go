package ws

// Message a ws message payload
type Message struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}
