package middleware

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// An adaptation inspired by this stackoverflow thread: https://stackoverflow.com/questions/26744873/converting-map-to-struct
func setStructField(target interface{}, fieldName string, fieldValue any) error {
	// Capitalize first letter in the name to match the exported schema field
	fieldName = strings.ToUpper(fieldName[:1]) + fieldName[1:]

	// The struct
	structValue := reflect.ValueOf(target).Elem()

	// The struct's field
	structFieldValue := structValue.FieldByName(fieldName)
	structFieldType := structFieldValue.Type()

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no field: %s in obj", fieldName)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	// The reflect value we want to set
	newValue := reflect.ValueOf(fieldValue)

	//  ---- Main conversion logic starts here ----
	// Convert string (request param) -> target struct schema type
	switch structFieldType.Kind() {
	case reflect.String:
		// string -> string
		structFieldValue.Set(newValue)
	case reflect.Int:
		// string -> int
		intVal, err := strconv.Atoi(newValue.String())
		if err != nil {
			return fmt.Errorf("cannot convert %s to %s: %v", newValue.Type(), structFieldType, err)
		}
		structFieldValue.Set(reflect.ValueOf(intVal))
	case reflect.Slice:
		elemType := structFieldType.Elem()

		// string -> []
		switch elemType.Kind() {
		case reflect.Int:
			// 1,2,3 -> [1, 2, 3]
			stringVals := strings.Split(newValue.String(), ",")

			// string -> int[]
			var intVals []int
			for _, iv := range stringVals {
				v, err := strconv.Atoi(iv)
				if err != nil {
					return errors.New("error converting to int slice")
				}
				intVals = append(intVals, v)
			}
			structFieldValue.Set(reflect.ValueOf(intVals))
		default:
			return fmt.Errorf("unsupported struct field type: %s", structFieldType)
		}
	default:
		return fmt.Errorf("unsupported struct field type: %s", structFieldType)
	}

	return nil
}
