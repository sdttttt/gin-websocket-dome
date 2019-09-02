/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 13:16:23
 * @LastEditTime: 2019-09-02 21:57:11
 * @LastEditors: Please set LastEditors
 */
package configuration

import (
	"time"
)

const (
	WriterWait = 60 * time.Second

	MaxMessageSize = 8192

	//报时间隔
	TimeNoticeInterval = 1 * time.Second

	//当前人数通知
	ConnectCountNoticeInterval = 5 * time.Second

	// ping
	PingTime = 40 * time.Second

	//pong
	PongTime = 45 * time.Second
)
