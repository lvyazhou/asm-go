package vo

// MenuNode 树形菜单列表
type MenuNode struct {
	// 菜单ID
	ID int64 `json:"id,string"`

	// 菜单名称
	MenuName string `json:"menu_name"`

	// 菜单路径
	MenuPath string `json:"menu_path"`

	// 菜单ICON
	MenuIcon string `json:"menu_icon"`

	// 排序编号
	SortBy int8 `json:"sort_by"`

	// 父编号
	ParentId int64 `json:"parent_id,string"`

	// 是否导航栏(1否 2-是)
	NavigationBar int8 `json:"navigation_bar"`

	// 增加权限（0-无 1-有）
	AddPermission int8 `json:"add_permission"`

	// 编辑权限（0-无 1-有）
	EditPermission int8 `json:"edit_permission"`

	// 查看权限（0-无 1-有）
	DetailPermission int8 `json:"detail_permission"`

	// 删除权限（0-无 1-有）
	DeletePermission int8 `json:"delete_permission"`

	// 查询权限（0-无 1-有）
	SearchPermission int8 `json:"search_permission"`

	// 导入权限（0-无 1-有）
	ImportPermission int8 `json:"import_permission"`

	// 导出权限（0-无 1-有）
	ExportPermission int8 `json:"export_permission"`

	// 子节点
	Children []*MenuNode `json:"children"`
}
