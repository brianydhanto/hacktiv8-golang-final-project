package controllers

import (
	"hacktiv8-golang-final-project/database"
	"hacktiv8-golang-final-project/helpers"
	"hacktiv8-golang-final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediasCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          userId,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func SocialMediaGetAll(c *gin.Context) {
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

func SocialMediaUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}
	socialMedia.ID = socialMediaId

	Result := map[string]interface{}{}
	SqlStatement := "Update photos SET name = ?, social_media_url = ?, updated_at = ? WHERE id = ? RETURNING id, name, social_media_url, user_id, updated_at"
	err := db.Raw(
		SqlStatement,
		socialMedia.Name, socialMedia.SocialMediaUrl, time.Now(), socialMediaId,
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

func SocialMediaDelete(c *gin.Context) {
	db := database.GetDB()
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	socialMedia := models.SocialMedia{}
	err := db.Delete(socialMedia, uint(socialMediaId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
