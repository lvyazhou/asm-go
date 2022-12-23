package asset_router

import (
	"asm_platform/interfaces/handler/asset"
	"github.com/gin-gonic/gin"
)

func SetupAssetRouter(group *gin.RouterGroup) {
	assetGroup := group.Group("/asset")
	{
		// 实例化应用层接口
		assets := asset_handle.NewAssetHandle()

		// 保存资产
		assetGroup.POST("/save/", assets.SaveAsset)

		// 批量保存资产
		assetGroup.POST("/batch_save/", assets.BatchSaveAsset)

		// 查找资产
		assetGroup.GET("/get/:id", assets.GetAsset)

		// 查询资产
		assetGroup.POST("/find/", assets.FindAssetList)

		// 查询资产
		assetGroup.POST("/list/", assets.FindAssetListByPage)

	}
}
