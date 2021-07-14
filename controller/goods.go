/**
@auther: kid1999
@date: 2021/7/14 10:57
@desc: goods
**/

package controller

import (
	"esystem/database"
	models "esystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Goods struct {
	Basic
}


func (g *Goods) AddGoods(c *gin.Context){
	var goods models.Goods
	c.Bind(&goods)

	if goods.GoodsName == "" || goods.Description == "" || goods.OwnerStudentID == ""{
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		goods.State = "applying"
		database.DB.Create(&goods)
		c.JSON(http.StatusOK,goods)
	}
}


func (g *Goods) GoodsInfo(c *gin.Context){
	goods := models.Goods{}
	goodsID := c.Query("id")

	if goodsID == "" {
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Where("ID=?",goodsID).First(&goods)
		c.JSON(http.StatusOK,goods)
	}
}


func (g *Goods) FindAllGoods(c *gin.Context){
	goodsList := []models.Goods{}

	database.DB.Find(&goodsList)
	c.JSON(http.StatusOK,goodsList)
}


func (g *Goods) DelGoods(c *gin.Context){
	goodsID := c.Query("id")

	if goodsID == "" {
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Delete(&models.Goods{},goodsID)
		c.JSON(http.StatusOK,"Delete successful")
	}
}


func (g *Goods) UpdateGoods(c *gin.Context){
	var goods models.Goods
	c.Bind(&goods)

	if goods.GoodsName == "" || goods.Description == "" || goods.OwnerStudentID == ""{
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Model(&goods).Updates(goods)
		c.JSON(http.StatusOK,goods)
	}
}


func (g *Goods) FindGoodsByName(c *gin.Context){
	goodsList := []models.Goods{}
	goodsName := c.Query("goodsName")

	if goodsName == "" {
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Find(&goodsList,"GoodsName like '%?%'",goodsName)
		c.JSON(http.StatusOK,goodsList)
	}
}


func (g *Goods) FindGoodsByStudentID(c *gin.Context){
	goodsList := []models.Goods{}
	studentID := c.Query("studentID")

	if studentID == "" {
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Find(&goodsList,"OwnerStudentID = ?",studentID)
		c.JSON(http.StatusOK,goodsList)
	}
}