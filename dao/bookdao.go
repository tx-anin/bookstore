package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

//GetBooks 获取数据库中所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		//将book添加奥books
		books = append(books, book)
	}
	return books, nil
}

//AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//写sql
	sqlStr := "insert into books (title,author,price,sales,stock,img_path) values (?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 根据图书的ID删除一本书
func DeleteBook(bookID string) error {
	//写sql
	sqlStr := "delete from books where id = ?"
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的ID查询一本书
func GetBookByID(bookID string) (*model.Book, error) {
	//写Sql
	sqlStr := "select * from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	//创建book
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

//UpdateBook 更加图书ID 更新图书信息
func UpdateBook(b *model.Book) error {
	//sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBooksByPrice 获取带分页和价格的数据
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	//将页面转换int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中的图书总记录数
	strSql := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(strSql, minPrice, maxPrice)
	row.Scan(&totalRecord)

	//设置每页只显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页面中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?, ?"
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}

//GetPageBooks 获取带分页的数据
func GetPageBooks(pageNo string) (*model.Page, error) {
	//将页面转换int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中的图书总记录数
	strSql := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(strSql)
	row.Scan(&totalRecord)

	//设置每页只显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页面中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?, ?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, nil
}
