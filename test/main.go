package main

import (
	"fmt"

	"github.com/digitallumens/mapdevicetypes"
)

func main() {
	err := mapdevicetypes.WriteKnownDeviceTypes()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
