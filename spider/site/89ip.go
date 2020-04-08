package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

// IP89 returns a request to www.89ip.cn for proxies.
func IP89() *model.Request {
	req := model.NewRequest()
	req.WebName = "89ip"
	req.WebURL = "http://www.89ip.cn/index_"
	req.TrRegular = ".layui-table tbody tr"
	req.Pages = viper.GetInt("89ip.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.Trim = true
	req.Protocol = func(s string) string {
		return "http"
	}
	return req
}
