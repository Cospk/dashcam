package main

import "server/initialize"

func main() {
	initialize.InitViper()
	initialize.InitZap("debug")
	initialize.InitGorm()
	initialize.InitRedis()
	initialize.RunServer()
}
