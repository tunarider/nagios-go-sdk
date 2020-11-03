package nagios

import "testing"

func TestFull(t *testing.T) {
	p := Performance{
		Label: "test",
		Value: 1,
		Warning: 2,
		Critical: 3,
		Min: 0,
		Max: 3,
	}
	if p.String() != "test=1;2;3;0;3" {
		t.Error("Failed")
	}
}