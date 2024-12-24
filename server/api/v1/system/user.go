package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	systemReq "server/model/system/request"
)

type UserApi struct{}

// AddUser
// @Tags      UserApi
// @Summary   添加用户
// @Router    /user/addUser [post]
func (ua *UserApi) AddUser(c *gin.Context) {
	var addUser systemReq.AddUser
	if err := c.ShouldBindJSON(&addUser); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.AddUser(&addUser); err != nil {
		response.FailWithMessage("添加失败", c)
		global.Log.Error("添加失败", zap.Error(err))
	} else {
		response.Ok("", c)
	}
}
