package util

import (
	"encoding/json"
	"github.com/cmatrixprobe/proxygool/model"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

// CombAddr converts address to socket.
func CombAddr(address *model.Address) string {
	return address.Host + ":" + address.Port
}

// CombURL converts address to url.
func CombURL(address *model.Address) string {
	return address.Protocol + "://" + CombAddr(address)
}

// AddressMarshal returns the JSON encoding of address.
func AddressMarshal(address *model.Address) string {
	bytes, err := json.Marshal(address)
	if err != nil {
		logrus.WithField("address", address).Warn(err)
		return ""
	}
	return string(bytes)
}

// AddressUnMarshal parses the JSON-encoded data and stores the result.
func AddressUnMarshal(str string) *model.Address {
	address := model.NewAddress()
	err := json.Unmarshal([]byte(str), address)
	if err != nil {
		logrus.WithField("address", str).Warn(err)
		return nil
	}
	return address
}

// RandomElement returns a random element in address slice.
func RandomElement(addresses []*model.Address) *model.Address {
	addrLen := len(addresses)
	if addrLen <= 0 {
		return nil
	}
	return addresses[rand.New(rand.NewSource(time.Now().Unix())).Intn(addrLen)]
}

// ReplaceSpecialChar removes "\n" and "\t" in string s.
func ReplaceSpecialChar(s string) string {
	r := strings.NewReplacer("\n", "", "\t", "")
	return r.Replace(s)
}
