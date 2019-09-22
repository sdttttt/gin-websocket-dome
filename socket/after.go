/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 21:44:20
 * @LastEditTime: 2019-09-17 18:24:39
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"gin-web/configuration"
	"strconv"
	"time"
)

/**
 * @description 这里编写所有的任务列队
	- Support
	  -目前任务列队只支持2种方式
		 1. 定时向用户发送数据型的任务
		 2. 定时面向用户连接操作型的任务
扩展方式：
	- 使用 AddAfterEnable 加入 定时器 And 处理器
	- 处理器的返回值必须实现FullMessage
*/

func TimeNoticeEnable() {
	AddAfterEnable(configuration.TimeNoticeInterval,
		func() *FullMessage {

			return &FullMessage{Username: "【服务器君】", Message: time.Now().String()}
		})
}

func ConnectCountNoticeEnable() {
	AddAfterEnable(configuration.ConnectCountNoticeInterval,
		func() *FullMessage {
			return &FullMessage{Username: "【服务器君】",
				Message: ("Crrent Count : " + strconv.Itoa(GetConnectHub().GetCrrentCount()))}
		})
}

func AddAfterEnable(s time.Duration, after func() *FullMessage) {
	GetConnectHub().AddAfter(s, after)
}
