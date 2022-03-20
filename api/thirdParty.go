package api

import (
	"douban/model"
	"douban/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

var (
	githubOauthConfig = &model.Config{
		RedirectURL:  "http://121.4.229.95:8090/callback",
		ClientID:     "1a1ef437a61310f98d9e",
		ClientSecret: "152a00800c632b4d6a50dbfe4fe142bfe87e2708",
		Endpoint: model.URL{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	randomState = "xianye"
)

func home(c *gin.Context) {
	var html = `<html><body><a href="/loginByGit">Github Login</a></body></html>`
	_, err := fmt.Fprint(c.Writer, html)
	if err != nil {
		fmt.Println("set html failed:", err)
		return
	}
}

func loginByGit(c *gin.Context) {
	url := "https://github.com/login/oauth/authorize?client_id=" + githubOauthConfig.ClientID +
		"&redirect_uri=" + githubOauthConfig.RedirectURL +
		"&state=" + randomState

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func callback(c *gin.Context) {
	if c.Query("state") != randomState {
		fmt.Println("state is not valid")
		tool.RespInternalError(c)
		return
	}
	//oauth2Token, err := githubOauthConfig.Exchange(oauth2.NoContext, c.Query("code"))
	//if err != nil {
	//	fmt.Println("could not get token:", err)
	//	tool.RespInternalError(c)
	//	return
	//}
	//fmt.Println(oauth2Token)

	code := c.Query("code")
	postData := url.Values{}
	postData.Add("code", code)
	postData.Add("client_id", githubOauthConfig.ClientID)
	postData.Add("client_secret", githubOauthConfig.ClientSecret)
	body := strings.NewReader(postData.Encode())

	//发送请求
	var req *http.Request
	req, err := http.NewRequest(http.MethodPost, "https://github.com/login/oauth/access_token", body)
	if err != nil {
		fmt.Println("could not get token:", err)
		tool.RespInternalError(c)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "application/json")
	var resp *http.Response
	var client1 = http.Client{}
	if resp, err = client1.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	var token = make(map[string]interface{})
	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return
	}
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

	var userInfoUrl = "https://api.github.com/user" //github用户信息获取接口
	var req2 *http.Request
	if req2, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return
	}
	req2.Header.Set("accept", "application/json")
	req2.Header.Set("Authorization", fmt.Sprintf("token %s"))

	//发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}
	//将响应的数据写入 userInfo中并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return
	}
	//tool.RespSuccessfulWithDate(c, userInfo)
	defer res.Body.Close()
	//用户信息获取部分的参考博客
	//https://blog.csdn.net/qq_19018277/article/details/104935403?utm_source=app&app_version=5.0.1&code=app_1562916241&uLinkId=usr1mkqgl919blen

	iUsername := userInfo["login"]
	username := iUsername.(string)
	token1, err1 := CreateToken(username)
	if err1 != nil {
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessfulWithDate(c, token1)
	return

}
