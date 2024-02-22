package config

import (
	"os"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

var Wg sync.WaitGroup

func LoadConfig() {
	log.Info("Loading configuration")

	batchSize, set := os.LookupEnv("BATCH_SIZE")
	if !set || batchSize == "" {
		log.Info("BATCH_SIZE env is not set. Setting it to default")
		os.Setenv("BATCH_SIZE", "100")
	}

	batchInterval, set := os.LookupEnv("BATCH_INTERVAL")
	if !set || batchInterval == "" {
		log.Info("BATCH_INTERVAL env is not set. Setting it to default")
		os.Setenv("BATCH_INTERVAL", "600")

	}

	postEndpoint, set := os.LookupEnv("POST_ENDPOINT")
	if !set || postEndpoint == "" {
		log.Info("POST_ENDPOINT env is not set. Setting it to default")
		os.Setenv("POST_ENDPOINT", "http://0.0.0.0:3001/log-collector-srv/bulk/log")
	}
}
