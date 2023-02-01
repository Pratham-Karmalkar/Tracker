package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Pratham-Karmalkar/Tracker/pkg/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Transaction struct {
	TransactionID   string    `json:"transactionID"`
	UserID          string    `json:"userID"`
	Amount          int       `json:"amount"`
	TransactionDate time.Time `json:"transactionDate"`
	TransactionTime time.Time `json:"transactionTime"`
	Tag             string    `json:"tag"`
}

var tdb *sql.DB

func init() {

	config.Connect()
	tdb = config.GetDB()

}

//transaction done here
func (t *Transaction) DoTransaction(userid string) (*Transaction, error) {

	//user := User{}
	//t.UserID = user.ID
	//fmt.Println("userId: ", userid)

	_, err := tdb.Exec("Insert into transactions (userid_text  , amount , tags) values ($1, $2, $3)", userid, t.Amount, t.Tag)

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(t.Amount)
	return t, err

}

func (t *Transaction) TotalAmount(userid string) int {

	var totalAmount int

	if err := tdb.QueryRow("Select Sum(amount) from Transactions where userid_text = $1", userid).Scan(&totalAmount); err != nil {
		fmt.Println(err)
	}
	return totalAmount
}
