package mapper

// ErrUnmatchType error occured when source struct's field and target struct's field has unmatch data type
type ErrUnmatchType string

func (errUnmatchTy ErrUnmatchType) Error() string {
	return string(errUnmatchTy)
}

func errUnmatchType(fsrc, fdest string) ErrUnmatchType {
	return ErrUnmatchType(fsrc + " and " + fdest + " has unmatch type")
}

// ErrUnsupported error occured when trying to map non struct or pointer to struct
type ErrUnsupported string

func (errUnsupported ErrUnsupported) Error() string {
	return string(errUnsupported)
}

func errUnsupported(name string) ErrUnsupported {
	return ErrUnsupported(name + " is not supported, only struct or pointer to struct is supported")
}

// ErrNil error occured when trying to map nil struct
type ErrNil string

func (errNil ErrNil) Error() string {
	return string(errNil)
}

func errNil(name string) ErrNil {
	return ErrNil(name + " is nil")
}

// ErrDestNotPointer error occured when trying to map source to non pointer destination
type ErrDestNotPointer string

func (errDestNotPointer ErrDestNotPointer) Error() string {
	return string(errDestNotPointer)
}

func errDestNotPointer() ErrDestNotPointer {
	return ErrDestNotPointer("destination is not pointer")
}
