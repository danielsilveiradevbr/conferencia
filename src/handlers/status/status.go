package statusHandle

import (
	"net/http"

	statusService "github.com/danielsilveiradevbr/conferencia/src/services/status"
	"github.com/gin-gonic/gin"
)

func StatusRoutes(router *gin.Engine) {
	router.GET("/status", status)
}

func status(c *gin.Context) {

	c.JSON(http.StatusOK,
		statusService.GetStatus(),
	)

}
