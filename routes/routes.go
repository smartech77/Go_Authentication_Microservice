package routes

import (
	"github.com/gin-gonic/gin"
)

// Routes -> define endpoints
func Routes(router *gin.Engine) {

	router.GET("/api/v1/userservice/test", middlewares.IsAuthorized(controllers.Test))

}
