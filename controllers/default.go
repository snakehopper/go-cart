package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "http://" + this.Ctx.Request.Host
	this.Data["Project"] = beego.AppName
	this.Data["Email"] = beego.AppConfig.String("email")
	this.TplNames = "index.html"
}

type CartController struct {
	beego.Controller
}

func (this *CartController) Add() {
	id := this.Ctx.Params["0"]
	var s Shopping
	var c *Cart
	var msg string
	v := this.GetSession("20130805")

	if v == nil {
		c = new(Cart)
		c.Items = make(map[string]Item)
		msg = c.Add(id, 1)
		s = c
	} else {
		c = v.(*Cart)
		msg = c.Add(id, 1)
		s = c
	}
	this.SetSession("20130805", s)
	this.Data["Msg"] = msg
	this.Data["Cart"] = c
	this.TplNames = "cart.html"
}

func (this *CartController) View() {
	var s Shopping
	var c *Cart
	v := this.GetSession("20130805")
	if v == nil {
		c = new(Cart)
		c.Items = make(map[string]Item)
		s = c
	} else {
		c = v.(*Cart)
		s = c
	}
	this.SetSession("20130805", s)
	this.Data["Cart"] = c
	this.TplNames = "cart.html"
}

func (this *CartController) Remove() {
	id := this.Ctx.Params["0"]
	var s Shopping
	var c *Cart
	v := this.GetSession("20130805")
	if v != nil {
		c = v.(*Cart)
		c.Remove(id)
		s = c
	}
	this.SetSession("20130805", s)
	this.Redirect("/cart/view", 302)
}
