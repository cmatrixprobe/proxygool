package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

func PLP() *model.Request {
	req := model.NewRequest()
	req.WebName = "proxylistplus"
	req.WebURL = "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-"
	req.TrRegexp = ".bg tr.cells"
	req.Pages = viper.GetInt("plphttp.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.Protocol = func(s string) string {
		return "http"
	}
	return req
}

func PLPSSL() *model.Request {
	req := model.NewRequest()
	req.WebName = "proxylistplus"
	req.WebURL = "https://list.proxylistplus.com/SSL-List-"
	req.TrRegexp = ".bg tr.cells"
	req.Pages = viper.GetInt("plphttps.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.Protocol = func(s string) string {
		return "https"
	}
	return req
}
