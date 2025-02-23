package main

import (
	"os"

	"github.com/danielsilveiradevbr/conferencia/src/helper"
	"github.com/danielsilveiradevbr/conferencia/src/routers"

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
	router := routers.SetupRouter()

	porta := os.Getenv("PORT_LISTEN")
	if porta == "" {
		porta = "3001"
	}
	porta = ":" + porta
	println("Iniciou na porta: " + porta)
	helper.NewLog(1, "Iniciou na porta: "+porta)
	router.Run(porta)

}
