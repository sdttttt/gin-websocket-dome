/*
 * @Description: In User Settings Edit
 * @Author: SDTTTTT
 * @Date: 2019-08-31 19:54:47
 * @LastEditTime: 2019-09-02 17:47:34
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

	/** Tasks里存储所有你所定义的定时任务
	它们会在连接初始化的时候被执行
	*/
	tasks map[*time.Ticker]func(*websocket.Conn)
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
			group = &ConnectGroup{conns: make(map[*websocket.Conn]bool),
				tasks: make(map[*time.Ticker]func(*websocket.Conn))}
		})

	return group
}

const (
	writerWait       = 10 * time.Second
	maxMessageSize   = 8192
	closeGracePeriod = 12 * time.Second
)

func (group *ConnectGroup) GetConnCount() int {
	return len(group.conns)
}

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

/** @description 增加定时任务
 */
func (group *ConnectGroup) AddTickerTask(s time.Duration, connEvent func(*websocket.Conn)) {
	ticker := time.NewTicker(s)

	group.Lock()
	group.tasks[ticker] = connEvent
	group.Unlock()
}

/** @description  Task增加完成之后使用 ExecuteTasks 加载所有的任务
 */
func (group *ConnectGroup) ExecuteTasks() {

	if group.tasks == nil {
		return
	}

	log.Println("当前列队任务数 : ", len(group.tasks))

	for tiker, task := range group.tasks {
		go func(t *time.Ticker) {
			for {
				<-t.C
				for conn, status := range group.conns {
					if status {
						task(conn)
					}
				}
			}
		}(tiker)
	}
}
