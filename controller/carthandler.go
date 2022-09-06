package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	json2 "encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录

		//获取要添加的图书的ID
		bookID := r.FormValue("bookId")
		fmt.Println("id是:", bookID)
		//根据图书的ID获取图书的信息
		book, _ := dao.GetBookByID(bookID)
		fmt.Println("book是:", book)

		//获取用户ID
		userID := session.UserID
		//判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)
		fmt.Println("购物车信息：", cart)
		if cart != nil {
			//当前用户已经有购物车，判断购物车中是否有当前这本图书
			carItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if carItem != nil {
				//购物车的购物项中已经有该图书，只需要将购物项中的数量+1
				//1. 获取购物车切片中所有的购物项
				cts := cart.CartItems
				//2. 遍历
				for _, v := range cts {
					fmt.Println("当前购物项中的book:", v.Book)
					fmt.Println("333 carItem.book", carItem.Book)
					//3. 找到当前购物项
					if v.Book.ID == carItem.Book.ID {
						//将当前购物项中图书的数量+1
						v.Count = v.Count + 1
						//更新到数据库中
						dao.UpdateBookCount(v)
					}
				}

			} else {
				fmt.Println("当前购物车中还没有改图书对应的购物项")
				//购物车的购物项中还没有该图书
				//需要创建购物项，并添加到数据库中
				fmt.Println("当前要添加的图书信息：", book)
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到当前Cart切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的购物项保存到数据库
				dao.AddCartItem(cartItem)
			}
			//更新购物车的总数量和总金额
			dao.UpdateCart(cart)

		} else {
			//当前用户没有购物车，需要创建购物车，购物项
			//1. 创建购物车
			//生成购物车的ID
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			//2. 创建购物车中的购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			//将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			//3. 将购物项切片设置给购物车cart
			cart.CartItems = cartItems
			//4. 将购物车保存到数据库
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将《" + book.Title + "》加入到购物车中"))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}
}

//GetCartInfo 根据用户ID来获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	_, session := dao.IsLogin(r)
	//获取用户ID
	userID := session.UserID
	fmt.Println("当前用户ID：", userID)
	//根据用户的ID从数据库中获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		session.Cart = cart
		//解析模板文件
		t, err := template.ParseFiles("views/pages/cart/cart.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		//执行
		err2 := t.Execute(w, session)
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	} else {
		//改用户还没有购物车
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}
}

//DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除购物车的id
	cartID := r.FormValue("cartId")
	//清空购物车
	dao.DeleteCartByCartID(cartID)
	//调用GetCartInfo查询购物车信息
	GetCartInfo(w, r)
}

//DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除购物项的id
	cartItemID := r.FormValue("cartItemId")
	//购物项Id转换int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取改用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//获取购物车中的购物项
	cartItems := cart.CartItems
	//遍历得到的每个购物项
	for k, v := range cartItems {
		//寻找要删除的购物项
		if v.CartItemID == iCartItemID {
			//这个就是要删除的购物项
			//将当前购物项从切片中移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除之后的再次赋值给cart中的切片
			cart.CartItems = cartItems
			//将当前购物项从数据库中删除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	//更新购物车中的总数量，总金额
	dao.UpdateCart(cart)
	//获取购物车信息的函数，再次查询
	GetCartInfo(w, r)
}

//UpdateCartItem
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新购物项的id
	cartItemID := r.FormValue("cartItemId")
	//购物项Id转换int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取改用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//获取购物车中的购物项
	cartItems := cart.CartItems
	//遍历得到的每个购物项
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemID == iCartItemID {
			//这个就是要更新的购物项
			//将当前购物项中的图书的数量设置为用户输入的值
			v.Count = iBookCount
			//更新数据库中改购物项的数量和金额小计
			dao.UpdateBookCount(v)
		}
	}
	//更新购物车中的总数量，总金额
	dao.UpdateCart(cart)

	////获取购物车信息的函数，再次查询
	//GetCartInfo(w, r)

	//再次查询购物车
	cart, _ = dao.GetCartByUserID(userID)
	fmt.Println("购物车中的信息：", cart)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	var amount float64
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			//这个就是要找的购物项，获取当前购物项的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	////声明一个String切片
	//var data []string
	////将总数量，总金额，金额小计，转换为字符串
	//strTotalCount := strconv.FormatInt(totalCount, 10)
	//strTotalAmount := strconv.FormatFloat(totalAmount, 'f', 2, 64)
	//strAmount := strconv.FormatFloat(amount, 'f', 2, 64)
	//data = append(data, strTotalCount)
	//data = append(data, strTotalAmount)
	//data = append(data, strAmount)
	//将cart转换为Json
	json, _ := json2.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
