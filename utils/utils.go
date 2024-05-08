package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

var Validate = validator.New()

func GetIdFromRequest(r *http.Request) (string, error) {
	pathSegments := strings.Split(r.URL.Path, "/")

	if len(pathSegments) < 5 || pathSegments[4] == "" {
		return "", fmt.Errorf("missing ID")
	}

	id := pathSegments[4]
	return id, nil
}

func ParseAndValidateJSON(w http.ResponseWriter, r *http.Request, payload any) error {
	if r.Body == nil {
		WriteError(w, http.StatusBadRequest, fmt.Errorf("missing request body"))
		return nil
	}

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return err
	}

	// Validate the payload
	if err := Validate.Struct(payload); err != nil {
		var missingFields []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			if tag == "required" {
				field = strings.ToLower(field)
				missingFields = append(missingFields, field)
			}
		}
		errorMessage :=
			fmt.Sprintf("Missing required fields: %s", strings.Join(missingFields, ", "))
		WriteError(w, http.StatusBadRequest, fmt.Errorf(errorMessage))
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	message := "An unexpected error occurred. " +
		"Please try again later or contact support if the problem persists."

	WriteJSON(w,
		http.StatusInternalServerError,
		map[string]string{
			"error":   "Internal Server Error",
			"message": message,
		})
}

func PrepareJSONWithMessage(message string, payload interface{}) map[string]interface{} {
	response := map[string]interface{}{
		"message": message,
	}

	val := reflect.ValueOf(payload)
	switch val.Kind() {
	case reflect.Slice:
		response["data"] = handleSlicePayload(val)
	case reflect.Ptr:
		response["data"] = handlePointerPayload(val)
	default:
		response["data"] = createPayloadMap(payload)
	}

	jsonResponse := marshalAndUnmarshalJSON(response)
	return jsonResponse
}

func handleSlicePayload(val reflect.Value) interface{} {
	if val.Len() == 0 {
		return []interface{}{}
	}

	var payloadData []map[string]interface{}
	for i := 0; i < val.Len(); i++ {
		payloadData = append(payloadData, createPayloadMap(val.Index(i).Interface()))
	}
	return payloadData
}

func handlePointerPayload(val reflect.Value) interface{} {
	if val.IsNil() {
		return []interface{}{}
	}
	return createPayloadMap(val.Interface())
}

func createPayloadMap(payload interface{}) map[string]interface{} {
	payloadMap := make(map[string]interface{})
	val := reflect.ValueOf(payload)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typeOfPayload := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typeOfPayload.Field(i).Name
		snakeCaseName := toSnakeCase(fieldName)
		fieldValue := val.Field(i).Interface()

		if gormModel, ok := fieldValue.(gorm.Model); ok {
			handleGormModelField(gormModel, payloadMap)
		} else {
			payloadMap[snakeCaseName] = fieldValue
		}
	}

	return payloadMap
}

func handleGormModelField(gormModel gorm.Model, payloadMap map[string]interface{}) {
	payloadMap["id"] = gormModel.ID
	payloadMap["created_at"] = gormModel.CreatedAt
	payloadMap["updated_at"] = gormModel.UpdatedAt
	if gormModel.DeletedAt.Valid {
		payloadMap["deleted_at"] = gormModel.DeletedAt.Time
	}
}

func marshalAndUnmarshalJSON(response map[string]interface{}) map[string]interface{} {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		return nil
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(jsonResp, &jsonResponse)
	if err != nil {
		return nil
	}

	return jsonResponse
}

func toSnakeCase(str string) string {
	return strcase.ToSnake(str)
}
