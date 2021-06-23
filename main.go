package main

import (
	Mysql "shuyuxuan_blog_go/Databases"
	"shuyuxuan_blog_go/Router"
)

func main() {
	Mysql.Init()
	Router.Init()
}
