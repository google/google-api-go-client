// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"fmt"
	"reflect"
	"strings"
)

// SchemaToMap returns a map containing an entry for each selected field in schema.
// A field is selected if:
//   * it has a non-empty value, or
//     * its field name is present in mustInclude, and
//     * it is not a nil pointer or nil interface.
// The map key for each selected field is set to the name from the field's json: struct tag.
// The map is suitable for rendering as a JSON payload in a Google API request.
func SchemaToMap(schema interface{}, mustInclude map[string]struct{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	s := reflect.ValueOf(schema)
	st := s.Type()

	for i := 0; i < s.NumField(); i++ {
		jsonTag := st.Field(i).Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		tag, err := parseJsonTag(jsonTag)
		if err != nil {
			return nil, err
		}

		v := s.Field(i)
		f := st.Field(i)
		if !includeField(v, f, mustInclude) {
			continue
		}

		// nil maps are treated as empty maps.
		if f.Type.Kind() == reflect.Map && v.IsNil() {
			m[tag.apiName] = map[string]string{}
			continue
		}

		// nil slices are treated as empty slices.
		if f.Type.Kind() == reflect.Slice && v.IsNil() {
			m[tag.apiName] = []bool{}
			continue
		}

		if tag.stringFormat {
			m[tag.apiName] = formatAsString(v, f.Type.Kind())
		} else {
			m[tag.apiName] = v.Interface()
		}
	}
	return m, nil
}

// formatAsString returns a string representation of v, dereferencing it first if possible.
func formatAsString(v reflect.Value, kind reflect.Kind) string {
	if kind == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return fmt.Sprintf("%v", v.Interface())
}

// jsonTag represents a restricted version of the struct tag format used by encoding/json.
// It is used to describe the JSON encoding of fields in a Schema struct.
type jsonTag struct {
	apiName      string
	stringFormat bool
}

// parseJsonTag parses a restricted version of the struct tag format used by encoding/json.
// The format of the tag must match that generated by the Schema.writeSchemaStruct method
// in the api generator.
func parseJsonTag(val string) (jsonTag, error) {
	var tag jsonTag
	elems := strings.Split(val, ",")
	if len(elems) < 2 || len(elems) > 3 || elems[0] == "" || elems[1] != "omitempty" {
		return tag, fmt.Errorf("malformed json tag: %s", val)
	}
	tag = jsonTag{
		apiName:      elems[0],
		stringFormat: len(elems) == 3 && elems[2] == "string",
	}
	return tag, nil
}

// Reports whether the struct field "f" with value "v" should be included in JSON output.
func includeField(v reflect.Value, f reflect.StructField, mustInclude map[string]struct{}) bool {
	// The regular JSON encoding of a nil pointer is "null", which means "delete this field".
	// Therefore, we could enable field deletion by honoring pointer fields' presence in the mustInclude set.
	// However, many fields are not pointers, so there would be no way to delete these fields.
	// Rather than partially supporting field deletion, we ignore mustInclude for nil pointer fields.
	// Deletion will be handled by a separate mechanism.
	if f.Type.Kind() == reflect.Ptr && v.IsNil() {
		return false
	}

	// The "any" type is represented as an interface{}.  If this interface
	// is nil, there is no reasonable representation to send.  We ignore
	// these fields, for the same reasons as given above for pointers.  We
	// also ignore "any" fields which have been initialized with a nil
	// pointer.
	if f.Type.Kind() == reflect.Interface {
		if v.IsNil() {
			return false
		}
		if elem := v.Elem(); elem.Kind() == reflect.Ptr && elem.IsNil() {
			return false
		}
	}

	_, ok := mustInclude[f.Name]
	return ok || !isEmptyValue(v)
}

// isEmptyValue reports whether v is the empty value for its type.  This
// implementation is based on that of the encoding/json package, but its
// correctness does not depend on it being identical. What's important is that
// this function return false in situations where v should not be sent as part
// of a PATCH operation.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
