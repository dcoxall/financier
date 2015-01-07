package financier

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBudgetMetadataJSONUnarshal(t *testing.T) {
	jsonString := `{
		"relativeDataFolderName": "data2-69747A23",
		"TED": 17199000000000,
		"formatVersion": "2"
	}`
	expected := BudgetMetadata{
		RelativeDataFolderName: "data2-69747A23",
		TED:           17199000000000,
		FormatVersion: "2",
	}
	actual := BudgetMetadata{}
	json.Unmarshal([]byte(jsonString), &actual)
	assert.Equal(t, actual, expected)
}
