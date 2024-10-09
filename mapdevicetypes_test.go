package mapdevicetypes

// Test mapdevicetypes by running `go test` (or `go test -v`)

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCapability(t *testing.T) {
	b, err := HasCapability("SCN", capabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN should have Light")

	b, err = CapabilityIsTrue("SCN", capabilityLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Light should be true")

	b, err = HasCapability("TRH", capabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH should have Temp")
	b, err = HasCapability("TRH", capabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH should have Humidity")

	b, err = CapabilityIsTrue("TRH", capabilityTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH Temp should be true")
	b, err = CapabilityIsTrue("TRH", capabilityHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH Humidity should be true")

	// raw get value defaults to float64 because interface{}
	ctf64, err := GetCapabilityValue("WIO", capabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ctf64, 2.0, "WIO should have capabilityNumAnalogChannels = 2")

	ct, err := GetCapabilityIntValue("WIO", capabilityNumAnalogChannels)
	assert.Nil(t, err)
	assert.Equal(t, ct, 2, "WIO should have capabilityNumAnalogChannels = 2")
}
