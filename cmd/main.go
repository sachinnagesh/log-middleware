package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/log-middleware/cache"
	"github.com/sachinnagesh/log-middleware/config"
	"github.com/sachinnagesh/log-middleware/router"
)

func main() {

	serviceName := "log-collector-srv"

	log.Info("Starting Service : ", serviceName)

	//load default config if env not set
	config.LoadConfig()

	//init cache
	cache.InitLogCache()

	app, err := router.InitRouter()
	if err != nil {
		log.Fatal(fmt.Printf("error creating routes err : %s", err.Error()))
	}

	err = app.Listen(":3001")
	if err != nil {
		log.Fatal(fmt.Printf("Error starting server. err : %s", err.Error()))
	}
	log.Info("Server started successfully!!!")
	config.Wg.Wait()

}
