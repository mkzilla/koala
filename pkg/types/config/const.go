package config

import (
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"golang.org/x/oauth2"
)

var (
	GinEngine     = gin.Default()
	DBEngine      *xorm.Engine
	Version       = ""
	RsaPrivateKey *rsa.PrivateKey
	OAuthConfig   *OAuth
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

type OAuth struct {
	Config      oauth2.Config
	Name        string
	UsernameMap string
	NicknameMap string
	EmailMap    string
	API         string
}
