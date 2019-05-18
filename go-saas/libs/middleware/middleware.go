package middleware

import (
	"saas/go-saas/controllers"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type Middleware struct {

}

func (this Middleware) Handle(g *gin.Context){
	this.Before(g)
}

func (this *Middleware) Before(g *gin.Context){
	//预处理文件对象
	controllers.Ctx = g
	controllers.Session = sessions.Default(g)

	controllers.RedisClient = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialPassword(""),
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(5*time.Second),
				redis.DialWriteTimeout(5*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}