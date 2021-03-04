package pools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/lbryio/lbry.go/v2/extras/stop"
)

const url = "https://www2.coinmine.pl/lbc/index.php?page=api&action=getpoolstatus&api_key="

// CoinMineAPIKey is the api key to use to access coinmine metrics
var CoinMineAPIKey string

func monitorCoinmine(parent *stop.Group) {
	stopper := stop.New(parent)
	ticker := time.NewTicker(checkPeriod)
	for {
		select {
		case <-stopper.Ch():
			return
		case <-ticker.C:
			err := checkCoinMine()
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func checkCoinMine() error {
	logrus.Debug("Checking Coin Mine")
	apiURL := fmt.Sprintf("%s%s", url, CoinMineAPIKey)
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
	result := &CoinMineResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return errors.Err(err)
	}
	return nil
}

// CoinMineResult holds the data returned from the CoinMine API
type CoinMineResult struct {
	Getpoolstatus struct {
		Version string  `json:"version"`
		Runtime float64 `json:"runtime"`
		Data    struct {
			PoolName            string  `json:"pool_name"`
			Hashrate            float64 `json:"hashrate"`
			Efficiency          float64 `json:"efficiency"`
			Progress            float64 `json:"progress"`
			Workers             int     `json:"workers"`
			Currentnetworkblock int     `json:"currentnetworkblock"`
			Nextnetworkblock    int     `json:"nextnetworkblock"`
			Lastblock           int     `json:"lastblock"`
			Networkdiff         float64 `json:"networkdiff"`
			Esttime             float64 `json:"esttime"`
			Estshares           int64   `json:"estshares"`
			Timesincelast       int     `json:"timesincelast"`
			Nethashrate         float64 `json:"nethashrate"`
		} `json:"data"`
	} `json:"getpoolstatus"`
}
