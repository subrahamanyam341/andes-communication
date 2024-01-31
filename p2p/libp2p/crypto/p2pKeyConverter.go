package crypto

import (
	"fmt"

	libp2pCrypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/subrahamanyam341/andes-core-16/core"
	"github.com/subrahamanyam341/andes-core-16/core/check"
	crypto "github.com/subrahamanyam341/andes-crypto-123"
)

type p2pKeyConverter struct {
}

// NewP2PKeyConverter returns a new instance of p2pKeyConverter
func NewP2PKeyConverter() *p2pKeyConverter {
	return &p2pKeyConverter{}
}

// ConvertPeerIDToPublicKey will convert core peer id to common public key
func (converter *p2pKeyConverter) ConvertPeerIDToPublicKey(keyGen crypto.KeyGenerator, pid core.PeerID) (crypto.PublicKey, error) {
	libp2pPid, err := peer.IDFromBytes(pid.Bytes())
	if err != nil {
		return nil, err
	}

	pubk, err := libp2pPid.ExtractPublicKey()
	if err != nil {
		return nil, fmt.Errorf("cannot extract signing key: %s", err.Error())
	}

	pubKeyBytes, err := pubk.Raw()
	if err != nil {
		return nil, err
	}

	return keyGen.PublicKeyFromByteArray(pubKeyBytes)
}

// ConvertPublicKeyToPeerID will convert a public key to core.PeerID
func (converter *p2pKeyConverter) ConvertPublicKeyToPeerID(pk crypto.PublicKey) (core.PeerID, error) {
	if check.IfNil(pk) {
		return "", ErrNilPublicKey
	}

	pkBytes, err := pk.ToByteArray()
	if err != nil {
		return "", err
	}

	libp2pPk, err := libp2pCrypto.UnmarshalSecp256k1PublicKey(pkBytes)
	if err != nil {
		return "", err
	}

	pid, err := peer.IDFromPublicKey(libp2pPk)
	if err != nil {
		return "", err
	}

	return core.PeerID(pid), nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (converter *p2pKeyConverter) IsInterfaceNil() bool {
	return converter == nil
}
