package database

import (
	"database/sql"
	"testing"

	"github.com/aluferraz/goexpert-ex3/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestListAllOrders() {
	target_id := "999999"
	repo := NewOrderRepository(suite.Db)
	_, err := suite.Db.Exec("INSERT INTO  orders ( id,price,tax,final_price ) VALUES (?,100,0,1000)", target_id)
	suite.NoError(err)
	orders, err := repo.ListAll()
	foundOrder := false
	for _, order := range orders {
		if order.ID == target_id {
			foundOrder = true
			break
		}
	}
	suite.Equal(foundOrder, true)
	_, err = suite.Db.Exec("DELETE FROM orders WHERE id = ?", target_id)
	suite.NoError(err)
	foundOrder = false
	orders, err = repo.ListAll()
	suite.NoError(err)

	for _, order := range orders {
		if order.ID == target_id {
			foundOrder = true
			break
		}
	}
	suite.Equal(foundOrder, false)
}
