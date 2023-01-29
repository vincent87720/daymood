package routers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24,
		HttpOnly: true,
	})

	return sessions.Sessions("DAYMOODSESSID", store)
}

func AuthSession() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		sessionID := session.Get("secret")
		if sessionID == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "auth",
			})
			return
		}
		context.Next()
	}
}

func SaveSession(context *gin.Context, userID int64) {
	session := sessions.Default(context)
	session.Set("secret", userID)
	session.Save()
}

func GetSession(context *gin.Context) int64 {
	session := sessions.Default(context)
	sessionID := session.Get("secret")
	if sessionID == nil {
		return sessionID.(int64)
	}
	return 0
}

func ClearSession(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	session.Save()
}
