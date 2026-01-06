package user

import (
	"github.com/gin-gonic/gin"
	"url-shortener/database"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"url-shortener/models"
    "strings"
    "url-shortener/utils"
)

func Signup(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashed_password,_ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{
        Email:  input.Email,
        Password:  string(hashed_password),
    }

	if err := database.DB.Create(&user).Error; err != nil {
        if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "An account with this email already exists. Try signing in instead.",
			})
			return
		}
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Signin(c *gin.Context) {
    var input struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token, err := utils.GenerateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "message": "Login successful"
    })
}