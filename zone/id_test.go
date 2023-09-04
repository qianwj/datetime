package zone

import (
	"testing"
	"time"
)

func TestDefaultId(t *testing.T) {
	defaultZone, _ := time.LoadLocation("")
	t.Logf("zone: %s", defaultZone.String())
	id := DefaultId()
	t.Logf("id: %+v", id)
	loc, err := time.LoadLocation(id.val)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("got default loc: %+v", loc)
	}
}
