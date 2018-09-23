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
	Short: "Print the version number of go-fixtures",
	Long:  `All software has versions. This is go-fixtures's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-fixtures v0.1.0 -- HEAD")
	},
}
