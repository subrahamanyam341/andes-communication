package metrics_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/metrics"
	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-core-16/core"
)

func TestNewPrintConnectionsWatcher(t *testing.T) {
	t.Parallel()

	t.Run("invalid value for time to live parameter should error", func(t *testing.T) {
		t.Parallel()

		pcw, err := metrics.NewPrintConnectionsWatcher(metrics.MinTimeToLive-time.Nanosecond, &testscommon.LoggerStub{})
		assert.Nil(t, pcw)
		assert.True(t, errors.Is(err, metrics.ErrInvalidValueForTimeToLiveParam))
	})
	t.Run("nil logger should error", func(t *testing.T) {
		t.Parallel()

		pcw, err := metrics.NewPrintConnectionsWatcher(metrics.MinTimeToLive, nil)
		assert.Nil(t, pcw)
		assert.Equal(t, p2p.ErrNilLogger, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		pcw, err := metrics.NewPrintConnectionsWatcher(metrics.MinTimeToLive, &testscommon.LoggerStub{})
		assert.NotNil(t, pcw)
		assert.Nil(t, err)

		_ = pcw.Close()
	})
}

func TestPrintConnectionsWatcher_Close(t *testing.T) {
	t.Parallel()

	t.Run("no iteration has been done", func(t *testing.T) {
		t.Parallel()

		pcw, _ := metrics.NewPrintConnectionsWatcher(time.Hour, &testscommon.LoggerStub{})
		err := pcw.Close()

		assert.Nil(t, err)
		time.Sleep(time.Second) // allow the go routine to close
		assert.True(t, pcw.GoRoutineClosed())
	})
	t.Run("iterations were done", func(t *testing.T) {
		t.Parallel()

		pcw, _ := metrics.NewPrintConnectionsWatcher(time.Second, &testscommon.LoggerStub{})
		time.Sleep(time.Second * 4)
		err := pcw.Close()

		assert.Nil(t, err)
		time.Sleep(time.Second) // allow the go routine to close
		assert.True(t, pcw.GoRoutineClosed())
	})

}

func TestPrintConnectionsWatcher_NewKnownConnection(t *testing.T) {
	t.Parallel()

	t.Run("invalid connection", func(t *testing.T) {
		providedPid := core.PeerID("pid")
		connection := " "
		numCalled := 0

		handler := func(pid core.PeerID, conn string, log p2p.Logger) {
			numCalled++
		}
		pcw, _ := metrics.NewPrintConnectionsWatcherWithHandler(time.Hour, handler)

		pcw.NewKnownConnection(providedPid, connection)
		assert.Equal(t, 0, numCalled)
	})
	t.Run("valid connection", func(t *testing.T) {
		providedPid := core.PeerID("pid")
		connection := "connection"
		numCalled := 0

		handler := func(pid core.PeerID, conn string, log p2p.Logger) {
			numCalled++
			assert.Equal(t, providedPid, pid)
			assert.Equal(t, connection, conn)
		}
		pcw, _ := metrics.NewPrintConnectionsWatcherWithHandler(time.Hour, handler)

		pcw.NewKnownConnection(providedPid, connection)
		assert.Equal(t, 1, numCalled)
		pcw.NewKnownConnection(providedPid, connection)
		assert.Equal(t, 1, numCalled)
	})
}

func TestLogPrintHandler_shouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("should have not panic: %v", r))
		}
	}()

	metrics.LogPrintHandler("pid", "connection", &testscommon.LoggerStub{})
}

func TestPrintConnectionsWatcher_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	pcw, _ := metrics.NewPrintConnectionsWatcher(time.Second, nil)
	assert.True(t, pcw.IsInterfaceNil())

	pcw, _ = metrics.NewPrintConnectionsWatcher(time.Second, &testscommon.LoggerStub{})
	assert.False(t, pcw.IsInterfaceNil())
}
