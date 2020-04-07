package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider/fetch"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"sync"
)

var (
	plpUrl    = "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-"
	plpPages  = 1
	plpsUrl   = "https://list.proxylistplus.com/SSL-List-"
	plpsPages = 1
)

func init() {
	plpPages = viper.GetInt("proxylistplus.http.pages")
	plpsPages = viper.GetInt("proxylistplus.https.pages")
}

func PLP() (result []*model.Address) {
	logrus.Info("Call PLP.")
	var wg sync.WaitGroup
	for i := 1; i <= plpPages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			curUrl := plpUrl + strconv.Itoa(pageNum)
			PLPLogger := logrus.WithField("url", curUrl)

			resp, err := fetch.Fetch(curUrl)
			if err != nil {
				PLPLogger.Error(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				PLPLogger.WithField("StatusCode", resp.StatusCode).Error()
				return
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				PLPLogger.Error(err)
				return
			}

			doc.Find(".bg tr.cells").Each(func(i int, s *goquery.Selection) {
				address := model.NewAddress()

				td := s.Find("td")
				address.Host = td.Eq(1).Text()
				address.Port = td.Eq(2).Text()
				address.Protocol = "http"
				address.Origin = "proxylistplus"

				PLPLogger.WithField("address", address).Info()
				result = append(result, address)
			})
		}(i)
	}
	wg.Wait()
	logrus.Info("PLP done.")
	return
}

func PLPS() (result []*model.Address) {
	logrus.Info("Call PLPS.")
	var wg sync.WaitGroup
	for i := 1; i <= plpsPages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			curUrl := plpsUrl + strconv.Itoa(pageNum)
			PLPLogger := logrus.WithField("url", curUrl)

			resp, err := fetch.Fetch(curUrl)
			if err != nil || resp.StatusCode != http.StatusOK {
				PLPLogger.WithField("StatusCode", resp.StatusCode).Error(err)
				return
			}
			defer resp.Body.Close()

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				PLPLogger.Error(err)
				return
			}

			doc.Find(".bg tr.cells").Each(func(i int, s *goquery.Selection) {
				address := model.NewAddress()

				td := s.Find("td")
				address.Host = td.Eq(1).Text()
				address.Port = td.Eq(2).Text()
				address.Protocol = "https"
				address.Origin = "proxylistplus"

				PLPLogger.WithField("address", address).Info()
				result = append(result, address)
			})
		}(i)
	}
	wg.Wait()
	logrus.Info("PLPS done.")
	return
}