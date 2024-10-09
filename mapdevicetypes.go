// Package fixturecharacterizations defines an interface to access the
// contents of a fixtures.xml file

package mapdevicetypes

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
)

// EMBED THE mapdevicetypes.json FILE ••••••••••••••••••••••••••
//
//go:embed mapdevicetypes.json
var embeddedFile embed.FS

const DeviceTypesFile = "mapdevicetypes.json"

// Device capability attributes
const (
	Light     = 0x00000001
	Temp      = 0x00000002
	Humidity  = 0x00000004
	Pressure  = 0x00000008
	Power     = 0x00000010
	Keypad    = 0x00000020
	Flow      = 0x00000040
	Leak      = 0x00000080
	DigitalIO = 0x00000100
)

type DeviceType struct {
	Name       string `json:"name,omitempty"`
	ProdCode   uint32 `json:"prodcode,omitempty"`
	Capability uint32 `json:"capability,omitempty"`
}

type DeviceTypes struct {
	Version int
	Types   []DeviceType
}

func addCapabilityStr(caps, cap string) string {
	if caps == "" {
		caps = cap
	} else {
		caps = caps + ", " + cap
	}
	return caps
}

func CapabilityStr(cap uint32) string {
	capsStr := ""
	if cap&Light == Light {
		capsStr = addCapabilityStr(capsStr, "Light")
	}
	if cap&Temp == Temp {
		capsStr = addCapabilityStr(capsStr, "Temp")
	}
	if cap&Humidity == Humidity {
		capsStr = addCapabilityStr(capsStr, "Humidity")
	}
	if cap&Pressure == Pressure {
		capsStr = addCapabilityStr(capsStr, "Pressure")
	}
	if cap&Power == Power {
		capsStr = addCapabilityStr(capsStr, "Power")
	}
	if cap&Keypad == Keypad {
		capsStr = addCapabilityStr(capsStr, "Keypad")
	}
	if cap&Flow == Flow {
		capsStr = addCapabilityStr(capsStr, "Flow")
	}
	if cap&Leak == Leak {
		capsStr = addCapabilityStr(capsStr, "Leak")
	}
	if cap&DigitalIO == DigitalIO {
		capsStr = addCapabilityStr(capsStr, "Digital IO")
	}
	return capsStr
}

func GetCapabilities(name string) (uint32, error) {
	dts := GetKnownDeviceTypes()
	for _, dt := range dts.Types {
		if dt.Name == name {
			return dt.Capability, nil
		}
	}
	return 0, fmt.Errorf("Device %q not found", name)
}

func HasCapability(name string, capability uint32) (bool, error) {
	dts := GetKnownDeviceTypes()
	for _, dt := range dts.Types {
		if dt.Name == name {
			return (dt.Capability&capability != 0), nil
		}
	}
	return false, fmt.Errorf("Device %q not found", name)
}

func GetKnownDeviceTypes() DeviceTypes {
	var dt DeviceType
	var dts DeviceTypes

	dts.Version = 1

	dt = DeviceType{Name: "SCN", ProdCode: 0x00000000, Capability: Light}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "TRH", ProdCode: 0x00000000, Capability: (Temp | Humidity)}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "WIO", ProdCode: 0x00000000, Capability: DigitalIO}
	dts.Types = append(dts.Types, dt)

	dt = DeviceType{Name: "Generic Light", ProdCode: 0x00000000, Capability: Light}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Temp Sensor", ProdCode: 0x00000000, Capability: Temp}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Humidity Sensor", ProdCode: 0x00000000, Capability: Humidity}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Pressure Meter", ProdCode: 0x00000000, Capability: Pressure}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Power Meter", ProdCode: 0x00000000, Capability: Power}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Keypad", ProdCode: 0x00000000, Capability: Keypad}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Flow Meter", ProdCode: 0x00000000, Capability: Flow}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Leak Sensor", ProdCode: 0x00000000, Capability: Leak}
	dts.Types = append(dts.Types, dt)
	dt = DeviceType{Name: "Generic Leak Sensor", ProdCode: 0x00000000, Capability: Leak}
	dts.Types = append(dts.Types, dt)

	return dts
}

// Output a json file that can presumably be consumed by CommissionerX
func WriteKnownDeviceTypes() error {
	dts := GetKnownDeviceTypes()

	jsonData, err := json.MarshalIndent(dts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(DeviceTypesFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	fmt.Printf("Successfully wrote %q\n", DeviceTypesFile)
	return nil
}
