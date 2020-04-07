package main

import (
	"github.com/cmatrixprobe/proxygool/api"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider"
	"github.com/cmatrixprobe/proxygool/store"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	//logrus.SetReportCaller(true)
	//logrus.SetLevel(logrus.WarnLevel)

	// Start HTTP server
	go func() {
		api.Run()
	}()

	// Check proxies in DB
	ticker1 := time.NewTicker(time.Minute)
	go func() {
		for {
			store.CheckProxyDB()
			<-ticker1.C
		}
	}()

	addressChan := make(chan *model.Address, 1000)
	// Check proxies in channel
	for i := 0; i < 1000; i++ {
		go func() {
			for address := range addressChan {
				store.ValidateProxy(address)
			}
		}()
	}

	// Crawl proxies
	ticker2 := time.NewTicker(time.Minute)
	for {
		logrus.WithFields(logrus.Fields{
			"Channel": len(addressChan),
			"DB":      store.CountProxy(),
		}).Info()
		if len(addressChan) < cap(addressChan)>>1 {
			go spider.Run(addressChan)
		}
		<-ticker2.C
	}
}
