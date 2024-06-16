package routes

import (
	"github.com/adrianomonteiroweb/library-golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
    router.GET("/books", controllers.FindBooks)
}
