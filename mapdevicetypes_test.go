package mapdevicetypes

// Test mapdevicetypes by running `go test` (or `go test -v`)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCapability(t *testing.T) {
	b, err := HasCapability("SCN Family", CapabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family should have Light")

	b, err = CapabilityIsTrue("SCN Family", CapabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family Light should be true")

	b, err = HasCapability("SWN-TRH", CapabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH should have Temp")
	b, err = HasCapability("SWN-TRH", CapabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH should have Humidity")

	b, err = CapabilityIsTrue("SWN-TRH", CapabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH Temp should be true")
	b, err = CapabilityIsTrue("SWN-TRH", CapabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH Humidity should be true")

	// raw get value defaults to float64 because interface{}
	ctf64, err := GetCapabilityValue("SWN-WIO", CapabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ctf64, 2.0, "SWN-WIO should have CapabilityNumAnalogChannels = 2")

	ct, err := GetCapabilityIntValue("SWN-WIO", CapabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ct, 2, "SWN-WIO should have CapabilityNumAnalogChannels = 2")

	// test Unspecified Device
	b, err = CapabilityIsTrue("Unspecified Device", CapabilityTemperature)
	assert.NotNil(t, err)
	assert.Equal(t, b, false, "Unspecified Device Temp Capability should be false")
}
