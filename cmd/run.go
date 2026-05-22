package cmd

import (
	"fmt"
	"cronfacil/internal/models"
	"cronfacil/internal/repository"
	"net/http"

	"github.com/robfig/cron/v3"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [job]",
	Short: "Executa um job manualmente",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobName := args[0]
		scheduler.Start(jobName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}