package financier

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Represents a YNAB client device used to access, create or manage the budget
type Device struct {
	Knowledge string
	// A globally unique id for the device
	DeviceGUID                string
	KnowledgeInFullBudgetFile string
	// Matches the filename assigned by YNAB
	ShortDeviceID             string
	FormatVersion             string
	LastDataVersionFullyKnown string
	// Indicates the type of client (Android, Desktop, etc)
	DeviceType                 string
	HighestDataVersionImported string
	// The display name for the device
	FriendlyName string
	// Client version string
	YNABVersion      string
	HasFullKnowledge bool
}

// A collection of Devices
type DeviceList []Device

// When provided with a path to a budget file/directory it will locate all the
// registered devices known to the budget.
func AvailableDevices(budgetPath string) DeviceList {
	devicePaths, _ := filepath.Glob(
		filepath.Join(budgetPath, "devices", "*.ydevice"),
	)
	devices := make(DeviceList, 0)
	for _, path := range devicePaths {
		device := Device{}
		file, _ := os.Open(path)
		jsonDecoder := json.NewDecoder(file)
		jsonDecoder.Decode(&device)
		devices = append(devices, device)
	}
	return devices
}
