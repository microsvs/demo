package main

import (
	"github.com/graphql-go/graphql"
	"github.com/microsvs/base/pkg/errors"
)

type Error struct {
	Code    int    `json:"err_code"`
	Message string `json:"err_msg"`
}

func retErrors(p graphql.ResolveParams) (interface{}, error) {
	emap := errors.GetErrors()
	var errs = make([]*Error, 0, len(emap))
	for code, msg := range emap {
		errs = append(errs, &Error{
			Code:    int(code),
			Message: msg,
		})
	}
	return errs, nil
}
