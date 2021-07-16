/**
@auther: kid1999
@date: 2021/7/15 15:29
@desc: file
**/

package controller

import (
	"esystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct {
	Basic
}

func (f *FileController) UploadFile(c *gin.Context){
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := header.Filename
	fmt.Println(file, err, filename)

	fileUrl, err := util.UploadFileByStream(filename,file)

	c.JSON(http.StatusOK,fileUrl)
}

