package nagios

func ResolveState(states ...State) State {
	o := StateOk
	for _, s := range states {
		if s.Gt(o) {
			o = s
		}
	}
	return o
}
