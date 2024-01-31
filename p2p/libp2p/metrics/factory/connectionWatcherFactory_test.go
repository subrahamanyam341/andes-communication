package factory_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/metrics/factory"
	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestNewConnectionsWatcher(t *testing.T) {
	t.Parallel()

	t.Run("print connections watcher", func(t *testing.T) {
		t.Parallel()

		cw, err := factory.NewConnectionsWatcher(p2p.ConnectionWatcherTypePrint, time.Second, &testscommon.LoggerStub{})
		assert.Nil(t, err)
		assert.False(t, check.IfNil(cw))
		assert.Equal(t, "*metrics.printConnectionsWatcher", fmt.Sprintf("%T", cw))
	})
	t.Run("disabled connections watcher", func(t *testing.T) {
		t.Parallel()

		cw, err := factory.NewConnectionsWatcher(p2p.ConnectionWatcherTypeDisabled, time.Second, &testscommon.LoggerStub{})
		assert.Nil(t, err)
		assert.False(t, check.IfNil(cw))
		assert.Equal(t, "*metrics.disabledConnectionsWatcher", fmt.Sprintf("%T", cw))
	})
	t.Run("empty connections watcher", func(t *testing.T) {
		t.Parallel()

		cw, err := factory.NewConnectionsWatcher(p2p.ConnectionWatcherTypeEmpty, time.Second, &testscommon.LoggerStub{})
		assert.Nil(t, err)
		assert.False(t, check.IfNil(cw))
		assert.Equal(t, "*metrics.disabledConnectionsWatcher", fmt.Sprintf("%T", cw))
	})
	t.Run("unknown type", func(t *testing.T) {
		t.Parallel()

		cw, err := factory.NewConnectionsWatcher("unknown", time.Second, &testscommon.LoggerStub{})
		assert.True(t, errors.Is(err, factory.ErrUnknownConnectionWatcherType))
		assert.True(t, check.IfNil(cw))
	})
}
