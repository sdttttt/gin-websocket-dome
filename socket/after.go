/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 21:44:20
 * @LastEditTime: 2019-09-02 22:12:14
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"fmt"
	"gin-web/configuration"
	"time"
)

func TimeNoticeEnable() {
	GetConnectHub().AddAfter(configuration.TimeNoticeInterval, time.Now().String())
}

func ConnectCountNoticeEnable() {
	GetConnectHub().AddAfter(configuration.ConnectCountNoticeInterval,
		fmt.Sprintf("%d", GetConnectHub().GetCrrentCount()))
}
