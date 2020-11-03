package nagios

import (
	"fmt"
	"strings"
)

type Output struct {
	Message string
	State   State
}

func Ok(msg string) Output {
	return Output{Message: fmt.Sprintf("OK: %s", msg), State: StateOk}
}

func Warning(msg string) Output {
	return Output{Message: fmt.Sprintf("Warning: %s", msg), State: StateWarning}
}

func Critical(msg string) Output {
	return Output{Message: fmt.Sprintf("Critical: %s", msg), State: StateCritical}
}

func Unknown(msg string) Output {
	return Output{Message: fmt.Sprintf("Unknown: %s", msg), State: StateUnknown}
}

func MessageWithPerformance(msg string, performances []string) string {
	return fmt.Sprintf("%s | %s ", msg, strings.Join(performances, " "))
}
