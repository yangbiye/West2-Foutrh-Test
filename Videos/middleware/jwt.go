package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	util "github.com/videos/pkg/util/jwt"
)

var UserID uint
var ManagerID uint

// JWT JWT中间件
func JWT1() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取token
		token := c.Query("token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg4": "验证过期",
			})
			c.Abort()
			return
		}

		// 解析Token
		claims, err := util.ParseToken1(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg5": "验证过期",
			})
			c.Abort()
			return
		}

		// 验证
		now := time.Now()
		if now.After(claims.ExpiresAt.Time) || claims.Issuer != "Video" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg6": "验证过期",
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		UserID = claims.UserID
		c.Next()
	}
}

func JWT2() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取token
		token := c.Query("token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg1": "验证过期",
			})
			c.Abort()
			return
		}

		// 解析Token
		claims, err := util.ParseToken2(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg2": "验证过期",
			})
			c.Abort()
			return
		}

		// 验证
		now := time.Now()
		if now.After(claims.ExpiresAt.Time) || claims.Issuer != "Video" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg3": "验证过期",
			})
			c.Abort()
			return
		}

		c.Set("manager_id", claims.ManagerID)
		c.Next()
	}
}
