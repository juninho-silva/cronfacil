package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [job]",
	Short: "Executa um job manualmente",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		fmt.Println("Executando job:", jobName)

		// Aqui você chamaria sua camada de serviço
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
