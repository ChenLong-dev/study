package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Posts
//查询
func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	//var posts []Post
	//db.Limit(limit).Offset(offset).Find(&posts)
	posts := fmt.Sprintf("limit:%d, offset:%d", limit, offset)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    posts,
	})

}

// Show
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

// Store
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
	//db.Create(&post)
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
	//db.Unscoped().Delete(&post)
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

	//db.Save(&oldpost)

	c.JSON(http.StatusOK, gin.H{
		"messege": "Post has been updated",
		"data":    oldpost,
	})

}

// 根据id查询
func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	//db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"messege": "post not found",
			"data":    fmt.Sprintf("id:%s", id),
		})
	}
	return post
}
