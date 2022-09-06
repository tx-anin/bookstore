package main

import (
	"bookstore/controller"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

//Logging 记录每个URL请求的执行时长
func Logging() Middleware {

	//创建中间件
	return func(f http.HandlerFunc) http.HandlerFunc {
		//创建一个新的handler包装http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			//中间件的处理逻辑
			start := time.Now()
			//time.Since 返回自 t 以来经过的时间。
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			//调用下一个中间件或者最终的handler处理程序
			f(w, r)
		}
	}
}

// Method 验证请求用的是否是指定的Http Method，不是则返回400 Bad Request
func Method(m string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

// Chain 把应用到http.HandlerFunc处理器的中间件
// 		 按照先后顺序和处理器本身链起来供http.HandleFunc调用
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {

	fmt.Println(os.Getwd())

	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./views/pages"))))

	//去首页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注销
	http.HandleFunc("/logout", controller.Logout)
	//去注册
	http.HandleFunc("/regist", controller.Regist)

	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)

	//获取带分页的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//获取带分页和价格区间的图书
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	//添加图书
	//http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新图书
	//http.HandleFunc("/updateBook", controller.UpdateBook)
	//更新 或 添加 图书 （合并）
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)

	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//去结帐
	http.HandleFunc("/checkout", controller.Checkout)

	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)
	//获取订单详情及订单所对应的所有订单项
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//获取用户所有订单项
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)
	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//确认收获
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.HandleFunc("/hhh", Chain(Hello, Method("GET"), Logging()))

	http.ListenAndServe(":8080", nil)
}
