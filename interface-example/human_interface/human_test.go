package human_interface

import (
	"testing"
)

func TestPrinter(t *testing.T) {
	extraValue := "printer"
	humanExpected := Human{
		Name: "Pao",
	}
	if humanExpected.NamePrinter() != humanExpected.Name+extraValue {
		t.Errorf("Name was incorrect, got: %s, want: %s.", humanExpected.NamePrinter(), humanExpected.Name)
	}
}
