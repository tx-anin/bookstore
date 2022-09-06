package model

import "fmt"

//CartItem 购物项结构图
type CartItem struct {
	CartItemID int64   //购物项ID
	Book       *Book   //购物项中的图书信息
	Count      int64   //购物项中的图书数量
	Amount     float64 //购物项中图书的金额小计
	CartID     string  //当前购物项属于那个购物车 uuid生成
}

//GetAmount 获取购物项中的图书的金额小计
func (cartItem *CartItem) GetAmount() float64 {
	//获取当前购物项中的图书价格
	fmt.Println("购物项中获取的图书", cartItem.Book)
	price := cartItem.Book.Price
	fmt.Println("获取的图书价格：", price)
	return float64(cartItem.Count) * price
}
