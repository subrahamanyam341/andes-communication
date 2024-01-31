package websocket

import (
	"github.com/subrahamanyam341/andes-communication/websocket/data"
	"github.com/subrahamanyam341/andes-core-16/core/check"
	"github.com/subrahamanyam341/andes-core-16/marshal"
)

type webSocketsPayloadConverter struct {
	marshaller marshal.Marshalizer
}

// NewWebSocketPayloadConverter returns a new instance of websocketPayloadParser
func NewWebSocketPayloadConverter(marshaller marshal.Marshalizer) (*webSocketsPayloadConverter, error) {
	if check.IfNil(marshaller) {
		return nil, data.ErrNilMarshaller
	}

	return &webSocketsPayloadConverter{
		marshaller: marshaller,
	}, nil
}

// ExtractWsMessage will extract the provided payload in a *data.WsMessage
func (wpc *webSocketsPayloadConverter) ExtractWsMessage(payload []byte) (*data.WsMessage, error) {
	wsMessage := &data.WsMessage{}
	err := wpc.marshaller.Unmarshal(wsMessage, payload)
	if err != nil {
		return nil, err
	}

	return wsMessage, nil
}

// ConstructPayload will marshal the provided *data.WsMessage
func (wpc *webSocketsPayloadConverter) ConstructPayload(wsMessage *data.WsMessage) ([]byte, error) {
	return wpc.marshaller.Marshal(wsMessage)
}

// IsInterfaceNil -
func (wpc *webSocketsPayloadConverter) IsInterfaceNil() bool {
	return wpc == nil
}
