# buffalo-mw-rbac

## Installation

```bash
$ go get -u github.com/kgosse/buffalo-mw-rbac
```

## Usage

```go
// setup casbin auth rules
authEnforcer, err := casbin.NewEnforcerSafe("rbac_model.conf", "rbac_policy.csv")
if err != nil {
  log.Fatal(err)
}
roleFunc := func(c buffalo.Context) (string, error) {
  // implement your logic to get user's role
  role := "anonymous"
  return role, nil
}
app.Use(rbac.New(authEnforcer, roleFunc))
```

If you want a real example, check this project:
https://github.com/kgosse/shop-back
