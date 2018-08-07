// +build !cgo

package ledger_goclient

import (
	"fmt"
)

func ListDevices() {
}

func FindLedger() (*Ledger, error) {
	return nil, fmt.Errorf("cgo not available during build: no ledger support")
}
