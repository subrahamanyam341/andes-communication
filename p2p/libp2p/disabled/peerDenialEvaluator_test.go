package disabled_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/disabled"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestPeerDenialEvaluator_ShouldWork(t *testing.T) {
	t.Parallel()

	pde := &disabled.PeerDenialEvaluator{}

	assert.False(t, check.IfNil(pde))
	assert.Nil(t, pde.UpsertPeerID("", time.Second))
	assert.False(t, pde.IsDenied(""))
}
