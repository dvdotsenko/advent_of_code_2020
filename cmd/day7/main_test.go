package main

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	var n_should_be StackItem
	n := StackItem{}

	n.Append("asdf")
	n.Append("qwer")

	n_should_be = StackItem{"asdf", "qwer"}
	if !reflect.DeepEqual(n, n_should_be) {
		t.Errorf("StackItem.Append no worky: %v != %v", n, n_should_be)
	}

	n.Pop()
	n_should_be = StackItem{"asdf"}
	if !reflect.DeepEqual(n, n_should_be) {
		t.Errorf("StackItem.Append no worky: %v != %v", n, n_should_be)
	}

	n.Append("zxvc")
	n_should_be = StackItem{"asdf", "zxvc"}
	if !reflect.DeepEqual(n, n_should_be) {
		t.Errorf("StackItem.Append no worky: %v != %v", n, n_should_be)
	}
}
