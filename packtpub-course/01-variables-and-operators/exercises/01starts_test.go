package exercises

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

}
