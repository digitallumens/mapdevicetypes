package mapdevicetypes

// Test mapdevicetypes by running `go test` (or `go test -v`)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCapability(t *testing.T) {
	b, err := HasCapability("SCN Family", capabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family should have Light")

	b, err = CapabilityIsTrue("SCN Family", capabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family Light should be true")

	b, err = HasCapability("SWN-TRH", capabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH should have Temp")
	b, err = HasCapability("SWN-TRH", capabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH should have Humidity")

	b, err = CapabilityIsTrue("SWN-TRH", capabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH Temp should be true")
	b, err = CapabilityIsTrue("SWN-TRH", capabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SWN-TRH Humidity should be true")

	// raw get value defaults to float64 because interface{}
	ctf64, err := GetCapabilityValue("SWN-WIO", capabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ctf64, 2.0, "SWN-WIO should have capabilityNumAnalogChannels = 2")

	ct, err := GetCapabilityIntValue("SWN-WIO", capabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ct, 2, "SWN-WIO should have capabilityNumAnalogChannels = 2")
}
