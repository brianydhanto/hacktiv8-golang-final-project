package controllers

import (
	"fmt"
	"hacktiv8-golang-final-project/database"
	"hacktiv8-golang-final-project/helpers"
	"hacktiv8-golang-final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    userId,
		"created_at": Comment.CreatedAt,
	})
}

func CommentGetAll(c *gin.Context) {
	db := database.GetDB()

	Comments := []models.Comment{}
	err := db.Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comments)
}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Comments := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}
	Comments.ID = commentId

	fmt.Println(Comments)

	Result := map[string]interface{}{}
	SqlStatement := "Update comments SET message = ?, updated_at = ? WHERE id = ? RETURNING id, title, message, photo_id, user_id, updated_at"
	err := db.Raw(
		SqlStatement,
		Comments.Message, time.Now(), commentId,
	).Scan(&Result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Result)
}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Photo := models.Photo{}
	err := db.Delete(Photo, uint(commentId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
