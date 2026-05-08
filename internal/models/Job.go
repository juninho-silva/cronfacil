package models

import "time"

type Job struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Endpoint  string    `bson:"endpoint"`
	Method    string    `bson:"method"`
	Interval  string    `bson:"interval"`
	Active    bool      `bson:"active"`
	CreatedAt time.Time `bson:"created_at"`
}
