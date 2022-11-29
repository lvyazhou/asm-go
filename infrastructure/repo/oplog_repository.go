package repo

import (
	oplog_entity "asm_platform/domain/entity/oplog"
	"asm_platform/domain/repository"
	"asm_platform/infrastructure/pkg/database/es"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
)

type OpLogRepo struct {
	client     *elastic.Client `json:"client"`
	index      string          `json:"index"`
	indexes    string          `json:"indexes"`
	mappingTpl string          `json:"mapping_tpl"`
}

func NewOpLogRepository(index string) *OpLogRepo {
	var mappingTpl = `{
    "mappings":{
      "_source": {
        "enabled": false
      },
      "properties": {
        "timestamp": {
          "type": "date",
          "format": "yyyy-MM-dd HH:mm:ss||epoch_millis"
        },
        "id": {
          "type": "keyword"
        },
        "user": {
          "properties": {
            "id": {
              "type": "long"
            },
            "name": {
              "type": "keyword"
            },
            "email": {
              "type": "keyword"
            }
          }
        },
        "request_uri": {
          "type": "keyword"
        },
        "client_ip": {
          "type": "ip"
        },
        "request_args": {
          "type": "text"
        },
        "response": {
          "properties": {
            "code": {
              "type": "keyword"
            },
            "message": {
              "type": "keyword"
            }
          }
        },
        "customer": {
          "properties": {
            "id": {
              "type": "long"
            },
            "name": {
              "type": "keyword"
            }
          }
        }
      }
    }`
	repo := &OpLogRepo{
		client:     es.Client,
		index:      index,
		mappingTpl: mappingTpl,
	}
	repo.init()
	return repo
}

// 初始化加载
func (repo *OpLogRepo) init() {
	ctx := context.Background()

	exists, err := repo.client.IndexExists(repo.index).Do(ctx)
	if err != nil {
		slog.Errorf("oplog es template init exist failed err is %v\n", err)
		return
	}

	if !exists {
		_, err = repo.client.CreateIndex(repo.index).Body(repo.mappingTpl).Do(ctx)
		if err != nil {
			slog.Errorf("oplog es template init failed err is %v\n", err)
			return
		}
	}
}

var _ repository.OpLogRepository = &OpLogRepo{}

// SaveOpLog
// @Description: 保存操作日志
// @params opLog 操作日志实体信息
// @return err   错误码
//
func (repo *OpLogRepo) SaveOpLog(opLog *oplog_entity.OpLog) error {
	_id := strconv.FormatInt(opLog.ID, 10)
	_, err := repo.client.Index().Index(repo.index).OpType("create").Id(_id).BodyJson(opLog).Refresh("true").Do(context.Background())

	//_, err := repo.client.Index().Index(repo.index).Type("_doc").BodyJson(opLog).Do(context.Background())
	if err != nil {
		slog.Errorf("[es save]:%s index save doc failed,err: %v", repo.index, err)
		return err
	}
	//slog.Debugf("[es save]:%s index save doc success,data: %v", repo.index, opLog)
	return nil
}

// BatchSaveOpLog
// @Description: 批量保存操作日志
// @params opLog 操作日志实体信息
// @return err   错误码
//
func (repo *OpLogRepo) BatchSaveOpLog(opLogMap map[string]string) error {
	bulkService := repo.client.Bulk().Index(repo.index).Refresh("true")
	// 添加多个文档请求
	for id, doc := range opLogMap {
		bulkService.Add(elastic.NewBulkCreateRequest().
			Index(repo.index).
			Id(id).
			Doc(doc))
	}
	res, err := bulkService.Do(context.Background())
	if err != nil {
		return err
	}
	if len(res.Failed()) > 0 {
		return errors.New(res.Failed()[0].Error.Reason)
	}
	return nil
}

// DeleteOpLogById
// @Description: 删除操作日志
// @params opLog 操作日志ID
// @return err   错误码
//
func (repo *OpLogRepo) DeleteOpLogById(id string) error {
	_, err := repo.client.Delete().Index(repo.index).Id(id).Refresh("true").Do(context.Background())
	return err
}

