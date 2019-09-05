/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 19:41:42
 * @LastEditTime: 2019-09-04 14:10:51
 * @LastEditors: Please set LastEditors
 */
package service_test

import (
	"gin-web/dao"
	"gin-web/dao/service"
	"testing"
	"time"
)

func TestFindUser(b *testing.T) {

	user := &dao.User{Username: "sdttttt", Password: "7758258"}
	println(service.GetUserService().FindUser(user))

	println(user.ID)
	println(user.Username)
	println(user.Password)

}

func TestCreateUser(t *testing.T) {

	user := &dao.User{Username: "sdttttt", Password: "7758258"}
	println(service.GetUserService().CreateUser(user))

	println(user.ID)
	println(user.Username == "")
	println(user.Password)
}

func TestFindUserCount(t *testing.T) {

	var count int
	dao.GetDbInstance().Model(&dao.User{}).Where("username = ?", "sdttttt").Count(&count)

	println(count > 0)
}

func BenchmarkFindUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &dao.User{Username: "Test", Password: "7758258", CreateTime: time.Now(), UpdateTime: time.Now()}
		println(service.GetUserService().CreateUser(user))

		println(user.ID)
		println(user.Username)
		println(user.Password)
	}
}
