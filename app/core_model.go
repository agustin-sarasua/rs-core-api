package app

import (
	"time"
)

type Task struct {
	ID            int64
	Flow          *Flow
	StepName      string
	Running       bool
	Retries       int
	TransactionID int64
	CreatedAt     time.Time
	LastRunDate   time.Time
}

type Flow struct {
	ID    int64
	Name  string
	Steps []*Step
}

type Step struct {
	Name      string
	URL       string
	TTL       int64
	Sleep     int64
	NextSteps []*NextStep `json:"next_steps"`
}

type NextStep struct {
	Output string
	Step   string
}
