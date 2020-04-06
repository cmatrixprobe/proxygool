package main

import (
	"github.com/cmatrixprobe/proxygool/api"
	"github.com/cmatrixprobe/proxygool/global"
	"github.com/cmatrixprobe/proxygool/store"
)



func main() {
	//addressChan := make(chan *model.Address, 1000)

	//logrus.SetReportCaller(true)

	// Start HTTP server
	go func() {
		api.Run()
	}()

	// Check proxies in DB
	go func() {
		store.CheckProxyDB(global.GetStore())
	}()

	select {}
}
