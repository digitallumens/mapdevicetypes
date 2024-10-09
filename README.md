# map-device-types

This package tracks known/supported Zigbee3 Devices and their capabilities for SiteWorx.
Data comes from the `mapdevicetypes.json` file, which is consumed/served by the package.

`CommissionerZ` will use the json file for map creation options, and `site-editor` and `siteworx-cli` will use the information for map uploading and remote commissioning information respectively.

**Usage:**

`GetCapabilities(name string)` to get all capabilities for a device type, based on its name
`HasCapability(name string)` checks a specific capability of a device type, based on its name
`CapabilityStr(cap uint32)` returns a string representation of device's capabilities

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
