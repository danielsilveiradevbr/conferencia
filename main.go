package main

import (
	"net/http"
	"os"

	"github.com/danielsilveiradevbr/conferencia/src/helper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		helper.NewLog(2, "Passou pelo recover")
		recover()
	}()

	err := godotenv.Load()
	if err != nil {
		helper.NewLog(1, err.Error())
	}

	if os.Getenv("DEBUGANDO") != "T" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	porta := os.Getenv("PORT_LISTEN")
	if porta == "" {
		porta = "3001"
	}
	porta = ":" + porta
	println("Iniciou na porta: " + porta)
	http.ListenAndServe(porta, router)

}
