package config

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type config struct {
	EnableOAuth    bool `json:"enableOAuth"`
	EnablePassword bool `json:"enablePassword"`
}

func Echo(c *gin.Context) {
	cfg := config{}
	if os.Getenv("KOALA_NAME") != "" {
		cfg.EnableOAuth = true
	}
	if os.Getenv("KOALA_PASSWORD_LOGIN") == "enable" {
		cfg.EnablePassword = true
	}
	c.JSON(http.StatusOK, cfg)
}
