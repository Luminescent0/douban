package api

import (
	"douban/tool"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"time"
)

var (
	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://121.4.229.95:8090/callback",
		ClientID:     "1a1ef437a61310f98d9e",
		ClientSecret: "152a00800c632b4d6a50dbfe4fe142bfe87e2708",
		Endpoint:     github.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	randomState = "xianye"
	randomNonce = "xian"
)

//func randString(nByte int) (string, error) {
//	b := make([]byte, nByte)
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		return "", err
//	}
//	return base64.RawURLEncoding.EncodeToString(b), nil
//}
func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func home(c *gin.Context) {
	var html = `<html><body><a href="/loginByGit">Github Login</a></body></html>`
	_, err := fmt.Fprint(c.Writer, html)
	if err != nil {
		fmt.Println("set html failed:", err)
		return
	}
}

func loginByGit(c *gin.Context) {
	//url := "https://github.com/login/oauth/authorize?client_id=" + githubOauthConfig.ClientID +
	//	"&redirect_uri=" + githubOauthConfig.RedirectURL +
	//	"&state=" + randomState

	setCallbackCookie(c.Writer, c.Request, "state", randomState)
	setCallbackCookie(c.Writer, c.Request, "nonce", randomNonce)
	c.Redirect(http.StatusFound, githubOauthConfig.AuthCodeURL(randomState, oidc.Nonce(randomNonce)))
}

func callback(c *gin.Context) {
	if c.Query("state") != randomState {
		fmt.Println("state is not valid")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	oauth2Token, err := githubOauthConfig.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		fmt.Println("could not get token:", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Println(oauth2Token)
	ctx := context.Background()
	fmt.Println(ctx)
	provider, err := oidc.NewProvider(ctx, "https://token.actions.githubusercontent.com")
	if err != nil {
		fmt.Println(err)
	}
	oidcConfig := &oidc.Config{ClientID: "1a1ef437a61310f98d9e"}
	verifier := provider.Verifier(oidcConfig)

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		fmt.Println("No id_token field in oauth2 token.")
		c.Redirect(http.StatusInternalServerError, "/")
		return
	}
	fmt.Println(rawIDToken)
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		fmt.Println("failed to verify ID Token:", err)
		c.Redirect(http.StatusInternalServerError, "/")
		return
	}
	nonce, err := c.Request.Cookie("nonce")
	if err != nil {
		fmt.Println("nonce not found:", err)
		c.Redirect(http.StatusBadRequest, "/")
		return
	}
	if idToken.Nonce != nonce.Value {
		fmt.Println("nonce did not match", err)
		c.Redirect(http.StatusBadRequest, "/")
		return
	}
	resp := struct {
		Oauth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage
	}{oauth2Token, new(json.RawMessage)}
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/")
		return
	}
	data, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusInternalServerError, "/")
		return
	}
	//c.Writer.Write(data)
	tool.RespSuccessfulWithDate(c, data)
	//resp,err := http.Post("https://api.github.com/user?access_token="+token.AccessToken,"application/x-www-form-urlencoded",nil)
	//fmt.Println(token.AccessToken)
	//if err != nil {
	//	fmt.Println("could not create get request:",err)
	//	c.Redirect(http.StatusTemporaryRedirect,"/")
	//	return
	//}
	//defer resp.Body.Close()
	//content,err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("could not parse response:",err)
	//	c.Redirect(http.StatusTemporaryRedirect,"/")
	//	return
	//}
	//_,err = fmt.Fprintf(c.Writer,"Response:%s",content)
	//if err != nil {
	//	fmt.Println("could not fmt content:",err)
	//	return
	//}
	//tool.RespSuccessfulWithDate(c,content)
	//上面这么写会报 Must specify access token via Authorization header.
	//GitHub在2020.3起不允许将access_token作为url中的参数明文传输，要将其作为Authorization HTTP header中的参数传输

	//var userInfoUrl = "https://api.github.com/user" //github用户信息获取接口
	//var req *http.Request
	//if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
	//	return
	//}
	//req.Header.Set("accept", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("token %s", oauth2Token.AccessToken))
	//
	////发送请求并获取响应
	//var client = http.Client{}
	//var res *http.Response
	//if res, err = client.Do(req); err != nil {
	//	return
	//}
	////将响应的数据写入 userInfo中并返回
	//var userInfo = make(map[string]interface{})
	//if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
	//	return
	//}
	//tool.RespSuccessfulWithDate(c, userInfo)

	//用户信息获取部分的参考博客
	//https://blog.csdn.net/qq_19018277/article/details/104935403?utm_source=app&app_version=5.0.1&code=app_1562916241&uLinkId=usr1mkqgl919blen

	//iUsername := userInfo["login"]
	//username := iUsername.(string)
	//token1, err1 := CreateToken(username)
	//if err1 != nil {
	//	tool.RespInternalError(c)
	//	return
	//}
	//tool.RespSuccessfulWithDate(c, token1)
	//return

}
