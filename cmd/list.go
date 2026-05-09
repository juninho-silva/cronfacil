package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cron *cron.Cron

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os jobs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listando jobs...")

		jobs, err := repository.GetAllJobs()

		if err != nil {
			fmt.Printf("Erro ao listar jobs: %v\n", err)
			return
		}

		for _, job := range jobs {
			fmt.Printf("ID: %d, Nome: %s, Cron: %s\n", job.ID, job.Name, job.Cron)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
