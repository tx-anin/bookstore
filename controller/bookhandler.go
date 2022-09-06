package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}

	fmt.Println(pageNo)
	//调用bookdao获取带分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)

	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Execute(w, page)
}

//GetPageBooksByPrice 获取带分页和价格的图书信息
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	//条件判断
	maxPrice := r.FormValue("max")
	//条件判断

	if pageNo == "" {
		pageNo = "1"
	}

	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao获取带分页的图书函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao获取带分页和价格的图书函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//fmt.Println(pageNo)
	////调用bookdao获取带分页和价格的图书函数
	//page, _ := dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)

	////调用bookdao
	//books, _ := dao.GetBooks()

	////获取Cookie
	//cookie, _ := r.Cookie("user")
	//if cookie != nil {
	//	//获取cookie的Value
	//	cookieValue := cookie.Value
	//	//根据cookieValue查询对应的session
	//	session, _ := dao.GetSession(cookieValue)
	//
	//	fmt.Println("获取的session:", session)
	//
	//	if session.UserID > 0 {
	//		//已经登录，设置page中的Username 和islogin
	//		page.IsLogin = true
	//		page.Username = session.UserName
	//	}
	//}

	//调用IsLogin函数判断是否已经登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录，设置page中的Username 和islogin
		page.IsLogin = true
		page.Username = session.UserName
	}

	//解析模板
	t := template.Must(template.ParseFiles("./views/index.html"))
	//执行
	fmt.Println(page)
	err := t.Execute(w, page)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//GetPageBooks 获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}

	fmt.Println(pageNo)
	//调用bookdao获取带分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)

	////调用bookdao
	//books, _ := dao.GetBooks()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	fmt.Println(page)
	err := t.Execute(w, page)
	if err != nil {
		fmt.Println(err.Error())
	}
}

////GetBooks 获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	//调用bookdao
//	books, _ := dao.GetBooks()
//	//解析模板
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	fmt.Println(books)
//	//执行
//	err := t.Execute(w, books)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

////AddBook 添加图书
//func AddBook(w http.ResponseWriter, r *http.Request) {
//	//获取图书信息
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price := r.PostFormValue("price")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//
//	//将价格销量库存进行转换
//	fPrice, _ := strconv.ParseFloat(price, 64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	//创建book
//	book := &model.Book{
//		Title:   title,
//		Author:  author,
//		Price:   fPrice,
//		Sales:   int(iSales),
//		Stock:   int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	//调用dao添加函数
//	dao.AddBook(book)
//	//调用Getbook 处理器函数再 查询一次数据库
//	GetBooks(w, r)
//}

//DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//获取要删除图书的ID
	bookID := r.FormValue("bookID")
	//调用dao删除
	dao.DeleteBook(bookID)
	//刷新，查询一次
	GetPageBooks(w, r)
}

//ToUpdateBookPage 去 更新 或 添加 图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取要更新图书的ID
	bookID := r.FormValue("bookId")
	//获取图书的函数
	book, _ := dao.GetBookByID(bookID)
	fmt.Println("获取的图书信息是：", book)

	//
	if book.ID > 0 {
		//更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}
}

//UpdateOrAddBook 更新 或 添加 图书的信息
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取图书信息
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	//将价格销量库存进行转换
	iBookID, _ := strconv.ParseInt(bookID, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	//创建book
	book := &model.Book{
		ID:      int(iBookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}

	if book.ID > 0 {
		//更新
		//调用dao更新函数
		dao.UpdateBook(book)

	} else {
		//添加
		//调用dao添加函数
		dao.AddBook(book)
	}
	//调用Getbook 处理器函数再 查询一次数据库
	GetPageBooks(w, r)
}
