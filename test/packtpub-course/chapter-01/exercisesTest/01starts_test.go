package exercisesTest

import (
	"pocGolang/packtpub-course/01-variables-and-operators/exercises"
	"testing"
)

func TestSum(t *testing.T) {
	total := exercises.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

}
