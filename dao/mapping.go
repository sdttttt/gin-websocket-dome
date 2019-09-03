/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 11:10:33
 * @LastEditTime: 2019-09-03 20:24:45
 * @LastEditors: Please set LastEditors
 */
package dao

import (
	"time"
)

/** @description User table Mappping
 */
type User struct {
	ID int `gorm:"column:id;primary_key"`

	Username string `gorm:"type:varchar(255);column:username;unique"`

	Password string `gorm:"type:varchar(255);column:password"`

	CreateTime time.Time

	UpdateTime time.Time
}

func (User) TableName() string {
	return "user"
}
