package model

//OrderItem 订单项
type OrderItem struct {
	OrderItemID int64   //订单项ID
	Count       int64   //订单项中图书的数量
	Amount      float64 //订单项中的金额小计
	Title       string  //订单项中图书的书名
	Author      string  //作者
	Price       float64 //价格
	ImgPath     string  //封面
	OrderID     string  //订单项所属订单
}
