package routers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/example.v1/server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/authorize", &controllers.AuthorizeController{})
	beego.Router("/token", &controllers.TokenController{})
}
