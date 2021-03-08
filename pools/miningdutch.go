package pools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/extras/stop"
	"github.com/sirupsen/logrus"
)

// MiningDutchAPIKey api key for polling the mining dutch mining pool
var MiningDutchAPIKey string
var lastMiningDutchResult *MiningDutchResponse

var mdurl = "https://www.mining-dutch.nl/pools/lbrycredits.php?page=api&id=66009&action=getpoolhashrate&api_key="

func monitorMiningDutch(parent *stop.Group) {
	stopper := stop.New(parent)
	ticker := time.NewTicker(checkPeriod)
	for {
		select {
		case <-stopper.Ch():
			return
		case <-ticker.C:
			err := checkMiningDutch()
			if err != nil {
				logrus.Error(errors.FullTrace(err))
			}
		}
	}
}

func checkMiningDutch() error {
	logrus.Debug("Checking Mining Dutch")
	apiURL := fmt.Sprintf("%s%s", mdurl, MiningDutchAPIKey)
	req, err := http.NewRequest(http.MethodGet, apiURL, bytes.NewReader(nil))
	if err != nil {
		return errors.Err(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Err(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Err(err)
	}
	result := &MiningDutchResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return errors.Err(err)
	}
	lastMiningDutchResult = result
	return nil
}

// MiningDutchResponse holds the data returned from the Mining Dutch API
type MiningDutchResponse struct {
	Getpoolhashrate struct {
		Version string  `json:"version"`
		Runtime float64 `json:"runtime"`
		Data    float64 `json:"data"`
	} `json:"getpoolhashrate"`
}
