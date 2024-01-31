package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-communication/websocket/data"
	factoryHost "github.com/subrahamanyam341/andes-communication/websocket/factory"
	"github.com/subrahamanyam341/andes-core-16/marshal/factory"
	logger "github.com/subrahamanyam341/andes-logger-123"
)

var (
	marshaller, _ = factory.NewMarshalizer("json")
	log           = logger.GetOrCreate("client")
	url           = ":12345"
)

func main() {
	_ = logger.SetLogLevel("*:DEBUG")
	args := factoryHost.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        url,
			Mode:                       data.ModeClient,
			RetryDurationInSec:         1,
			WithAcknowledge:            true,
			BlockingAckOnError:         false,
			DropMessagesIfNoConnection: false,
			AcknowledgeTimeoutInSec:    10,
		},
		Marshaller: marshaller,
		Log:        log,
	}

	wsClient, err := factoryHost.CreateWebSocketHost(args)
	if err != nil {
		log.Error("cannot create WebSocket client", "error", err)
		return
	}

	defer func() {
		_ = wsClient.Close()
	}()

	err = wsClient.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte, topic string, version uint32) error {
			log.Info("received", "topic", topic, "payload", string(payload), "version", fmt.Sprint(version))
			return nil
		},
	})
	log.LogIfError(err)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
	err = wsClient.Close()
	log.LogIfError(err)
}
