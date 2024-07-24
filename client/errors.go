package client

import "errors"

var (
	ErrNotAProperty = errors.New("Given URL is not a Trustpilot property/domain")
)