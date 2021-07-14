package main

import (
	"esystem/conf"
	"esystem/database"
	"esystem/router"
	"fmt"
)

func main() {
	if err := config.Load("conf/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	_, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}

