package mapdevicetypes

// Test mapdevicetypes by running `go test` (or `go test -v`)

import (
	// "reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCapability(t *testing.T) {
	b, err := HasCapability("SCN", Light)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN should be Light")
	b, err = HasCapability("TRH", Temp)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH should have Temp")
	b, err = HasCapability("TRH", Humidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH should have Humidity")
	b, err = HasCapability("TRH", (Temp | Humidity))
	assert.Nil(t, err)
	assert.Equal(t, b, true, "TRH should have Temp and Humidity")
}

func TestGetCapabilities(t *testing.T) {
	cap, err := GetCapabilities("SCN")
	assert.Nil(t, err)
	assert.Equal(t, cap, Light, "SCN should be Light")
	cap, err = GetCapabilities("TRH")
	assert.Nil(t, err)
	assert.Equal(t, cap, true, "TRH should have Temp and Humidity")
}
