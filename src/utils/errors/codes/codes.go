package codes

import "math"

type Code uint64

const (
	NoCode = math.MaxUint64
)
const (
	CodeInvalidValue = (iota + Code(1))

	// database error
	CodeDBNotSupported
	CodeDBConnectionError

	// files
	CodeFileNotExist
	CodeCannotReadFile

	// json marshal/unmarshal
	CodeJSONUnmarshalError
	CodeJSONMarshalError
)
