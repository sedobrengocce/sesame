// +build !windows

package configdefaults

import (
	"os"
	"path/filepath"
)

var (
    StorePath = filepath.Join(os.Getenv("HOME"), ".sesame-store", "store")
    PubKeyPath = filepath.Join(os.Getenv("HOME"), ".sesame-store", "pub.key")
    PrivKeyPath = filepath.Join(os.Getenv("HOME"), ".config", "sesame", "priv.key")
    ConfigPath = filepath.Join(os.Getenv("HOME"), ".config", "sesame", "config.yaml")
)
