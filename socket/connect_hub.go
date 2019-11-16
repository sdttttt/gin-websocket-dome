/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:18:55
 * @LastEditTime: 2019-09-17 18:26:55
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"log"
	"sync"
	"time"
)

type Task map[*time.Ticker]func(*Client)

type PushMessage map[*time.Ticker]func() *FullMessage

type ConnectHub struct {
	clients map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan []byte

	tasks Task

	afters PushMessage
}

var connectHub *ConnectHub

var once sync.Once

/**
 * @description 客户端连接池引用
	一个Application只会有一个连接池，如果需要开放多个连接池
	需要同时开启多个Application
*/
func GetConnectHub() *ConnectHub {
	once.Do(
		func() {
			connectHub = &ConnectHub{
				clients:    make(map[*Client]bool),
				register:   make(chan *Client),
				unregister: make(chan *Client),
				broadcast:  make(chan []byte),
				tasks:      make(Task),
				afters:     make(PushMessage),
			}
		})

	return connectHub
}

/**
 * @description 启动Websocket服务中心
 */
func (hub *ConnectHub) RunAndListen() {

	log.Println("Hub Starting ...")

	hub.LoadTask()
	hub.ExecuteAfter()

	log.Println("Hub OK...")

	// 服务端启动监听
	for {
		select {

		case client := <-hub.register:
			hub.clients[client] = true
			log.Println("人数 : ", hub.GetCrrentCount())

		case client := <-hub.unregister:
			delete(hub.clients, client)

		case message := <-hub.broadcast:
			for client, status := range hub.clients {
				if status {
					client.messageBuffer <- message
				} else {
					close(client.messageBuffer)
					delete(hub.clients, client)
				}
			}
		}
	}
}

/**
 * @description 获取当前所有客户端的数量
 */
func (hub *ConnectHub) GetCrrentCount() int {
	return len(hub.clients)
}

/**
 * @description 增加定时消息通知
 */
func (hub *ConnectHub) AddAfter(s time.Duration, after func() *FullMessage) {
	hub.afters[time.NewTicker(s)] = after
}

/**
 * @description 增加定时任务队列
 */
func (hub *ConnectHub) AddTask(s time.Duration, task func(*Client)) {
	hub.tasks[time.NewTicker(s)] = task
}

/**
 * @description 执行所有的定时消息通知
 */
func (hub *ConnectHub) ExecuteAfter() {
	if hub.afters == nil {
		return
	}

	for ticker, after := range hub.afters {
		time.Sleep(1 * time.Second)
		go func(t *time.Ticker, af func() *FullMessage) {
			for {
				<-t.C
				hub.broadcast <- af().GetFullMessage()
			}
		}(ticker, after)
	}
}

/**
 * @description 加载所有任务队列
 */
func (hub *ConnectHub) LoadTask() {
	if hub.tasks == nil {
		return
	}

	for ticker, task := range hub.tasks {
		go func(t *time.Ticker, k func(*Client)) {
			for {
				<-t.C
				for client, status := range hub.clients {
					if status {
						k(client)
					}
				}
			}
		}(ticker, task)
	}
}
