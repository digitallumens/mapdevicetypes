# map-device-types

This package tracks known/supported Zigbee3 Devices and their capabilities for SiteWorx

`CommissionerZ` will use the json output for map creation, and `site-editor` and `siteworx-cli` will use the information for map uploading and remote commissioning information

**Usage:**

`GetCapabilities(name string)` to get all capabilities for a device type, based on its name
`HasCapability(name string)` checks a specific capability of a device type, based on its name
`CapabilityStr(cap uint32)` returns a string representation of device's capabilities
`GetKnownDeviceTypes()` returns all known device types

```
type DeviceType struct {
	Name       string `json:"name,omitempty"`
	ProdCode   uint32 `json:"prodcode,omitempty"`
	Capability uint32 `json:"capability,omitempty"`
}

type DeviceTypes struct {
	Version int
	Types   []DeviceType
}
```
