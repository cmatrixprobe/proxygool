package spider

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider/parse"
	"github.com/cmatrixprobe/proxygool/spider/site"
	"github.com/sirupsen/logrus"
	"sync"
)

var requests []*model.Request

func init() {
	requests = []*model.Request{
		site.Xici(),
		site.Kuai(),
		site.IP3366(),
		site.Qiyun(),
		//site.PLP(),
		//site.PLPSSL(),
		site.IP66(),
		site.IP89(),
	}
}

// Run spider
func Run(addressChan chan<- *model.Address, done chan<- bool) {
	logrus.Info("Crawl started.")

	var wg sync.WaitGroup
	for _, request := range requests {
		wg.Add(1)
		go func(req *model.Request) {
			defer wg.Done()
			parse.Parse(req, addressChan)
		}(request)
	}
	wg.Wait()

	logrus.Info("Crawl finished.")
	done <- true
}
