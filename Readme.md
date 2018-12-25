## Substitute Environment Variables (subenv)
This package was developed to improve the workflow when using the environmental variables.

### Use case
1) When you want to use default value for listening port but want the value to be overridable by ENV
2) Using localhost for database connection but overridable in prod by ENV
3) Any other value that u want to be controllable from outside the config files

### Available methods

Env(string key,default value)string - Get the string equivalent of the `key` or `value` will be returned if the `key` is null or non present

EnvI(string key,default value)int - Get the int equivalent of the `key` or `value` will be returned if the `key` is null or non present

EnvB(string key,default value)bool - Get the bool equivalent of the `key` or `value` will be returned if the `key` is null or non present

Override(key string, value interface{}) - Override any values that might be fetched. This is usually great for testing