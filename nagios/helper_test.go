package nagios

import "testing"

func TestResolveState(t *testing.T) {
	if ResolveState(StateOk, StateWarning) != StateWarning {
		t.Error("Failed")
	}
	if ResolveState(StateWarning, StateCritical) != StateCritical {
		t.Error("Failed")
	}
	if ResolveState(StateCritical, StateUnknown) != StateUnknown {
		t.Error("Failed")
	}
}
