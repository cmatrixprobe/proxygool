package store

import (
	"encoding/json"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/util"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// ValidateProxy checks addresses and decides whether to store.
func ValidateProxy(address *model.Address) {
	if CheckAddress(address) {
		insertProxy(address)
	}
}

// CheckAddress checks the address work or not.
func CheckAddress(address *model.Address) bool {
	// test by speedtest.cn
	testAddr := util.CombURL(address)
	targetURL := address.Protocol + "://forge.speedtest.cn/api/location/geo?ip=" + address.Host

	reqLogger := logrus.WithFields(logrus.Fields{
		"testAddr":  testAddr,
		"targetURL": targetURL,
	})
	reqLogger.Info()

	begin := time.Now()
	// get targetURL by test proxy
	resp, _, errs := gorequest.New().Proxy(testAddr).Get(targetURL).End()
	if errs != nil {
		reqLogger.Warn(errs)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		reqLogger.WithField("StatusCode", resp.StatusCode).Info()
		return false
	}

	// verify json
	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		reqLogger.Warn(err)
		return false
	}
	var tr model.TestResponse
	err = json.Unmarshal(info, &tr)
	if err != nil {
		reqLogger.Warn(err)
		return false
	}
	reqLogger.Info(tr)

	// speed(ms)
	end := time.Now()
	address.Speed = end.Sub(begin).Nanoseconds() / 1e6
	reqLogger.WithField("speed", address.Speed).Info()

	return true
}

// CheckProxyDB checks proxy addresses in DB.
func CheckProxyDB() {
	addresses, err := storage.GetAll()
	if err != nil {
		logrus.Warn(err)
		return
	}

	// check and delete unavailable records
	count, err := storage.Count()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.WithField("record", count).Info("Before check")
	var wg sync.WaitGroup
	for _, address := range addresses {
		wg.Add(1)
		go func(v *model.Address) {
			if !CheckAddress(v) {
				deleteProxy(v)
			} else {
				syncSpeed(v)
			}
		}(address)
	}
	wg.Wait()
	count, err = storage.Count()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.WithField("record", count).Info("After check")
}

func syncSpeed(address *model.Address) {
	addr := util.CombAddr(address)
	err := storage.Update(address)
	if err != nil {
		logrus.WithField("addr", addr).Warn(err)
	}
}

func deleteProxy(address *model.Address) {
	addr := util.CombAddr(address)
	err := storage.Delete(addr)
	if err != nil {
		logrus.WithField("addr", addr).Warn(err)
	}
}

func insertProxy(address *model.Address) {
	err := storage.Set(address)
	if err != nil {
		logrus.Error(err)
	}
}

// RandomOne returns a random address in storage.
func RandomOne() *model.Address {
	address, err := storage.GetRandOne()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return address
}

// RandomHTTPS returns a random address based on https protocol.
func RandomHTTPS() *model.Address {
	address, err := storage.GetRandHTTPS()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return address
}

// CountProxy returns the count of proxies in storage.
func CountProxy() int64 {
	count, err := storage.Count()
	if err != nil {
		logrus.Error(err)
		return -1
	}
	return count
}
