package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"server/serializer"
	"server/service/user"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var loginService user.LoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	currentUser := CurrentUser(c)
	res := serializer.BuildUserResponse(*currentUser)
	c.JSON(200, res)
}

// ChangePass 修改密码
func ChangePass(c *gin.Context) {
	changePassService := user.ChangePassService{}
	if err := c.ShouldBind(&changePassService); err == nil {
		res := changePassService.Change(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ResrtPass 重置密码
func ResrtPass(c *gin.Context) {
	resetPassService := user.ResetPassService{}
	if err := c.ShouldBind(&resetPassService); err == nil {
		res := resetPassService.Reset(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserAdd 用户创建接口
func UserAdd(c *gin.Context) {
	var registerService user.AddUserService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserDelete 用户删除接口
func UserDelete(c *gin.Context) {
	deleteUserService := user.DeleteUserService{}
	res := deleteUserService.Delete(c.Param("id"))
	c.JSON(200, res)
}

// UserUpdate 用户更新接口
func UserUpdate(c *gin.Context) {
	updateUserService := user.UpdateUserService{}
	if err := c.ShouldBind(&updateUserService); err == nil {
		res := updateUserService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListUsers 产品列表接口
func UsersList(c *gin.Context) {
	listUserService := user.ListUsersService{}
	if err := c.ShouldBind(&listUserService); err == nil {
		res := listUserService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
