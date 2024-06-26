package router

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.GET("/dhd", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})
}
