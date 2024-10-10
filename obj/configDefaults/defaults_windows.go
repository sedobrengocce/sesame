package configdefaults

import (
	"os"
	"path/filepath"
)

var (
    StorePath = filepath.Join(os.Getenv("APPDATA"), ".sesame-store", "store")
    PubKeyPath = filepath.Join(os.Getenv("APPDATA"), ".sesame-store", "pub.key")
    PrivKeyPath = filepath.Join(os.Getenv("APPDATA"), ".config", "sesame", "priv.key")
    ConfigPath = filepath.Join(os.Getenv("APPDATA"), ".config", "sesame", "config.yaml")
)
