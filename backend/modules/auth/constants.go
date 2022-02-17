package auth

import (
	"net/http"

	"mchat.com/api/lib"
)

const (
	EmailExistsErr = iota + 1
	EmailNotFoundErr
	UserNotFoundErr
	WrongPasswordErr
	TokenGenerateErr
	HashingPassErr
	ResetPasswordLinkExpErr
	SomethingWentWrongErr

	ResetPassEmailSentSucccess
	PassChangedSuccess
)

var HttpErrors = map[int]*lib.HttpResponseStruct{
	EmailExistsErr: lib.HttpResponse(http.StatusBadRequest).Errors(lib.H{
		"email": "Email already exists",
	}),
	UserNotFoundErr: lib.HttpResponse(http.StatusNotFound).Message("User not found"),
	EmailNotFoundErr: lib.HttpResponse(http.StatusNotFound).Errors(lib.H{
		"email": "Email doesn't exists",
	}),
	WrongPasswordErr: lib.HttpResponse(http.StatusBadRequest).Errors(lib.H{
		"password": "Invalid credentials",
	}),
	TokenGenerateErr:        lib.HttpResponse(http.StatusInternalServerError).Message("Failed to create user please contact customer care"),
	ResetPasswordLinkExpErr: lib.HttpResponse(http.StatusBadRequest).Message("Reset password link expired"),
	SomethingWentWrongErr:   lib.HttpResponse(http.StatusInternalServerError).Message("Something went wrong"),
}

var HttpSuccess = map[int]*lib.HttpResponseStruct{
	ResetPassEmailSentSucccess: lib.HttpResponse(http.StatusOK).Message("An email has been sent to your email address if your account exists"),
	PassChangedSuccess:         lib.HttpResponse(200).Message("Password changed successfully"),
}
