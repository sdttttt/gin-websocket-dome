/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 19:54:47
 * @LastEditTime: 2019-09-01 17:42:12
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var group = &ConnectGroup{conns: make(map[*websocket.Conn]bool)}

func GetGroup() *ConnectGroup {
	return group
}

type ConnectGroup struct {
	sync.RWMutex
	conns map[*websocket.Conn]bool
}

const (
	writerWait = 10 * time.Second

	maxMessageSize = 8192

	closeGracePeriod = 12 * time.Second
)

func (group *ConnectGroup) AddConnect(conn *websocket.Conn) {

	group.Lock()
	group.conns[conn] = true
	group.Unlock()
	log.Println("当前连接池数量 -> ", len(group.conns))

}

func (group *ConnectGroup) ChangeConnStatus(conn *websocket.Conn) {
	group.Lock()
	group.conns[conn] = !group.conns[conn]
	group.Unlock()
}

func (group *ConnectGroup) DelAndCloseConn(conn *websocket.Conn) {
	group.Lock()
	delete(group.conns, conn)
	group.Unlock()

	conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	conn.Close()
}

func (group *ConnectGroup) DelConnect(conn *websocket.Conn) {
	group.Lock()
	delete(group.conns, conn)
	group.Unlock()
	conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func (group *ConnectGroup) IsLive() {

}

func (group *ConnectGroup) Broadcast(messageType int, message []byte) error {

	log.Println("准备使用广播...")

	var err error = nil
	for conn, status := range group.conns {
		if status {
			if err = conn.WriteMessage(messageType, message); err != nil {
				return err
			}
		}
	}
	return err
}
