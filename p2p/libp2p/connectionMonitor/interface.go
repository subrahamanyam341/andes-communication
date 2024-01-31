package connectionMonitor

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/subrahamanyam341/andes-core-16/core"
)

// Sharder defines the eviction computing process of unwanted peers
type Sharder interface {
	ComputeEvictionList(pidList []peer.ID) []peer.ID
	Has(pid peer.ID, list []peer.ID) bool
	SetSeeders(addresses []string)
	IsSeeder(pid core.PeerID) bool
	IsInterfaceNil() bool
}
