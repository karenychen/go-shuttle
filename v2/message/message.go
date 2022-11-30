package message

import (
	"fmt"

	servicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// Message is the wrapping type of service bus message with a type
type Message struct {
	msg         *servicebus.Message
	messageType string
}

// NewMessage creates a message, validating type first
func NewMessage(msg *servicebus.Message) (*Message, error) {
	messageType, ok := msg.ApplicationProperties["type"]
	if !ok {
		return nil, fmt.Errorf("message did not include a \"type\" in ApplicationProperties")
	}
	return &Message{msg, messageType.(string)}, nil
}

// Message returns the message as received by the SDK
func (m *Message) Message() *servicebus.Message {
	return m.msg
}

// Type returns the message type as a string, passed on the message UserProperties
func (m *Message) Type() string {
	return m.messageType
}

// Body returns the message data as a string. Can be unmarshalled to the original struct
func (m *Message) Body() string {
	return string(m.msg.Body)
}
