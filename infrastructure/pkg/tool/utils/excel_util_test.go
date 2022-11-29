// @Title  Title
// @Description:
// @Author: lvyazhou
// @Date: 2022/6/24 15:56

package utils_tool

import (
	"fmt"
	"testing"
)

type Order struct {
	GoodsName string `json:"goods_name"`
	Price     string `json:"price"`
	Num       string `json:"num"`
}

func TestDemo(t *testing.T) {
	// sheet1:表头数据
	title := []SheetHeaderInfo{
		{
			ColumnWidth:   25,
			OccupyCellNum: 1,
			HeaderName:    "商品",
			Field:         "goods_name",
		},
		{
			ColumnWidth:   30,
			OccupyCellNum: 2,
			HeaderName:    "价格",
			Field:         "price",
		},
		{
			ColumnWidth:   40,
			OccupyCellNum: 1,
			HeaderName:    "库存",
			Field:         "num",
		},
	}
	// sheet1:内容数据
	content := []Order{
		{
			GoodsName: "AAA",
			Price:     "56.2",
			Num:       "786456456",
		},
		{
			GoodsName: "BBB",
			Price:     "78.6",
			Num:       "45643734521",
		},
		{
			GoodsName: "cccc",
			Price:     "992.99",
			Num:       "4531754322",
		},
	}

	excel := NewExcel([]string{"test1", "test2"})
	excel.WriteHeader(1, "test1", title, nil)
	excel.WriteContent("test1", 2, title, content, nil)
	path := "D:\\export\\soc-file"
	t.Log("输出到 -> " + path)
	_, err := excel.ExportToPath("测试", path)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}
