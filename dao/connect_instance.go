/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 12:24:20
 * @LastEditTime: 2019-09-03 13:18:04
 * @LastEditors: Please set LastEditors
 */
package dao

import (
	"gin-web/configuration"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDbInstance() *gorm.DB {
	db, err := gorm.Open(configuration.DataBaseType, configuration.GDBC)

	if err != nil {
		log.Println(err)
	}

	return db
}
