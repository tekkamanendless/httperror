# `httperror`

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tekkamanendless/httperror?label=version&logo=version&sort=semver)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tekkamanendless/httperror)](https://pkg.go.dev/github.com/tekkamanendless/httperror)

This package contains an `error` value for each error-level HTTP status.

If you write code that returns the appropriate error, it's easy to handle those errors.

For example:

```
func YourFunction() error {
	// ...

	err := httperror.ErrorFromStatus(response.StatusCode)
	return err
}

// ...
err := YourFunction()
if err != nil {
	if errors.Is(err, httperror.ErrStatusNotFound) {
		// The item couldn't be found; do something.
	} else if errors.Is(err, httperror.ErrStatusUnauthorized) {
		// The user isn't signed in; do something else.
	} else {
		// Fail in a different way.
	}
}
```

