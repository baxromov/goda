package views

import (
	"github.com/gin-gonic/gin"
	"goda/config"
	"gorm.io/gorm"
	"net/http"
)

// ModelViewSet handles CRUD operations for a specific model.
type ModelViewSet struct {
	APIView
	Model      interface{} // GORM model type
	Serializer interface{} // Serializer instance
	DB         *gorm.DB    // Database connection
}

// List retrieves a list of all objects.
// @Summary List all objects
// @Description Retrieve all records for the given model.
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (v *ModelViewSet) List(c *gin.Context) {
	var results []interface{}
	if err := config.DB.Model(v.Model).Find(&results).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to fetch objects")
		return
	}
	v.RespondSuccess(c, results)
}

// Retrieve retrieves a single object by ID.
func (v *ModelViewSet) Retrieve(c *gin.Context) {
	id := c.Param("id")
	var result interface{}
	if err := config.DB.Model(v.Model).First(&result, id).Error; err != nil {
		v.RespondError(c, http.StatusNotFound, "Object not found")
		return
	}
	v.RespondSuccess(c, result)
}

// Create a new object.
// @Summary Create an object
// @Description Create a new record for the given model.
// @Tags users
// @Accept json
// @Produce json
// @Param data body serializers.UserSerializer true "User Data"
// @Success 201 {object} models.User
// @Router /users [post]
func (v *ModelViewSet) Create(c *gin.Context) {
	if err := c.ShouldBindJSON(v.Serializer); err != nil {
		v.RespondError(c, http.StatusBadRequest, "Invalid input data")
		return
	}

	// Convert to model and save
	model := v.Serializer.(interface{ ToModel() interface{} }).ToModel()
	if err := config.DB.Create(model).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to create object")
		return
	}

	v.RespondCreated(c, model)
}

// Update updates an object by ID.
func (v *ModelViewSet) Update(c *gin.Context) {
	id := c.Param("id")

	var result interface{}
	if err := config.DB.Model(v.Model).First(&result, id).Error; err != nil {
		v.RespondError(c, http.StatusNotFound, "Object not found")
		return
	}

	if err := c.ShouldBindJSON(v.Serializer); err != nil {
		v.RespondError(c, http.StatusBadRequest, "Invalid input data")
		return
	}

	// Convert serializer to model
	model := v.Serializer.(interface{ ToModel() interface{} }).ToModel()

	// Update object
	if err := config.DB.Model(result).Updates(model).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to update object")
		return
	}

	v.RespondSuccess(c, result)
}

// Delete deletes an object by ID.
func (v *ModelViewSet) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := v.DB.Delete(v.Model, id).Error; err != nil {
		v.RespondError(c, http.StatusInternalServerError, "Failed to delete object")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
