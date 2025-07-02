package mapdevicetypes

// Test mapdevicetypes by running `go test` (or `go test -v`)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAttribute(t *testing.T) {
	dt, err := GetDeviceType("SCN Family")
	assert.Nil(t, err)
	b := HasAttribute(dt, AttributeLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family should have Light")
	assert.Equal(t, dt.NodeType, 0, "SCN Family should have NodeType 0")

	b = AttributeIsTrue(dt, AttributeLight)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "SCN Family Light should be true")

	dt, err = GetDeviceType("Wiser Temperature/Humidity Sensor (CCT593012)")
	assert.Nil(t, err)
	b = HasAttribute(dt, AttributeTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "Wiser Temperature/Humidity Sensor (CCT593012) should have Temp")
	b = HasAttribute(dt, AttributeHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "Wiser Temperature/Humidity Sensor (CCT593012) should have Humidity")
	b = AttributeIsTrue(dt, AttributeTemperature)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "Wiser Temperature/Humidity Sensor (CCT593012) AttributeTemperature should be true")
	b = AttributeIsTrue(dt, AttributeHumidity)
	assert.Nil(t, err)
	assert.Equal(t, b, true, "Wiser Temperature/Humidity Sensor (CCT593012) AttributeHumidity should be true")
	assert.Equal(t, dt.NodeType, 13, "Wiser Temperature/Humidity Sensor (CCT593012) should have NodeType 13")
}
