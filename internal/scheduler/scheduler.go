package scheduler

import (
	"cronfacil/internal/models"
	"cronfacil/internal/repository"
	"fmt"
	"net/http"

	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

func StartAll() {
	Cron = cron.New()
	jobs, err := repository.ListJobs()

	if err != nil {
		panic(err)
	}

	for _, job := range jobs {
		if !job.Active {
			continue
		}
		registerJob(job)
	}

	Cron.Start()
	fmt.Println("🚀 Scheduler iniciado")
}

func Start(jobName string) (*models.Job, error) {
	job, err := repository.GetJobByName(jobName)

	if err != nil {
		return nil, err
	}

	if !job.Active {
		fmt.Printf("O job '%s' está inativo. Ative-o para executá-lo.\n", jobName)
		return nil, fmt.Errorf("job '%s' is inactive", jobName)
	}

	Cron = cron.New()

	registerJob(*job)

	Cron.Start()
	return job, nil
}

func registerJob(job models.Job) {
	_, err := Cron.AddFunc(job.Interval, func() {
		fmt.Printf("▶ Executando job: %s\n", job.Name)

		resp, err := http.Get(job.Endpoint)

		log := models.Log{
			JobName: job.Name,
		}

		if err != nil {
			log.Status = "ERROR"
			log.Response = err.Error()
			fmt.Println("❌ Erro:", err)
		} else {
			log.Status = "SUCCESS"
			log.Response = resp.Status
			fmt.Println("✅ Sucesso:", resp.Status)
		}

		repository.CreateLog(log)
	})

	if err != nil {
		fmt.Println("Erro ao registrar job:", err)
	}
}
