package websocket

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-communication/websocket/data"
	"github.com/subrahamanyam341/andes-core-16/data/outport"
)

func TestNewWebSocketPayloadConverter(t *testing.T) {
	t.Parallel()

	payloadConverter, err := NewWebSocketPayloadConverter(nil)
	require.Nil(t, payloadConverter)
	require.Equal(t, data.ErrNilMarshaller, err)

	payloadConverter, _ = NewWebSocketPayloadConverter(&testscommon.MarshallerMock{})
	require.NotNil(t, payloadConverter)
	require.False(t, payloadConverter.IsInterfaceNil())
}

func TestWebSocketPayloadConverter_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	addrGroup, _ := NewWebSocketPayloadConverter(nil)
	require.True(t, addrGroup.IsInterfaceNil())

	addrGroup, _ = NewWebSocketPayloadConverter(&testscommon.MarshallerMock{})
	require.False(t, addrGroup.IsInterfaceNil())
}

func TestWebSocketsPayloadConverter_ConstructPayload(t *testing.T) {
	t.Parallel()

	payloadConverter, _ := NewWebSocketPayloadConverter(&testscommon.MarshallerMock{})

	wsMessage := &data.WsMessage{
		WithAcknowledge: true,
		Payload:         []byte("test"),
		Topic:           outport.TopicSaveAccounts,
		Counter:         10,
		Type:            data.PayloadMessage,
	}

	payload, err := payloadConverter.ConstructPayload(wsMessage)
	require.Nil(t, err)

	newWsMessage, err := payloadConverter.ExtractWsMessage(payload)
	require.Nil(t, err)
	require.Equal(t, wsMessage, newWsMessage)
}
