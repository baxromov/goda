// serializers/base.go
package serializers

import (
	"encoding/json"
	"errors"
	"goda/models "
	"reflect"
)

// Serializer is the base interface for all serializers
type Serializer interface {
	Validate() error
	ToRepresentation(model interface{}) (map[string]interface{}, error)
	ToInternal(data map[string]interface{}) (interface{}, error)
}

// BaseSerializer provides common functionality for serializers
type BaseSerializer struct {
	Data   map[string]interface{}
	Model  interface{}
	Fields []string
}

// Validate performs validation on the serializer data
func (s *BaseSerializer) Validate() error {
	if s.Data == nil {
		return errors.New("serializer data is nil")
	}
	return nil
}

// ToRepresentation converts a model to its representation (map/JSON)
func (s *BaseSerializer) ToRepresentation(model interface{}) (map[string]interface{}, error) {
	if model == nil {
		return nil, errors.New("model is nil")
	}

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	// Filter fields if specified
	if len(s.Fields) > 0 {
		filteredResult := make(map[string]interface{})
		for _, field := range s.Fields {
			if value, exists := result[field]; exists {
				filteredResult[field] = value
			}
		}
		return filteredResult, nil
	}

	return result, nil
}

// ToInternal converts a data map to a model instance
func (s *BaseSerializer) ToInternal(data map[string]interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create a new instance of the model type
	modelType := reflect.TypeOf(s.Model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	newInstance := reflect.New(modelType).Interface()
	if err := json.Unmarshal(jsonData, &newInstance); err != nil {
		return nil, err
	}

	return newInstance, nil
}

// ModelSerializer provides functionality for serializing models
type ModelSerializer struct {
	BaseSerializer
	ModelType reflect.Type
}

// NewModelSerializer creates a new ModelSerializer for the specified model
func NewModelSerializer(model models.Model, fields []string) *ModelSerializer {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return &ModelSerializer{
		BaseSerializer: BaseSerializer{
			Model:  model,
			Fields: fields,
		},
		ModelType: modelType,
	}
}

// ValidateCreate performs validation on create operations
func (s *ModelSerializer) ValidateCreate(data map[string]interface{}) error {
	s.Data = data
	return s.Validate()
}

// ValidateUpdate performs validation on update operations
func (s *ModelSerializer) ValidateUpdate(instance models.Model, data map[string]interface{}) error {
	s.Data = data
	s.Model = instance
	return s.Validate()
}

// Create converts data to a model and validates it for creation
func (s *ModelSerializer) Create(data map[string]interface{}) (models.Model, error) {
	if err := s.ValidateCreate(data); err != nil {
		return nil, err
	}

	instance, err := s.ToInternal(data)
	if err != nil {
		return nil, err
	}

	model, ok := instance.(models.Model)
	if !ok {
		return nil, errors.New("failed to convert to Model interface")
	}

	if err := model.Validate(); err != nil {
		return nil, err
	}

	return model, nil
}

// Update applies the provided data to an existing model instance
func (s *ModelSerializer) Update(instance models.Model, data map[string]interface{}) (models.Model, error) {
	if err := s.ValidateUpdate(instance, data); err != nil {
		return nil, err
	}

	// Convert existing instance to map
	instanceMap, err := s.ToRepresentation(instance)
	if err != nil {
		return nil, err
	}

	// Update the map with new data
	for k, v := range data {
		instanceMap[k] = v
	}

	// Convert back to model
	updatedInstance, err := s.ToInternal(instanceMap)
	if err != nil {
		return nil, err
	}

	updatedModel, ok := updatedInstance.(models.Model)
	if !ok {
		return nil, errors.New("failed to convert to Model interface")
	}

	if err := updatedModel.Validate(); err != nil {
		return nil, err
	}

	return updatedModel, nil
}
