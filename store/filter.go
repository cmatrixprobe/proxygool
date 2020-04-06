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

// CheckProxy
func CheckProxy(store Store, address *model.Address) {
	if CheckAddress(address) {
		InsertProxy(store, address)
	}
}

// CheckAddress checks the address work or not.
func CheckAddress(address *model.Address) bool {
	var testAddr, targetUrl string

	// test by speedtest.cn
	switch address.Protocol {
	case "https":
		testAddr = "https://" + util.CombAddr(address.Host, address.Port)
		targetUrl = "https://forge.speedtest.cn/api/location/geo?ip=" + address.Host
	case "http":
		testAddr = "http://" + util.CombAddr(address.Host, address.Port)
		targetUrl = "http://forge.speedtest.cn/api/location/geo?ip=" + address.Host
	}

	reqLogger := logrus.WithFields(logrus.Fields{
		"testAddr":  testAddr,
		"targetUrl": targetUrl,
	})
	reqLogger.Info()

	begin := time.Now()
	// get targetUrl by test proxy
	resp, _, errs := gorequest.New().Proxy(testAddr).Get(targetUrl).End()
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
	//_, err := simplejson.NewFromReader(resp.Body)
	//if err != nil {
	//	reqLogger.Error(err)
	//	return false
	//}
	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		reqLogger.Warn(err)
		return false
	}
	var tr TestResponse
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

// CheckProxyDB checks proxy addresses in DB
func CheckProxyDB(store Store) {
	addresses, err := store.GetAll()
	if err != nil {
		logrus.Warn(err)
		return
	}

	// check and delete unavailable records
	count, err := store.Count()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.WithField("record", count).Info("Before check")
	var wg sync.WaitGroup
	for _, address := range addresses {
		wg.Add(1)
		go func(v *model.Address) {
			if !CheckAddress(v) {
				DeleteProxy(store, v)
			} else {
				SyncSpeed(store, v)
			}
		}(address)
	}
	wg.Wait()
	count, err = store.Count()
	if err != nil {
		logrus.Warn(err)
	}
	logrus.WithField("record", count).Info("After check")
}

// SyncSpeed
func SyncSpeed(store Store, address *model.Address) {
	addr := util.CombAddr(address.Host, address.Port)
	err := store.Update(address)
	if err != nil {
		logrus.WithField("addr", addr).Warn(err)
	}
}

// DeleteProxy
func DeleteProxy(store Store, address *model.Address) {
	addr := util.CombAddr(address.Host, address.Port)
	err := store.Delete(addr)
	if err != nil {
		logrus.WithField("addr", addr).Warn(err)
	}
}

// RandomHttp
func RandomOne(store Store) *model.Address {
	address, err := store.GetRandOne()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return address
}

// RandomHttps
func RandomHttps(store Store) *model.Address {
	address, err := store.GetRandHttps()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return address
}

// InsertProxy
func InsertProxy(store Store, address *model.Address) {
	err := store.Set(address)
	if err != nil {
		logrus.Error(err)
	}
}
