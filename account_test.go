package financier

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountJSONUnmarshalling(t *testing.T) {
	exampleJSON := `{
		"onBudget": true,
		"accountName": "HSBC Current Account",
		"isTombstone": true,
		"entityVersion": "A-1035",
		"lastEnteredCheckNumber": -1,
		"lastReconciledDate": "2014-01-23",
		"entityId": "B76D4242-4860-7FAC-3630-DAB0057EF88F",
		"lastReconciledBalance": 0,
		"sortableIndex": 0,
		"entityType": "account",
		"accountType": "Checking",
		"hidden": true
	}`
	expected := Account{
		OnBudget:               true,
		AccountName:            "HSBC Current Account",
		IsTombstone:            true,
		EntityVersion:          "A-1035",
		LastEnteredCheckNumber: -1,
		LastReconciledDate: Date(
			time.Date(2014, time.January, 23, 0, 0, 0, 0, time.UTC),
		),
		EntityId:              "B76D4242-4860-7FAC-3630-DAB0057EF88F",
		LastReconciledBalance: 0,
		SortableIndex:         0,
		EntityType:            "account",
		AccountType:           "Checking",
		Hidden:                true,
	}
	actual := Account{}
	json.Unmarshal([]byte(exampleJSON), &actual)
	assert.Equal(t, actual, expected)
}
