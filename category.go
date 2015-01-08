package financier

// A category is a budgeting group/category within YNAB. They are often
// organised in collections consisting of a master category (indicated by the
// entity type) with one or more sub categories.
type Category struct {
	// A collection of categories managed within this category
	SubCategories CategoryList
	// The display name of the category
	Name          string
	IsTombstone   bool
	EntityVersion string
	EntityId      string
	// An index which can be used to sort the categories in a standard way
	SortableIndex uint64
	// Can be one of "masterCategory" or "category"
	EntityType       string
	Expanded         bool
	Deleteable       bool
	Type             string
	MasterCategoryId string
	// Any notes left on the category
	Note          string
	CachedBalance int64
}

type CategoryList []Category
