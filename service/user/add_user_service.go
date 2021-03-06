package user

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
	"server/util"
)

// AddUserService 增加用户服务
type AddUserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=2,max=10"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=16"`
	Authority uint8  `form:"authority" json:"authority" binding:"required,numeric"`
}

// valid 验证表单
func (service *AddUserService) valid() *serializer.Response {
	count := 0
	_ = database.DB.QueryRow(`SELECT  COUNT(username) FROM user WHERE username = ?`, service.UserName).Scan(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已被占用",
		}
	}
	if service.Authority > 127 {
		return &serializer.Response{
			Code: 40007,
			Msg:  "Authority只能为小于128的正整数",
		}
	}
	return nil
}

// Create 用户创建
func (service *AddUserService) Create() serializer.Response {
	user := model.User{
		UserName:  service.UserName,
		Password:  service.Password,
		Authority: service.Authority,
		CreatedAt: time.Now(),
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	res, err := database.DB.Exec(`INSERT INTO dbks.user(username,password,authority,create_at )
		VALUES (?,?,?,?)`, user.UserName, user.Password, user.Authority, user.CreatedAt)
	if res != nil {
		tmp, _ := res.LastInsertId()
		user.ID = uint64(tmp)
	}

	if err != nil {
		util.Log().Panic("创建用户失败", err)
	}

	return serializer.BuildUserResponse(user)
}
