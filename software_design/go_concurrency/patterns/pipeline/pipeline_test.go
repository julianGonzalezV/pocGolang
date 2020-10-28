package pipeline

import "testing"

/// Table o tests
func TestLaunchPipeline(t *testing.T) {
	//slice of slices: es unamatriz  de enteros
	tableTests := [][]int{
		{3, 14},
		{5, 55},
	}
	var res int
	for _, test := range tableTests {
		res = LaunchPipeline(test[0])
		t.Log("Iter => ", test)
		if res != test[1] {

			t.Logf("NOT EQUAL %d != %d\n", res, test[1])
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}
