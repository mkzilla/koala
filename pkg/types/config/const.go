package config

import (
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

var (
	GinEngine     = gin.Default()
	Version       = ""
	DBEngine      *xorm.Engine
	RsaPrivateKey *rsa.PrivateKey
	Configs       *Config
)

const (
	User       = "X-USER"
	PageSize   = "pageSize"
	PageNumber = "pageNumber"
	PageOffset = "pageOffset"
)

type QueryParams struct {
	PageSize   int
	PageNumber int
}
