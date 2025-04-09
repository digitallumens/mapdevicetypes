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
const UnspecifiedDeviceTypeName = "Unspecified Device"

type DeviceType struct {
	Name       string                 `json:"name"`
	ProdCode   string                 `json:"prod_code"`
	Attributes map[string]interface{} `json:"attributes"`
}

type DeviceTypes struct {
	Version int // must be > 0
	Types   []DeviceType
}

// Device attributes
const (
	AttributeSleepy            = "sleepy"
	AttributeLight             = "light"
	AttributeOccupancy         = "occupancy sensing"
	AttributeDaylight          = "daylight harvesting"
	AttributeTemperature       = "temperature sensing"
	AttributeHumidity          = "humidity sensing"
	AttributeDigitalIO         = "digital io"
	AttributePower             = "power metering"
	AttributeFlow              = "flow metering"
	AttributeLeak              = "leak sensing"
	AttributePressure          = "pressure sensing"
	AttributeVibration         = "vibration sensing"
	AttributeAssetTracking     = "asset tracking"
	AttributePushButton        = "push button"
	AttributeLoadControl       = "load control"
	AttributeNumButtons        = "num buttons"
	AttributeNumAnalogChannels = "num analog channels"
	AttributeNumADChannels     = "num adc channels"
)

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

func GetAllDeviceTypes() (dts []DeviceType, err error) {
	if !Init() {
		return nil, fmt.Errorf("Init failed")
	}
	return deviceTypes.Types, nil
}

func GetAttributes(name string) (interface{}, error) {
	if !Init() {
		return nil, fmt.Errorf("Init failed")
	}
	for _, dt := range deviceTypes.Types {
		if dt.Name == name {
			return dt.Attributes, nil
		}
	}
	return nil, fmt.Errorf("Device %q not found", name)
}

func GetAllKnownAttributes() (map[string]struct{}, error) {
	if !Init() {
		return nil, fmt.Errorf("Init failed")
	}
	attrsMap := map[string]struct{}{}
	for _, dt := range deviceTypes.Types {
		for attr, _ := range dt.Attributes {
			attrsMap[attr] = struct{}{}
		}
	}
	return attrsMap, nil
}

func HasAttribute(dt DeviceType, attribute string) bool {
	if !Init() {
		return false
	}
	_, exists := dt.Attributes[attribute]
	return exists
}

func AttributeIsTrue(dt DeviceType, attribute string) bool {
	if !Init() {
		return false
	}
	val, exists := dt.Attributes[attribute]
	if !exists {
		return false
	}
	b, ok := val.(bool)
	if !ok {
		return false // conversion to bool failed
	}
	return b
}
