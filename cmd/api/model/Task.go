package model

import "time"

type Task struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	Priority    uint32    `json:"priority,omitempty"`
	Done        bool      `json:"done,omitempty"`
	Assignee    string    `json:"assignee,omitempty"`
}

func (task Task) Validate() error {
	return nil
}
