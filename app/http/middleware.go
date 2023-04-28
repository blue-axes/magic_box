package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"magic_box/pkg/errors"
	"magic_box/server/entity"
	"reflect"
	"strings"
)

type (
	Binder struct {
	}
)

func (Binder) Bind(i interface{}, c echo.Context) error {
	b := echo.DefaultBinder{}
	err := b.Bind(i, c)
	if err != nil {
		return err
	}
	// validate
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	valid := validator.New()
	switch val.Kind() {
	case reflect.Struct:
		err = valid.Struct(val)
		switch verr := err.(type) {
		case validator.FieldError:
			return errors.New(entity.CodeInvalidArgs, fmt.Sprintf("field:%s is invalid. %s", verr.Field(), verr.Error()))
		case validator.ValidationErrors:
			var (
				fields = make([]string, 0)
				errMsg = make([]string, 0)
			)
			for _, v := range verr {
				fields = append(fields, v.Field())
				errMsg = append(errMsg, v.Error())
			}
			return errors.New(entity.CodeInvalidArgs, fmt.Sprintf("fields:%s is invalid. %s", strings.Join(fields, ","), strings.Join(errMsg, ",")))
		default:
			return err
		}

	default:
		return nil
	}
}

func Pre(next echo.HandlerFunc) echo.HandlerFunc {
	traceID := uuid.New().String()
	return func(c echo.Context) error {
		c.Set(entity.CtxTraceID, traceID)

		return next(c)
	}
}
