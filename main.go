package main

import (
	"fmt"
	"os"
	"shuyuxuan_blog_go/Router"
)

func main() {
	fmt.Printf("env:%s", os.Getenv("GO_ENV"))
	Router.Init()
}
