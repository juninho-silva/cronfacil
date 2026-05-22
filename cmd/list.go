package cmd

import (
	"cronfacil/internal/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os jobs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listando jobs...")

		jobs, err := repository.ListJobs()

		if err != nil {
			fmt.Println("Falha!")
			panic(err)
		}

		for _, job := range jobs {
			fmt.Printf("* name: %s\n createdAt: %s", job.Name, job.CreatedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
