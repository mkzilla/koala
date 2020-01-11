package initial

import (
	"github.com/mkzilla/koala/pkg/types/config"
)

func InitConfig(filepath string) {
	cfg := new(config.Config)

	err := cfg.LoadFromJson(filepath)
	if err != nil {
		panic(err)
	}
	cfg.BasicRegister.Template = cfg.Read(cfg.BasicRegister.Template)
	if cfg.TokenMaxAge < 1000 {
		cfg.TokenMaxAge = 86400
	}
	config.Configs = cfg
}
