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
	"strconv"
)

type Goods struct {
	Basic
}


func (g *Goods) AddGoods(c *gin.Context){
	var goods models.Goods
	c.ShouldBind(&goods)

	if goods.GoodsName == "" || goods.Description == "" || goods.OwnerStudentID == ""{
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		goods.State = ""
		database.DB.Model(&models.Goods{}).Create(&goods)
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
	pageindex, _ := strconv.Atoi(c.Param("pageindex"))
	pagesize,_  := strconv.Atoi(c.Param("pagesize"))

	var count int64
	database.DB.Model(&models.Goods{}).Where("state <> ?","").Count(&count)

	goodsList := []models.Goods{}

	database.DB.Offset((pageindex-1)*pagesize).Limit(pagesize).Find(&goodsList).Where("state <> ?","")
	res := make(map[string]interface{})
	res["goods"] = goodsList
	res["count"] = count
	res["pageindex"] = pageindex
	res["pagesize"] = pagesize
	c.JSON(http.StatusOK,res)
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
	pageindex, _ := strconv.Atoi(c.Query("pageindex"))
	pagesize,_  := strconv.Atoi(c.Query("pagesize"))
	text := c.Query("search")

	var count int64
	database.DB.Model(&models.Goods{}).Where("state <> ? And goods_name like ?","","%" + text + "%").Count(&count)

	goodsList := []models.Goods{}

	database.DB.Offset((pageindex-1)*pagesize).Limit(pagesize).Where("state <> '' And  goods_name like ?","%" + text + "%").Find(&goodsList)

	res := make(map[string]interface{})
	res["goods"] = goodsList
	res["count"] = count
	res["pageindex"] = pageindex
	res["pagesize"] = pagesize
	c.JSON(http.StatusOK,res)
}


func (g *Goods) FindGoodsByStudentID(c *gin.Context){
	goodsList := []models.Goods{}
	studentID := c.Query("studentID")

	if studentID == "" {
		c.JSON(http.StatusBadRequest,"参数请求错误！")
	}else{
		database.DB.Find(&goodsList,"OwnerStudentID = ?",studentID).Where("state <> ?","")
		c.JSON(http.StatusOK,goodsList)
	}
}