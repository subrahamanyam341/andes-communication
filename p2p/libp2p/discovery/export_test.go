package discovery

import (
	"context"
	"time"

	"github.com/subrahamanyam341/andes-communication/p2p"
)

const KadDhtName = kadDhtName
const OptimizedKadDhtName = optimizedKadDhtName
const NullName = nilName

// ------- ContinuousKadDhtDiscoverer

func (ckdd *continuousKadDhtDiscoverer) ConnectToOnePeerFromInitialPeersList(
	durationBetweenAttempts time.Duration,
	initialPeersList []string) <-chan struct{} {

	return ckdd.connectToOnePeerFromInitialPeersList(durationBetweenAttempts, initialPeersList)
}

func (ckdd *continuousKadDhtDiscoverer) StopDHT() error {
	ckdd.mutKadDht.Lock()
	err := ckdd.stopDHT()
	ckdd.mutKadDht.Unlock()

	return err
}

// NewOptimizedKadDhtDiscovererWithInitFunc -
func NewOptimizedKadDhtDiscovererWithInitFunc(
	arg ArgKadDht,
	createFunc func(ctx context.Context) (KadDhtHandler, error),
) (*optimizedKadDhtDiscoverer, error) {
	sharder, err := prepareArguments(arg)
	if err != nil {
		return nil, err
	}

	if arg.SeedersReconnectionInterval < minIntervalForSeedersReconnection {
		return nil, p2p.ErrInvalidSeedersReconnectionInterval
	}

	okdd := &optimizedKadDhtDiscoverer{
		sharder:                     sharder,
		peersRefreshInterval:        arg.PeersRefreshInterval,
		seedersReconnectionInterval: arg.SeedersReconnectionInterval,
		protocolID:                  arg.ProtocolID,
		initialPeersList:            arg.InitialPeersList,
		bucketSize:                  arg.BucketSize,
		routingTableRefresh:         arg.RoutingTableRefresh,
		status:                      statNotInitialized,
		chanInit:                    make(chan struct{}),
		errChanInit:                 make(chan error),
		chanConnectToSeeders:        make(chan struct{}),
		log:                         arg.Logger,
	}

	okdd.createKadDhtHandler = createFunc
	argConnectionManagement := ArgsHostWithConnectionManagement{
		ConnectableHost:    arg.Host,
		Sharder:            okdd.sharder,
		ConnectionsWatcher: arg.ConnectionWatcher,
	}
	okdd.hostConnManagement, err = NewHostWithConnectionManagement(argConnectionManagement)
	if err != nil {
		return nil, err
	}

	go okdd.processLoop(arg.Context)

	return okdd, nil
}
