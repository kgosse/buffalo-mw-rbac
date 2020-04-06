package rbac_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/gobuffalo/buffalo"
	rbac "github.com/herla97/buffalo-mw-rbac"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func initApp(r string) *buffalo.App {
	h := func(c buffalo.Context) error {
		return c.Render(200, nil)
	}
	a := buffalo.New(buffalo.Options{})
	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcer("./rbac_model.conf", "./rbac_policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	roleFunc := func(c buffalo.Context) (string, error) {
		return r, nil
	}

	mwCashbin := rbac.New(authEnforcer, roleFunc)
	a.Use(mwCashbin)
	a.GET("/home", h)
	a.GET("/admin", h)
	return a
}

func TestMiddlewareAnonymous(t *testing.T) {
	r := require.New(t)
	w := willie.New(initApp("anonymous"))

	res := w.Request("/admin").Get()
	r.Equal(http.StatusUnauthorized, res.Code)

	res = w.Request("/home").Get()
	r.Equal(http.StatusOK, res.Code)
}

func TestMiddlewareAdmin(t *testing.T) {
	r := require.New(t)
	w := willie.New(initApp("admin"))

	res := w.Request("/home").Get()
	r.Equal(http.StatusOK, res.Code)

	res = w.Request("/admin").Get()
	r.Equal(http.StatusOK, res.Code)
}
