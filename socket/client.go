/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:19:09
 * @LastEditTime: 2019-09-02 20:50:25
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"bytes"
	"gin-web/configuration"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn          *websocket.Conn
	messageBuffer chan []byte
}

func (client *Client) reader() {

	defer func() {
		client.conn.Close()
	}()

	client.conn.SetReadLimit(configuration.MaxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(configuration.PongTime))
		return nil
	})

	for {
		_, message, err := client.conn.ReadMessage()
		log.Println(message)
		if err != nil {
			return
		}

		message = bytes.TrimSpace(message)
		log.Println(message)
		GetConnectHub().broadcast <- message
	}

}

func (client *Client) writer() {

	ticker := time.NewTicker(configuration.PingTime)

	defer func() {
		ticker.Stop()
		client.conn.Close()
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

func ProcessConnect(c *websocket.Conn) {

	client := &Client{
		conn:          c,
		messageBuffer: make(chan []byte),
	}

	GetConnectHub().register <- client

	log.Println("启动玄冥二老")
	go client.reader()
	go client.writer()

	select {}
}
