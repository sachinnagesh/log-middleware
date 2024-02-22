package cache

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/log-middleware/config"
	"github.com/sachinnagesh/log-middleware/internal/model"
	"github.com/sachinnagesh/log-middleware/internal/service"
)

var cacheStore *logCache

type logCache struct {
	mu   sync.Mutex
	logs []model.LogPayload
}

func GetLogCache() *logCache {
	if cacheStore == nil {
		cacheStore = &logCache{logs: []model.LogPayload{}}
	}
	return cacheStore
}

func InitLogCache() {
	log.Info("Intializing log cache")
	cache := GetLogCache()
	//start ticker
	batchInterval, err := strconv.Atoi(os.Getenv("BATCH_INTERVAL"))
	if err != nil {
		log.Error("Error getting BATCH_INTERVAL : ", err.Error())
	}
	ticker := time.NewTicker(time.Duration(batchInterval) * time.Second)

	config.Wg.Add(1)
	//start goroutine for ticker
	go func(ticker *time.Ticker) {
		defer config.Wg.Done()
		for range ticker.C {

			if len(cache.logs) > 0 {
				cache.mu.Lock()
				config.Wg.Add(1)
				go service.PostLog(cache.logs)
				cache.logs = []model.LogPayload{}
				cache.mu.Unlock()
			} else {
				log.Info("No logs to post")
			}

		}
	}(ticker)

}

func (lc *logCache) ForwardCacheData() {

}

func (lc *logCache) AddLog(logPayload model.LogPayload) error {
	lc.mu.Lock()

	lc.logs = append(lc.logs, logPayload)

	if os.Getenv("BATCH_SIZE") != "" {
		batchSize, err := strconv.Atoi(os.Getenv("BATCH_SIZE"))
		if err != nil {
			log.Error("Error getting BATCH_SIZE : ", err.Error())
		}
		if lc.GetLogsCount() == batchSize {
			config.Wg.Add(1)
			go service.PostLog(lc.logs)
			lc.logs = []model.LogPayload{}
		}
	}
	lc.mu.Unlock()
	return nil
}

func (lc *logCache) GetLogsCount() int {
	return len(lc.logs)

}

func (lc *logCache) ClearCache() bool {
	return true

}
