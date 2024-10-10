package cmd

import (
	"fmt"

	"github.com/sedobrengocce/sesame/obj/config"
	"github.com/sedobrengocce/sesame/obj/store"
	"github.com/spf13/cobra"
)

var (
    name string
)

func init() {
  rootCmd.AddCommand(storeCmd)
  storeCmd.AddCommand(saveCmd)
  saveCmd.Flags().StringVarP(&host, "host", "H", "", "host to knock")
  saveCmd.Flags().StringVarP(&name, "name", "n", "", "name of the sequence")
  saveCmd.Flags().IntSliceVarP(&sequence, "ports", "p", []int{}, "sequence of ports to knock")
  saveCmd.Flags().IntVarP(&delay, "delay", "d", 10, "delay between knocks")
  saveCmd.Flags().BoolVarP(&udp, "udp", "u", false, "use UDP instead of TCP")
  saveCmd.MarkFlagRequired("ports")
  saveCmd.MarkFlagRequired("host")
  saveCmd.MarkFlagRequired("name")
}

var storeCmd = &cobra.Command{
    Use:   "store",
    Short: "Load, save and manage sequences",
    Long: `Load, save and manage sequences`,
}

var saveCmd = &cobra.Command{
    Use:   "save",
    Short: "Save a sequence",
    Long: `Save a sequence in the store if it doesn't exist`,
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.Get()
        if err != nil {
            fmt.Printf("Error getting the config: %v\n", err)
            panic(err)
        }
        s := store.NewStore(cfg.Paths.StorePath, cfg.Paths.PubKeyPath, cfg.Paths.PrivKeyPath)
        s.Save(name, host, sequence, delay, udp)
    },
}
