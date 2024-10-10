package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
    Paths struct {
        StorePath  string `yaml:"store"`
        PubKeyPath string `yaml:"pubkey"`
        PrivKeyPath string `yaml:"privkey"`
    } `yaml:"paths"`
}

var once sync.Once
var instance *Config

func Get(cfgPath string) (*Config, error) {
    var err error = nil
    once.Do(func() {
        yamlfile, e := os.ReadFile(cfgPath)
        if e != nil {
            err = e
            return
        }
        err = yaml.Unmarshal(yamlfile, &instance)
        if err != nil {
            return
        }
    })
    return instance, err
}
