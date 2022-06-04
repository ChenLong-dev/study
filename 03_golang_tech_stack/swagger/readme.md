# Golang使用swagger生成api接口文档

---
现在大部分应用都是前后端分离的项目，那么，前端和后端交互只有通过api接口文档来实现。swagger可以根据注释来生成api接口文档。
### 本课程预备知识

---
- gin
- gorm
- mysql

### 添加库

---
```bigquery
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get gorm.io/gorm/logger
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/gin-swagger/swaggerFiles
```

### gin+gorm创建一个crud项目

---
#### model.go
```bigquery
package main

import (
	"gorm.io/gorm"
)

//this model represent a database table
type Post struct {
	gorm.Model
	Title  string `gorm:"type:varchar(100);" json:"title" example:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" example:"desc" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status" example:"Active"`
}

type Response struct {
	Msg  string
	Data interface{}
}
```

#### api.go
```bigquery
package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

//查询
func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    posts,
	})

}

// @Summary      查询
// @Description  查询
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "pid"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts/{id} [get]
func Show(c *gin.Context) {
	post := getById(c)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    post,
	})
}

// @Summary      添加post
// @Description  添加post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param content body string true "json"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts [post]
func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	post.Status = "Active"
	db.Create(&post)
	c.JSON(http.StatusOK, Response{
		Msg:  "",
		Data: post,
	})
}

//删除
func Delete(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		return
	}
	db.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"messege": "deleted successfuly",
		"data":    "",
	})

}

//更新
func Update(c *gin.Context) {
	oldpost := getById(c)
	var newpost Post
	if err := c.ShouldBindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	oldpost.Title = newpost.Title
	oldpost.Des = newpost.Des
	if newpost.Status != "" {
		oldpost.Status = newpost.Status
	}

	db.Save(&oldpost)

	c.JSON(http.StatusOK, gin.H{
		"messege": "Post has been updated",
		"data":    oldpost,
	})

}

// 根据id查询
func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"messege": "post not found",
			"data":    "",
		})
	}
	return post
}
```
#### main.go
```bigquery
package main

import (
	_ "pro02/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//making an instance of the type DB from the gorm package
var db *gorm.DB = nil
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
	//establishing connection with mysql database 'CRUD'
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Post{})

	server := gin.Default()

	//set up the different routes
	server.GET("/posts", Posts)
	server.GET("/posts/:id", Show)
	server.POST("/posts", Store)
	server.PATCH("/posts/:id", Update)
	server.DELETE("/posts/:id", Delete)

	server.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	server.Run(":8000")
}
```

#### 生成接口文档
```bigquery
swag init
```

---
---

#### 在main.go文件中添加注释
```bigquery
// @title gin+gorm crud 测试swagger（必填）
// @version 1.0 （必填）
// @description gin+gorm crud 测试swagger
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
```

#### 在api.go文件中添加接口的注释
```bigquery
// @Summary      查询
// @Description  查询
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "pid"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts/{id} [get]
func Show(c *gin.Context) {
    ...
}
```

```bigquery
// @Summary      添加post
// @Description  添加post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param content body string true "json"
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
// @Router       /posts [post]
func Store(c *gin.Context) {
    ...
}
```

#### 在main.go文件中引入gin-swagger渲染文档数据
```bigquery
import (
    "docs" // 导入swagger文档用的
    gs "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)
```
#### 在main.go文件中，在注册http路由的地方，注册swagger api相关路由
```bigquery
r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
```

#### 访问测试
```bigquery
http://localhost:8000/swagger/index.html
```


