package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/log-middleware/internal/handler"
)

func InitRouter() (*fiber.App, error) {
	log.Info("Creating router")
	app := fiber.New()
	basePath := "/log-collector-srv"

	app.Get(basePath+"/ping", handler.Ping)
	app.Post(basePath+"/log", handler.StoreLog)
	app.Get(basePath+"/log/count", handler.GetLogsCount)
	app.Post(basePath+"/bulk/log", handler.BulkLogProcessor)

	return app, nil

}
