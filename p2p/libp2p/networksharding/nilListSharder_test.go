package networksharding_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/networksharding"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestNilListSharderSharder(t *testing.T) {
	nls := networksharding.NewNilListSharder()

	assert.False(t, check.IfNil(nls))
	assert.Equal(t, 0, len(nls.ComputeEvictionList(nil)))
	assert.False(t, nls.Has("", nil))
	assert.Nil(t, nls.SetPeerShardResolver(nil))
}
