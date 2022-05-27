package serialize

import (
	"reflect"
	"strings"
)

func numericalValue(value reflect.Value) (float64, bool) {
	switch value.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int()), true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.Uint()), true

	case reflect.Float32, reflect.Float64:
		return value.Float(), true

	default:
		return 0, false
	}
}

func lessValue(a, b reflect.Value) bool {
	aValue, aNumerical := numericalValue(a)
	bValue, bNumerical := numericalValue(b)

	if aNumerical && bNumerical {
		return aValue < bValue
	}

	if !aNumerical && !bNumerical {
		// In theory this should mean they are both strings. In reality
		// they could be any other type and the String() representation
		// will be something like "<bool>" if it is not a string. Since
		// distinct values of non-strings still return the same value
		// here that's what makes the ordering undefined.
		return strings.Compare(a.String(), b.String()) < 0
	}

	// Numerical values are always treated as less than other types
	// (including strings that might represent numbers themselves). The
	// inverse is also true.
	return aNumerical && !bNumerical
}
