package util

import (
	"encoding/json"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func CombAddr(host, port string) string {
	return host + ":" + port
}

func AddressMarshal(address *model.Address) string {
	bytes, err := json.Marshal(address)
	if err != nil {
		logrus.WithField("address", address).Warn(err)
		return ""
	}
	return string(bytes)
}

func AddressUnMarshal(str string) *model.Address {
	address := model.NewAddress()
	err := json.Unmarshal([]byte(str), address)
	if err != nil {
		logrus.WithField("address", str).Warn(err)
		return nil
	}
	return address
}

func RandomElement(addresses []*model.Address) *model.Address {
	addrLen := len(addresses)
	if addrLen <= 0 {
		return nil
	}
	return addresses[rand.New(rand.NewSource(time.Now().Unix())).Intn(addrLen)]
}
