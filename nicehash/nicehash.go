package nicehash

import (
	"time"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/extras/stop"

	"github.com/sirupsen/logrus"
)

// CheckPeriod time between checking on nicehash
var CheckPeriod = 60
var checkPeriod = time.Duration(CheckPeriod) * time.Second

// Monitor kicks off the monitoring of nice hash apis
func Monitor(parent *stop.Group) {
	stopper := stop.New(parent)
	ticker := time.NewTicker(checkPeriod)
	for {
		select {
		case <-stopper.Ch():
			return
		case <-ticker.C:
			err := checkNiceHash()
			if err != nil {
				logrus.Error(errors.FullTrace(err))
			}
		}
	}
}

func checkNiceHash() error {
	return nil
}
