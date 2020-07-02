package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"server/database"
	"server/model"
	"server/serializer"
)

// LoginService 管理用户登录的服务
type LoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=2,max=10"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// setSession 设置session
func (service *LoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("id", user.ID)
	_ = s.Save()
}

// Login 用户登录函数
func (service *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	var count = 0

	_ = database.DB.QueryRow(`SELECT  COUNT(username) FROM user WHERE username = ?`, service.UserName).Scan(&count)
	_ = database.DB.QueryRow(`SELECT * FROM user WHERE username = ?`, service.UserName).Scan(&user.ID, &user.UserName, &user.Password, &user.Authority, &user.CreatedAt)
	if count == 0 {
		return serializer.ParamErr("用户名或密码错误", nil)
	}
	if user.Password != service.Password {
		return serializer.ParamErr("用户名或密码错误", nil)
	}
	// 设置session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
