package main

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	src2 "golang_tech_stack/special/swagger/src"
	_ "golang_tech_stack/swagger/docs"
)

//making an instance of the type DB from the gorm package
//var db *gorm.DB = nil
var err error

// @title gin+gorm crud 测试swagger（必填）
// @version 1.0 （必填）
// @description gin+gorm crud 测试swagger
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
func main() {
	////establishing connection with mysql database 'CRUD'
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Silent),
	//})
	//handle the error comes from the connection with DB
	//if err != nil {
	//	panic(err.Error())
	//}
	server := gin.Default()

	//set up the different routes
	server.GET("/posts", src2.Posts)
	server.GET("/posts/:id", src2.Show)
	server.POST("/posts", src2.Store)
	server.PATCH("/posts/:id", src2.Update)
	server.DELETE("/posts/:id", src2.Delete)

	server.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	server.Run(":8000")
}