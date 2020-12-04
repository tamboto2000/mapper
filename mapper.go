// Package mapper map struct to another struct with same or similiar fields.
package mapper

import (
	"reflect"
)

// Map map src to dest.
// src can be struct or pointer to struct, dest MUST be pointer to struct.
// src and dest fields must have same name and same type, but not necessarily have same fields count, examples:
//  type Struct1 struct {
// 	 Str string
// 	 Num int
// 	 Float float64
//  }
//  
//  type Struct2 struct {
// 	 Str string
// 	 Num int
//  }
//  
//  struct1 := Struct1{Str: "Hello world", Num: 1, Float: 2.5}
//  struct2 := new(Struct2)
//  err := mapper.Map(struct1, struct2)
//  // check error...
func Map(src, dest interface{}) error {
	// convert src and dest to reflect.Value and reflect.Type
	srcVal := reflect.ValueOf(src)
	srcType := srcVal.Type()
	destVal := reflect.ValueOf(dest)
	destType := destVal.Type()

	// check if src is pointer
	if srcType.Kind() == reflect.Ptr {
		// is src nil? If nil, return error
		if srcVal.IsNil() {
			return errNil(srcType.Name())
		}

		srcVal = srcVal.Elem()
		srcType = srcVal.Type()
		// is src is struct? If not, return error
		if srcType.Kind() != reflect.Struct {
			return errUnsupported(srcType.Name())
		}
	}

	// if src is not struct, return err
	if srcVal.Kind() != reflect.Struct {
		return errUnsupported(srcType.Name())
	}

	// check if dest is pointer
	if destType.Kind() != reflect.Ptr {
		// if dest not pointer, return err
		return errDestNotPointer()
	}

	// if dest is nil, return err
	if destVal.IsNil() {
		return errNil("destination")
	}

	// if dest is not pointer to struct, return err
	if destVal.Elem().Kind() != reflect.Struct {
		return errUnsupported(destType.Name())
	}

	destVal = destVal.Elem()
	destType = destVal.Type()

	// iterate dest fields
	for i := 0; i < destVal.NumField(); i++ {
		destF := destVal.Field(i)
		destFName := destType.Field(i).Name

		// find field in src that has the same name as dest from i index
		srcF := srcVal.FieldByName(destFName)
		// if not found, continue iterate
		if !srcF.IsValid() {
			continue
		}

		// check if destF has same type with srcF, if not, return error
		if destF.Type().String() != srcF.Type().String() {
			srcFType, _ := srcType.FieldByName(destFName)
			srcFName := srcFType.Name
			return errUnmatchType(srcType.Name()+"."+srcFName, destType.Name()+"."+destFName)
		}

		// set destF value
		destF.Set(srcF)
	}

	return nil
}
