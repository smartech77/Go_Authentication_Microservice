package controllers

import (
	gin "github.com/helios/go-sdk/proxy-libs/heliosgin"

	middlewares "tracio.com/userservice/handlers"
)

var mySigningKey = []byte(middlewares.DotEnvVariable("JWT_SECRET"))

func Test(c *gin.Context) {
	middlewares.SuccessMessageResponse("Congratulations... It's working.", c.Writer)
}
