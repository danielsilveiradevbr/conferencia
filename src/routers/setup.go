package routers

import (
	statusHandle "github.com/danielsilveiradevbr/conferencia/src/handlers/status"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	statusHandle.StatusRoutes(router)
	return router
}
