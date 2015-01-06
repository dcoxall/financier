package financier

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBudgetMetadataAttributes(t *testing.T) {
	meta := BudgetMetadata{
		RelativeDataFolderName: "data2-69747A23",
		TED:           17199000000000,
		FormatVersion: "2",
	}
	assert.Equal(t, meta.RelativeDataFolderName, "data2-69747A23")
	assert.Equal(t, meta.TED, 17199000000000)
	assert.Equal(t, meta.FormatVersion, "2")
}

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
