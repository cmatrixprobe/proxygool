package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

// Xici returns a request to www.xicidaili.com for proxies.
func Xici() *model.Request {
	req := model.NewRequest()
	req.WebName = "xicidaili"
	req.WebURL = "http://www.xicidaili.com/nn/"
	req.TrRegular = "#ip_list tbody tr"
	req.Pages = viper.GetInt("xicidaili.pages")
	req.HostIndex = 1
	req.PortIndex = 2
	req.ProtIndex = 5
	return req
}
