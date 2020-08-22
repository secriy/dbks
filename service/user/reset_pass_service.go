package user

import (
	"golang.org/x/crypto/bcrypt"
	"server/database"
	"server/serializer"
	"server/util"
)

// ResetPassService 密码重置服务
type ResetPassService struct {
}

type DefaultPass struct {
	Password string `json:"password"`
}

func (service *ResetPassService) Reset(id string) serializer.Response {
	// 生成随机字符串
	pass := util.RandomPass()

	// 加密密码
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 11)
	if err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 更新密码
	_, err = database.DB.Exec(`
				UPDATE user SET password = ?
				WHERE id = ?`, string(bytes), id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "密码更新失败",
			Error: err.Error(),
		}
	}

	// 返回参数
	return serializer.Response{
		Data: DefaultPass{
			Password: pass,
		},
		Msg: "密码重置成功",
	}
}
