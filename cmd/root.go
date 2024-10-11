/*Print the version number of Sesame
Copyright © 2024 Giuseppe Tufo <giuseppe.tufo@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.2.1"

var rootCmd = &cobra.Command{
	Use:   "sesame",
    Version: version,
	Short: "A simple port knocking tool",
	Long: `Sesame is a simple port knocking tool 
    that allows you to open ports on a remote server
    by sending a sequence of packets to a specific 
    set of ports. This tool allow you to save the
    sequences on a crypted database to improve the
    security of the sequences.`,
	Run: func(cmd *cobra.Command, args []string) {
        // TODO: Implement GUI
        fmt.Println("GUI not implemented yet")
        fmt.Println("Please run sesame -h for help")
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.SetVersionTemplate("Sesame version: {{.Version}}\n")
}

