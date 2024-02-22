package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	app, err := InitRouter()
	assert.Nil(t, err)
	assert.NotNil(t, app)
}
