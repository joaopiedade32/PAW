package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get users",
		"users":   service.GetAllUsers(),
	})
}

func Profile(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.GetUserProfile(userId)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "get user",
		"user":    user,
	})
}

func Register(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	user = service.RegisterUser(user)
	c.JSON(200, gin.H{
		"message": "register user",
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	user.ID = userId
	user, err = service.UpdateUserProfile(user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "update user profile",
		"user":    user,
	})
}

func DeleteAccount(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.DeleteUser(userId)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}
