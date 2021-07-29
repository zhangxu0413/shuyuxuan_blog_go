package Router

import (
	"github.com/gin-gonic/gin"
	"log"
	"shuyuxuan_blog_go/Services"
	"shuyuxuan_blog_go/Types"
)

var toolsService Services.ToolsService = Services.ToolsService{}
func initToolsRouter()  {
	r.POST("/api/tools/translate", func(c *gin.Context) {
		var t Types.TranslateReq = Types.TranslateReq{}
		c.BindJSON(&t)
		log.Printf("%v",&t)
		toolsService.Translate(t)
		c.JSON(200, gin.H{
			"msg": "更新用户信息失败",
		})
	})
}