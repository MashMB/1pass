// Implementation of data type enumeration.
//
// @author TSS

package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	DataTypeEnum = newDataTypeRegistry()
)

type DataType struct {
	name string
}

type dataTypeRegistry struct {
	Address   *DataType
	Date      *DataType
	MonthYear *DataType
	values    []*DataType
}

func newDataType(name string) *DataType {
	return &DataType{name}
}

func newDataTypeRegistry() *dataTypeRegistry {
	address := newDataType("address")
	date := newDataType("date")
	monthYear := newDataType("monthYear")

	return &dataTypeRegistry{
		Address:   address,
		Date:      date,
		MonthYear: monthYear,
		values:    []*DataType{address, date, monthYear},
	}
}

func (dt *DataType) GetName() string {
	return dt.name
}

func (r *dataTypeRegistry) GetValues() []*DataType {
	return r.values
}

func (r *dataTypeRegistry) FromName(name string) (*DataType, error) {
	name = strings.ToLower(name)
	var dt *DataType

	for _, value := range r.values {
		if strings.ToLower(value.name) == name {
			dt = value
			break
		}
	}

	if dt == nil {
		return nil, ErrUnknownDataType
	}

	return dt, nil
}

func (r *dataTypeRegistry) parseAddress(jsonValue map[string]interface{}) string {
	var parsed string

	for key, value := range jsonValue {
		val := value.(string)
		val = strings.TrimSpace(val)
		val = strings.ReplaceAll(val, "\n", " ")
		formatted := fmt.Sprintf("\n\t\t%v: %v", key, val)
		parsed = parsed + formatted
	}

	return parsed
}

func (r *dataTypeRegistry) parseDate(value string) string {
	unix, _ := strconv.ParseInt(value, 10, 64)
	timestamp := time.Unix(unix, 0)

	return timestamp.Format("2006-01-02")
}

func (r *dataTypeRegistry) parseMonthYear(value string) string {
	if len(value) > 4 {
		year := value[:4]
		month := value[4:]
		value = fmt.Sprintf("%v-%v", year, month)
	}

	return value
}

func (r *dataTypeRegistry) ParseValue(dataType *DataType, value string, jsonValue map[string]interface{}) string {
	switch dataType {
	case r.Address:
		value = r.parseAddress(jsonValue)

	case r.Date:
		value = r.parseDate(value)

	case r.MonthYear:
		value = r.parseMonthYear(value)
	}

	return value
}
