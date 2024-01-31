package mock

import (
	"time"

	"github.com/subrahamanyam341/andes-core-16/core"
)

// PeerDenialEvaluatorStub -
type PeerDenialEvaluatorStub struct {
	UpsertPeerIDCalled func(pid core.PeerID, duration time.Duration) error
	IsDeniedCalled     func(pid core.PeerID) bool
}

// UpsertPeerID -
func (pdes *PeerDenialEvaluatorStub) UpsertPeerID(pid core.PeerID, duration time.Duration) error {
	if pdes.UpsertPeerIDCalled != nil {
		return pdes.UpsertPeerIDCalled(pid, duration)
	}

	return nil
}

// IsDenied -
func (pdes *PeerDenialEvaluatorStub) IsDenied(pid core.PeerID) bool {
	if pdes.IsDeniedCalled != nil {
		return pdes.IsDeniedCalled(pid)
	}
	return false
}

// IsInterfaceNil -
func (pdes *PeerDenialEvaluatorStub) IsInterfaceNil() bool {
	return pdes == nil
}
