/*
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-05 09:46:48
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-05 09:57:50
 * @FilePath: \study\03_golang_tech_stack\special\jwt\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var MySecret = []byte("secret")

func GenToken(username string, password string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "laoguo",                             // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 注意这个地方一定要是字节切片不能是字符串
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return MySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	s, err := GenToken("ghz", "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GenToken s: %v\n", s)

	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjQ4NzAwNTUyLCJpc3MiOiJsYW9ndW8ifQ.eD4c_s5tminPbKJgmCr3n9jUnp0LT2I4t0_Fd5gml7U"
	token := s
	mc, err := ParseToken(token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mc.Password: %v\n", mc.Password)
	fmt.Printf("mc.Username: %v\n", mc.Username)
}
