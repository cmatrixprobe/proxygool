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

func AddressUnMarshal(str string) (address *model.Address) {
	err := json.Unmarshal([]byte(str), address)
	if err != nil {
		logrus.WithField("address", str).Warn(err)
		return nil
	}
	return
}

func RandomElement(arr ...interface{}) interface{} {
	return arr[rand.New(rand.NewSource(time.Now().Unix())).Intn(len(arr))]
}
