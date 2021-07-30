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
		res := toolsService.Translate(t)
		log.Printf("%v",&res)
		c.JSON(200, res)
	})
}