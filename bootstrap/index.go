package bootstrap

import (
	"sample-api/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(":8000")
}
