## go-rest-template
### Golang rest API template with gin using a mvc pattern derivation 

#### Stack:
* Gin
* Gorm
* Logrus  
* Godotenv

#### Obs:

* Easy to change log style (severity, style and format) for different environments, <b>check .env.example</b>;
* Dependency injection is being done by hand (at compile time), but in a very easy way to handle, <b>check /app/router/injection.go</b>
