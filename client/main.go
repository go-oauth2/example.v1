package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-oauth2/example.v1/client/routers"
)

func main() {
	beego.Run()
}
