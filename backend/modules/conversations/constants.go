package conversations

import "mchat.com/api/lib"

const (
	NotFoundErr = iota + 1
	UserNotFoundErr
	AlreadyExistsErr
)

var HttpErrors = map[int]*lib.HttpResponseStruct{
	NotFoundErr:      lib.HttpResponse(404).Message("Entity not found"),
	UserNotFoundErr:  lib.HttpResponse(404).Message("User not found"),
	AlreadyExistsErr: lib.HttpResponse(400).Message("Conversation already exists"),
}
