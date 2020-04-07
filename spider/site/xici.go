package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

func Xici() *model.Request {
	req := model.NewRequest()
	req.WebName = "xicidaili"
	req.WebURL = "http://www.xicidaili.com/nn/"
	req.TrRegexp = "#ip_list tbody tr"
	req.Pages = viper.GetInt("xicidaili.pages")
	req.HostIndex = 1
	req.PortIndex = 2
	req.ProtIndex = 5
	return req
}
