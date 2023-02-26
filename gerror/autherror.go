package gerror

import "github.com/pkg/errors"

var (
	ErrorAuthPassword = errors.New("Password error")
	ErrorAuthToken    = errors.New("Token error")
)
