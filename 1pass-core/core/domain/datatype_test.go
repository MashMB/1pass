// Unit tests for data type enumeration.
//
// @author TSS

package domain

import (
	"testing"
)

func TestDTGetName(t *testing.T) {
	expected := "date"
	dataType := DataTypeEnum.Date

	if expected != dataType.GetName() {
		t.Errorf("GetName() = %v; expected = %v", dataType.GetName(), expected)
	}
}

func TestDTGetValues(t *testing.T) {
	expected := 3
	values := DataTypeEnum.GetValues()

	if len(values) != expected {
		t.Errorf("GetValues() = %d; expected = %d", len(values), expected)
	}
}

func TestDTFromName(t *testing.T) {
	expected := DataTypeEnum.MonthYear
	dataType, err := DataTypeEnum.FromName("monthYear")

	if err != nil {
		t.Error("FromName() should pass because of valid name")
	}

	if dataType != expected {
		t.Errorf("FromName() = %v; expected = %v", dataType.GetName(), expected.GetName())
	}

	dataType, err = DataTypeEnum.FromName("")

	if err == nil {
		t.Error("FromName() should fail beause of empty name")
	}
}

func TestDTParseValue(t *testing.T) {
	expected := "2021-04-28"
	result := DataTypeEnum.ParseValue(DataTypeEnum.Date, "1619621164", nil)

	if result != expected {
		t.Errorf("ParseValue() = %v; expected = %v", result, expected)
	}

	expected = "2020-11"
	result = DataTypeEnum.ParseValue(DataTypeEnum.MonthYear, "202011", nil)

	if result != expected {
		t.Errorf("ParseValue() = %v; expected = %v", result, expected)
	}

	expected = "2021"
	result = DataTypeEnum.ParseValue(DataTypeEnum.MonthYear, "2021", nil)

	if result != expected {
		t.Errorf("ParseValue() = %v; expected = %v", result, expected)
	}

	address := make(map[string]interface{})
	address["country"] = "Poland"
	address["city"] = "Warsaw"
	expected = "\n\t\tcountry: Poland\n\t\tcity: Warsaw"
	result = DataTypeEnum.ParseValue(DataTypeEnum.Address, "", address)

	if result != expected {
		t.Errorf("ParseValue() = %v; expected = %v", result, expected)
	}
}
