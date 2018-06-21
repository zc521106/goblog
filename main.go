package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/images","static/images")
	beego.SetStaticPath("/upload","static/upload")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")
	beego.Run()
}

