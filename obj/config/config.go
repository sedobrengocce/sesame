package config

import (
	"os"
	"sync"

	"github.com/sedobrengocce/sesame/obj/configDefaults"
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

func Get() (*Config, error) {
    var err error = nil
    once.Do(func() {
        file, e := os.Open(configdefaults.ConfigPath)
        if e != nil {
            err = e
            return
        }
        defer file.Close()
        decoder := yaml.NewDecoder(file)
        err = decoder.Decode(instance)
        instance = &Config{}
    })
    return instance, err
}
