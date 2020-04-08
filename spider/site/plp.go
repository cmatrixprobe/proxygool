package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

// PLP returns a request to list.proxylistplus.com for http proxies.
func PLP() *model.Request {
	req := model.NewRequest()
	req.WebName = "proxylistplus"
	req.WebURL = "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-"
	req.TrRegular = ".bg tr.cells"
	req.Pages = viper.GetInt("plphttp.pages")
	req.HostIndex = 1
	req.PortIndex = 2
	req.Protocol = func(s string) string {
		return "http"
	}
	return req
}

// PLPSSL returns a request to list.proxylistplus.com for https proxies.
func PLPSSL() *model.Request {
	req := model.NewRequest()
	req.WebName = "proxylistplus"
	req.WebURL = "https://list.proxylistplus.com/SSL-List-"
	req.TrRegular = ".bg tr.cells"
	req.Pages = viper.GetInt("plphttps.pages")
	req.HostIndex = 1
	req.PortIndex = 2
	req.Protocol = func(s string) string {
		return "https"
	}
	return req
}
