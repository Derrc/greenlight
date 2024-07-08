package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// customize JSON encoding for Runtime field by adding MarshalJSON()
// makes Runtime type satisfy Marshaler interface
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
