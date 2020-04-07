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
	ip3366Url   = "http://www.ip3366.net/free/?stype=1&page="
	ip3366Pages = 1
)

func init() {
	ip3366Pages = viper.GetInt("ip3366.pages")
}

func IP3366() (result []*model.Address) {
	logrus.Info("Call IP3366.")
	var wg sync.WaitGroup
	for i := 1; i <= ip3366Pages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			curUrl := ip3366Url + strconv.Itoa(pageNum)
			ip3366Logger := logrus.WithField("url", curUrl)

			resp, err := fetch.Fetch(curUrl)
			if err != nil || resp.StatusCode != http.StatusOK {
				ip3366Logger.WithField("StatusCode", resp.StatusCode).Error(err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				ip3366Logger.Error(err)
				return
			}

			doc.Find(".table tbody tr").Each(func(i int, s *goquery.Selection) {
				address := model.NewAddress()

				td := s.Find("td")
				address.Host = td.Eq(0).Text()
				address.Port = td.Eq(1).Text()
				address.Protocol = strings.ToLower(td.Eq(3).Text())
				address.Origin = "ip3366"

				ip3366Logger.WithField("address", address).Info()
				result = append(result, address)
			})
		}(i)
	}
	wg.Wait()
	logrus.Info("IP3366 done.")
	return
}
