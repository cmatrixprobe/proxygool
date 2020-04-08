package api

import (
	"github.com/cmatrixprobe/proxygool/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

// Run HTTP server.
func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", RandomProxyHandler)
	r.GET("/https", HTTPSProxyHandler)

	server := viper.Sub("server")
	ip := server.GetString("host")
	port := server.GetString("port")
	address := ip + ":" + port

	logrus.Fatal(r.Run(address))
}

// RandomProxyHandler returns a random JSON-encoded address.
func RandomProxyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, store.RandomOne())
}

// HTTPSProxyHandler returns a random JSON-encoded https protocol address.
func HTTPSProxyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, store.RandomHTTPS())
}
