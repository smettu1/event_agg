package parse

import (
	"github.com/events/models"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestParseData(t *testing.T) {
	t.Run("TestParseData", func(t *testing.T) {
		i := models.Input{}
		resp,err := ParseData("" )
		assert.Equal(t, resp, i)
		assert.Equal(t, err.Error(), "unexpected end of JSON input", "DevUser object not as expected")
	})
}