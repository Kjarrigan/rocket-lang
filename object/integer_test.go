package object_test

import (
	"github.com/flipez/rocket-lang/object"
	"testing"
)

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func TestIntegerObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`2.plz_s()`, "2"},
		{`10.plz_s(2)`, "1010"},
		{`10.type()`, "INTEGER"},
		{`2.nope()`, "Failed to invoke method: nope"},
		{`(2.wat().lines().size() == 2.methods().size() + 1).plz_s()`, "true"},
	}

	testInput(t, tests)
}

func TestIntegerHashKey(t *testing.T) {
	int1_1 := &object.Integer{Value: 1}
	int1_2 := &object.Integer{Value: 1}
	int2 := &object.Integer{Value: 2}

	if int1_1.HashKey() != int1_2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if int1_1.HashKey() == int2.HashKey() {
		t.Errorf("integer with different content have same hash keys")
	}
}

func TestIntegerInspect(t *testing.T) {
	int1 := &object.Integer{Value: 1}

	if int1.Inspect() != "1" {
		t.Errorf("integer inspect does not match value")
	}
}
