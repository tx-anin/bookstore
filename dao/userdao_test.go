package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	//fmt.Println("测试userdao中的函数")
	//t.Run("验证用户名或密码：", testLogin)
	//t.Run("验证用户名：", testRegist)
	//t.Run("保存用户：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@atguigu.com")
}

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的方法")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBook)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取图书", TestGetBookByIDBook)
	//t.Run("测试更新图书", TestUpdateBook)
	//t.Run("测试带分页的图书", TestGetPageBooks)
	//t.Run("测试带分页和价格的图书", TestGetPageBooksByPrice)
	//t.Run("测试Session", TestAddSession)
	t.Run("测试删除Session", TestDeleteSession)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//遍历得到的每一本书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	//调用
	err := AddBook(book)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func testDeleteBook(t *testing.T) {
	//调用删除图书的函数
	err := DeleteBook("47")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestGetBookByIDBook(t *testing.T) {
	//调用删除图书的函数
	book, err := GetBookByID("49")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("获取的图书信息是：", book)
}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      49,
		Title:   "三个女人和一百零五个男人的故事",
		Author:  "罗贯中",
		Price:   66.66,
		Sales:   500,
		Stock:   500,
		ImgPath: "/static/img/default.jpg",
	}

	UpdateBook(book)
}

//
func TestGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书：", page.Books)
	for _, v := range page.Books {
		fmt.Println("图书的信息是：", v)
	}
}

func TestGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("2", "10", "30")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书：", page.Books)
	for _, v := range page.Books {
		fmt.Println("图书的信息是：", v)
	}
}

func TestAddSession(t *testing.T) {
	cart := &model.Cart{}
	s := &model.Session{
		"123456789",
		"admin",
		1,
		cart,
		"",
		nil,
	}

	AddSession(s)
}

func TestDeleteSession(t *testing.T) {
	DeleteSession("123456789")
}

func TestGetSession(t *testing.T) {
	sess, _ := GetSession("b116ba9c-147a-4e18-6728-7385769f25f4")
	fmt.Println("Session的信息是：", sess)
}

func TestCart(t *testing.T) {
	fmt.Println("测试购物车的相关函数")
	t.Run("测试添加购物车", testAddCart)
}

func testAddCart(t *testing.T) {
	//设置第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}

	//购物项切片
	var cartItems []*model.CartItem
	//两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    26,
	}

	//将购物车插入数据库
	AddCart(cart)
}

func TestGetCartItemByBookID(t *testing.T) {
	cartItem, err := GetCartItemByBookIDAndCartID("1", "66668888")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cartItem)
}
func TestGetCartItemsByCartID(t *testing.T) {
	var cartItems []*model.CartItem
	cartItems, err := GetCartItemsByCartID("66668888")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range cartItems {
		fmt.Println(v)
	}
}

func TestGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(26)
	fmt.Println(cart)
}

func TestUpdateBookCount(t *testing.T) {

	book := &model.Book{
		ID: 2,
	}

	cartItem := &model.CartItem{
		Count:  100,
		CartID: "66668888",
	}

	cartItem.Book = book
	err := UpdateBookCount(cartItem)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestDeleteCartByCartID(t *testing.T) {
	err := DeleteCartByCartID("66668888")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestDeleteCartItemByID(t *testing.T) {

	err := DeleteCartItemByID("111")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	t.Run("测试添加订单和订单项", testAddOrder)
	t.Run("获取所有订单", testGetOrders)
	t.Run("获取所有订单项", testGetOrderItems)
	t.Run("获取用户的所有订单项", testGetOrdersByUser)
	t.Run("更新订单状态", testUpdateOrderState)
}

func testAddOrder(t *testing.T) {

	orderID := utils.CreateUUID()

	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      26,
	}

	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国",
		Author:  "罗三胖",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}

	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "西游",
		Author:  "吴二傻",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}

	//保存订单
	AddOrder(order)
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println("订单信息：", v)
	}
}

func testGetOrderItems(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("df90a33f-f3d5-4366-41aa-d8306243ae9c")
	for _, v := range orderItems {
		fmt.Println("订单项信息：", v)
	}
}

func testGetOrdersByUser(t *testing.T) {
	orders, _ := GetMyOrders(26)
	for _, v := range orders {
		fmt.Println("我的订单信息：", v)
	}
}

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("df90a33f-f3d5-4366-41aa-d8306243ae9c", 1)
}
