package views

import (
	"goda/auth"
	"goda/config"
	"goda/models"
	"goda/serializers"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user in the database from posted JSON input.
func CreateUser(c *gin.Context) {
	var userSerializer serializers.UserSerializer
	userSerializer.DB = config.DB // Provide DB connection to serializer

	// Bind JSON input to the serializer
	if err := c.ShouldBindJSON(&userSerializer); err != nil {
		c.JSON(http.StatusBadRequest, serializers.NewSerializerError(err))
		return
	}

	// Validate input
	if err := userSerializer.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user to the database
	user := userSerializer.ToModel()
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user as a JSON response
	userSerializer.FromModel(user)
	c.JSON(http.StatusCreated, userSerializer)
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Goda REST Framework!"})
}

func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON input to loginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
