package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

// TokenResponse 令牌响应参数
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// CallbackController 认证回执
type CallbackController struct {
	beego.Controller
}

func (cc *CallbackController) Get() {
	if v := cc.GetSession("IsAuthorize"); v != nil {
		cc.Redirect("/", http.StatusFound)
		return
	}
	code, state := cc.GetString("code"), cc.GetString("state")
	if code == "" {
		cc.Redirect("/", http.StatusFound)
		return
	}
	if state != "index" {
		panic("未知的State")
	}
	tokenValues := url.Values{}
	tokenValues.Add("grant_type", "authorization_code")
	tokenValues.Add("code", code)
	tokenValues.Add("redirect_uri", url.QueryEscape(beego.AppConfig.String("ClientHost")))
	tokenValues.Add("client_id", beego.AppConfig.String("ClientID"))
	tokenValues.Add("client_secret", beego.AppConfig.String("ClientSecret"))
	res, err := http.PostForm(beego.AppConfig.String("ServerHost")+"/token", tokenValues)
	if err != nil {
		panic(err)
	}
	var tokenRes TokenResponse
	err = json.NewDecoder(res.Body).Decode(&tokenRes)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	cc.SetSession("Token", tokenRes)
	cc.SetSession("IsAuthorize", true)
	cc.Redirect("/", http.StatusFound)
}
