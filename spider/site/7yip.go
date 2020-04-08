package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

// Qiyun returns a request to www.7yip.cn for proxies.
func Qiyun() *model.Request {
	req := model.NewRequest()
	req.WebName = "7yip"
	req.WebURL = "https://www.7yip.cn/free/?action=china&page="
	req.TrRegular = ".table tbody tr"
	req.Pages = viper.GetInt("7yip.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.ProtIndex = 3
	return req
}
