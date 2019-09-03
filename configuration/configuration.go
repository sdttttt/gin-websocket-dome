/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 13:16:23
 * @LastEditTime: 2019-09-03 13:18:59
 * @LastEditors: Please set LastEditors
 */
package configuration

import (
	"time"
)

const (

	/**
	 * ****************************************
	 *
	 * 	Socket 配置
	 * ********************************************
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

	GDBC = "root:root@tcp(192.168.0.104:3306)/jojo?charset=utf8&parseTime=True&loc=Local"
)
