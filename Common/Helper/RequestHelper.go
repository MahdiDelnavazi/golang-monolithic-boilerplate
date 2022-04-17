package Helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(request *http.Request, val interface{}) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	if decodeError := decoder.Decode(val); decodeError != nil {
		fmt.Errorf("decodeError: %w", decodeError)
		return decodeError
	}
	return nil
}
