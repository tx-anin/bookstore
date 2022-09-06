package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//sql
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//UpdateBookCount 根据图书ID 和购物车ID 以及图书数量更新购物项
func UpdateBookCount(cartItem *model.CartItem) error {
	//sql
	sqlStr := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookIDAndCartID 根据图书ID和购物车的ID获取对应购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {

	//写sql
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)

	//创建CartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//根据图书ID，查询图书信息
	book, _ := GetBookByID(bookID)
	cartItem.Book = book

	return cartItem, nil
}

//GetCartItemsByCartID 根据购物车的ID获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {

	//sql
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接受bookID
		var bookID string
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		//根据bookid获取图书信息
		book, _ := GetBookByID(bookID)
		//将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//DeleteCartItemsByCartID 根据购物车的ID删除所有购物项
func DeleteCartItemsByCartID(cartID string) error {
	//sql
	sqlStr := "delete from cart_items where cart_id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根据购物项的ID删除购物项
func DeleteCartItemByID(cartItemID string) error {
	//sql
	sqlStr := "delete from cart_items where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
