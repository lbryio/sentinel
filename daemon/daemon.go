package daemon

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lbryio/sentinel/nicehash"

	"github.com/lbryio/lbry.go/v2/extras/stop"

	"github.com/lbryio/sentinel/pools"
)

var stopper = stop.New(nil)

// Start starts the daemon that runs collecting information and watching the blockchain
func Start() {
	//Start daemon jobs
	go pools.Monitor(stopper)
	go nicehash.Monitor(stopper)

	//Wait for shutdown signal, then shutdown api server. This will wait for all connections to finish.
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-interruptChan
	stopper.StopAndWait()
}
