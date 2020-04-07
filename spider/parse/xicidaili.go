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
	xiciUrl   = "http://www.xicidaili.com/nn/"
	xiciPages = 1
)

func init() {
	xiciPages = viper.GetInt("xicidaili.pages")
}

// Xici gets proxies from xicidaili.com
func Xici() (result []*model.Address) {
	logrus.Info("Call Xici.")
	var wg sync.WaitGroup
	for i := 1; i <= xiciPages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			curUrl := xiciUrl + strconv.Itoa(pageNum)
			xiciLogger := logrus.WithField("url", curUrl)

			resp, err := fetch.Fetch(curUrl)
			if err != nil || resp.StatusCode != http.StatusOK {
				xiciLogger.WithField("StatusCode", resp.StatusCode).Error(err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				xiciLogger.Error(err)
				return
			}

			doc.Find("#ip_list tbody tr").Each(func(i int, s *goquery.Selection) {
				address := model.NewAddress()

				td := s.Find("td")
				address.Host = td.Eq(1).Text()
				address.Port = td.Eq(2).Text()
				address.Protocol = strings.ToLower(td.Eq(5).Text())
				address.Origin = "xicidaili"

				xiciLogger.WithField("address", address).Info()
				result = append(result, address)
			})
		}(i)
	}
	wg.Wait()
	logrus.Info("Xici done.")
	return
}
