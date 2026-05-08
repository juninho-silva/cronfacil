package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cronfacil",
	Short: "Gerenciador de cron jobs simples",
	Long:  "Uma CLI para gerenciar e executar cron jobs manualmente",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
