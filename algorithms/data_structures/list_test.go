package data_structure

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	inputSlice, expected := []int{1, 2, 3}, []int{3, 2, 1}
	sliceR := Reverse(inputSlice)
	if !reflect.DeepEqual(sliceR, expected) {
		t.Errorf("Reverse doens't work, reverse of *%d must be %d but was %d  \n", inputSlice,
			expected, sliceR)
	}
}

func TestReverseAbstractInt(t *testing.T) {
	inputSlice, expected := []interface{}{1, 2, 3}, []interface{}{3, 2, 1}
	sliceR := ReverseAbstract(inputSlice)
	if !reflect.DeepEqual(sliceR, expected) {
		t.Errorf("Reverse doens't work, reverse of *%d must be %d but was %d  \n", inputSlice,
			expected, sliceR)
	}
}

func TestReverseAbstractStr(t *testing.T) {
	inputStrSlice, expectedStrSlice := []interface{}{"a", "b", "c"}, []interface{}{"c", "b", "a"}
	sliceRstrSlice := ReverseAbstract(inputStrSlice)
	if !reflect.DeepEqual(sliceRstrSlice, expectedStrSlice) {
		t.Errorf("Reverse doens't work, reverse of *%s must be %s but was %s  \n", inputStrSlice,
			expectedStrSlice, sliceRstrSlice)
	}
}

func TestReverseAbstractStrSwap(t *testing.T) {
	inputStrSlice, expectedStrSlice := []interface{}{"a", "b", "c"}, []interface{}{"c", "b", "a"}
	sliceRstrSlice := ReverseAbstractSwap(inputStrSlice)
	if !reflect.DeepEqual(sliceRstrSlice, expectedStrSlice) {
		t.Errorf("Reverse doens't work, reverse of *%s must be %s but was %s  \n", inputStrSlice,
			expectedStrSlice, sliceRstrSlice)
	}
}
