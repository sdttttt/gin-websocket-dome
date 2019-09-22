/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 11:19:17
 * @LastEditTime: 2019-09-05 17:46:04
 * @LastEditors: Please set LastEditors
 */
package service

import (
	"gin-web/dao"
	"gin-web/util"
	"sync"
	"time"

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

/*
	A Instance
*/
func GetUserService() *IUserService {
	userServiceOnce.Do(func() {
		aIUserService = &IUserService{DbConnect: dao.GetDbInstance()}
	})

	return aIUserService
}

func (service *IUserService) ExistsUser(username string) bool {
	var count int
	service.DbConnect.Model(&dao.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

/**
 * @test OK
 */
func (service *IUserService) CreateUser(user *dao.User) bool {
	if service.ExistsUser(user.Username) {
		user.Username = ""
		return false
	}

	user.Password = util.ToSha1(user.Password)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if service.DbConnect.Create(user); user.ID > 0 {
		return true
	}

	return false
}

/**
 * @test OK
 */
func (service *IUserService) FindUser(user *dao.User) bool {
	user.Password = util.ToSha1(user.Password)
	service.DbConnect.Where("username = ? AND password = ?", user.Username, user.Password).Find(user)
	if user.Password == "" || user.Username == "" || user.ID == 0 {
		return false
	}
	return true
}
