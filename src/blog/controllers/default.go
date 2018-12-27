package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type RegisteController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	userName := c.GetString("username")
	passwd := c.GetString("password")
	if userName == passwd {
		c.Ctx.WriteString("登录成功！")
	} else {
		c.Ctx.WriteString("登录失败！")
	}

}

func (c *RegisteController) Get() {
	c.TplName = "registe.html"
}

func (c *RegisteController) Post() {
	c.Ctx.WriteString("注册成功！")
	c.TplName = "login.html"
}
