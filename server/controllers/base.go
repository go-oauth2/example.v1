package controllers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/example.v1/server/models"
	"gopkg.in/oauth2.v1"
)

var (
	oAuthManager *oauth2.OAuthManager
)

func init() {
	err := models.Init(beego.AppConfig.String("MongoURL"), beego.AppConfig.String("MongoURLDBName"))
	if err != nil {
		panic(err)
	}
	cli := &models.Client{
		ID:     "999999",
		Name:   "OAuth2测试应用",
		Secret: "999999",
		Domain: beego.AppConfig.String("ClientHost"),
	}
	if err := cli.Create(); err != nil {
		panic(err)
	}
	mgoConfig := &oauth2.MongoConfig{
		URL:    beego.AppConfig.String("MongoURL"),
		DBName: beego.AppConfig.String("MongoURLDBName"),
	}
	oauthConfig := &oauth2.OAuthConfig{
		ACConfig: &oauth2.ACConfig{
			ATExpiresIn: 60 * 60 * 24,
		},
	}
	manager, err := oauth2.NewDefaultOAuthManager(oauthConfig, mgoConfig, "Client", "")
	if err != nil {
		panic(err)
	}
	manager.SetACGenerate(oauth2.NewDefaultACGenerate())
	manager.SetACStore(oauth2.NewACMemoryStore(0))

	oAuthManager = manager
}
