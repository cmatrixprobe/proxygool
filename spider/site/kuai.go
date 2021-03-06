package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

// Kuai returns a request to www.kuaidaili.com for proxies.
func Kuai() *model.Request {
	req := model.NewRequest()
	req.WebName = "kuaidaili"
	req.WebURL = "http://www.kuaidaili.com/free/inha/"
	req.TrRegular = ".table tbody tr"
	req.Pages = viper.GetInt("kuaidaili.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.ProtIndex = 3
	return req
}
