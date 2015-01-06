package financier

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

var exampleMetadata BudgetMetadata = BudgetMetadata{
	RelativeDataFolderName: "data2-69747A23",
	TED:           17199000000000,
	FormatVersion: "2",
}

func TestBudgetAttributes(t *testing.T) {
	budget := Budget{
		Title:    "Household Budget",
		Key:      "NCC1701D",
		Path:     "/home/foo/YNAB/Household Budget~NCC1701D.ynab4",
		Metadata: exampleMetadata,
	}
	assert.Equal(t, budget.Title, "Household Budget")
	assert.Equal(t, budget.Key, "NCC1701D")
	assert.Equal(t, budget.Path, "/home/foo/YNAB/Household Budget~NCC1701D.ynab4")
	assert.Equal(t, budget.Metadata, exampleMetadata)
}

func TestGetBudgets(t *testing.T) {
	root, _ := filepath.Abs("./fixtures")
	budgets := AvailableBudgets(root)
	expected := BudgetList{
		Budget{
			Title:    "Household Budget",
			Key:      "NCC1701D",
			Path:     filepath.Join(root, "Household Budget~NCC1701D.ynab4"),
			Metadata: exampleMetadata,
		},
		Budget{
			Title:    "Test",
			Key:      "E8570C74",
			Path:     filepath.Join(root, "Test~E8570C74.ynab4"),
			Metadata: exampleMetadata,
		},
	}
	assert.Equal(t, expected, budgets)
}
