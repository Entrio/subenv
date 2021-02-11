## Substitute Environment Variables (subenv)
This package was developed to simplify the usage of environmental variables, mainly when working with containers. I personally do not like using config files and prefer my apps not to be dependent on any external configurations to be able to function.


### Use case
Instead of hard-coding variables into your application or configuration files, simply use this library to read the required configuration from environment variable! If there is no variable present, a default value can be also set.
In some cases, you might wish to override a variable, this can be done by using the `Override(key, value)` function.


#### Common uses
1) Setting application listening port. 
2) Changing database host address and listening port.
3) Any other value that should have the ability to be changed.

### Installation
```sh
go get github.com/Entrio/subenv
```

### Example
```go
package main

import (
	"github.com/Entrio/subenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, world")
	})
	e.Logger.Fatal(e.Start(subenv.Env("LISTEN_ENDPOINT", ":1234")))
}

```

### Available functions

`Env(string key,default value) string` - Get the **string** equivalent of `key`. Default `value` will be returned if the `key` is null or not present.

`EnvI(string key,default value) int` - Get the **int** equivalent of the `key`. Default `value` will be returned if the `key` is null or non present

`EnvB(string key,default value) bool` - Get the **bool** equivalent of the `key`. Default `value` will be returned if the `key` is null or non present

`Override(key string, value interface{})` - Override any values that might be fetched. This is usually great for testing. This should be used right at the start of application launch.