package main

import (
	// "go_fabric/cache"
	"go_fabric/conf"
	"go_fabric/routes"
)

func main() {
	conf.Init()
	// cache.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
