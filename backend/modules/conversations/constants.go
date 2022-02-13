package conversations

import "mchat.com/api/lib"

const (
	NotFoundErr = iota + 1
)

var HttpErrors = map[int]*lib.HttpResponseStruct{
	NotFoundErr: lib.HttpResponse(404).Message("Entity not found"),
}
