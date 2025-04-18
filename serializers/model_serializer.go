package serializers

import (
	"gorm.io/gorm"
)

// ModelSerializer is directly tied to a GORM model, handling CRUD serialization.
type ModelSerializer struct {
	BaseSerializer
	Model interface{} // Model reference
	DB    *gorm.DB    // Database instance for validation queries
}
