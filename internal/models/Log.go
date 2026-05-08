package models

import "time"

type Log struct {
	ID        string    `bson:"_id,omitempty"`
	JobName   string    `bson:"job_name"`
	Status    string    `bson:"status"`
	Response  string    `bson:"response"`
	CreatedAt time.Time `bson:"created_at"`
}
