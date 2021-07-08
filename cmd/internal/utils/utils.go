package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// DecodeJSON is a convenience function to create a JSON decoder
// set it up to disallow unknown fields and then decode into the
// given value.
func DecodeJSON(data io.Reader, out interface{}) error {
	if data == nil {
		return io.EOF
	}

	decoder := json.NewDecoder(data)
	decoder.DisallowUnknownFields()
	return decoder.Decode(&out)
}

// RenderJSON render encodes given data to JSON and writes it to response.
func RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// JSON render encodes given data to JSON and writes it to response.
func JSON(w http.ResponseWriter, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

// Empty represents an empty response.
// swagger:response emptyResponse
type Empty struct{}

// Error represents a JSON encoded API error.
// swagger:response errorResponse
type Error struct {
	Message string `json:"message"`
}

// BoolResponse represents a JSON encoded API for bool responses.
// swagger:response boolResponse
type BoolResponse struct {
	Status bool `json:"status"`
}

// BadRequest Error Helper
func BadRequestError(w http.ResponseWriter, msg string) {
	JSON(w, http.StatusBadRequest, Error{Message: msg})
}
