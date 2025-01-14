package test

import "github.com/flipped-aurora/gin-vue-admin/server/test/gorm"

var TestInstance = new(TestInstanceDef)

type TestInstanceDef struct {
	Api gorm.TestApi
}

type RouterGroup struct {
}
