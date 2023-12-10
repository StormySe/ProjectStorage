package config

import (
	"io"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
  DBName string `yaml:"dbfile"`
  DemoMode bool `yaml:"demomode"`
}

func (c *Config) Parse() {
  cfgFile, _ := os.Open("config.yml")
  defer cfgFile.Close()
  conf, _ := io.ReadAll(cfgFile)
  err := yaml.Unmarshal(conf, &c)
  if err != nil {
    panic(err)
  }
}

func (c *Config) Update() {
  cfgFile, _ := os.OpenFile("config.yml", os.O_WRONLY, os.ModeAppend)
  defer cfgFile.Close()
  newConf, err := yaml.Marshal(c)
  if err != nil {
    panic(err)
  }
  cfgFile.Write(newConf)
}

var AppConfig Config = Config{}

func ConfigInit() {
  AppConfig.Parse()
}
