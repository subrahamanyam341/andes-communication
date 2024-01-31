package mock

import "github.com/subrahamanyam341/andes-core-16/core"

// PeerTopicNotifierStub -
type PeerTopicNotifierStub struct {
	NewPeerFoundCalled func(pid core.PeerID, topic string)
}

// NewPeerFound -
func (stub *PeerTopicNotifierStub) NewPeerFound(pid core.PeerID, topic string) {
	if stub.NewPeerFoundCalled != nil {
		stub.NewPeerFoundCalled(pid, topic)
	}
}

// IsInterfaceNil -
func (stub *PeerTopicNotifierStub) IsInterfaceNil() bool {
	return stub == nil
}
