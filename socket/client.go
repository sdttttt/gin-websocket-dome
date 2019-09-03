/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:19:09
 * @LastEditTime: 2019-09-03 18:24:17
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
	conn          *websocket.Conn
	messageBuffer chan []byte

	done chan struct{}
}

func (client *Client) reader() {

	defer func() {
		client.conn.Close()
		GetConnectHub().broadcast <- ([]byte)(client.conn.RemoteAddr().String() + "be away")
		GetConnectHub().unregister <- client
		println("reader下线")
	}()

	client.conn.SetReadLimit(configuration.MaxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
		return nil
	})

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			return
		}

		message = bytes.TrimSpace(message)
		GetConnectHub().broadcast <- message
	}

}

func (client *Client) writer() {

	ticker := time.NewTicker(configuration.PingTime)

	defer func() {
		ticker.Stop()
		println("writer下线")
		close(client.done)
	}()

	for {
		select {
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
