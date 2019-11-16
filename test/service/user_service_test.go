/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 19:41:42
 * @LastEditTime: 2019-09-04 14:10:51
 * @LastEditors: Please set LastEditors
 */
package service_test

import (
	"context"
	"gin-web/dao"
	"gin-web/dao/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMangGuo(b *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.0.103:27017"))
	if err != nil {
		println(err.Error())
		return
	}

	conn := client.Database("fuckyou").Collection("fuckfuck")
	res, err := conn.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	println(id)

}

func TestFindUser(b *testing.T) {

	user := &dao.User{Username: "sdttttt", Password: "7758258"}
	println(service.GetUserService().FindUser(user))

	println(user.ID)
	println(user.Username)
	println(user.Password)

}

func TestCreateUser(t *testing.T) {

	user := &dao.User{Username: "TestG", Password: "7758258"}
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

/*
	测试GDBC文本是否正确
*/
func TestGetGDBCString(t *testing.T) {
	//println(dao.ParseConnectParam())
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
