// +build cgo

package ledger_goclient

import (
	"errors"
	"fmt"

	"github.com/brejski/hid"
)

func ListDevices() {
	devices, err := hid.Devices()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	if len(devices) == 0 {
		fmt.Printf("No devices")
	}

	for _, d := range devices {
		fmt.Printf("============ %s\n", d.Path)
		fmt.Printf("Manufacturer  : %s\n", d.Manufacturer)
		fmt.Printf("Product       : %s\n", d.Product)
		fmt.Printf("ProductID     : %x\n", d.ProductID)
		fmt.Printf("VendorID      : %x\n", d.VendorID)
		fmt.Printf("VersionNumber : %x\n", d.VersionNumber)
		fmt.Printf("UsagePage     : %x\n", d.UsagePage)
		fmt.Printf("Usage         : %x\n", d.Usage)
		fmt.Printf("\n")
	}
}

func FindLedger() (*Ledger, error) {
	devices, err := hid.Devices()
	if err != nil {
		return nil, err
	}
	for _, d := range devices {
		if d.VendorID == VendorLedger && d.UsagePage == UsagePageLedger {
			device, err := d.Open()
			if err != nil {
				return nil, err
			}
			return NewLedger(device), nil
		}
	}
	return nil, errors.New("no ledger connected")
}
