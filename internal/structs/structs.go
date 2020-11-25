package structs

import "time"

//SampleTable for Cassandra
type SampleTable struct {
	ID   int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Description   string    `json:"description"`
	Status int       `json:"status"`
}
