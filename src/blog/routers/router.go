package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index.html", &controllers.IndexController{})
	beego.Router("/about.html", &controllers.AboutController{})
	beego.Router("/album.html", &controllers.AlbumController{})
	beego.Router("/details.html", &controllers.DetailsController{})
	beego.Router("/leacots.html", &controllers.LeacotsController{})
	beego.Router("/whisper.html", &controllers.WhisperController{})
}
