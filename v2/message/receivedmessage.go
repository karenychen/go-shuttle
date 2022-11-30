package message

import (
	"fmt"

	servicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/Azure/go-shuttle/v2/marshal"
)

type ReceivedMessage struct {
	msg         *servicebus.ReceivedMessage
	messageType string
}

// NewReceivedMessage creates a ReceivedMessage, validating type first
func NewReceivedMessage(msg *servicebus.ReceivedMessage) (*ReceivedMessage, error) {
	messageType, ok := msg.ApplicationProperties["type"]
	if !ok {
		return nil, fmt.Errorf("message did not include a \"type\" in UserProperties")
	}
	return &ReceivedMessage{msg, messageType.(string)}, nil
}

// Message returns the message as received by the SDK
func (m *ReceivedMessage) Message() *servicebus.ReceivedMessage {
	return m.msg
}

// Type returns the message type as a string, passed on the message ApplicationProperties
func (m *ReceivedMessage) Type() string {
	return m.messageType
}

// Body returns the message data as a string. Can be unmarshalled to the original struct
func (m *ReceivedMessage) Body() string {
	return string(m.msg.Body)
}

func (m *ReceivedMessage) Unmarshal(data []byte, v interface{}) error {
	contentType := *m.msg.ContentType
	marshaller, ok := marshal.DefaultMarshallerRegistry[contentType]
	if !ok {
		return fmt.Errorf("no marshaller registered for content-type %s", contentType)
	}
	return marshaller.Unmarshal(data, v)
}
