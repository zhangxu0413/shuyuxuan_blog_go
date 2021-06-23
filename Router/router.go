package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shuyuxuan_blog_go/Services"
	"shuyuxuan_blog_go/Types"
)

var r = gin.Default()
var userServer Services.UserServer

func Init() {
	r.POST("/upsertUser", func(c *gin.Context) {
		var user Types.UserReq
		c.BindJSON(&user)
		log.Printf("%v",&user)
		err := userServer.UpsertUser(user)
		if err != nil {
			fmt.Printf("upsertUser Err %v", err)
		}
		c.JSON(200, gin.H{
			"id": user.Id,
			"name":   user.Name,
			"slogan": "一个文艺的前端工程师",
		})
	})
	r.Run(":9090")
}
