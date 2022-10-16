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

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    userId,
		"created_at": Photo.CreatedAt,
	})
}

func PhotoGetAll(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	err := db.Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Photos := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Photos)
	} else {
		c.ShouldBind(&Photos)
	}
	Photos.ID = photoId

	fmt.Println(Photos)

	Result := map[string]interface{}{}
	SqlStatement := "Update photos SET title = ?, caption = ?, photo_url = ?, updated_at = ? WHERE id = ? RETURNING id, title, caption, photo_url, user_id, updated_at"
	err := db.Raw(
		SqlStatement,
		Photos.Title, Photos.Caption, Photos.PhotoUrl, time.Now(), photoId,
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

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}
	err := db.Delete(Photo, uint(photoId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
