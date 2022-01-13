package object_test

import (
	"testing"

	"github.com/flipez/rocket-lang/object"
)

func TestNullType(t *testing.T) {
	n := &object.Null{}

	if n.Type() != object.NULL_OBJ {
		t.Errorf("null.Type() returns wrong type")
	}
	if n.Inspect() != "null" {
		t.Errorf("null.Inspect() returns wrong type")
	}
}
func TestNullObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`[1][1].nope()`, "Failed to invoke method: nope"},
	}

	testInput(t, tests)
}