/*
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-05 00:08:11
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-05 00:14:58
 * @FilePath: \study\03_golang_tech_stack\special\viper\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadIni() {
	v := viper.New()
	v.AddConfigPath("./conf") // 路径
	v.SetConfigName("config") // 名称
	v.SetConfigType("ini")    // 类型

	err := v.ReadInConfig() // 读配置
	if err != nil {
		panic(err)
	}

	s := v.GetString("db.username")
	fmt.Printf("is: %v\n", s)
}

func ReadYml() {
	v := viper.New()
	v.AddConfigPath("./conf") // 路径
	v.SetConfigName("config") // 名称
	v.SetConfigType("yaml")   // 类型

	err := v.ReadInConfig() // 读配置
	if err != nil {
		panic(err)
	}

	s := v.GetString("db.username")
	fmt.Printf("ys: %v\n", s)
}

func main() {
	ReadIni()

	ReadYml()
	
}
