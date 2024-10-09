package cmd

import (
	"fmt"

	"github.com/sedobrengocce/sesame/obj/knocker"
	"github.com/spf13/cobra"
)

var (
    host string
    sequence []int
    delay int
    udp bool
)

func init() {
  rootCmd.AddCommand(knockCmd)
  knockCmd.Flags().StringVarP(&host, "host", "H", "", "host to knock")
  knockCmd.Flags().IntSliceVarP(&sequence, "ports", "p", []int{}, "sequence of ports to knock")
  knockCmd.Flags().IntVarP(&delay, "delay", "d", 10, "delay between knocks")
  knockCmd.Flags().BoolVarP(&udp, "udp", "u", false, "use UDP instead of TCP")
  knockCmd.MarkFlagRequired("ports")
  knockCmd.MarkFlagRequired("host")
}

var knockCmd = &cobra.Command{
    Use:   "knock",
    Short: "knock a seqence to an host",
    Long: `knock a sequence to an host you 
    can specify the sequence and the host`,
    Run: func(cmd *cobra.Command, args []string) {
        knocker := knocker.NewKnocker(host)
        knocker.WithDelay(delay)
        knocker.WithUdp(udp)
        err := knocker.WithSequence(sequence)
        if err != nil {
            fmt.Println(err)
            return
        }
        err = knocker.Knock()
        if err != nil {
            fmt.Println(err)
            return
        }
    },
}

