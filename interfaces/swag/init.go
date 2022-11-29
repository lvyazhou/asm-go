package swag

import (
	"asm_platform/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine, conf *viper.Viper) {
	env := conf.GetString("server.env")
	host := conf.GetString("server.host")
	port := conf.GetString("server.port")

	if env == "dev" {
		docs.SwaggerInfo.Title = "Swagger asm_platform API"
		docs.SwaggerInfo.Description = "This is a sample server Petstore server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = host + ":" + port
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		docs.SwaggerInfo.BasePath = "/asm_platform"
		r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	}
}
