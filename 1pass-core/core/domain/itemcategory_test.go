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
		t.Errorf("[LOGIN] GetCode() = %v; expected = %v", category.GetCode(), expected)
	}

	category = ItemCategoryEnum.BankAccount
	expected = "101"

	if category.GetCode() != expected {
		t.Errorf("[BANK] GetCode() = %v; expected = %v", category.GetCode(), expected)
	}
}

func TestGetName(t *testing.T) {
	category := ItemCategoryEnum.Login
	expected := "Login"

	if category.GetName() != expected {
		t.Errorf("[LOGIN] GetName() = %v; expected = %v", category.GetName(), expected)
	}

	category = ItemCategoryEnum.Passport
	expected = "Passport"

	if category.GetName() != expected {
		t.Errorf("[PASSPORT] GetName() = %v; expected = %v", category.GetName(), expected)
	}
}

func TestGetValues(t *testing.T) {
	values := ItemCategoryEnum.GetValues()
	expected := 18

	if len(values) != expected {
		t.Errorf("len(GetValues()) = %d; expected = %d", len(values), expected)
	}
}

func TestFromCode(t *testing.T) {
	code := "001"
	expected := ItemCategoryEnum.Login
	category, err := ItemCategoryEnum.FromCode(code)

	if err != nil {
		t.Error("[LOGIN] FromCode() should not fail because of existing category for code")
	}

	if expected != category {
		t.Errorf("[LOGIN] FromCode() = %v; expected = %v", category, expected)
	}

	code = "099"
	expected = ItemCategoryEnum.Tombstone
	category, err = ItemCategoryEnum.FromCode(code)

	if err != nil {
		t.Error("[TOMBSTONE] FromCode() should not fail because of existing category for code")
	}

	if expected != category {
		t.Errorf("[TOMBSTONE] FromCode() = %v; expected = %v", category, expected)
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

	name = "Secure Note"
	expected = ItemCategoryEnum.SecureNote
	category, err = ItemCategoryEnum.FromName(name)

	if err != nil {
		t.Error("[NOTE] FromCode() should not fail because of existing category for name")
	}

	if expected != category {
		t.Errorf("[NOTE] FromCode() = %v; expected = %v", category, expected)
	}
}
