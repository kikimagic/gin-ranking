package controllers

import (
	"gin-ranking/api/models"
	"gin-ranking/api/pkg/logger"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, err := models.GetUserByName(username)
	if err != nil {
		logger.Error(map[string]interface{}{"get user error": err.Error()})
		ReturnError(c, 4004, "获取用户信息失败")
		return
	}
	if user.Id == 0 || user.Password != EncryptPassword(password) {
		ReturnError(c, 4005, "用户名或密码不正确")
		return
	}

	data := UserApi{Id: user.Id, Username: user.Username}
	session := sessions.Default(c)
	//fmt.Println(session)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	//fmt.Println(session)
	session.Save()

	ReturnSuccess(c, 0, "success", data, 1)
}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4002, "两次密码不一致")
		return
	}

	user, err := models.GetUserByName(username)
	if user.Id != 0 {
		ReturnError(c, 4003, "用户名已存在")
		return
	}

	_, err = models.AddUser(username, EncryptPassword(password))
	if err != nil {
		logger.Error(map[string]interface{}{"add user error": err.Error()})
		ReturnError(c, 4004, "注册失败")
		return
	}

	ReturnSuccess(c, 200, "注册成功", nil, 0)
}
