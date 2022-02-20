package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"path"
	"strconv"
	"time"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err :", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	}
	token, err1 := CreateToken(username)
	if err1 != nil {
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, gin.H{"msg": token})
	tool.RespSuccessful(ctx)
	return
}
func register(ctx *gin.Context) {
	username, password := verify(ctx)
	user := model.User{
		Username: username,
		Password: password,
	}
	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithDate(ctx, "用户名已经存在")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespErrorWithDate(ctx, "注册失败")
}
func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("oldPassword")
	newPassword := ctx.PostForm("newPassword")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string) //接口断言

	//检验旧密码是否正确
	flag, err := service.IsPasswordCorrect(username, oldPassword)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	if !flag {
		tool.RespErrorWithDate(ctx, "旧密码错误")
		return
	}

	//修改新密码
	err = service.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("change password err: ", err)
		tool.RespSuccessfulWithDate(ctx, "修改失败")
		return
	}

	tool.RespSuccessfulWithDate(ctx, "修改成功")
}
func verify(ctx *gin.Context) (string, string) { //验证非法输入
	validate := validator.New() //创建验证器
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	u := model.User{Id: 0, Username: username, Password: password}

	err := validate.Struct(u)
	fmt.Println(err)
	if err != nil {
		return "存在非法输入", ""
	}
	return username, password

}

func introduction(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	user, err := service.CheckIntroduction(username)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithDate(ctx, user.Introduction)
}

func changeIntroduction(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	introduction := ctx.PostForm("introduction")
	err := service.ChangeIntroduction(username, introduction)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessful(ctx)
}

func uploadAvatar(c *gin.Context) {
	iUsername, _ := c.Get("username")
	username := iUsername.(string)
	//解析上传的参数,file username
	file, err := c.FormFile("avatar")
	if err != nil {
		tool.RespErrorWithDate(c, "参数解析失败")
		return
	}
	if file.Size > 1024*1024*5 {
		tool.RespErrorWithDate(c, "文件过大")
		return
	}
	fileSuffix := path.Ext(file.Filename)
	if !(fileSuffix == ".jpg" || fileSuffix == ".png") {
		tool.RespErrorWithDate(c, "文件格式错误")
		return
	}
	//file保存到本地
	fileName := "./uploadfile" + strconv.FormatInt(time.Now().Unix(), 10) + username + fileSuffix
	fileAddress := "/opt/gocode/src/douban" + fileName[1:]
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.RespErrorWithDate(c, "保存头像失败")
		return
	}
	//将保存后的文件本地路径保存到用户表中的头像字段
	loadString := "http:121.4.229.95:8080/picture/" + fileName[13:]
	err = service.UploadAvatar(username, loadString, fileAddress)
	if err != nil {
		tool.RespErrorWithDate(c, "上传失败")
		return
	}
	tool.RespSuccessfulWithDate(c, "上传成功")
}
