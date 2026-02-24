package blockscout

import (
	"fmt"
	"reflect"
	"strings"
)

// buildPath replaces {param} placeholders in a URL path template with actual values.
func buildPath(template string, pathParams map[string]string) string {
	result := template
	for key, value := range pathParams {
		result = strings.ReplaceAll(result, "{"+key+"}", value)
	}
	return result
}

// structToQueryParams converts a struct to map[string]string using the "query" tag.
// Fields with omitempty that are zero-valued or nil pointers are skipped.
func structToQueryParams(obj any) map[string]string {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	params := make(map[string]string)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		tag := field.Tag.Get("query")
		if tag == "" || tag == "-" {
			continue
		}

		name, opts := parseTag(tag)
		omitempty := false
		for _, opt := range opts {
			if opt == "omitempty" {
				omitempty = true
			}
		}

		if omitempty && isZeroValue(fieldVal) {
			continue
		}

		// Dereference pointer
		if fieldVal.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				continue
			}
			fieldVal = fieldVal.Elem()
		}

		params[name] = fmt.Sprintf("%v", fieldVal.Interface())
	}

	return params
}

func parseTag(tag string) (string, []string) {
	parts := strings.Split(tag, ",")
	if len(parts) == 1 {
		return parts[0], nil
	}
	return parts[0], parts[1:]
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	default:
		return false
	}
}
