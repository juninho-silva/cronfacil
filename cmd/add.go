package cmd

import (
	"cronfacil/internal/models"
	"cronfacil/internal/repository"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Job struct {
	Name     string
	Endpoint string
	Method   string
	Interval string
	Active   bool
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Cria um novo job interativamente",
	Run: func(cmd *cobra.Command, args []string) {

		job := Job{}

		// Nome
		promptName := promptui.Prompt{
			Label: "Nome do job",
		}
		job.Name, _ = promptName.Run()

		// Endpoint
		promptEndpoint := promptui.Prompt{
			Label: "Endpoint",
		}
		job.Endpoint, _ = promptEndpoint.Run()

		// Método HTTP
		methods := []string{"GET", "POST", "PUT", "DELETE"}
		promptMethod := promptui.Select{
			Label: "Método HTTP",
			Items: methods,
		}
		_, job.Method, _ = promptMethod.Run()

		// Intervalo
		promptInterval := promptui.Prompt{
			Label: "Intervalo (ex: 5m, 1h)",
		}
		job.Interval, _ = promptInterval.Run()

		// Ativo?
		activePrompt := promptui.Select{
			Label: "Ativar job?",
			Items: []string{"Sim", "Não"},
		}
		_, activeResult, _ := activePrompt.Run()
		job.Active = activeResult == "Sim"

		fmt.Println("\n📦 Job criado:")
		fmt.Printf("%+v\n", job)

		model := models.Job{
			Name:     job.Name,
			Endpoint: job.Endpoint,
			Method:   job.Method,
			Interval: job.Interval,
			Active:   job.Active,
		}

		repository.CreateJob(model)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
