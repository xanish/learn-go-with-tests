package _2_reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// The "reflect" package has a function ValueOf which returns us a Value of a given variable.
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		// We then make some very optimistic assumptions about the value passed in:
		// We check all values one by one. However, there may be no fields at all, which would cause a panic.
		field := val.Field(i)

		// We then call String(), which returns the underlying value as a string. However, this
		// would be wrong if the field was something other than a string.
		fn(field.String())
	}
}
