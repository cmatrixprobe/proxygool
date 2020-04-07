package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider/fetch"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	kuaiUrl   = "http://www.kuaidaili.com/free/inha/"
	kuaiPages = 1
)

func init() {
	kuaiPages = viper.GetInt("kuaidaili.pages")
}

func Kuai() (result []*model.Address) {
	logrus.Info("Call Kuai.")
	var wg sync.WaitGroup
	for i := 1; i <= kuaiPages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			curUrl := kuaiUrl + strconv.Itoa(pageNum)
			kuaiLogger := logrus.WithField("url", curUrl)

			resp, err := fetch.Fetch(curUrl)
			if err != nil || resp.StatusCode != http.StatusOK {
				kuaiLogger.WithField("StatusCode", resp.StatusCode).Error(err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				kuaiLogger.Error(err)
				return
			}

			doc.Find(".table tbody tr").Each(func(i int, s *goquery.Selection) {
				address := model.NewAddress()

				td := s.Find("td")
				address.Host = td.Eq(0).Text()
				address.Port = td.Eq(1).Text()
				address.Protocol = strings.ToLower(td.Eq(3).Text())
				address.Origin = "kuaidaili"

				kuaiLogger.WithField("address", address).Info()
				result = append(result, address)
			})
		}(i)
	}
	wg.Wait()
	logrus.Info("Kuai done.")
	return
}
