/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 15:15:14
 * @LastEditTime: 2019-09-03 18:12:21
 * @LastEditors: Please set LastEditors
 */
package socket

import (
	"flag"
	"gin-web/dao/service"
	"gin-web/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

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

func echo(writer http.ResponseWriter, request *http.Request, username string) {

	c := changeProtocol(writer, request)

	ProcessConnect(c, username)
}

/*
*	為 Gin 框架提供快速啟動,請把它註冊到您的路由中
 */
const GinEchoUrl = "/ws"

func GinEcho(context *gin.Context) {
	username := util.GetSession(context, "token")
	if result, ok := username.(string); ok && result != "" {
		if service.GetUserService().ExistsUser(result) {
			log.Println(username)
			echo(context.Writer, context.Request, result)
		}
	}
	context.AbortWithStatus(http.StatusBadRequest)

}

/*
*	如果你想單獨啟動這個 WebSocket 服務器的話 ， 可以使用這個
	这段代码目前不可用，WebSocket 服務器和Gin目前有耦合性
*/
//func Run() {
//	flag.Parse()
//	log.SetFlags(0)
//	http.HandleFunc("/ws", echo)
//
//	println("socket server GO GO GO!!!")
//	log.Fatal(http.ListenAndServe(*addr, nil))
//}
