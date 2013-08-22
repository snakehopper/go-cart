package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Bootstrap() {
	this.TplNames = "test-bootstrap.html"
}
