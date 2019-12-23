package initial

import (
	"github.com/mkzilla/koala/pkg/types/config"
)

func InitOAuth() {
	if config.Configs.GithubLogin.Enable {
		config.Configs.GithubLogin.Config.ClientID = config.Configs.GithubLogin.ClientID
		config.Configs.GithubLogin.Config.ClientSecret = config.Configs.GithubLogin.Secret
		config.Configs.GithubLogin.Config.Endpoint.AuthURL = "https://github.com/login/oauth/authorize"
		config.Configs.GithubLogin.Config.Endpoint.TokenURL = "https://github.com/login/oauth/access_token"
		config.Configs.GithubLogin.API = "https://api.github.com/user"
		config.Configs.GithubLogin.Name = "github"
	}
	if config.Configs.OAuthV2BasicLogin.Enable {
		config.Configs.OAuthV2BasicLogin.Config.ClientID = config.Configs.OAuthV2BasicLogin.ClientID
		config.Configs.OAuthV2BasicLogin.Config.ClientSecret = config.Configs.OAuthV2BasicLogin.Secret
		config.Configs.OAuthV2BasicLogin.Config.Endpoint.AuthURL = config.Configs.OAuthV2BasicLogin.BasicOauthUserMap.AuthURL
		config.Configs.OAuthV2BasicLogin.Config.Endpoint.TokenURL = config.Configs.OAuthV2BasicLogin.BasicOauthUserMap.TokenURL
		config.Configs.OAuthV2BasicLogin.Name = "basic"
	}
}
