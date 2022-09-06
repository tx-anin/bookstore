package model

type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量
	TotalAmount float64     //购物车中图书的总金额
	UserID      int         //当前购物车所属哪个用户
	UserName    string      //用户名
}

//GetTotalCount 获取购物车中的图书总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中购物项切片
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取购物车中的图书总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中购物项切片
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
