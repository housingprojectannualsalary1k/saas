package controllers

import (
	"github.com/gin-gonic/gin"
	"saas/go-saas/models"
)

type CustomerController struct {
	CommonController
}

func NewCustomerController() *CustomerController{
	o := CustomerController{}
	return &o
}

func (this *CustomerController) View(g *gin.Context){
	//搜索条件
	_search := models.CustomerSearch{}

	//绑定搜索数据
	err := g.Bind(&_search)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
	}

	//调用查询方法
	_customer := models.NewCustomerModel()
	_customer.View(&_search)
}

//添加客户
func (this *CustomerController) Create(g *gin.Context){
	_customer := models.NewCustomerModel()

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