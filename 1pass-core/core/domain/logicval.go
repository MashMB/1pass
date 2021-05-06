// Implementation of logic value enumeration.
//
// @author TSS

package domain

import (
	"strings"
)

var (
	LogicValEnum = newLogicValRegistry()
)

type LogicVal struct {
	name  string
	value bool
}

type logicValRegistry struct {
	No     *LogicVal
	Yes    *LogicVal
	values []*LogicVal
}

func newLogicVal(name string, value bool) *LogicVal {
	return &LogicVal{
		name:  name,
		value: value,
	}
}

func newLogicValRegistry() *logicValRegistry {
	no := newLogicVal("n", false)
	yes := newLogicVal("y", true)

	return &logicValRegistry{
		No:     no,
		Yes:    yes,
		values: []*LogicVal{no, yes},
	}
}

func (lv *LogicVal) GetName() string {
	return lv.name
}

func (lv *LogicVal) GetValue() bool {
	return lv.value
}

func (r *logicValRegistry) GetValues() []*LogicVal {
	return r.values
}

func (r *logicValRegistry) FromName(name string) (*LogicVal, error) {
	name = strings.ToLower(name)
	var lv *LogicVal

	for _, val := range r.values {
		if strings.ToLower(val.name) == name {
			lv = val
			break
		}
	}

	if lv == nil {
		return nil, ErrUnknownLogicVal
	}

	return lv, nil
}

func (r *logicValRegistry) FromValue(value bool) *LogicVal {
	var lv *LogicVal

	for _, val := range r.values {
		if value == val.value {
			lv = val
		}
	}

	return lv
}
