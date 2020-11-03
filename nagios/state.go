package nagios

type State int

const (
    StateOk       State = 0
    StateWarning  State = 1
    StateCritical State = 2
    StateUnknown  State = 3
)

func (s State) Int() int {
    return int(s)
}

func (s State) Gt(c State) bool {
    return s.Int() > c.Int()
}
