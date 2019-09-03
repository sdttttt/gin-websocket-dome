/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-02 18:18:55
 * @LastEditTime: 2019-09-02 23:05:07
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

func (hub *ConnectHub) GetCrrentCount() int {
	return len(hub.clients)
}

func (hub *ConnectHub) AddAfter(s time.Duration, after func() []byte) {
	hub.afters[time.NewTicker(s)] = after
}

func (hub *ConnectHub) AddTask(s time.Duration, task func(*Client)) {
	hub.tasks[time.NewTicker(s)] = task
}

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
