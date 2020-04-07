package fetch

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"time"
)

var transport *http.Transport

func init() {
	proxy := viper.GetString("fetch.proxy")
	if proxy != "" {
		urlproxy, err := new(url.URL).Parse(proxy)
		if err != nil {
			logrus.WithField("proxy", proxy).Error(err)
		}
		transport.Proxy = http.ProxyURL(urlproxy)
	}
}

// Fetch returns target pages
func Fetch(webUrl string) (*http.Response, error) {
	client := &http.Client{
		//Transport: transport,
		Timeout:   time.Second * 10,
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
