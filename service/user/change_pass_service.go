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
	var old string
	var createAt time.Time

	userC := model.User{
		ID: CurrentUser(c).ID,
	}

	_ = database.DB.QueryRow(`SELECT password,create_at FROM user WHERE id = ?`, userC.ID).Scan(&old, &createAt)
	if old != service.OldPass {
		return serializer.Response{
			Code: 50009,
			Msg:  "原密码错误",
		}
	}

	userC.Password = service.Password
	userC.CreatedAt = createAt

	_, err := database.DB.Exec(`
				UPDATE user SET password = ?
				WHERE id = ?`, service.Password, userC.ID)
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
