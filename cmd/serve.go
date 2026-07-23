package cmd

import (
	"cronfacil/internal/infra"
	"cronfacil/internal/scheduler"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o scheduler para executar os jobs agendados",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := infra.Connect()

		if err != nil {
			panic(err)
		}

		scheduler.StartAll()
		select {}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
