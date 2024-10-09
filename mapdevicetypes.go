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
	capabilityLight             = "light"
	capabilityOccupancy         = "occupancy"
	capabilityTemperature       = "temperature"
	capabilityHumidity          = "humidity"
	capabilityPressure          = "pressure"
	capabilityPower             = "power"
	capabilityKeypad            = "keypad"
	capabilityFlow              = "flow"
	capabilityLeak              = "leak"
	capabilityDigitalIO         = "digital_io"
	capabilityNumAnalogChannels = "num_analog_channels"
	capabilityNumADCChannels    = "num_adc_channels"
)

type DeviceType struct {
	Name         string                 `json:"name"`
	ProdCode     string                 `json:"prod_code"`
	Capabilities map[string]interface{} `json:"capabilities"`
}

type DeviceTypes struct {
	Version int
	Types   []DeviceType
}

func GetDeviceType(name string) (dt DeviceType, err error) {
	if !Init() {
		return dt, fmt.Errorf("Init failed")
	}
	for _, dt := range deviceTypes.Types {
		if dt.Name == name {
			return dt, nil
		}
	}
	return dt, fmt.Errorf("Device %q not found", name)
}

func GetCapabilities(name string) (interface{}, error) {
	if !Init() {
		return nil, fmt.Errorf("Init failed")
	}
	for _, dt := range deviceTypes.Types {
		if dt.Name == name {
			return dt.Capabilities, nil
		}
	}
	return nil, fmt.Errorf("Device %q not found", name)
}

func HasCapability(deviceName, capability string) (bool, error) {
	caps, err := GetCapabilities(deviceName)
	if err != nil {
		return false, err
	}
	capsMap, ok := caps.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("Capabilities for %q is not a map", deviceName)
	}

	_, exists := capsMap[capability]
	return exists, nil
}

func CapabilityIsTrue(deviceName, capability string) (bool, error) {
	caps, err := GetCapabilities(deviceName)
	if err != nil {
		return false, err
	}
	capsMap, ok := caps.(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("Capabilities for %q is not a map", deviceName)
	}
	ret := false
	if !Init() {
		return ret, fmt.Errorf("Init failed")
	}
	val, exists := capsMap[capability]
	if !exists {
		return ret, fmt.Errorf("Capability does not exist for %q", deviceName)
	}
	if ret, ok = val.(bool); !ok {
		return false, fmt.Errorf("Capability does not exist as bool for %q", deviceName)
	}
	return ret, nil
}

func GetCapabilityValue(deviceName, capability string) (interface{}, error) {
	caps, err := GetCapabilities(deviceName)
	if err != nil {
		return nil, err
	}
	capsMap, ok := caps.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Capabilities for %q is not a map", deviceName)
	}
	val, exists := capsMap[capability]
	if !exists {
		return nil, fmt.Errorf("Capability does not exist for %q", deviceName)
	}
	return val, nil
}

func GetCapabilityIntValue(deviceName, capability string) (int, error) {
	val, err := GetCapabilityValue(deviceName, capability)
	if err != nil {
		return -1, err
	}
	if floatVal, ok := val.(float64); ok {
		if float64(int(floatVal)) == floatVal {
			return int(floatVal), nil
		} else {
			return -1, fmt.Errorf("Capability %q for %q could not convert from float64 to int", capability, deviceName)
		}
	}
	return -1, fmt.Errorf("Capability %q for %q could not convert to float64", capability, deviceName)
}

func Init() bool {
	// Only if we haven't already done the global init
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
