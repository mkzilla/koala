package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c2 "github.com/mkzilla/koala/pkg/types/config"
)

type config struct {
	EnableGithub       bool `json:"enableGithub"`
	EnableBasicOAuthV2 bool `json:"enableBasicOAuthV2"`
	EnablePassword     bool `json:"enablePassword"`
}

func Echo(c *gin.Context) {
	cfg := config{
		c2.Configs.GithubLogin.Enable,
		c2.Configs.OAuthV2BasicLogin.Enable,
		c2.Configs.BasicLogin.Enable,
	}
	c.JSON(http.StatusOK, cfg)
}
