package cmd

import (
	"fmt"

	"github.com/sedobrengocce/sesame/obj/config"
	"github.com/sedobrengocce/sesame/obj/configDefaults"
	"github.com/sedobrengocce/sesame/obj/store"
	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(storeCmd)
    storeCmd.AddCommand(saveCmd)
    storeCmd.AddCommand(loadCmd)
    storeCmd.AddCommand(setKeyCmd)

    saveCmd.Flags().IntSliceVarP(&sequence, "ports", "p", []int{}, "sequence of ports to knock")
    saveCmd.Flags().IntVarP(&delay, "delay", "d", 10, "delay between knocks")
    saveCmd.Flags().BoolVarP(&udp, "udp", "u", false, "use UDP instead of TCP")
    saveCmd.MarkFlagRequired("ports")

    loadCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show the sequence")
}

var storeCmd = &cobra.Command{
    Use:   "store",
    Short: "Load, save and manage sequences",
    Long: `Load, save and manage sequences`,
}

var saveCmd = &cobra.Command{
    Use:   "save [name of sequence] [host]",
    Short: "Save a sequence",
    Long: `Save a sequence in the store if it doesn't exist`,
    Args: cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.Get(configdefaults.ConfigPath)
        if err != nil {
            fmt.Printf("Error getting the config: %v\n", err)
            panic(err)
        }
        if args[0] == "" {
            fmt.Println("Error: no sequence name provided")
            return
        }
        name := args[0] + ".yaml"
        if args[1] == "" {
            fmt.Println("Error: no host provided")
            return
        }
        host := args[1]
        s := store.NewStore(cfg.Paths.StorePath, cfg.Paths.PubKeyPath, cfg.Paths.PrivKeyPath)
        s.Save(name, host, sequence, delay, udp)
    },
}

var loadCmd = &cobra.Command{
    Use:   "load [name of sequence]",
    Short: "Load a sequence",
    Long: `Load a sequence in the store and knock it if it exists`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.Get(configdefaults.ConfigPath)
        if err != nil {
            fmt.Printf("Error getting the config: %v\n", err)
            panic(err)
        }
        if args[0] == "" {
            fmt.Println("Error: no sequence name provided")
            return
        }
        name := args[0] + ".yaml"
        s := store.NewStore(cfg.Paths.StorePath, cfg.Paths.PubKeyPath, cfg.Paths.PrivKeyPath)
        s.Load(name, verbose)
    },
}

var setKeyCmd = &cobra.Command{
    Use:   "setkey [path to public key] [path to private key]",
    Short: "Set the public and private keys",
    Long: `Set the public and private keys for the store`,
    Args: cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.Get(configdefaults.ConfigPath)
        if err != nil {
            fmt.Printf("Error getting the config: %v\n", err)
            panic(err)
        }
        if args[0] == "" {
            fmt.Println("Error: no public key path provided")
            return
        }
        if args[1] == "" {
            fmt.Println("Error: no private key path provided")
            return
        }
        pubKeyPath := args[0]
        privKeyPath := args[1]
        s := store.NewStore(cfg.Paths.StorePath, cfg.Paths.PubKeyPath, cfg.Paths.PrivKeyPath)
        s.SetKeys(pubKeyPath, privKeyPath)
    },
}

