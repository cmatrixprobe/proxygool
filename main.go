package main

import (
	"github.com/cmatrixprobe/proxygool/api"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider"
	"github.com/cmatrixprobe/proxygool/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

func init() {
	//logrus.SetReportCaller(true)
	logFile := viper.GetString("logger.filename")
	logLevel := viper.GetUint32("logger.level")

	logrus.SetLevel(logrus.Level(logLevel))
	if viper.GetBool("docker") == true {
		logrus.SetLevel(logrus.FatalLevel)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, file))
}

func main() {
	// start HTTP server
	go func() {
		api.Run()
	}()

	// check proxies in DB
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			store.CheckProxyDB()
			<-ticker.C
		}
	}()

	addressChan := make(chan *model.Address, 2000)
	// check proxies in channel
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				if address, ok := <-addressChan; ok {
					store.ValidateProxy(address)
				}
			}
		}()
	}

	// crawl proxies
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
