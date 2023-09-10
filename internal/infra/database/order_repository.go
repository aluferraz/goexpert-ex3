package database

import (
	"database/sql"
	"fmt"

	"github.com/aluferraz/goexpert-ex3/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) ListAll() ([]entity.Order, error) {
	var orders []entity.Order
	rows, err := r.Db.Query("Select * from orders")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order entity.Order
		fmt.Println(order.ID)
		if err := rows.Scan(&order.ID, &order.Tax, &order.Price, &order.FinalPrice); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
