package repository

import (
	"context"
	"time"

	"cronfacil/internal/infra"
	"cronfacil/internal/models"
)

func CreateLog(log models.Log) error {
	collection := infra.DB.Collection("logs")

	log.CreatedAt = time.Now()

	_, err := collection.InsertOne(context.Background(), log)
	return err
}

func GetLogsByJob(jobName string) ([]models.Log, error) {
	collection := infra.DB.Collection("logs")

	filter := map[string]interface{}{
		"job_name": jobName,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var logs []models.Log
	if err := cursor.All(context.Background(), &logs); err != nil {
		return nil, err
	}

	return logs, nil
}
