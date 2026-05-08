package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os jobs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listando jobs...")

		// Aqui viria do storage ou config
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
