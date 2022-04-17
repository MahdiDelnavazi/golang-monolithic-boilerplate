package Common

import "encoding/json"

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type FieldErrors []FieldError

func (fieldError FieldErrors) Error() string {
	details, err := json.Marshal(fieldError)
	if err != nil {
		return err.Error()
	}
	return string(details)
}

func (fieldError FieldErrors) Fields() map[string]string {
	fieldWithErrors := make(map[string]string)
	for _, field := range fieldError {
		fieldWithErrors[field.Field] = field.Error
	}
	return fieldWithErrors
}
