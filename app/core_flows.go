package app

import (
	m "github.com/agustin-sarasua/rs-model"
	tx "github.com/agustin-sarasua/rs-transaction-api/app"
)

var FlowsMap = map[string]map[string]func(*m.Transaction) (string, error){
	"RENT_FLOW": tx.RentFlowMap}
