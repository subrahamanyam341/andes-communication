package mock

import (
	"github.com/subrahamanyam341/andes-core-16/core"
)

// PeerShardResolverStub -
type PeerShardResolverStub struct {
	GetPeerInfoCalled func(pid core.PeerID) core.P2PPeerInfo
}

// GetPeerInfo -
func (psrs *PeerShardResolverStub) GetPeerInfo(pid core.PeerID) core.P2PPeerInfo {
	if psrs.GetPeerInfoCalled != nil {
		return psrs.GetPeerInfoCalled(pid)
	}
	return core.P2PPeerInfo{}
}

// IsInterfaceNil -
func (psrs *PeerShardResolverStub) IsInterfaceNil() bool {
	return psrs == nil
}
