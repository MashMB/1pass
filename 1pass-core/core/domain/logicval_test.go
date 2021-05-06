// Unit tests for logic value enumeration.
//
// @author TSS

package domain

import (
	"testing"
)

func TestLVGetName(t *testing.T) {
	yes := LogicValEnum.Yes
	expected := "y"

	if yes.GetName() != expected {
		t.Errorf("GetName() = %v; expected = %v", yes.GetName(), expected)
	}
}

func TestLVGetValue(t *testing.T) {
	no := LogicValEnum.No
	expected := false

	if no.GetValue() != expected {
		t.Errorf("GetValue() = %v; expected = %v", no.GetValue(), expected)
	}
}

func TestLVGetValues(t *testing.T) {
	values := LogicValEnum.GetValues()
	expected := 2

	if len(values) != expected {
		t.Errorf("GetValues() = %v; expected = %v", len(values), expected)
	}
}

func TestLVFromName(t *testing.T) {
	expected := false
	result, err := LogicValEnum.FromName("n")

	if err != nil {
		t.Error("FromName() should pass because of valid name")
	}

	if result.GetValue() != expected {
		t.Errorf("GetValue() = %v; expected = %v", result.GetValue(), expected)
	}

	result, err = LogicValEnum.FromName("e")

	if err == nil {
		t.Error("FromName() should fail because of invalid name")
	}
}

func TestLVFromValue(t *testing.T) {
	expected := "y"
	result := LogicValEnum.FromValue(true)

	if result.GetName() != expected {
		t.Errorf("FromValue() = %v; expected = %v", result.GetName(), expected)
	}
}
