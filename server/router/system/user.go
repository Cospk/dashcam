package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct {
}

func (router *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")

	userApi := v1.ApiGroupApp.System.UserApi
	{
		userRouter.POST("addUser", userApi.AddUser)
	}
}
