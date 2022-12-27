package asset_handle

import (
	"asm_platform/application/app/asset_app"
	oplog_dto "asm_platform/application/dto"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/interfaces/handler"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AssetHandle struct {
	ah asset_app.AssetAppInterface
}

func NewAssetHandle() *AssetHandle {
	return &AssetHandle{ah: asset_app.NewAssetApp()}
}

func (ah *AssetHandle) SaveAsset(c *gin.Context) {
	ah.ah.SaveAsset()
	handler.ReturnFormat(c, constapicode.Success, nil)
	return
}

func (ah *AssetHandle) BatchSaveAsset(c *gin.Context) {
	ah.ah.BatchSaveAsset()
	handler.ReturnFormat(c, constapicode.Success, nil)
	return
}

func (ah *AssetHandle) GetAsset(c *gin.Context) {
	aId, a := c.Params.Get("id")
	if !a {
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	id, _ := strconv.ParseInt(aId, 10, 64)
	asset, err := ah.ah.GetAssetById(id)
	handler.ReturnFormat(c, err, asset)
	return
}

func (ah *AssetHandle) FindAssetList(c *gin.Context) {
	assets, err := ah.ah.FindAssetList()
	handler.ReturnFormat(c, err, assets)
	return
}

func (ah *AssetHandle) FindAssetListByPage(c *gin.Context) {
	var param = &oplog_dto.AssetQueryDto{}
	if err := c.ShouldBindJSON(param); err != nil {
		slog.Errorf("[asset][request][/asset/list/ [post]] FindAssetListByPage valid error %v.", err.Error())
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	assets, counts, err := ah.ah.FindAssetListByPage(param)
	handler.ReturnPageFormat(c, err, assets, counts)
	return
}

func (ah *AssetHandle) SaveAssetDomain(context *gin.Context) {

}
