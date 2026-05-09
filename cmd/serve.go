package cmd

import (
	"cronfacil/infra"
	"cronfacil/internal/scheduler"
)

var runCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia o scheduler para executar os jobs agendados",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command) {
		err := infra.Connect()

		if err != nil {
			panic(err)
		}

		scheduler.Start()
		select {}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}