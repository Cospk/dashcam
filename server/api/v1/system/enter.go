package system

import "server/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.System
)
