/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 11:19:17
 * @LastEditTime: 2019-09-03 17:12:16
 * @LastEditors: Please set LastEditors
 */
package service

import (
	"gin-web/dao"
	"gin-web/util"
	"sync"

	"github.com/jinzhu/gorm"
)

type UserService interface {
	CreateUser(*dao.User) bool

	FindUser(*dao.User) *dao.User
}

type IUserService struct {
	DbConnect *gorm.DB
}

var aIUserService *IUserService

var userServiceOnce sync.Once

func GetUserService() *IUserService {
	userServiceOnce.Do(func() {
		aIUserService = &IUserService{DbConnect: dao.GetDbInstance()}
	})

	return aIUserService
}

func (service *IUserService) CreateUser(user *dao.User) bool {
	user.Password = util.ToSha1(user.Password)
	if service.DbConnect.Create(user); user.ID > 0 {
		return true
	}

	return false
}

func (service *IUserService) FindUser(user *dao.User) bool {
	user.Password = util.ToSha1(user.Password)
	service.DbConnect.Find(user)
	if user.Password == "" || user.Username == "" || user.ID == 0 {
		return false
	}
	return true
}