// DeleteOpLogByIds
// @Description: 批量删除操作日志
// @params ids 操作日志IDs
// @return err   错误码
//
func (repo *OpLogRepo) DeleteOpLogByIds(ids []string) (int, error) {
	bulkService := repo.client.Bulk().Index(repo.index).Refresh("true")
	for i := range ids {
		req := elastic.NewBulkDeleteRequest().Id(ids[i])
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

// DeleteOpLogByQuery
// @Description: 删除操作日志
// @params query 操作日志条件实体
// @return err   错误码
//
func (repo *OpLogRepo) DeleteOpLogByQuery(opQuery *oplog_entity.OpLogQuery) (int64, error) {
	if opQuery.IsEmpty() {
		return 0, errors.New("es query is null")
	}
	// 组装查询条件
	esQuery := elastic.NewBoolQuery()

	// ID 查询
	if opQuery.ID != 0 {
		esQuery.Filter(elastic.NewTermQuery("id", opQuery.ID))
	}

	// 执行删除语句
	rsp, err := repo.client.DeleteByQuery(repo.index).Query(esQuery).Refresh("true").Do(context.Background())
	if err != nil {
		return 0, err
	}
	return rsp.Deleted, nil
}

// FindOpLogList
// @Description: 查询分页操作日志
// @params param 查询条件
// @return total 总记录数
// @return date  数据
// @return err   错误码
//
func (repo *OpLogRepo) FindOpLogList(query *oplog_entity.OpLogQuery) ([]oplog_entity.OpLog, int64, error) {
	if query.Size == 0 {
		query.Size = 10
	}
	// es分页处理
	switch {
	case query.Page <= 1:
		query.Page = 0
	case query.Page > 1:
		query.Page = (query.Page - 1) * query.Size
	}

	// 组装查询条件
	boolQuery := elastic.NewBoolQuery()

	if len(query.RequestUri) != 0 {
		// 请求uri
		boolQuery.Must(elastic.NewMatchQuery("request_uri", query.RequestUri))
	}
	// 按照操作时间排序
	sortByTime := elastic.NewFieldSort("timestamp").Desc()

	// 获取索引名称 - 查询当月
	indexNames := repo.index

	// 查询结果
	data, esErr := repo.client.Search().
		TrackTotalHits(true). // 取消最大10000的限制,否则那怕2w条数据,带回的记录数量也是1w
		Index(indexNames).
		Query(boolQuery).
		From(query.Page).
		Size(query.Size).
		SortBy(sortByTime).
		Pretty(true).
		Do(context.Background())
	if esErr != nil {
		slog.Errorf("[soc op log query]:query es result failed,err: %v", esErr)
		return nil, 0, esErr
	}
	// 循环处理结果
	var result []oplog_entity.OpLog
	for _, item := range data.Each(reflect.TypeOf(oplog_entity.OpLog{})) {
		alarm := item.(oplog_entity.OpLog)
		result = append(result, alarm)
	}

	return result, data.Hits.TotalHits.Value, nil
}

//func FindOpLogList(param *apioplog.ListQueryParam, customId string) (total int64, result []oplog_entity.OpLog, err error) {
//	// 分页
//	modelUtil.EsPageHelper(&param.ApiCommon)
//	// 组装查询条件
//	boolQuery := elastic.NewBoolQuery()
//
//	if len(param.OpUser) != 0 {
//		// 操作用户
//		boolQuery.Must(elastic.NewMatchQuery("user.name", param.OpUser))
//	}
//	if len(param.OpStartTime) != 0 && len(param.OpEndTime) != 0 {
//		// 操作时间
//		s, e := utils.GetDateTime(param.OpStartTime, param.OpEndTime)
//		boolQuery.Must(elastic.NewRangeQuery("timestamp").Gte(s).Lte(e))
//	}
//	if len(param.RequestUri) != 0 {
//		// 操作用户
//		boolQuery.Must(elastic.NewMatchQuery("request_uri", param.RequestUri))
//	}
//	if len(param.ClientIp) != 0 {
//		// 客户端IP
//		boolQuery.Must(elastic.NewMatchQuery("client_ip", param.ClientIp))
//	}
//
//	// 获取索引名称
//	indexNames := getIndexNames(customId)
//
//	// 选定时间内，不存在对应索引
//	if len(indexNames) == 0 {
//		return 0, result, constapicode.Success
//	}
//
//	// 按照操作时间排序
//	sortByTime := elastic.NewFieldSort("timestamp").Desc()
//
//	// 查询结果
//	data, esErr := es.Client.Search().
//		TrackTotalHits(true). // 取消最大10000的限制,否则那怕2w条数据,带回的记录数量也是1w
//		Index(indexNames...).
//		Query(boolQuery).
//		From(param.Page).
//		Size(param.Size).
//		SortBy(sortByTime).
//		Pretty(true).
//		Do(context.Background())
//	if esErr != nil {
//		slog.Errorf("[soc op log query]:query es result failed,err: %v", esErr)
//		return 0, make([]OpLog, 0), constapicode.ESSearchError
//	}
//	// 循环处理结果
//	for _, item := range data.Each(reflect.TypeOf(OpLog{})) {
//		alarm := item.(OpLog)
//		result = append(result, alarm)
//	}
//
//	return data.Hits.TotalHits.Value, result, constapicode.Success
//}
