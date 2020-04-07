package fetch

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/cmatrixprobe/proxygool/store"
	"github.com/cmatrixprobe/proxygool/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"time"
)

// Fetch downloads target pages to local.
func Fetch(webUrl string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// fetch by proxy
	if viper.GetBool("fetch.proxy") == true && store.CountProxy() > 0 {
		proxy := util.CombUrl(store.RandomOne())
		urlproxy, err := new(url.URL).Parse(proxy)
		if err != nil {
			logrus.WithField("proxy", proxy).Error(err)
		} else {
			logrus.WithField("proxy", proxy).Info("Fetch by proxy")
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		}
	}

	request, err := http.NewRequest("GET", webUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", browser.Computer())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
