package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version Info:")
		fmt.Println("  Version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}