package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Sesame",
  Long:  `Print the version number of Sesame`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Sesame knocking tool v0.0")
  },
}
