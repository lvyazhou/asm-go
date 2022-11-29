# SWAG

访问地址
```
http://localhost:9000/swagger/index.html#/
```

安装包
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

引入pkg
```
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
import  docs "GinTest/swag/docs" // your docs

```

代码片段
```
func main()  {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld",Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

```

所在方法中增加swag注释
```
// GetMenu godoc
// @Summary      菜单查看
// @Description  菜单查看
// @Tags         菜单管理
// @Accept       json
// @Produce      json
// @Param        menuID   path      int  true  "菜单ID"
// @Success      200  {json} {api.Response}
// @Failure      400  {json} {api.Response}
// @Failure      404  {json} {api.Response}
// @Failure      500  {json} {api.Response}
// @Router       /menu/:menuID [get]
```

最后执行更新命令
```
使用swag init将会更新文档中的信息

PS: 
     如果 swag init 保存，提示 "cannot find type definition: XXXXX"，可使用以下命令
     
     //优先使用此命令
     swag init
  
     swag init --parseDependency --parseInternal
     
```