package nagios

import (
	"fmt"
)

const (
	PerformanceFormatFull   string = "%s=%d;%d;%d;%d;%d"
	PerformanceFormatSimple string = "%s=%d;;;%d;%d"
)

type Performance struct {
	Label    string
	Value    int
	Warning  int
	Critical int
	Min      int
	Max      int
}

func (p Performance) String() string {
	if p.Warning == 0 || p.Critical == 0 {
		return fmt.Sprintf(PerformanceFormatSimple, p.Label, p.Value, p.Min, p.Max)
	} else {
		return fmt.Sprintf(PerformanceFormatFull, p.Label, p.Value, p.Warning, p.Critical, p.Min, p.Max)
	}
}
