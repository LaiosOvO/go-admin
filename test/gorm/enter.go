package gorm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type TestApi struct {
}

// customer test
func (e *TestApi) GetExaCustomerList() {

	var CustomerList []example.ExaCustomer
	db := global.GVA_DB.Model(&example.ExaCustomer{})
	err := db.Preload("SysUser").Find(&CustomerList).Error
	if err != nil {
		panic(err)
	}

	for _, ite := range CustomerList {
		fmt.Println(ite.ID)
		fmt.Println(ite.SysUser)
	}

	err = db.Find(&CustomerList).Error
	for _, ite := range CustomerList {
		fmt.Println(ite)
		fmt.Println(ite.SysUser)
	}
}
