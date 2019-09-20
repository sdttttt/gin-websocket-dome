/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:19:09
 * @LastEditTime: 2019-09-20 15:34:32
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"bytes"
	"gin-web/configuration"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {

	/**
	* 客户端链接
	 */
	conn *websocket.Conn

	/**
	* 消息缓冲区
	 */
	messageBuffer chan []byte

	/**
	* 客户端结束接受信道
	 */
	done chan struct{}
}

/**
 * @description 读取器
	用户IO中存在数据，Reader将数据读出
	并broadcast
*/
func (client *Client) reader() {

	defer func() {
		client.conn.Close()
		GetConnectHub().broadcast <- ([]byte)(client.conn.RemoteAddr().String() + "be away")
		GetConnectHub().unregister <- client
		println("reader下线")
	}()

	// 消息最大Byte
	client.conn.SetReadLimit(configuration.MaxMessageSize)

	//读超时
	client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
		return nil
	})

	// 监听
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			return
		}

		message = bytes.TrimSpace(message)
		GetConnectHub().broadcast <- message
	}

}

/**
 * @description 写入器
	外部消息会从这里进入
*/
func (client *Client) writer() {

	ticker := time.NewTicker(configuration.PingTime)

	defer func() {
		ticker.Stop()
		println("writer下线")
		close(client.done)
	}()

	for {
		select {
		// 监听用户缓冲区，每个用户的缓冲区由ConnectHub管理
		case message, ok := <-client.messageBuffer:
			client.conn.SetWriteDeadline(time.Now().Add(configuration.WriterWait))

			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := client.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			writer.Write(message)
			writer.Close()

		// 检查用户是否已经发生写超时
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(configuration.WriterWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}

/**
 * @description 客户端处理机制
 */
func ProcessConnect(c *websocket.Conn) {

	client := &Client{
		conn:          c,
		messageBuffer: make(chan []byte),
		done:          make(chan struct{}),
	}

	GetConnectHub().register <- client

	go client.reader()
	go client.writer()

	<-client.done

	println("the end")
}
