
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

func SuccessMessageResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	temp := &errdata{Message: "success", Status: msg}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

func ErrorResponse(error string, writer http.ResponseWriter) {
	type errdata struct {
		Message string `json:"message"`
	}
	temp := &errdata{Message: error}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

func SuccessOneRespond(fields interface{}, modelType string, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}
	temp := &data{Data: fields, Message: "success"}
	if err != nil {
		color.Red("Marshal Data Failed in SuccessOneRespond() for Type(%v)...", modelType)
		ServerErrResponse(err.Error(), writer)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	switch modelType {
	case "User":
		temp.Data = fields.(models.User)
	case "Role":
		temp.Data = fields.(models.Role)
	default:
		color.Red("Invalid Model Type in SuccessOneRespond() for Type ( %v )...", modelType)
	}

	json.NewEncoder(writer).Encode(temp)
}
