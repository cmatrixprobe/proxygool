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
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			store.CheckProxyDB()
			<-ticker.C
		}
	}()

	addressChan := make(chan *model.Address, 2000)
	// Check proxies in channel
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				if address, ok := <-addressChan; ok {
					store.ValidateProxy(address)
				}
			}
		}()
	}

	// Crawl proxies
	done := make(chan bool)
	tick := time.NewTicker(time.Minute)
	for {
		logrus.WithFields(logrus.Fields{
			"Channel": len(addressChan),
			"DB":      store.CountProxy(),
		}).Info()
		if len(addressChan) < cap(addressChan)>>1 {
			go spider.Run(addressChan, done)
		}
		select {
		case <-done:
		case <-tick.C:
		}
	}
}
