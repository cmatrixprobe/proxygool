package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

func IP66() *model.Request {
	req := model.NewRequest()
	req.WebName = "66ip"
	req.WebURL = "http://www.66ip.cn/"
	req.TrRegexp = "table:last-child tbody tr:nth-child(n+2)"
	req.Pages = viper.GetInt("66ip.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.Protocol = func(s string) string {
		return "http"
	}
	return req
}
