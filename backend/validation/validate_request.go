package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mchat.com/api/lib"
)

func ValidateReq(s interface{}, c *gin.Context) bool {
	if err := c.ShouldBind(s); err != nil {
		errs := err.(validator.ValidationErrors)
		lib.HttpResponse(400).Errors(FormatErrors(errs)).Send(c)
		return false
	}
	return true
}
