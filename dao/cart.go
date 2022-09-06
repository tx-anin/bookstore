package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, cartItem := range cartItems {
		//保存购物项插入到数据库
		AddCartItem(cartItem)
	}

	return nil
}

//GetCartByUserID 根据用户ID从数据库中查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	//写sql
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, userID)
	//创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//获取当前购物车里的购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//将所有购物项设置到购物车中
	cart.CartItems = cartItems

	return cart, nil
}

//UpdateCart 更新购物车的图书数量和金额
func UpdateCart(cart *model.Cart) error {
	//sql
	sqlStr := "update carts set total_count = ? , total_amount = ? where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartByCartID 根据购物车ID删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车之前，先删除所有购物项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where id = ?"
	_, err2 := utils.Db.Exec(sqlStr, cartID)
	if err2 != nil {
		return err2
	}
	return nil
}
