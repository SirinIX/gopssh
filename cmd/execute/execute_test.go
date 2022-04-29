package execute

import (
	"testing"
)

func TestExecute(t *testing.T) {
	op := &option{
		command: "ls -l",
		labels: "group=1",
	}
	if err := execute(op); err != nil {
		t.Error(err)
	}
}
