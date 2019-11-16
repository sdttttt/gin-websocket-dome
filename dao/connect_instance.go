/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 12:24:20
 * @LastEditTime: 2019-09-04 17:43:57
 * @LastEditors: Please set LastEditors
 */
package dao

import (
	"gin-web/configuration"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * TODO:
	- [x] Able Work ?
*/
func parseConnectParam() string {

	GDBC := configuration.DB_Username + ":"
	GDBC += configuration.DB_Password
	GDBC += "@" + configuration.DB_Protocol
	GDBC += "(" + configuration.DB_Address + ":" + configuration.DB_Post + ")"
	GDBC += "/" + configuration.DB_Name + "?" + "charset=" + configuration.DB_Charset
	if configuration.DB_ParseTime {
		GDBC += "&parseTime=True"
	} else {
		GDBC += "&parseTime=False"
	}

	GDBC += "&loc=Local"
	return GDBC
}

/**
 * TODO:
	- [x] Able Work ?
 * @description get a DataBaseConnest
*/
func GetDbInstance() *gorm.DB {
	db, err := gorm.Open(configuration.DataBaseType, parseConnectParam())

	if err != nil {
		log.Println(err)
	}

	return db
}
