package initial

import (
	"os"

	"github.com/mkzilla/koala/pkg/types/config"
)

func InitOAuth() {

	cfg := new(config.OAuth)

	cfg.Name = os.Getenv("KOALA_NAME")
	cfg.Config.ClientID = os.Getenv("KOALA_CLIENT_ID")
	cfg.Config.ClientSecret = os.Getenv("KOALA_CLIENT_SECRET")
	cfg.Config.Endpoint.AuthURL = os.Getenv("KOALA_AUTH_URL")
	cfg.Config.Endpoint.TokenURL = os.Getenv("KOALA_TOKEN_URL")
	cfg.API = os.Getenv("KOALA_API")
	cfg.NicknameMap = os.Getenv("KOALA_NICKNAME_MAP")
	cfg.UsernameMap = os.Getenv("KOALA_USERNAME_MAP")
	cfg.EmailMap = os.Getenv("KOALA_EMAIL_MAP")

	config.OAuthConfig = cfg
}
