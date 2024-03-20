package configmodels

import "reflect"

func ValidateFields(cfg any) {
	t := reflect.TypeOf(cfg)
	v := reflect.ValueOf(cfg)

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		// Get the field
		fieldValue := v.Field(i).Interface()
		if isEmpty(fieldValue) {
			panic(t.Field(i).Name + " must not be empty")
		}
	}
}

// isEmpty checks if a value is considered empty
func isEmpty(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}
