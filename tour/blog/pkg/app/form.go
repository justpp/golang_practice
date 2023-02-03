package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	key     string
	Message string
}

func (e *ValidError) Error() string {
	return e.Message
}

type ValidErrors []*ValidError

func (e ValidErrors) Errors() []string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return errs
}

func (e ValidErrors) Error() string {
	return strings.Join(e.Errors(), ",")
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for k, v := range errors.Translate(trans) {
			errs = append(errs, &ValidError{k, v})
		}
		return false, errs
	}
	return true, nil
}
