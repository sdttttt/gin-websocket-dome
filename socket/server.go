/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 15:15:14
 * @LastEditTime: 2019-09-01 17:34:25
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/**
* 连接池
 */

var addr = flag.String("addr", "localhost:8080", "http service address")

/**
*	允许跨域
 */
var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

/**
* Http 协议升级为 WebSocket
 */
func changeProtocol(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		println("upgrade Error -> ", err)
	}
	return c
}

func echo(writer http.ResponseWriter, request *http.Request) {

	c := changeProtocol(writer, request)

	ProcessConnect(c)
}

func Run() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", echo)

	println("socket server GO GO GO!!!")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
