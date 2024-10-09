# map-device-types

This package tracks known/supported Zigbee3 Devices and their capabilities for SiteWorx.
Data comes from the `mapdevicetypes.json` file, which is consumed/served by the package.

`CommissionerZ` will use the json file for map creation, and `site-editor` and `siteworx-cli` will use the information for map uploading and remote commissioning information respectively.

**Structure:**

```
// Device capability attributes
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
```

**Usage:**

`GetCapabilities` returns all capabilities for a device in an `interface{}`

`HasCapability` returns whether a specific capability exists for a device as a `bool`

`CapabilityIsTrue` returns whether a specific capability is true for a device as a `bool`

`GetCapabilityValue` returns a specific capability in an `interface{}`

`GetCapabilityIntValue` returns a specific capability as an `int`
