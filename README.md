# map-device-types

This package tracks known/supported Zigbee3 Devices and their Attributes for SiteWorx.
Data comes from the `mapdevicetypes.json` file, which is consumed/served by the package.

`CommissionerZ` will use the json file for map creation, and `site-editor` and `siteworx-cli` will use the information for map uploading and remote commissioning information respectively.

**Structure:**

```
// Device attributes
const (
	attributeLight             = "light"
	attributeOccupancy         = "occupancy"
	attributeTemperature       = "temperature"
	attributeHumidity          = "humidity"
	attributePressure          = "pressure"
	attributePower             = "power"
	attributeKeypad            = "keypad"
	attributeFlow              = "flow"
	attributeLeak              = "leak"
	attributeDigitalIO         = "digital_io"
	attributeNumAnalogChannels = "num_analog_channels"
	attributeNumADCChannels    = "num_adc_channels"
)

type DeviceType struct {
	Name         string                 `json:"name"`
	ProdCode     string                 `json:"prod_code"`
	Attributes map[string]interface{}   `json:"attributes"`
}

type DeviceTypes struct {
	Version int
	Types   []DeviceType
}
```

**Usage:**

`Init` loads all DeviceTypes from mapdevicetypes.json if it's not been done yet

`GetDeviceType` returns a specific DeviceType by name

`GetAllDeviceTypes` returns all known DeviceTypes

`GetAttributes` returns the Attributes for a given DeviceType

`GetAllKnownAttributes` returns all known Attributes

`HasAttribute` returns true if a given DeviceType has an attribute (whether true or false)

`AttributeIsTrue` returns true if a given DeviceType's attribute is true
