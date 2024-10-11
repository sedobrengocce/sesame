package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
	"github.com/sedobrengocce/sesame/obj/knocker"
	"golang.org/x/term"
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
            privKeyFile, err := os.Open(s.privKeyPath)
            if err != nil {
                fmt.Println("Error: failed to open private key")
                return
            }
            defer privKeyFile.Close()
            privKeyContent, err := io.ReadAll(privKeyFile)
            if err != nil {
                fmt.Println("Error: failed to read private key")
                return
            }
            fmt.Print("Enter passphrase: ")
            passphrase, err := term.ReadPassword(syscall.Stdin)
            if err != nil {
                fmt.Println("Error: failed to read passphrase")
                return
            }
            fmt.Println()
            privKey, err := crypto.NewPrivateKeyFromArmored(string(privKeyContent), passphrase)
            if err != nil {
                fmt.Println("Error: failed to create private key")
                return
            }
            defer privKey.ClearPrivateParams()
            pgp := crypto.PGP()
            decHandle, err := pgp.Decryption().DecryptionKey(privKey).New()
            if err != nil {
                fmt.Println("Error: failed to create decryption handle")
                fmt.Println(err)
                return
            }
            file, err := os.Open(filepath.Join(s.storePath, name))
            if err != nil {
                fmt.Println("Error: failed to open sequence file")
                return
            }
            defer file.Close()
            seq := make(map[string]interface{})
            encryptedData, err := io.ReadAll(file)
            if err != nil {
                fmt.Println("Error: failed to read sequence file")
                return
            }
            yamlData, err := decHandle.Decrypt(encryptedData, crypto.Armor)
            err = yaml.Unmarshal(yamlData.Bytes(), &seq)
            if err != nil {
                fmt.Println("Error: failed to unmarshal sequence")
                return
            }
            host := seq["host"].(string)
            sequence := []int{}
            for _, s := range seq["ports"].([]interface{}) {
                sequence = append(sequence, int(s.(int)))
            }
            delay := seq["delay"].(int)
            udp := seq["udp"].(bool)
            k := knocker.NewKnocker(host)
            k.WithSequence(sequence)
            k.WithDelay(delay)
            k.WithUdp(udp)
            k.WithVerbose(verbose)
            k.Knock()
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
            pubKeyFile, err := os.Open(s.pubKeyPath)
            if err != nil {
                fmt.Println("Error: failed to open public key")
                return
            }
            defer pubKeyFile.Close()
            pubKeyContent, err := io.ReadAll(pubKeyFile)
            if err != nil {
                fmt.Println("Error: failed to read public key")
                return
            }
            pubKey, err := crypto.NewKeyFromArmored(string(pubKeyContent))
            if err != nil {
                fmt.Println("Error: failed to create public key")
                return
            }
            pgp := crypto.PGP()
            encHanle, err := pgp.Encryption().Recipient(pubKey).New()
            if err != nil {
                fmt.Println("Error: failed to create encryption handle")
                return
            }
            seq := make(map[string]interface{})
            seq["name"] = name
            seq["host"] = host
            seq["ports"] = sequence
            seq["delay"] = delay
            seq["udp"] = udp
            yamlData, err := yaml.Marshal(&seq)
            if err != nil {
                fmt.Println("Error: failed to marshal sequence")
                return
            }
            encData, err := encHanle.Encrypt([]byte(yamlData))
            if err != nil {
                fmt.Println("Error: failed to encrypt sequence")
                return
            }
            armorData, err := encData.Armor()
            if err != nil {
                fmt.Println("Error: failed to armor sequence")
                return
            }
            file, err := os.Create(filepath.Join(s.storePath, name))
            if err != nil {
                fmt.Println("Error: failed to create sequence file")
                return
            }
            defer file.Close()
            _, err = io.WriteString(file, string(armorData))
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

func (s *Store) SetKeys(pubKeyPath, privKeyPath string) {
    sourcePubKeyStat, err := os.Stat(pubKeyPath)
    if err != nil {
        fmt.Println("Error: public key not found")
        return
    }
    if !sourcePubKeyStat.Mode().IsRegular() {
        fmt.Println("Error: public key is not a regular file")
        return
    }
    sourcePrivKeyStat, err := os.Stat(privKeyPath)
    if err != nil {
        fmt.Println("Error: private key not found")
        return
    }
    if !sourcePrivKeyStat.Mode().IsRegular() {
        fmt.Println("Error: private key is not a regular file")
        return
    }
    sourcePubKey, err := os.Open(pubKeyPath)
    if err != nil {
        fmt.Println("Error: failed to open public key")
        return
    }
    defer sourcePubKey.Close()
    sourcePrivKey, err := os.Open(privKeyPath)
    if err != nil {
        fmt.Println("Error: failed to open private key")
        return
    }
    defer sourcePrivKey.Close()
    destPubKey, err := os.Create(s.pubKeyPath)
    if err != nil {
        fmt.Println("Error: failed to create public key")
        return
    }
    defer destPubKey.Close()
    destPrivKey, err := os.Create(s.privKeyPath)
    if err != nil {
        fmt.Println("Error: failed to create private key")
        return
    }
    defer destPrivKey.Close()
    _, err = io.Copy(destPubKey, sourcePubKey)
    if err != nil {
        fmt.Println("Error: failed to copy public key")
        return
    }
    _, err = io.Copy(destPrivKey, sourcePrivKey)
    if err != nil {
        fmt.Println("Error: failed to copy private key")
        return
    }
    fmt.Println("Keys set")
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
            return false

        } 
    } else if errors.Is(err, os.ErrNotExist) {
        return false

    }
    return false
}



