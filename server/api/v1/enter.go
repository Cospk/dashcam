package v1

import "server/api/v1/system"

type ApiGroup struct {
	System system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
