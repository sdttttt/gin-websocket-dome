/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 17:01:17
 * @LastEditTime: 2019-09-03 17:02:20
 * @LastEditors: Please set LastEditors
 */
package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func ToSha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}
