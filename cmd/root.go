/*Print the version number of Sesame
Copyright © 2024 Giuseppe Tufo <giuseppe.tufo@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

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
        // TODO: Implement TUI
        fmt.Println("TUI not implemented yet")
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sesame.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
    rootCmd.SetVersionTemplate("Sesame version: {{.Version}}\n")
}

