package financier

// Stores the contents of the Budget.ymeta files indicating further information
// regarding a particular Budget making it easier to correctly associate the
// data files.
type BudgetMetadata struct {
	// The relative directory name of the data folder containing the data files
	// for a particular Budget.
	RelativeDataFolderName string
	TED                    uint64
	FormatVersion          string
}
