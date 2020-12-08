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
	srcVal, destVal, err := validate(src, dest)
	if err != nil {
		return err
	}

	return matchAndAssign(srcVal, destVal, FieldOption{})
}

// MapWithOpt is same as Map, but with option, see FieldOption for detailed information
func MapWithOpt(src, dest interface{}, opt FieldOption) error {
	srcVal, destVal, err := validate(src, dest)
	if err != nil {
		return err
	}

	return matchAndAssign(srcVal, destVal, opt)
}

//FieldOption determine how field being mapped
type FieldOption struct {
	// set to true if you don't want to override assigned dest's field
	SkipAssigned bool
	// set to true if you want to allow unmatch type. Fields with unmatch type will be skipped on mapping
	IsLoose bool
}

func validate(src, dest interface{}) (srcVal, destVal reflect.Value, err error) {
	// convert src and dest to reflect.Value and reflect.Type
	srcVal = reflect.ValueOf(src)
	srcType := srcVal.Type()
	destVal = reflect.ValueOf(dest)
	destType := destVal.Type()

	// check if src is pointer
	if srcType.Kind() == reflect.Ptr {
		// is src nil? If nil, return error
		if srcVal.IsNil() {
			return srcVal, destVal, errNil(srcType.Name())
		}

		srcVal = srcVal.Elem()
		srcType = srcVal.Type()
		// is src is struct? If not, return error
		if srcType.Kind() != reflect.Struct {
			return srcVal, destVal, errUnsupported(srcType.Name())
		}
	}

	// if src is not struct, return err
	if srcVal.Kind() != reflect.Struct {
		return srcVal, destVal, errUnsupported(srcType.Name())
	}

	// check if dest is pointer
	if destType.Kind() != reflect.Ptr {
		// if dest not pointer, return err
		return srcVal, destVal, errDestNotPointer()
	}

	// if dest is nil, return err
	if destVal.IsNil() {
		return srcVal, destVal, errNil("destination")
	}

	// if dest is not pointer to struct, return err
	if destVal.Elem().Kind() != reflect.Struct {
		return srcVal, destVal, errUnsupported(destType.Name())
	}

	destVal = destVal.Elem()

	return srcVal, destVal, nil
}

func matchAndAssign(srcVal, destVal reflect.Value, opt FieldOption) error {
	srcType := srcVal.Type()
	destType := destVal.Type()
	// iterate dest fields
	for i := 0; i < destVal.NumField(); i++ {
		destF := destVal.Field(i)
		destFName := destType.Field(i).Name

		// find field in src that has the same name as dest from i index
		srcF := srcVal.FieldByName(destFName)
		// if not found, continue iterate
		if srcF.IsZero() {
			continue
		}

		// if field is pointer and has nil value, skip
		if srcF.Kind() == reflect.Ptr && srcF.IsNil() {
			continue
		}

		// skip assigned dest's field if SkipAssigned enabled
		if opt.SkipAssigned {
			if !destF.IsZero() {
				continue
			}
		}

		// check if destF has same type with srcF
		if destF.Type().String() != srcF.Type().String() {
			// just continue if mapping is loose
			if opt.IsLoose {
				continue
			}

			// return err if mapping is not loose
			srcFType, _ := srcType.FieldByName(destFName)
			srcFName := srcFType.Name
			return errUnmatchType(srcType.String()+"."+srcFName, destType.String()+"."+destFName)
		}

		// set destF value
		destF.Set(srcF)
	}

	return nil
}
