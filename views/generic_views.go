package views

import (
	"goda/config"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAPIView handles retrieving multiple objects.
type ListAPIView struct {
	APIView
	Model     interface{}                // GORM model type
	DB        *gorm.DB                   // Database connection
	QueryFunc func(db *gorm.DB) *gorm.DB // Optional custom query function
}

// List handles GET requests to list objects.
func (v *ListAPIView) List(c *gin.Context) {
	query := v.DB.Model(v.Model)
	if v.QueryFunc != nil {
		query = v.QueryFunc(query)
	}

	var results []interface{}
	if err := query.Find(&results).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to fetch objects")
		return
	}

	v.RespondSuccess(c, results)
}

// CreateAPIView handles creating a new object.
type CreateAPIView struct {
	APIView
	Model      interface{} // GORM model type
	Serializer any         // Serializer instance
}

// Create handles POST requests to create objects.
func (v *CreateAPIView) Create(c *gin.Context) {
	if err := c.ShouldBindJSON(v.Serializer); err != nil {
		v.RespondError(c, http.StatusBadRequest, "Invalid input data")
		return
	}

	// Convert serializer to model
	model := v.Serializer.(interface{ ToModel() interface{} }).ToModel()

	// Save to database
	if err := config.DB.Create(model).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to create object")
		return
	}

	// Return created object
	v.RespondCreated(c, model)
}
