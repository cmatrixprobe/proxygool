package site

import (
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/spf13/viper"
)

func IP3366() *model.Request {
	req := model.NewRequest()
	req.WebName = "ip3366"
	req.WebURL = "http://www.ip3366.net/free/?stype=1&page="
	req.TrRegular = ".table tbody tr"
	req.Pages = viper.GetInt("ip3366.pages")
	req.HostIndex = 0
	req.PortIndex = 1
	req.ProtIndex = 3
	return req
}
