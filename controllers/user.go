package controllers

import (
	"fmt"
	"hacktiv8-golang-final-project/database"
	"hacktiv8-golang-final-project/helpers"
	"hacktiv8-golang-final-project/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Users := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&Users)
	} else {
		c.ShouldBind(&Users)
	}
	Users.ID = userId

	fmt.Println(Users)

	Result := map[string]interface{}{}
	SqlStatement := "Update users SET username = ?, email = ?, updated_at = ? WHERE id = ? RETURNING id, email, username, age, updated_at"
	err := db.Raw(
		SqlStatement,
		Users.Username, Users.Email, time.Now(), uint(userId),
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

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))
	User := models.User{}
	err := db.Delete(User, uint(userId)).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
