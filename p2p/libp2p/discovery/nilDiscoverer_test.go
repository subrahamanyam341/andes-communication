package discovery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/discovery"
	"github.com/subrahamanyam341/andes-core-16/core/check"
)

func TestNilDiscoverer(t *testing.T) {
	t.Parallel()

	nd := discovery.NewNilDiscoverer()

	assert.False(t, check.IfNil(nd))
	assert.Equal(t, discovery.NullName, nd.Name())
	assert.Nil(t, nd.Bootstrap())
}
