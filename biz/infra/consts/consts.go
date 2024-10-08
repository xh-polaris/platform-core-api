package consts

import (
	"errors"
)

var ErrNotAuthentication = errors.New("not authentication")
var ErrForbidden = errors.New("forbidden")
var PageSize int64 = 10
