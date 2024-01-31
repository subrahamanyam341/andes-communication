package factory

import (
	"fmt"
	"time"

	"github.com/subrahamanyam341/andes-communication/p2p"
	"github.com/subrahamanyam341/andes-communication/p2p/libp2p/metrics"
)

// NewConnectionsWatcher creates a new ConnectionWatcher instance based on the input parameters
func NewConnectionsWatcher(connectionsWatcherType string, timeToLive time.Duration, logger p2p.Logger) (p2p.ConnectionsWatcher, error) {
	switch connectionsWatcherType {
	case p2p.ConnectionWatcherTypePrint:
		return metrics.NewPrintConnectionsWatcher(timeToLive, logger)
	case p2p.ConnectionWatcherTypeDisabled, p2p.ConnectionWatcherTypeEmpty:
		return metrics.NewDisabledConnectionsWatcher(), nil
	default:
		return nil, fmt.Errorf("%w %s", ErrUnknownConnectionWatcherType, connectionsWatcherType)
	}
}
