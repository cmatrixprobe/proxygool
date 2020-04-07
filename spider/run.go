package spider

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider/parse"
	"github.com/sirupsen/logrus"
	"sync"
)

func Run(addressChan chan<- *model.Address) {
	logrus.Info("Crawl started.")
	funs := []func() []*model.Address{
		//parse.Xici,
		parse.Kuai,
		parse.PLP,
		parse.PLPS,
		parse.IP3366,
	}

	var wg sync.WaitGroup
	for _, fun := range funs {
		wg.Add(1)
		go func(f func() []*model.Address) {
			temp := f()
			for _, v := range temp {
				addressChan <- v
			}
			wg.Done()
		}(fun)
	}
	wg.Wait()

	logrus.Info("Crawl finished.")
}
