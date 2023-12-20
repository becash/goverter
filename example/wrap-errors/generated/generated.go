// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	"fmt"
	wraperrors "github.com/jmattheis/goverter/example/wrap-errors"
	"strconv"
)

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source wraperrors.Input) (wraperrors.Output, error) {
	var exampleOutput wraperrors.Output
	xint, err := strconv.Atoi(source.PostalCode)
	if err != nil {
		return exampleOutput, fmt.Errorf("error setting field PostalCode: %w", err)
	}
	exampleOutput.PostalCode = xint
	return exampleOutput, nil
}