package main

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
