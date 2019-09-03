/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 21:44:20
 * @LastEditTime: 2019-09-02 23:09:43
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"gin-web/configuration"
	"strconv"
	"time"
)

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
