package service

import "server/service/system"

type ServiceGroup struct {
	System system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
