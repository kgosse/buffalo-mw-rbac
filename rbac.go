package rbac

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

var (
	// ErrUnauthorized is returned when the user is not allowed to perform a certain action
	ErrUnauthorized = errors.New("you are unauthorized to perform the requested action")
)

// RoleGetter must return the role of the user who made the request
type RoleGetter func(buffalo.Context) (string, error)

// New enables basic authentication
func New(e *casbin.Enforcer, r RoleGetter) buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			role, err := r(c)
			if err != nil {
				return errors.WithStack(err)
			}

			res, err := e.EnforceSafe(role, c.Request().URL.Path, c.Request().Method)
			if err != nil {
				return errors.WithStack(err)
			}
			if res {
				return next(c)
			}

			return c.Error(http.StatusUnauthorized, ErrUnauthorized)
		}
	}
}
