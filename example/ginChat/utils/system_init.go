package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("example/ginChat/config")
	viper.SetConfigName("app")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("config init err")
	}
	fmt.Println("config app inited")
	fmt.Println(viper.GetString("mysql.dns"))
}
