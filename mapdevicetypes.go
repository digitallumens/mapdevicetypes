package mapdevicetypes

import (
	"embed"
	"encoding/json"
	"fmt"
)

// Global to hold json file contents after Init()
var deviceTypes DeviceTypes

// EMBED THE mapdevicetypes.json FILE ••••••••••••••••••••••••••
//
//go:embed mapdevicetypes.json
var embeddedFile embed.FS

const DeviceTypesFileName = "mapdevicetypes.json"

// Device capability attributes bitfield
const (
	Light uint32 = 1 << iota
	Temp
	Humidity
	Pressure
	Power
	Keypad
	Flow
	Leak
	DigitalIO
)

type DeviceType struct {
	Name       string `json:"name"`
	ProdCode   uint32 `json:"prod_code"`
	Capability uint32 `json:"capability"`
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
	if !Init() {
		return 0, fmt.Errorf("Init failed")
	}
	for _, dt := range deviceTypes.Types {
		if dt.Name == name {
			return dt.Capability, nil
		}
	}
	return 0, fmt.Errorf("Device %q not found", name)
}

func HasCapability(name string, capability uint32) (bool, error) {
	if !Init() {
		return false, fmt.Errorf("Init failed")
	}
	for _, dt := range deviceTypes.Types {
		if dt.Name == name {
			return (dt.Capability&capability != 0), nil
		}
	}
	return false, fmt.Errorf("Device %q not found", name)
}

//
// Programatically creates a slice with all known device types
//

func Init() bool {
	if deviceTypes.Version == 0 {
		var err error
		var bytes []byte
		bytes, err = embeddedFile.ReadFile(DeviceTypesFileName)
		if err != nil {
			fmt.Printf("%s\n", err)
			return false
		}
		err = json.Unmarshal(bytes, &deviceTypes)
		if err != nil {
			fmt.Printf("%s\n", err)
			return false
		}
	}
	return true
}
