// Unit tests for item category enumeration.
//
// @author TSS

package domain

import (
	"testing"
)

func TestGetCode(t *testing.T) {
	category := ItemCategoryEnum.Login
	expected := "001"

	if category.GetCode() != expected {
		t.Errorf("GetCode() = %v; expected = %v", category.GetCode(), expected)
	}
}

func TestGetName(t *testing.T) {
	category := ItemCategoryEnum.Login
	expected := "Login"

	if category.GetName() != expected {
		t.Errorf("GetName() = %v; expected = %v", category.GetName(), expected)
	}
}

func TestGetValues(t *testing.T) {
	values := ItemCategoryEnum.GetValues()
	expected := 1

	if len(values) != expected {
		t.Errorf("len(GetValues()) = %d; expected = %d", len(values), expected)
	}
}

func TestFromCode(t *testing.T) {
	code := "001"
	expected := ItemCategoryEnum.Login
	category, err := ItemCategoryEnum.FromCode(code)

	if err != nil {
		t.Error("FromCode() should not fail because of existing category for code")
	}

	if expected != category {
		t.Errorf("FromCode() = %v; expected = %v", category, expected)
	}
}

func TestFromName(t *testing.T) {
	name := "Login"
	expected := ItemCategoryEnum.Login
	category, err := ItemCategoryEnum.FromName(name)

	if err != nil {
		t.Error("FromCode() should not fail because of existing category for name")
	}

	if expected != category {
		t.Errorf("FromCode() = %v; expected = %v", category, expected)
	}
}
