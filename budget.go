// Package financier provides a library that can parse YNAB data files allowing
// programmatic access to budget details and attributes.
package financier

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
)

// Budget stores the information for a particular budget setup within YNAB. This
// is effectively the root of everything you do within YNAB.
type Budget struct {
	// The budget name/title. This is the value displayed in YNAB.
	Title    string
	// The key is an additional attribute that is part of the budget directory
	// name.
	Key      string
	// The path is the fully qualified path to this particular budget
	Path     string
	// The budgets particular metadata including the relative name of the data
	// directory.
	Metadata BudgetMetadata
}

// A collection of one or more Budgets. This is to support installs of YNAB with
// multiple Budgets.
type BudgetList []Budget

var budgetDetails *regexp.Regexp = regexp.MustCompile(`\A(.*?)~(.{8}).ynab4\z`)

// Provide a BudgetList containing all the budgets within a particular
// directory. It is expected that the path argument is the root of the YNAB
// data directory.
func AvailableBudgets(path string) BudgetList {
	budgets := make(BudgetList, 0)
	paths, _ := filepath.Glob(filepath.Join(path, "*~*.ynab4"))
	for _, budgetPath := range paths {
		_, loc := filepath.Split(budgetPath)
		matches := budgetDetails.FindAllStringSubmatch(loc, -1)
		budgets = append(
			budgets,
			Budget{
				Title:    matches[0][1],
				Key:      matches[0][2],
				Path:     budgetPath,
				Metadata: readMetadata(filepath.Join(budgetPath, "Budget.ymeta")),
			},
		)
	}
	return budgets
}

func readMetadata(path string) BudgetMetadata {
	metadata := BudgetMetadata{}
	file, _ := os.Open(path)
	jsonDecoder := json.NewDecoder(file)
	jsonDecoder.Decode(&metadata)
	return metadata
}
