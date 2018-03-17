package controllers

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/beego/social-auth"
	"github.com/beego/social-auth/apps"

	// just use mysql driver for example
	_ "github.com/go-sql-driver/mysql"

	"main/models"

	"github.com/astaxie/beego/plugins/cors"
)

func IsUserLogin(ctx *context.Context) (int, bool) {
	if id, ok := ctx.Input.CruSession.Get("login_user").(int); ok && id == 1 {
		return id, true
	}
	return 0, false
}

func Logout(ctx *context.Context) {
	ctx.Input.CruSession.Delete("login_user")
	types := social.GetAllTypes()
	for _, t := range types {
		ctx.Input.CruSession.Delete(t.NameLower())
	}
}

func SetInfoToSession(ctx *context.Context, userSocial *social.UserSocial) {
	ctx.Input.CruSession.Set(userSocial.Type.NameLower(),
		fmt.Sprintf("Identify: %s, AccessToken: %s", userSocial.Identify, userSocial.Data.AccessToken))
}

func HandleRedirect(ctx *context.Context) {
	redirect, err := SocialAuth.OAuthRedirect(ctx)
	if err != nil {
		beego.Error("SocialAuth.handleRedirect", err)
	}

	if len(redirect) > 0 {
		ctx.WriteString("{\"redirect_uri\": \"" + redirect + "\"}")
		//ctx.WriteString("<a href=\"" + redirect + "\" target=_blank>asdasdasad</a>")
	} else {
		ctx.WriteString("{\"error\": \"Error for redirect\"}")
	}
}

func HandleAccess(ctx *context.Context) {
	redirect, userSocial, err := SocialAuth.OAuthAccess(ctx)
	if err != nil {
		beego.Error("SocialAuth.handleAccess", err)
	}

	st, ok := SocialAuth.ReadyConnect(ctx)
	if !ok {
		ctx.Redirect(302, beego.AppConfig.String("autorization_redirect_uri"))
		return
	}
	_, userSocial, err = SocialAuth.ConnectAndLogin(ctx, st, 1)
	if err != nil {
		beego.Error(err)
	} else {

		err = models.CheckUserRegistration(userSocial.Id, userSocial.Data.AccessToken)
		if err != nil {

			SetInfoToSession(ctx, userSocial)
			Logout(ctx)
		}
	}
	if len(redirect) > 0 {
		ctx.Redirect(302, beego.AppConfig.String("autorization_redirect_uri"))
	}
}

type AuthController struct {
	beego.Controller
}

// @router /available
func (this *AuthController) Login() {

	types := social.GetAllTypes()
	this.Data["json"] = make(map[string]interface{})

	for _, t := range types {
		if t.Available() {
			(this.Data["json"].(map[string]interface{}))[t.NameLower()] = true
		}
	}
	this.ServeJSON()
}

func (this *AuthController) Connect() {

	st, ok := SocialAuth.ReadyConnect(this.Ctx)
	if !ok {
		this.Data["json"] = errors.New("Couldnt connect to authentication service")
		this.ServeJSON()
		return
	}
	loginRedirect, userSocial, err := SocialAuth.ConnectAndLogin(this.Ctx, st, 1)
	if err != nil {
		this.Data["json"] = err
	} else {
		SetInfoToSession(this.Ctx, userSocial)
		this.Data["json"] = map[string]string{"redirect_uri": loginRedirect}
	}

	this.ServeJSON()
}

type socialAuther struct {
}

func (p *socialAuther) IsUserLogin(ctx *context.Context) (int, bool) {
	return IsUserLogin(ctx)
}

func (p *socialAuther) LoginUser(ctx *context.Context, uid int) (string, error) {
	// fake login the user
	if uid == 1 {
		ctx.Input.CruSession.Set("login_user", 1)
	}
	return "/login", nil
}

var SocialAuth *social.SocialAuth

func init() {
	var err error

	// OAuth
	var clientId, secret string

	appURL := beego.AppConfig.String("social_auth_url")
	if len(appURL) > 0 {
		social.DefaultAppUrl = appURL
	}

	clientId = beego.AppConfig.String("github_client_id")
	secret = beego.AppConfig.String("github_client_secret")
	err = social.RegisterProvider(apps.NewGithub(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("google_client_id")
	secret = beego.AppConfig.String("google_client_secret")
	err = social.RegisterProvider(apps.NewGoogle(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("weibo_client_id")
	secret = beego.AppConfig.String("weibo_client_secret")
	err = social.RegisterProvider(apps.NewWeibo(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("qq_client_id")
	secret = beego.AppConfig.String("qq_client_secret")
	err = social.RegisterProvider(apps.NewQQ(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("dropbox_client_id")
	secret = beego.AppConfig.String("dropbox_client_secret")
	err = social.RegisterProvider(apps.NewDropbox(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("facebook_client_id")
	secret = beego.AppConfig.String("facebook_client_secret")
	err = social.RegisterProvider(apps.NewFacebook(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://*", "http://petrosyan.in:8000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	SocialAuth = social.NewSocial("/v1/login/", new(socialAuther))
	beego.InsertFilter("/v1/login/*/access", beego.FinishRouter, HandleAccess)
	beego.InsertFilter("/v1/login/*", beego.FinishRouter, HandleRedirect)
	// http://127.0.0.1:8080/v1/login/google
}
