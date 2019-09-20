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
		 2. 面向用户连接操作型的任务
扩展方式：
	通过 GetConnectHub 获取用户链接池对象
	使用 AddAfter 输入定时器 And Handler
*/

func TimeNoticeEnable() {
	GetConnectHub().AddAfter(configuration.TimeNoticeInterval,
		func() []byte {
			return ([]byte)(time.Now().String())
		})
}

func ConnectCountNoticeEnable() {
	GetConnectHub().AddAfter(configuration.ConnectCountNoticeInterval,
		func() []byte {
			return ([]byte)("Crrent Count : " + strconv.Itoa(GetConnectHub().GetCrrentCount()))
		})
}
