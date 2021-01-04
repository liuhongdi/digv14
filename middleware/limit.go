package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv14/global"
	"github.com/liuhongdi/digv14/pkg/result"
	"github.com/liuhongdi/digv14/pkg/util"
)
//限流器
func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//得到ip地址
		ipAddr:=util.GetRealIp(c)
		fmt.Println("current ip:"+ipAddr)
		//ipAddr:="127.0.0.1"
		limiter := global.RateLimiter.GetLimiter(ipAddr)
		if !limiter.Allow() {
			fmt.Println("not allow,will return")
			resultRes := result.NewResult(c)
			resultRes.Error(2004,"访问超出限制")
			return
		} else {
			fmt.Println("allow,next")
			c.Next()
		}
	}
}

