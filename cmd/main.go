package main

import (
	_ "go-case/docs"
	"go-case/internal/api"
)

// @title In Memory Key Value API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email srknctnts@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	api.NewServer()
}