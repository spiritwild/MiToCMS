package routers

import (
	"github.com/spiritwild/MiToCMS/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
