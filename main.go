package main

import (
	"github.com/cmatrixprobe/proxygool/cache"
	_ "github.com/cmatrixprobe/proxygool/config"
	"log"
)

func main() {
	c := cache.Pool().Get()
	reply, err := c.Do("PING")
	if err != nil {
		panic(err)
	}
	log.Println(reply)
}
