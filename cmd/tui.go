/*Print the version number of Sesame
Copyright © 2024 Giuseppe Tufo <giuseppe.tufo@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Long: `Start the TUI`,
	Run: func(cmd *cobra.Command, args []string) {
        // TODO: Implement GUI
        fmt.Println("TUI not implemented yet")
        fmt.Println("Please run sesame -h for help")
    },
}

func init() {
    rootCmd.AddCommand(tuiCmd)
}

