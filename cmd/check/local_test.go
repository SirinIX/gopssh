package check

import (
	"testing"
)

func TestCheck(t *testing.T) {
	op := &option{
		configFile: "/Users/tokienohara/Documents/Project/gopssh/config/inventory.yaml",
	}
	if err := execute(op); err != nil {
		t.Error(err)
	}

}
