package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	tx "github.com/agustin-sarasua/rs-transaction-api/app"
)

func nextFlowStep(t *Task, output string) *NextStep {
	log.Printf("Caluclating NextStep %v output %v", t.StepName, output)
	for _, s := range t.Flow.Steps {
		if s.Name == t.StepName {
			for _, ns := range s.NextSteps {
				if ns.Output == output {
					log.Printf("NextStep Found %v", ns.Step)
					return ns
				}
			}
		}
	}
	log.Printf("NextStep not found for output %v, TaskID: %v", output, t.ID)
	return nil
}

func loadFlows(files []string) []*Flow {
	var flows []*Flow

	for _, f := range files {
		var flow Flow
		raw, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &flow)
		flows = append(flows, &flow)
	}

	return flows
}

func RunTask(t *Task) (int64, error) {
	log.Printf("Running Task: %v, FlowName: %v", t.StepName, t.Flow.Name)
	m, _ := FlowsMap[t.Flow.Name]

	txv := tx.LoadTransaction(t.TransactionID)
	fn, _ := m[t.StepName]
	output, err := fn(txv)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	ns := nextFlowStep(t, output)
	nt := SaveNextTask(t, ns)
	return nt.ID, nil
}

func SaveNextTask(t *Task, ns *NextStep) *Task {
	var nt = Task{
		StepName:      ns.Step,
		CreatedAt:     time.Now(),
		Flow:          t.Flow,
		Retries:       3,
		Running:       false,
		TransactionID: t.TransactionID}
	log.Printf("Saving new Task %v", nt.StepName)
	return &nt
}
