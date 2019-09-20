/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-17 15:54:16
 * @LastEditTime: 2019-09-17 16:04:00
 * @LastEditors: Please set LastEditors
 */
package file

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FileUploadUrl = "/file/upload"
)

func FileUpload(context *gin.Context) {

	if file, err := context.FormFile("yourFile"); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": "文件传输出现了问题",
		})
	} else {
		fileName := file.Filename

		println("============>", fileName)

	}

}
