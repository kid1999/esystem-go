/**
@auther: kid1999
@date: 2021/7/13 22:53
@desc: user
**/

package controller

import (
	"crypto/md5"
	"encoding/hex"
	"esystem/database"
	models "esystem/model"
	"esystem/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Basic
}


func (u *User) Login(c *gin.Context){
	var user models.User
	c.Bind(&user)
	password := user.Password

	database.DB.Model(&user).Where("student_id = ?",user.StudentID).First(&user)

	if Encoded(password) == user.Password{
		response := make(map[string]interface{})
		response["token"],_ = util.GenerateToken(user.StudentID,user.Password)
		user.Password = ""
		response["user"] = user
		c.JSON(http.StatusOK,response)
	}else{
		c.JSON(http.StatusBadRequest,"账号密码不正确")
	}
}

func (u *User) Register(c *gin.Context){
	var user models.User
	var count int64
	c.Bind(&user)

	if user.Username == "" || user.Password == "" || user.Phone == "" || user.StudentID == ""{
		c.JSON(http.StatusBadRequest,"请求参数不完整！")
	}

	database.DB.Model(&user).Where("student_id = ?",user.StudentID).Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest,"用户已存在！")
	}else{
		user.Password = Encoded(user.Password)
		user.State = "applying"
		database.DB.Model(&models.User{}).Create(&user)
		c.JSON(http.StatusOK,user)
	}
}


func (u *User) FindAllUsers(c *gin.Context){
	userList := []models.User{}

	database.DB.Find(&userList)
	c.JSON(http.StatusOK,userList)
}


func (u *User) Test(c *gin.Context){
	var b models.User
	c.Bind(&b)
	c.JSON(http.StatusOK,b)
}


// md5 Encoded
func Encoded(text string) string {
	password := []byte(text)
	md5Ctx := md5.New()
	md5Ctx.Write(password)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}