package metrics_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/metrics"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestDisabledConnectionsWatcher_MethodsShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("should have not panic: %v", r))
		}
	}()

	dcw := metrics.NewDisabledConnectionsWatcher()
	assert.False(t, check.IfNil(dcw))
	dcw.NewKnownConnection("", "")
	err := dcw.Close()
	assert.Nil(t, err)
}
