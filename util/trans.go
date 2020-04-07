package util

import (
	"encoding/json"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

// CombAddr
func CombAddr(address *model.Address) string {
	return address.Host + ":" + address.Port
}

// CombUrl
func CombUrl(address *model.Address) string {
	return address.Protocol + "://" + CombAddr(address)
}

// AddressMarshal
func AddressMarshal(address *model.Address) string {
	bytes, err := json.Marshal(address)
	if err != nil {
		logrus.WithField("address", address).Warn(err)
		return ""
	}
	return string(bytes)
}

// AddressUnMarshal
func AddressUnMarshal(str string) *model.Address {
	address := model.NewAddress()
	err := json.Unmarshal([]byte(str), address)
	if err != nil {
		logrus.WithField("address", str).Warn(err)
		return nil
	}
	return address
}

// RandomElement
func RandomElement(addresses []*model.Address) *model.Address {
	addrLen := len(addresses)
	if addrLen <= 0 {
		return nil
	}
	return addresses[rand.New(rand.NewSource(time.Now().Unix())).Intn(addrLen)]
}

func ReplaceSpecialChar(s string) string {
	r := strings.NewReplacer("\n", "", "\t", "")
	return r.Replace(s)
}
