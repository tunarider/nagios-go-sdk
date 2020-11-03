package nagios

import "testing"

func TestInt(t *testing.T) {
    if StateOk.Int() != 0 {
        t.Error("Failed")
    }
}

func TestGreaterThan(t *testing.T) {
    if !StateUnknown.Gt(StateCritical) {
        t.Error("Failed")
    }
    if !StateCritical.Gt(StateWarning) {
        t.Error("Failed")
    }
    if !StateWarning.Gt(StateOk) {
        t.Error("Failed")
    }
}