package socket

import "encoding/json"

/*
	所有发送给客户端的消息都要实现这个结构体
	和执行下面的方法
*/
type FullMessage struct {
	Username string `json:"username"`

	Message string `json:"message"`
}

/*
	消息转化为JSON
*/
func (message *FullMessage) GetFullMessage() []byte {
	if result, err := json.Marshal(message); err == nil {
		return result
	}
	return nil
}
