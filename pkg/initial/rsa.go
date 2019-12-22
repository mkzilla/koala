package initial

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/mkzilla/koala/pkg/types/config"
)

func InitRsaKey() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	config.RsaPrivateKey = privateKey
}
