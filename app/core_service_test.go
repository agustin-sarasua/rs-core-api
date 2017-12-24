package app

import (
	"log"
	"testing"
	"time"
)

func TestLoadFLows(t *testing.T) {
	f := loadFlows([]string{"./../resources/rent-flow-test.json"})
	log.Printf("FlowName: %v", f[0].Name)
	if len(f) != 1 {
		t.Errorf("Error loading flows")
	}
}

func TestNextFlowStep(t *testing.T) {
	f := loadFlows([]string{"./../resources/rent-flow-test.json"})
	var nt = Task{
		StepName:      "START",
		CreatedAt:     time.Now(),
		Flow:          f[0],
		Retries:       3,
		Running:       false,
		TransactionID: 123}
	log.Printf("Flow %v\n", f[0].Name)
	ns := nextFlowStep(&nt, "OK")
	log.Printf("NextStep: %v", ns.Step)
	if ns.Step != "TENANT_DOCUMENTATION_UPLOAD" {
		t.Errorf("Error calculating next step")
	}
}
