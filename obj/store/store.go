package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Store struct {
    storePath string
    pubKeyPath string
    privKeyPath string
}

func NewStore(storePath, pubKeyPath, privKeyPath string) *Store {
    return &Store{storePath: storePath, pubKeyPath: pubKeyPath, privKeyPath: privKeyPath}
}

func (s *Store) Load(name string, verbose bool) {
    if checkStore(s.storePath) && checkKey(s.pubKeyPath, s.privKeyPath) {
        if _, err := os.Stat(filepath.Join(s.storePath, name)); err == nil {
            fmt.Println("Sequence found")
            if verbose {
                fmt.Println("Verbose output")
            }
        } else if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Error: sequence not found")
        } else {
            fmt.Println("Error: store or keys not found")
        }
    } else {
        fmt.Println("Error: store or keys not found")
    }
}

func (s *Store) Save(name, host string, sequence []int, delay int, udp bool) {
    if checkStore(s.storePath) && checkKey(s.pubKeyPath, s.privKeyPath) {
        if _, err := os.Stat(filepath.Join(s.storePath, name)); err == nil {   
            fmt.Println("Error: sequence already exists")
            return
        } else if errors.Is(err, os.ErrNotExist) {
            seq := make(map[string]interface{})
            seq["name"] = name
            seq["host"] = host
            seq["ports"] = sequence
            seq["delay"] = delay
            seq["udp"] = udp
            fmt.Println(seq)
            yamlData, err := yaml.Marshal(&seq)
            if err != nil {
                fmt.Println("Error: failed to marshal sequence")
                return
            }
            fmt.Println(string(yamlData))
            file, err := os.Create(filepath.Join(s.storePath, name + ".yaml"))
            if err != nil {
                fmt.Println("Error: failed to create sequence file")
                return
            }
            defer file.Close()
            _, err = io.WriteString(file, string(yamlData))
            if err != nil {
                fmt.Println("Error: failed to write sequence to file")
                return
            }
        } else {
            fmt.Println("Error: store or keys not found")
            return
        }
        fmt.Println("Sequence saved")
    } else {
        fmt.Println("Error: store or keys not found")
    }
}

func (s *Store) Show() {
    // Show the store
}

func (s *Store) Delete() {
    // Delete the store
}

func (s *Store) Encrypt() {
    // Encrypt the store
}

func (s *Store) Decrypt() {
    // Decrypt the store
}

func checkStore(storePath string) bool {
    if _, err := os.Stat(storePath); err == nil {
        return true
    } else if errors.Is(err, os.ErrNotExist) {
        os.MkdirAll(storePath, 0755)
        return true
    } 
    return false
}

func checkKey(pubKeyPath, privKeyPath string) bool {
    if _, err := os.Stat(pubKeyPath); err == nil {
        if _, err := os.Stat(privKeyPath); err == nil {
            return true
        } else if errors.Is(err, os.ErrNotExist) {
            return true

        } 
    } else if errors.Is(err, os.ErrNotExist) {
        return true

    }
    return true
}



