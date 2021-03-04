package pools

import (
	"time"

	"github.com/lbryio/lbry.go/v2/extras/stop"
)

// CheckPeriod time between checking on a pool
var CheckPeriod = 60
var checkPeriod = time.Duration(CheckPeriod) * time.Second

var stopper *stop.Group

// MonitorPools kicks off the monitors for the different pools for mining LBRY
func MonitorPools(parent *stop.Group) {
	stopper := stop.New(parent)
	go monitorCoinmine(stopper)
}
