package middleware

import (
	"book/initalize/database/redis"
	"book/pkg/grf/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		//解析token
		claims, err := jwt.JWT().ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		//验证
		if err = claims.Valid(); err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		//检查token是否被释放
		if _, err = redis.Redis().Get(context.Background(), "jwt:"+claims.Id).Result(); err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
