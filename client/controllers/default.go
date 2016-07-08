package controllers

import (
	"encoding/json"
	"net/url"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["ClientName"] = beego.AppConfig.String("ClientName")
	if v := c.GetSession("IsAuthorize"); v != nil {
		tokenValue := c.GetSession("Token")
		jsonValue, err := json.MarshalIndent(tokenValue, "", "\t")
		if err != nil {
			panic(err)
		}
		c.Data["IsAuthorize"] = v
		c.Data["Token"] = string(jsonValue)
	} else {
		serverURL, err := url.Parse(beego.AppConfig.String("ServerHost"))
		if err != nil {
			panic(err)
		}
		serverURL.Path = "authorize"
		serverValues := url.Values{}
		serverValues.Set("response_type", "code")
		serverValues.Set("client_id", beego.AppConfig.String("ClientID"))
		serverValues.Set("redirect_uri", url.QueryEscape(beego.AppConfig.String("ClientHost")))
		serverValues.Set("state", "index")
		serverValues.Set("scope", "all")
		serverURL.RawQuery = serverValues.Encode()
		c.Data["Href"] = serverURL.String()
	}
	c.TplName = "index.tpl"
}
