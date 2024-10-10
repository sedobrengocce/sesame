package configdefaults

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckConfigPath() {
    if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
        os.MkdirAll(filepath.Dir(ConfigPath), 0755)
        yamlFile, err := os.Create(ConfigPath)
        if err != nil {
            fmt.Printf("Error creating the config file: %v\n", err)
            panic(err)
        }
        defer yamlFile.Close()
        yamlFile.WriteString("---\n")
        yamlFile.WriteString("paths:\n")
        yamlFile.WriteString("  store: " + StorePath + "\n")
        yamlFile.WriteString("  pubkey: " + PubKeyPath + "\n")
        yamlFile.WriteString("  privkey: " + PrivKeyPath + "\n")
    }
}
