package test

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/test/ws/chatroom/core/service"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (e *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	// 测试的handler
	testHandler := func(c *gin.Context) {
		response.OkWithMessage("test gin integeration", c)
	}

	testRouter := Router.Group("test")
	{
		testRouter.PUT("customer", testHandler)    // 更新客户
		testRouter.DELETE("customer", testHandler) // 删除客户
	}

	customerRouterWithoutRecord := Router.Group("ws")
	{
		customerRouterWithoutRecord.GET("chat", service.Chat) // 私聊例子
		//customerRouterWithoutRecord.GET("customerList", testHandler) // 获取客户列表
	}

	fmt.Println("************************")
	fmt.Println("init the ws router")
	fmt.Println("************************")
}
