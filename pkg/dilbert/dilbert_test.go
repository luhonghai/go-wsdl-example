package dilbert

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTodayDilbert(t *testing.T) {
	soap := NewDilbertSoap("", false, nil)
	resp, err := soap.TodaysDilbert(nil)
	if err != nil {
		t.Error("Could not call soap method", err)
	} else {
		assert.Equal(t, resp.TodaysDilbertResult, "something?")
	}
}
