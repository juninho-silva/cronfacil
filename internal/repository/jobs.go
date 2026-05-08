package repository

import (
	"context"
	"time"

	"cronfacil/internal/infra"
	"cronfacil/internal/models"
)

func CreateJob(job models.Job) error {
	collection := infra.DB.Collection("jobs")

	job.CreatedAt = time.Now()

	_, err := collection.InsertOne(context.Background(), job)
	return err
}

func ListJobs() ([]models.Job, error) {
	collection := infra.DB.Collection("jobs")

	cursor, err := collection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var jobs []models.Job
	if err := cursor.All(context.Background(), &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}
