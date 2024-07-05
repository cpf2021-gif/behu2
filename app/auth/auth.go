package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"behu2/app/auth/internal/config"
	"behu2/app/auth/internal/handler"
	"behu2/app/auth/internal/svc"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"gopkg.in/yaml.v2"
)

var configFile = flag.String("f", "etc/auth-api.yaml", "the config file")
var GlobalConfig *Config

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("http://localhost:3000", "http://127.0.0.1:3000"))
	defer server.Stop()

	err := loadConfig("app.yaml")
	if err != nil {
		panic(err)
	}

	initAuthConfig()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

type ServerConfig struct {
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Organization string `yaml:"organization"`
	Application  string `yaml:"application"`
}

type Config struct {
	Certificate string       `yaml:"certificate"`
	Server      ServerConfig `yaml:"server"`
}

func loadConfig(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	GlobalConfig = &cfg

	return nil
}

func initAuthConfig() {
	casdoorsdk.InitConfig(
		GlobalConfig.Server.Endpoint,
		GlobalConfig.Server.ClientID,
		GlobalConfig.Server.ClientSecret,
		GlobalConfig.Certificate,
		GlobalConfig.Server.Organization,
		GlobalConfig.Server.Application,
	)
}
