package routers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/example.v1/client/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/callback", &controllers.CallbackController{})
}
