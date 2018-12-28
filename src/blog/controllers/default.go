package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type AboutController struct {
	beego.Controller
}

type AlbumController struct {
	beego.Controller
}

type DetailsController struct {
	beego.Controller
}

type LeacotsController struct {
	beego.Controller
}

type WhisperController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *AboutController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "about.html"
}

func (c *AlbumController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "album.html"
}

func (c *DetailsController) Get() {
	c.Data["name"] = "Jay"
	c.Data["cont"] = "87678"
	c.TplName = "details.html"
}

func (c *LeacotsController) Get() {
	c.Data["name"] = "Jay"
	c.Data["cont"] = "87678"
	c.TplName = "leacots.html"
}

func (c *WhisperController) Get() {
	c.Data["name"] = "Jay"
	c.Data["cont"] = "87678"
	c.TplName = "whisper.html"
}
