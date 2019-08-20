package renderer

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

func GetRenderFunctionAndContentType(r *http.Request) (func(v interface{}) ([]byte, error), string, error) {
	switch r.Header.Get("Accept") {
	case "application/json":
		return json.Marshal, "application/json", nil
	case "application/xml":
		return xml.Marshal, "application/xml", nil
	default:
		return nil, "", errors.New("content-type not supported")
	}
}