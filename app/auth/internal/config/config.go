package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Certificate string     `yaml:"Certificate"`
	CasdoorCfg  CasdoorCfg `yaml:"Casdoor"`
}

type CasdoorCfg struct {
	Endpoint     string `yaml:"endpoint"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Organization string `yaml:"organization"`
	Application  string `yaml:"application"`
	FrontEndURL  string `yaml:"frontend_url"`
}
