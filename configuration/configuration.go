/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 13:16:23
 * @LastEditTime: 2019-09-09 22:50:52
 * @LastEditors: Please set LastEditors
 */
package configuration

import (
	"time"
)

const (

	/**
	 * ****************************************
	 * 	Socket 配置
	 * *****************************************
	 */
	WriterWait = 60 * time.Second

	MaxMessageSize = 8192

	//报时间隔
	TimeNoticeInterval = 30 * time.Second

	//当前人数通知
	ConnectCountNoticeInterval = 30 * time.Second

	// ping
	PingTime = 40 * time.Second

	//pong
	PongTime = 45 * time.Second

	/**
	 * ****************************************
	 *
	 * 	Database 配置
	 * ********************************************
	 */

	DataBaseType = "mysql"

	DB_Address = "192.168.0.103"

	DB_Post = "3306"

	DB_Protocol = "tcp"

	DB_Name = "jojo"

	DB_Charset = "utf8"

	DB_ParseTime = true

	DB_Username = "root"

	DB_Password = "root"

	//
	GDBC = DB_Username + ":" + DB_Password + "@tcp(192.168.0.103:3306)/jojo?charset=utf8&parseTime=True&loc=Local"
)
