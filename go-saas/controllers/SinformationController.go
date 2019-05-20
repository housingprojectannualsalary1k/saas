package controllers

import (
	"github.com/gin-gonic/gin"
	"saas/go-saas/models"
)

type SinformationController struct {
	CommonController
}

func NewSinformationController() *SinformationController{
	o := SinformationController{}
	return &o
}

func (this *SinformationController) View(g *gin.Context){
	//搜索条件
	_search := models.SinformationSearch{}

	//绑定搜索数据
	err := g.Bind(&_search)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
	}

	//调用查询方法
	_customer := models.NewSinformationModel()
	_customer.View(&_search)
}

//添加客户
func (this *SinformationController) Create(g *gin.Context){
	_customer := models.NewSinformationModel()

	//绑定添加
	err := g.Bind(&_customer)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//调用查询方法
	ID, err := _customer.Create()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}