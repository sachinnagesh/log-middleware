package cache

import (
	"testing"
	"time"

	"github.com/sachinnagesh/log-middleware/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetCache(t *testing.T) {
	cache := GetLogCache()
	assert.NotEqual(t, nil, cache)

}

func TestAddLog(t *testing.T) {
	log := model.LogPayload{
		UserId: 12345,
		Total:  1.65,
		Title:  "test-title",
		Meta: model.Meta{
			Logins: []model.Login{
				{Time: time.Now(), IP: "127.0.0.1"},
			},
		},
		Completed: false,
	}
	cache := GetLogCache()
	err := cache.AddLog(log)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, cache.GetLogsCount())

	log1 := model.LogPayload{
		UserId: 12345,
		Total:  1.65,
		Title:  "test-title",
		Meta: model.Meta{
			Logins: []model.Login{
				{Time: time.Now(), IP: "127.0.0.1"},
			},
		},
		Completed: false,
	}

	err = cache.AddLog(log1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, cache.GetLogsCount())
}

func TestGetLogsCount(t *testing.T) {
	cache := GetLogCache()
	logcount := cache.GetLogsCount()
	assert.Equal(t, 0, logcount)
}

func TestClearCache(t *testing.T) {

	cache := GetLogCache()
	status := cache.ClearCache()
	assert.Equal(t, true, status)

}
