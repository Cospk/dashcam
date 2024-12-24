package system

import (
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
)

type UserService struct{}

// AddUser 添加用户
func (us *UserService) AddUser(instance *systemReq.AddUser) (err error) {
	var userModel systemModel.UserModel
	userModel.Username = instance.Username
	userModel.Password = instance.Password
	userModel.Phone = instance.Phone
	userModel.Email = instance.Email
	userModel.Active = instance.Active
	userModel.RoleModelID = instance.RoleModelID
	return global.DB.Create(&userModel).Error
}
