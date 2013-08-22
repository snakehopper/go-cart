package main

import (
	"github.com/astaxie/beego"
	"github.com/snakehopper/go-cart/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.TestController{})
	beego.AutoRouter(&controllers.CartController{})
	beego.Run()
}
