package config

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name              string     `json:"name" yaml:"name"`
	Homepage          string     `json:"homepage" yaml:"homepage"`
	Addr              string     `json:"addr" yaml:"addr"`
	DataSource        DataSource `json:"datasource" yaml:"datasource"`
	BasicLogin        BasicLogin `json:"basicLogin" yaml:"basicLogin"`
	GithubLogin       OAuthIF    `json:"githubLogin" yaml:"githubLogin"`
	GoogleLogin       OAuthIF    `json:"googleLogin" yaml:"googleLogin"`
	WechatLogin       OAuthIF    `json:"wechatLogin" yaml:"wechatLogin"`
	OAuthV2BasicLogin OAuthIF    `json:"oauthV2BasicLogin" yaml:"oauthV2BasicLogin"`
}

func (c *Config) LoadFromYaml(path string) error {
	yamlS, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	err := yaml.Unmarshal(yamlS, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) LoadFromJson(path string) error {
	jsonS, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	err := json.Unmarshal(jsonS, &c)
	if err != nil {
		return err
	}
	return nil
}

type DataSource struct {
	Driver string `json:"driver" yaml:"driver"`
	Source string `json:"source" yaml:"source"`
	Debug  bool   `json:"debug" yaml:"debug"`
	Init   bool   `json:"init" yaml:"init"`
}

type BasicLogin struct {
	Enable bool `json:"enable" yaml:"enable"`
}

type OAuthIF struct {
	Enable   bool          `json:"enable" yaml:"enable"`
	ClientID string        `json:"client_id" yaml:"client_id"`
	Secret   string        `json:"secret" yaml:"secret"`
	API      string        `json:"-" yaml:"-"`
	Name     string        `json:"-" yaml:"-"`
	Config   oauth2.Config `json:"-" yaml:"-"`
}

type GithubUser struct {
	Login   string `json:"login"`
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Company string `json:"company"`
}
