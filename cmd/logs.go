package cmd

import (
	"cronfacil/internal/repository"
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

		logs, err := repository.GetLogsByJob(jobName)

		if err != nil {
			fmt.Println("Falha!")
			panic(err)
		}

		for _, log := range logs {
			fmt.Printf("* status: %s\n", log.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
