/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 13:16:23
 * @LastEditTime: 2019-09-02 17:17:01
 * @LastEditors: Please set LastEditors
 */
package configuration

import (
	"time"
)

const (

	//报时间隔
	TimeNoticeInterval = 30 * time.Second

	//当前人数通知
	ConnectCountNoticeInterval = 30 * time.Second

	// ping
	PingTime = 5 * time.Second

	//pong
	PongTime = 45 * time.Second
)
