package middleware

import (
	"errors"
	"gin-starter/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		adminInfo, ok := session.Get(public.AdminSessionInfoKey).(string)
		if !ok || adminInfo == "" {
			ResponseError(c, 2001, errors.New("用户未登录"))
			return
		}
		c.Next()
	}
}
