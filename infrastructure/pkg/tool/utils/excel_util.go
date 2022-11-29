// @Title  excel工具
// @Description: 使用excelize操作excel,官网文档 https://www.bookstack.cn/read/excelize-v2.0/0.md
// @Author: lvyazhou
// @Date: 2022/6/24 9:07

package utils_tool

import (
	constapicode2 "asm_platform/infrastructure/pkg/constants/api_code"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"net/url"
	"time"
)

var (
	// 默认'表头'样式
	defaultTBHeader = &excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Family: "宋体",
			Size:   12,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
	}

	// 默认表'内容'单元格样式
	defaultTBContent = &excelize.Style{
		Font: &excelize.Font{
			Family: "宋体",
			Size:   12,
			Color:  "#000000",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
	}

	// 默认行高度
	defaultHeight = 20.0
)

type SheetHeaderInfo struct {
	// 单元格宽度
	ColumnWidth float64

	// 横向占用单元格数量
	OccupyCellNum int

	// 表头名称
	HeaderName string

	// 表头对应数据变量名
	Field string
}

type SocExcelExport struct {
	file      *excelize.File
	sheetName string //可定义默认sheet名称
}

//
//  创建excel
//  @Description:
//  @params sheetNames sheet页名称
//  @return *SocExcelExport
//
func NewExcel(sheetNames []string) *SocExcelExport {
	return &SocExcelExport{file: createFile(sheetNames), sheetName: sheetNames[0]}
}

// ExportToPath 导出到本地
func (s *SocExcelExport) ExportToPath(name string, path string) (string, error) {
	if name == "" || len(name) == 0 {
		name = createFileName()
	} else {
		name = name + ".xlsx"
	}
	filePath := path + "/" + name
	err := s.file.SaveAs(filePath)
	return filePath, err
}

// Export 导出到浏览器
//
//  Export
//  @Description:
//  @receiver s
//  @params params 表头
//  @params data   数据
//  @params ctx    gin上下文环境
//
//func (s *SocExcelExport) Export(params []SheetHeaderInfo, data []map[string]interface{}, ctx *gin.Context) {
//	s.SetDataToSheet(params, nil)
//	s.ExportToWeb(createFileName(), s.file, ctx)
//}

//
//  导出
//  @Description:
//  @receiver s
//  @params fileName 文件名
//  @params file     文件
//  @params ctx      gin环境
//
func (s *SocExcelExport) ExportToWeb(fileName string, file *excelize.File, ctx *gin.Context) {
	buffer, _ := file.WriteToBuffer()
	//设置文件类型
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

//
//  写入表头
//  @Description:
//  @receiver s
//  @params startRow  表头起始行 0默认从1开始
//  @params sheetName sheet名
//  @params params    表头数据
//  @params style     样式,如果传空用默认
//
func (s *SocExcelExport) WriteHeader(startRow int32, sheetName string, params []SheetHeaderInfo, style *excelize.Style) {
	if style == nil {
		style = defaultTBHeader
	}
	topStyle, _ := s.file.NewStyle(style)
	if startRow == 0 {
		startRow = 1
	}
	column := 'A'

	// 首行写入
	var endIndex = 0
	for headerNum := 0; headerNum < len(params); headerNum++ {
		// 表头文字
		title := params[headerNum].HeaderName

		start := endIndex
		endIndex = endIndex + params[headerNum].OccupyCellNum
		for i := start; i < endIndex; i++ {
			line := fmt.Sprintf("%c%d", column, startRow)
			// 设置标题
			_ = s.file.SetCellValue(sheetName, line, title)
			// 列宽
			_ = s.file.SetColWidth(sheetName, fmt.Sprintf("%c", column), fmt.Sprintf("%c", column), params[headerNum].ColumnWidth)
			// 设置样式
			_ = s.file.SetCellStyle(sheetName, line, line, topStyle)
			column++
		}
		// 合并单元格
		if params[headerNum].OccupyCellNum != 1 {
			sc := fmt.Sprintf("%c%d", column-int32(params[headerNum].OccupyCellNum), startRow)
			ec := fmt.Sprintf("%c%d", column-1, startRow)
			s.file.MergeCell(sheetName, sc, ec)
		}
	}
}

//
//  WriteContent
//  @Description:
//  @receiver s
//  @params sheetName          sheet名称
//  @params contentStartIndex  内容开始行数
//  @params params             表头数据
//  @params obj                内容数据
//  @params style
//
func (s *SocExcelExport) WriteContent(sheetName string, contentStartIndex int16, headers []SheetHeaderInfo, obj interface{}, style *excelize.Style) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("obj convert to json fail, %v", err)
		return
	}
	data, convertErr := Json2SliceMap(string(jsonStr))
	if convertErr != constapicode2.Success {
		return
	}

	if style == nil {
		style = defaultTBContent
	}
	lineStyle, _ := s.file.NewStyle(style)
	//数据写入
	for i, val := range data {
		// 设置行高
		_ = s.file.SetRowHeight(sheetName, i+1, defaultHeight)
		// 逐列写入
		var column = 'A'

		// 首行写入
		var endIndex = 0
		for _, header := range headers {
			field := header.Field

			line := fmt.Sprintf("%c%v", column, contentStartIndex)
			value := val[field]

			start := endIndex
			endIndex = endIndex + header.OccupyCellNum

			for i := start; i < endIndex; i++ {
				// 设置值
				_ = s.file.SetCellValue(sheetName, line, value)
				// 设置样式
				_ = s.file.SetCellStyle(sheetName, line, line, lineStyle)
				column++
			}
			// 合并单元格
			if header.OccupyCellNum != 1 {
				sc := fmt.Sprintf("%c%d", column-int32(header.OccupyCellNum), i+2) // i+2 : i从0开始，而excel单元格从1开始
				ec := fmt.Sprintf("%c%d", column-1, i+2)
				s.file.MergeCell(sheetName, sc, ec)
			}
		}
		contentStartIndex++
	}
	// 设置行高 尾行
	_ = s.file.SetRowHeight(s.sheetName, len(data)+1, defaultHeight)
}

// 写入sheet内容
//func (s *SocExcelExport) SetDataToSheet(params []SheetHeaderInfo, style *excelize.Style) {
//	s.WriteHeader(params, style)
//	s.WriteContent(params, data, nil)
//}

//
//  创建excel
//  @Description:
//  @params sheetNames
//  @return *excelize.File
//
func createFile(sheetNames []string) *excelize.File {
	if len(sheetNames) == 0 {
		return nil
	}
	f := excelize.NewFile()
	for _, name := range sheetNames {
		f.NewSheet(name)
	}
	// 删除默认的sheet
	f.DeleteSheet("Sheet1")
	// 设置工作簿的默认工作表
	f.SetActiveSheet(0)
	return f
}

func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}
