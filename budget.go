package financier

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
)

type Budget struct {
	Title    string
	Key      string
	Path     string
	Metadata BudgetMetadata
}

type BudgetList []Budget

var budgetDetails *regexp.Regexp = regexp.MustCompile(`\A(.*?)~(.{8}).ynab4\z`)

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
