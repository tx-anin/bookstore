package model

//Page 结构
type Page struct {
	Books       []*Book //每页查询出来的图书切片
	PageNo      int64   //当前页
	PageSize    int64   //每页条数
	TotalPageNo int64   //总页数
	TotalRecord int64   //总记录数
	MinPrice    string
	MaxPrice    string
	IsLogin     bool   //是否登录
	Username    string //用户名
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNO 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetNextPageNO 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
