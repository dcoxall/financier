package financier

// Account is a structure that stores data regarding any single YNAB account
// such as a credit card, bank account, etc.
type Account struct {
	// Indicates if the account counts towards the budget
	OnBudget bool
	// The display name of the account
	AccountName            string
	IsTombstone            bool
	EntityVersion          string
	LastEnteredCheckNumber int
	// The last date that the account was reconciled
	LastReconciledDate    Date
	EntityId              string
	LastReconciledBalance int64
	SortableIndex         int64
	EntityType            string
	AccountType           string
	Hidden                bool
}
