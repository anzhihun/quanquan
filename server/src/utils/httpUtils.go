package utils

import (
	"net/http"
	"strings"
)

func GetLanguageId(header http.Header) string {
	return strings.TrimSpace(strings.Split(header["Accept-Language"][0], ",")[0])
}
