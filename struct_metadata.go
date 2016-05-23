package structmeta

import (
	//"fmt"
	"reflect"
	"strings"
)

type fieldMeta struct {
	Field string
	// This may need to be further evaluated if a pointer
	Value interface{}
	// First element of tag
	Key string
	// remainder tag elements
	Args []string
	// Field type
	Kind reflect.Kind
	// The actualy type (when  a pointer)
	Type reflect.Type
}

func (fm *fieldMeta) HasArg(arg string) (h bool) {
	for _, v := range fm.Args {
		if v == arg {
			h = true
			break
		}
	}
	return
}

//
// Ordered list of struct metadata
//
type StructMetadata []fieldMeta

func (sm StructMetadata) HasArg(arg string) StructMetadata {
	m := StructMetadata{}
	for _, v := range sm {
		if v.HasArg(arg) {
			m = append(m, v)
		}
	}
	return m
}

func (sm StructMetadata) NotHasArg(arg string) StructMetadata {
	m := StructMetadata{}
	for _, v := range sm {
		if !v.HasArg(arg) {
			m = append(m, v)
		}
	}
	return m
}

func (sm StructMetadata) FieldByName(name string) (fm *fieldMeta) {
	for _, v := range sm {
		if v.Field == name {
			fm = &v
			break
		}
	}
	return
}

func (sm StructMetadata) FieldByKey(key string) (fm *fieldMeta) {
	for _, v := range sm {
		if v.Key == key {
			fm = &v
			break
		}
	}
	return
}

func (sm StructMetadata) FieldNames() []string {
	names := make([]string, len(sm))

	for i, v := range sm {
		names[i] = v.Field
	}
	return names
}

// Return all tag keys
func (sm StructMetadata) Keys() []string {
	keys := make([]string, len(sm))
	for i, v := range sm {
		keys[i] = v.Key
	}
	return keys
}

func (sm StructMetadata) Values() []interface{} {
	values := make([]interface{}, len(sm))

	for i, v := range sm {
		values[i] = v.Value
	}
	return values

}

// includeTagless: include fields that do not have a tag
func ParseStructMetadata(f interface{}, tagname string, includeTagless bool) (sm StructMetadata) {
	sm = StructMetadata{}

	val := reflect.ValueOf(f).Elem()
	//name = val.Type().Name()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag.Get(tagname)

		fm := fieldMeta{
			Field: typeField.Name,
			Value: val.Field(i).Addr().Interface(),
			Kind:  val.Field(i).Kind(),
			Type:  typeField.Type,
		}

		if tag != "" {
			if parts := strings.Split(tag, ","); len(parts) > 0 {
				fm.Key = parts[0]
				fm.Args = parts[1:]

				if fm.Key == "" {
					fm.Key = fm.Field
				}
			}
		}

		if tag != "" || includeTagless {
			sm = append(sm, fm)
		}

	}
	return
}

func GetStructName(f interface{}) string {
	return reflect.ValueOf(f).Elem().Type().Name()
}
