/**
@auther: kid1999
@date: 2021/7/14 15:09
@desc: jwt_auth
**/

package middleware

import (
	"esystem/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		//token := c.Query("token")
		token := c.GetHeader("token")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusBadRequest
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusRequestTimeout
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : "authorization failed",
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
