package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/cmatrixprobe/proxygool/spider/fetch"
	"github.com/cmatrixprobe/proxygool/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"strconv"
	"sync"
)

// Parse parses the request and send result to channel.
func Parse(request *model.Request, result chan<- *model.Address) {
	logrus.Info("Call " + request.WebName)
	var wg sync.WaitGroup
	for i := 1; i <= request.Pages; i++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()
			url := request.WebURL + strconv.Itoa(pageNum)
			logger := logrus.WithField("url", url)

			resp, err := fetch.Fetch(url)
			if err != nil {
				logger.Error(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				logger.WithField("StatusCode", resp.StatusCode).Error()
				return
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				logger.Error(err)
				return
			}

			// parse by css selector
			doc.Find(request.TrRegular).Each(func(i int, s *goquery.Selection) {
				td := s.Find("td")

				address := model.NewAddress()
				address.Host = td.Eq(request.HostIndex).Text()
				address.Port = td.Eq(request.PortIndex).Text()
				address.Origin = request.WebName
				address.Protocol = request.Protocol(td.Eq(request.ProtIndex).Text())

				if request.Trim {
					trim(address)
				}

				logger.WithField("address", address).Info("Parse finished")
				result <- address
			})
		}(i)
	}
	wg.Wait()
	logrus.Info(request.WebName + " done")
	return
}

func trim(trimStruct interface{}) {
	rType := reflect.TypeOf(trimStruct)
	rValue := reflect.ValueOf(trimStruct)
	if rType.Kind() == reflect.Ptr && rType.Elem().Kind() == reflect.Struct {
		rType = rType.Elem()
		rValue = rValue.Elem()
	} else {
		logrus.Panic("Must be ptr to struct.")
	}
	for i := 0; i < rType.NumField(); i++ {
		val := rValue.Field(i)
		if val.Kind() == reflect.String {
			str := val.Interface().(string)
			str = util.ReplaceSpecialChar(str)
			val.SetString(str)
		}
	}
}
