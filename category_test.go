package financier

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	exampleJSON string = `{
		"subCategories": [
			{
				"masterCategoryId": "A7",
				"name": "Grocery",
				"isTombstone": true,
				"entityVersion": "A-1088",
				"note": "Often £100/week",
				"entityId": "786C36FB-7AAB-9D69-FCDB-DACCD553A312",
				"sortableIndex": 2146435071,
				"entityType": "category",
				"type": "OUTFLOW",
				"cachedBalance": 0
			}
		],
		"name": "Monthly Bills",
		"isTombstone": true,
		"entityVersion": "A-1100",
		"entityId": "A7",
		"sortableIndex": 2013265919,
		"entityType": "masterCategory",
		"expanded": true,
		"deleteable": true,
		"type": "OUTFLOW"
	}`
)

func TestCategoryJSONUnmarshalling(t *testing.T) {
	expected := Category{
		SubCategories: CategoryList{
			Category{
				MasterCategoryId: "A7",
				Name:             "Grocery",
				IsTombstone:      true,
				EntityVersion:    "A-1088",
				Note:             "Often £100/week",
				EntityId:         "786C36FB-7AAB-9D69-FCDB-DACCD553A312",
				SortableIndex:    2146435071,
				EntityType:       "category",
				Type:             "OUTFLOW",
				CachedBalance:    0,
			},
		},
		Name:          "Monthly Bills",
		IsTombstone:   true,
		EntityVersion: "A-1100",
		EntityId:      "A7",
		SortableIndex: 2013265919,
		EntityType:    "masterCategory",
		Expanded:      true,
		Deleteable:    true,
		Type:          "OUTFLOW",
	}
	actual := Category{}
	json.Unmarshal([]byte(exampleJSON), &actual)
	assert.Equal(t, actual, expected)
}
