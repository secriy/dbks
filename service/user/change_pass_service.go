package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"server/database"
	"server/model"
	"server/serializer"
)

// ChangePassService 密码更新服务
type ChangePassService struct {
	OldPass  string `form:"old_pass" json:"old_pass" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// Change 密码更新
func (service *ChangePassService) Change(c *gin.Context) serializer.Response {
	var createAt time.Time

	userC := model.User{
		ID: CurrentUser(c).ID,
	}

	// 判断原密码是否正确
	_ = database.DB.QueryRow(`SELECT password,create_at FROM user WHERE id = ?`, userC.ID).Scan(&userC.Password, &createAt)

	// 验证密码
	if userC.CheckPassword(service.OldPass) == false {
		return serializer.Response{
			Code: 50009,
			Msg:  "原密码错误",
		}
	}

	// 加密密码
	if err := userC.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 时间戳
	userC.CreatedAt = createAt

	// 更新密码
	_, err := database.DB.Exec(`
				UPDATE user SET password = ?
				WHERE id = ?`, userC.Password, userC.ID)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "密码更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(userC),
		Msg:  "密码更新成功",
	}
}
