/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-01 15:58:00
 * @LastEditTime: 2019-09-02 15:59:02
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	"github.com/gorilla/websocket"
)

/**
 * @description 将发生错误的连接从连接池中删除
 * @param *websocket.Conn
 * @return
 */
func messageErrorProcessor(c *websocket.Conn, err error) bool {
	if err != nil {
		fmt.Println("Error Message : ", err)
		GetGroup().DelConnect(c)
		return true
	}
	return false
}

/**
 * @description

	About This Function
		我姑且将这个函数定义为 ProcessConnect 但是其实你们发现，
		我里面基本只对Connect的ReadBuffer的输出做了判断。
		原因如下：
			我测试多遍发现，大部分时候，无论是客户端意外关闭应用，还是
			主动关闭连接。Buffer都能接受到一个err Object，里面会含有错
			误码，不过这个错误码是混在别的字符中，我使用的Editor不支持我
			直接寻找Go的源码，我没有办法很好的提取出出这个错误码，庆幸的是
			err 的错误码并不含有复杂的情况，我们都可以草草的判断这个Connect
			已经处于断开状态
 * @param {@link *websocket.Conn}
*/
func ProcessConnect(c *websocket.Conn) {

	GetGroup().AddConnect(c)

	/**
	* 堵塞 and 监听
	 */
	for {
		messageType, message, err := c.ReadMessage()

		if messageErrorProcessor(c, err) {
			break
		}

		log.Println(time.Now(), " => ", messageType)

		/**
		*	字节转换文本
		 */
		decodeMessage := (*string)(unsafe.Pointer(&message))
		println(*decodeMessage)

		log.Println("Message :", *decodeMessage)

		if err = group.Broadcast(messageType, message); err != nil {
			fmt.Println("Message Writer Error :", err)
			break
		}
	}
}
