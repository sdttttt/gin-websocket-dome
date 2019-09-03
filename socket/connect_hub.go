/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:18:55
 * @LastEditTime: 2019-09-03 18:23:34
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"log"
	"sync"
	"time"
)

type ConnectHub struct {
	clients map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan []byte

	tasks map[*time.Ticker]func(*Client)

	afters map[*time.Ticker]func() []byte
}

var connectHub *ConnectHub

var once sync.Once

/**
 * @description 客户端连接池引用
 */
func GetConnectHub() *ConnectHub {
	once.Do(
		func() {
			connectHub = &ConnectHub{
				clients:    make(map[*Client]bool),
				register:   make(chan *Client),
				unregister: make(chan *Client),
				broadcast:  make(chan []byte),
				tasks:      make(map[*time.Ticker]func(*Client)),
				afters:     make(map[*time.Ticker]func() []byte),
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
func (hub *ConnectHub) AddAfter(s time.Duration, after func() []byte) {
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
		go func(t *time.Ticker, af func() []byte) {
			for {
				<-t.C
				hub.broadcast <- af()
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
