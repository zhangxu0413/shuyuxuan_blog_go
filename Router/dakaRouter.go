package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shuyuxuan_blog_go/Services"
	"shuyuxuan_blog_go/Types"
)

var dakaService Services.DakaService = Services.DakaService{}

func initDakaRouter()  {
	r.POST("/api/daka/add", func(c *gin.Context) {
		var dakaRecord Types.DakaReq
		c.BindJSON(&dakaRecord)
		log.Printf("%v",&dakaRecord)
		err := dakaService.AddRecord(dakaRecord)
		if err != nil {
			fmt.Printf("addDaka Err %v", err)
			c.JSON(200, gin.H{
				"msg": "创建打卡信息失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "创建打卡信息成功",
			})
		}

	})
	r.GET("/api/daka/getList", func(c *gin.Context) {
		list := dakaService.GetList()
		c.JSON(200, list)
	})
}
