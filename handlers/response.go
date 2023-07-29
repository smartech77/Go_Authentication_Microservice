
package middlewares

import (
	"encoding/json"
	http "github.com/helios/go-sdk/proxy-libs/helioshttp"

	"github.com/fatih/color"
	"tracio.com/userservice/models"
)
func AuthorizationResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Message string `json:"message"`
	}
	temp := &errdata{Message: msg}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(writer).Encode(temp)
}
