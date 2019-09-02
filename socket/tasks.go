/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 22:46:55
 * @LastEditTime: 2019-09-02 17:21:28
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"gin-web/configuration"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

/** @description 为所有的客户端连接定期报时
 */
func TimeNoticeEnable() {
	log.Println("TimeNotice Load ...")
	GetGroup().AddTickerTask(configuration.TimeNoticeInterval, func(c *websocket.Conn) {
		c.WriteMessage(websocket.TextMessage, ([]byte)(time.Now().String()))
	})

}

func ConnectCountNiticeEnable() {
	log.Println("ConnectCountNitice Load ...")
	GetGroup().AddTickerTask(configuration.ConnectCountNoticeInterval, func(c *websocket.Conn) {
		c.WriteMessage(websocket.TextMessage, ([]byte)("当前人数"+strconv.Itoa(GetGroup().GetConnCount())))
	})
}

func KeepAliveEnable() {

}

func RunAllTask() {
	GetGroup().ExecuteTasks()
}
