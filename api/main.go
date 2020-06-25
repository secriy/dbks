package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"server/model"
	"server/serializer"
)

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", e.Field, e.Tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}