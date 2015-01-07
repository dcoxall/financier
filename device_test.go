package financier

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDeviceJSONUnmarshal(t *testing.T) {
	jsonString := `{
		"knowledge": "A-1754,C-61,D-101",
		"deviceGUID": "F3435299-8707-71D4-2AB2-693010E45C61",
		"knowledgeInFullBudgetFile": "A-1754,C-61,D-101",
		"shortDeviceId": "A",
		"formatVersion": "1.2",
		"lastDataVersionFullyKnown": "4.2",
		"deviceType": "Desktop (AIR), OS:Mac OS 10.10.1",
		"highestDataVersionImported": null,
		"friendlyName": "My-Laptop",
		"YNABVersion": "Desktop version: YNAB 4 v4.3.655 (com.ynab.YNAB4.LiveCaptive), AIR Version: 4.0.0.1390",
		"hasFullKnowledge": true
	}`
	expected := Device{
		Knowledge:                  "A-1754,C-61,D-101",
		DeviceGUID:                 "F3435299-8707-71D4-2AB2-693010E45C61",
		KnowledgeInFullBudgetFile:  "A-1754,C-61,D-101",
		ShortDeviceID:              "A",
		FormatVersion:              "1.2",
		LastDataVersionFullyKnown:  "4.2",
		DeviceType:                 "Desktop (AIR), OS:Mac OS 10.10.1",
		HighestDataVersionImported: "",
		FriendlyName:               "My-Laptop",
		YNABVersion:                "Desktop version: YNAB 4 v4.3.655 (com.ynab.YNAB4.LiveCaptive), AIR Version: 4.0.0.1390",
		HasFullKnowledge:           true,
	}
	actual := Device{}
	json.Unmarshal([]byte(jsonString), &actual)
	assert.Equal(t, actual, expected)
}

func TestAvailableDevices(t *testing.T) {
	budgetPath, _ := filepath.Abs("./fixtures/Test~E8570C74.ynab4")
	devices := AvailableDevices(budgetPath)
	expected := DeviceList{
		Device{
			Knowledge:                  "A-1754,C-61,D-101",
			DeviceGUID:                 "F3435299-8707-71D4-2AB2-693010E45C61",
			KnowledgeInFullBudgetFile:  "A-1754,C-61,D-101",
			ShortDeviceID:              "A",
			FormatVersion:              "1.2",
			LastDataVersionFullyKnown:  "4.2",
			DeviceType:                 "Desktop (AIR), OS:Mac OS 10.10.1",
			HighestDataVersionImported: "",
			FriendlyName:               "My-Laptop",
			YNABVersion:                "Desktop version: YNAB 4 v4.3.655 (com.ynab.YNAB4.LiveCaptive), AIR Version: 4.0.0.1390",
			HasFullKnowledge:           true,
		},
		Device{
			YNABVersion:                "Android build 3.0.9",
			DeviceGUID:                 "3ED76EDE-C36D-3EAB-BF7F-4FC97AD1DEE5",
			DeviceType:                 "Android",
			FormatVersion:              "1.2",
			FriendlyName:               "GT-I8190N",
			HasFullKnowledge:           false,
			HighestDataVersionImported: "4.2",
			Knowledge:                  "A-1754,C-61,D-101",
			KnowledgeInFullBudgetFile:  "",
			LastDataVersionFullyKnown:  "4.2",
			ShortDeviceID:              "B",
		},
	}
	assert.Equal(t, expected, devices)
}
