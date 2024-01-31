package factory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-communication/websocket/data"
)

func createArgs() ArgsWebSocketHost {
	return ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                "localhost:1234",
			WithAcknowledge:    false,
			Mode:               data.ModeClient,
			RetryDurationInSec: 1,
			BlockingAckOnError: false,
		},
		Marshaller: &testscommon.MarshallerMock{},
		Log:        &testscommon.LoggerMock{},
	}
}

func TestCreateClient(t *testing.T) {
	t.Parallel()

	args := createArgs()
	webSocketsClient, err := CreateWebSocketHost(args)
	require.Nil(t, err)
	require.Equal(t, "*client.client", fmt.Sprintf("%T", webSocketsClient))
}

func TestCreateServer(t *testing.T) {
	t.Parallel()

	args := createArgs()
	args.WebSocketConfig.Mode = data.ModeServer
	webSocketsClient, err := CreateWebSocketHost(args)
	require.Nil(t, err)
	require.Equal(t, "*server.server", fmt.Sprintf("%T", webSocketsClient))
}
