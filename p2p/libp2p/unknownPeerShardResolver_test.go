package libp2p_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p"
	"github.com/subrahamanyam341/andes-core-16/core"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestUnknownPeerShardResolver_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	upsr := libp2p.NewUnknownPeerShardResolver()
	assert.False(t, check.IfNil(upsr))
}

func TestUnknownPeerShardResolver_GetPeerInfoShouldReturnUnknownId(t *testing.T) {
	t.Parallel()

	upsr := libp2p.NewUnknownPeerShardResolver()
	expectedPeerInfo := core.P2PPeerInfo{
		PeerType: core.UnknownPeer,
		ShardID:  0,
	}

	assert.Equal(t, expectedPeerInfo, upsr.GetPeerInfo(""))
}
