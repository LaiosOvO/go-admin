package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/example/ginChat/service/user"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	// 首页

	// 用户
	r.POST("/user/getUserList", user.GetUserList)
	r.POST("/user/createUser", user.CreateUser)
	r.POST("/user/updateUser", user.UpdateUser)
	r.POST("/user/deleteUser", user.DeleteUser)
	r.POST("/user/login", user.Login)

	return r
}
