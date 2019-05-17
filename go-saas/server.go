package main

import (
	"fmt"
	"saas/controllers"
	"saas/libs/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func ResgiterRoute(){
	fmt.Println("server ResgiterRoute action success...")

	route := gin.Default()

	//初始化session
	//session_config := config.Get().Session
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	route.Use(sessions.Sessions("STAR", store))

	//创建中间件
	route.Use(middleware.Middleware{}.Handle)

	route.GET("/create/tables", controllers.CommonController{}.CreateTables)
	//新闻类
	offline_store := route.Group("/offline_store")
	{
		c := controllers.NewOffLineStoreController()
		offline_store.GET("/view", c.Copy().View)
	}


	route.Run()
}