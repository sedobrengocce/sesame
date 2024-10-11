package cmd

import (
	"fmt"

	"github.com/sedobrengocce/sesame/obj/knocker"
	"github.com/spf13/cobra"
)

var (
    sequence []int
    delay int = 10
    udp bool = false
    verbose bool = false
    fake bool = false
)

func init() {
  rootCmd.AddCommand(knockCmd)
  knockCmd.Flags().IntSliceVarP(&sequence, "ports", "p", []int{}, "sequence of ports to knock")
  knockCmd.Flags().IntVarP(&delay, "delay", "d", 10, "delay between knocks")
  knockCmd.Flags().BoolVarP(&udp, "udp", "u", false, "use UDP instead of TCP")
  knockCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
  knockCmd.Flags().BoolVarP(&fake, "fake", "f", false, "fake verbose output")
  knockCmd.MarkFlagRequired("ports")
}

var knockCmd = &cobra.Command{
    Use:   "knock [host]",
    Args: cobra.ExactArgs(1),
    Short: "knock a seqence to an host",
    Long: `knock a sequence to an host you 
    can specify the sequence and the host`,
    Run: func(cmd *cobra.Command, args []string) {
        if args[0] == "" {
            fmt.Println("Error: no host provided")
            return
        }
        host := args[0]
        knocker := knocker.NewKnocker(host)
        knocker.WithDelay(delay)
        knocker.WithUdp(udp)
        knocker.WithVerbose(verbose)
        knocker.WithFake(fake)
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

