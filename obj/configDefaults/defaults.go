package configdefaults

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sedobrengocce/sesame/obj/config"
	"gopkg.in/yaml.v3"
)

func CheckConfigPath() {
    if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
        os.MkdirAll(filepath.Dir(ConfigPath), 0755)
        cfg := config.Config{
            Paths: struct {
                StorePath  string `yaml:"store"`
                PubKeyPath string `yaml:"pubkey"`
                PrivKeyPath string `yaml:"privkey"`
            }{
                StorePath:  StorePath,
                PubKeyPath: PubKeyPath,
                PrivKeyPath: PrivKeyPath,
            },
        }
        yamlFile, err := yaml.Marshal(&cfg)
        if err != nil {
            fmt.Printf("Error marshalling the config file: %v\n", err)
            panic(err)
        }
        cfgFile , err := os.Create(ConfigPath)
        if err != nil {
            fmt.Printf("Error creating the config file: %v\n", err)
            panic(err)
        }
        defer cfgFile.Close()
        _, err = io.WriteString(cfgFile, string(yamlFile))
        if err != nil {
            fmt.Printf("Error writing the config file: %v\n", err)
            panic(err)
        }
    }
}
