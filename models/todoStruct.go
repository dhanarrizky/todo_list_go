package models

import "time"

type TodoList struct {
	Id        uint
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Status    bool
}
