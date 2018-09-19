## Substitute Environment Variables (subenv)
This package was developed to improve the workflow when using the environmental variables.

### Use case
1) When you want to use default value for listening port but want the value to be overwritable by ENV
2) using localhost for database connection but overwritable in prod by ENV

### Available methods

Env(string key,default value)string - Get the string equivelent of the `key` or `value` will be returned if the `key` is null or non present