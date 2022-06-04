package main

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "golang_tech_stack/swagger/docs"
	"golang_tech_stack/swagger/src"
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
	server.GET("/posts", src.Posts)
	server.GET("/posts/:id", src.Show)
	server.POST("/posts", src.Store)
	server.PATCH("/posts/:id", src.Update)
	server.DELETE("/posts/:id", src.Delete)

	server.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	server.Run(":8000")
}