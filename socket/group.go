/*
 * @Description: In User Settings Edit
 * @Author: SDTTTTT
 * @Date: 2019-08-31 19:54:47
 * @LastEditTime: 2019-09-02 14:46:21
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type ConnectGroup struct {
	sync.RWMutex
	conns map[*websocket.Conn]bool
}

var once sync.Once

/**
 *
 * 连接池
 */
var group *ConnectGroup

/**
 * @description 你可以使用这个方法来获取连接池的引用
 * @return {@link *ConnectGroup}
 */
func GetGroup() *ConnectGroup {
	once.Do(
		func() {
			log.Println(time.Now().String(), "Connect Pool initializer ...")
			group = &ConnectGroup{conns: make(map[*websocket.Conn]bool)}
		})

	return group
}

const (
	writerWait       = 10 * time.Second
	maxMessageSize   = 8192
	closeGracePeriod = 12 * time.Second
)

func (group *ConnectGroup) AddConnect(conn *websocket.Conn) {
	group.Lock()
	group.conns[conn] = true
	group.Unlock()
	log.Println("crrent ConnectPool Count : ", len(group.conns))
}

func (group *ConnectGroup) ChangeConnStatus(conn *websocket.Conn) {
	group.Lock()
	group.conns[conn] = !group.conns[conn]
	group.Unlock()
}

/**
 * @description 大部分情况下我们只需要关注连接是否在 ConnectPool 中被删除 <br />
 * 	这个方法不是那么被推荐的
 * @param {@link *websocket.Conn}
 */
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

	group.Broadcast(websocket.TextMessage, ([]byte)(conn.RemoteAddr().String()+"離開了"))
}

func (group *ConnectGroup) IsLive() {

}

/**
 * @description 你肯定非常想看到这个方法,这个方法会向ConnectPool 所有的连接ReadBuffer中写入数据
 * @param int , [] byte
 * @return error
 */
func (group *ConnectGroup) Broadcast(messageType int, message []byte) error {

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
