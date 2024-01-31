package messagecheck

import "github.com/subrahamanyam341/andes-core-16/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	IsInterfaceNil() bool
}
