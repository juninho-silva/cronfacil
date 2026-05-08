package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [job]",
	Short: "Mostra logs de um job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		fmt.Println("Logs do job:", jobName)

		// Buscar logs
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
