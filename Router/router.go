package Router

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Init() {
	initUserRouter()
	initDakaRouter()
	initToolsRouter()
	r.Run(":9090")
}
