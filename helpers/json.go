package helpers

import (
	"encoding/json"
	"net/http"
)

// DecodeJSONBody decodes a JSON payload into a struct
func DecodeJSONBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

// EncodeJSONResponse encodes a struct into a JSON payload
func EncodeJSONResponse(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}
