package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Store struct {
    storePath string
    pubKeyPath string
    privKeyPath string
}

func NewStore(storePath, pubKeyPath, privKeyPath string) *Store {
    return &Store{storePath: storePath, pubKeyPath: pubKeyPath, privKeyPath: privKeyPath}
}

func (s *Store) Load() {
    // Load the store
}

func (s *Store) Save(name, host string, sequence []int, delay int, udp bool) {
    if checkStore(s.storePath) && checkKey(s.pubKeyPath, s.privKeyPath) {
        if _, err := os.Stat(filepath.Join(s.storePath, name)); err == nil {   
            fmt.Println("Error: sequence already exists")
            return
        } else if errors.Is(err, os.ErrNotExist) {
            file, err := os.Create(filepath.Join(s.storePath, fmt.Sprintf("%s.yaml", name)))
            if err != nil {
                fmt.Println("Error: creating the sequence file")
                return
            }
            defer file.Close()
            file.WriteString("---\n")
            file.WriteString("host: "+ host + "\n")
            file.WriteString("sequence:\n")
            for _, s := range sequence {
                port := strconv.Itoa(s)
                file.WriteString("  - " + port + "\n")
            }
            delayStr := strconv.Itoa(delay)
            file.WriteString("delay: " + delayStr + "\n")
            if udp {
                file.WriteString("udp: " + "true" + "\n")
            } else {
                file.WriteString("udp: " + "false" + "\n")
            }
        } else {
            fmt.Println("Error: store or keys not found")
            return
        }
        fmt.Println("Sequence saved")
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



