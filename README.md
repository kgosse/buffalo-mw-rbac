# buffalo-mw-rbac

## Installation

```bash
$ go get -u github.com/herla97/buffalo-mw-rbac
```

## Usage

```go
// setup casbin auth rules.
authEnforcer, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
if err != nil {
  log.Fatal(err)
}

// Create role func.
roleFunc := func(c buffalo.Context) (string, error) {
  // implement your logic to get user's role
  role := "anonymous"
  return role, nil
}
app.Use(rbac.New(authEnforcer, roleFunc))
```

This is forked from: https://github.com/kgosse/buffalo-mw-rbac
