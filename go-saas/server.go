package main

import (
	"fmt"
	"saas/go-saas/controllers"
	"saas/go-saas/libs/config"
	"saas/go-saas/libs/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func ResgiterRoute(){
	fmt.Println("server ResgiterRoute action success...")

	route := gin.Default()

	//初始化session
	session_config := config.Get().Session
	store, _ := redis.NewStore(10, "tcp", fmt.Sprintf("%s:%d", session_config.Server, session_config.Port), fmt.Sprintf("%s", session_config.Password), []byte("secret"))
	route.Use(sessions.Sessions("STAR", store))

	//创建中间件
	route.Use(middleware.Middleware{}.Handle)

	route.GET("/create/tables", controllers.CommonController{}.CreateTables)
	//门店管理
	offline_store := route.Group("/offline_store")
	{
		c := controllers.NewOffLineStoreController()
		offline_store.GET("/view", c.View)
	}

	//客户管理
	customer := route.Group("/customer")
	{
		c := controllers.NewCustomerController()
		customer.GET("/view", c.View)
		customer.GET("/view/:id", c.Find)
		customer.POST("/create", c.Create)
		customer.GET("/delete/:id", c.Delete)
		customer.POST("/update", c.Update)
	}
	//区域
	zsetting := route.Group("/zsetting")
	{
		c := controllers.NewZsettingController()
		zsetting.GET("/view", c.View)
	}
	//产品管理
	product := route.Group("/product")
	{
		c := controllers.NewProductController()
		product.GET("/view", c.View)
	}
	//技师管理
	technician := route.Group("/technician")
	{
		c := controllers.NewTechnicianController()
		technician.GET("/view", c.View)
	}
	//门店订单管理
	sinformation := route.Group("/sinformation")
	{
		c := controllers.NewSinformationController()
		sinformation.GET("/view", c.View)
	}
	//到店记录
	records := route.Group("/records")
	{
		c := controllers.NewRecordsController()
		records.GET("/view", c.View)
	}
<<<<<<< HEAD

=======
	//客户管理
	customer := route.Group("/customer")
	{
		c := controllers.NewCustomerController()
		customer.GET("/view", c.View)
	}
>>>>>>> 09792f842f529e7004990697e71cb49f928bdb33
	//疗程记录
	course := route.Group("/course")
	{
		c := controllers.NewCourseController()
		course.GET("/view", c.View)
		course.GET("/view/:id", c.Find)
		course.POST("/create", c.Create)
		course.GET("/delete/:id", c.Delete)
		course.POST("/update", c.Update)
	}
	//维护记录
	maintenance := route.Group("/maintenance")
	{
		c := controllers.NewMaintenanceController()
		maintenance.GET("/view", c.View)
	}
	//返现客户
	cashback := route.Group("/cashback")
	{
		c := controllers.NewCashbackController()
		cashback.GET("/view", c.View)
	}
	//全部串店信息
	informatime := route.Group("/informatime")
	{
		c := controllers.NewInformatimeController()
		informatime.GET("/view", c.View)
	}
	//权限
	node := route.Group("/node")
	{
		c := controllers.NewNodeController()
		node.GET("/view", c.View)
	}
	//部门管理
	role := route.Group("/role")
	{
		c := controllers.NewRoleController()
		role.GET("/view", c.View)
	}

	route.Run()
}