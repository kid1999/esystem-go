/**
@auther: kid1999
@date: 2021/7/14 14:05
@desc: admin
**/

package controller

import (
	"esystem/database"
	models "esystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
	Basic
}


func (a *Admin) CheckUser(c *gin.Context){
	user := models.User{}

	c.Bind(&user)

	database.DB.Model(&user).Updates(user)
	c.JSON(http.StatusOK,user)
}


func (a *Admin) CheckGoods(c *gin.Context){
	goods := models.Goods{}

	c.Bind(&goods)

	database.DB.Model(&goods).Updates(goods)
	c.JSON(http.StatusOK,goods)
}
