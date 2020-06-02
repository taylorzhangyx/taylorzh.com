package t

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	res := Run(1, 2)
	if res != 3 {
		t.Error(fmt.Sprintln("Should be 3 but is", res))
	}
}
