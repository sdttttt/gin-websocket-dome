/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-01 15:58:00
 * @LastEditTime: 2019-09-01 17:45:36
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/gorilla/websocket"
)

func messageErrorProcessor(c *websocket.Conn, err error) bool {
	if err != nil {
		fmt.Println("Error Message : ", err)
		GetGroup().DelConnect(c)
		return true
	}
	return false
}

func ProcessConnect(c *websocket.Conn) {

	GetGroup().AddConnect(c)

	defer c.Close()

	/**
	* 堵塞 and 监听
	 */
	for {
		messageType, message, err := c.ReadMessage()

		if messageErrorProcessor(c, err) {
			break
		}

		println(messageType)

		if messageType == websocket.CloseMessage {
			println("有人要走了！！")
		}

		/**
		*	字节转换文本
		 */
		decodeMessage := (*string)(unsafe.Pointer(&message))
		println(*decodeMessage)

		log.Println("Message -> ", *decodeMessage)

		if err = group.Broadcast(messageType, message); err != nil {
			fmt.Println("Message Writer Error ->", err)
			break
		}
	}
}
