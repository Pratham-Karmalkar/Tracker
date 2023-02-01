package models

//User model and User functions are declared here

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Pratham-Karmalkar/Tracker/pkg/config"
	"github.com/Pratham-Karmalkar/Tracker/pkg/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var udb *sql.DB

type User struct {
	ID             string    `json:"id"`
	Fname          string    `json:"firstName"`
	Lname          string    `json:"lastName"`
	Email          string    `json:"email"`
	Password       string    `json:"pass"`
	TotalAmount    int       `json:"totalAmount"`
	IsActive       bool      `json:"isActive"`
	AccountCreated time.Time `json:"date"`
}

func init() {

	config.Connect()
	//udb.AutoMigrate(&User{})
	udb = config.GetDB()

}

//Creates New User on a sign in Request

func (u *User) CreateUser() (*User, int) {

	pass := &u.Password
	hash := utils.HashPassword(*pass)
	var errStat int

	if len(*pass) < 7 {
		errStat = 2
		fmt.Println("Small Password error")

	} else {
		//first name , lastname, hashed password and email are the only values passed here
		_, err := udb.Query("Insert into Users (id_bin  , isActive , accountCreated , firstName, lastName ,email , pass ) values (unhex(replace(uuid(),'-','')),true,curdate(),?,?,?,?)", u.Fname, u.Lname, u.Email, hash)

		if err != nil {
			errStat = 1
		}
	}

	return u, errStat
}

//Login user must generate a jwt token
func (u *User) LoginUser(email, pass string) (string, string, bool) {
	var emailFromDB, passFromDB, id string

	//For returning email  not found and and user should signup
	if err := udb.QueryRow("Select id_text , email,pass from Users where email = ? ", email).Scan(&id, &emailFromDB, &passFromDB); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
		}
	}
	pwd := utils.CheckHashPassword(pass, passFromDB)
	if email == emailFromDB && pwd == true {
		return id, emailFromDB, true
	} else {
		return "", "", false
	}

}

//Password is verified here
func (u *User) VerifyPassword(name, pass string) bool {
	var passFromDB string

	if err := udb.QueryRow("Select pass from Users where firstName = ? ", name).Scan(&passFromDB); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
		}
	}

	fmt.Println(passFromDB)
	torf := utils.CheckHashPassword(pass, passFromDB)

	return torf
}

/*func (u *User) GetUsers() []User {
	rows, err := udb.Query("Select  id_text, firstName,lastName,email,pass, totalAmount from users")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		user := User{}
		//err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.TotalAmount)
		if err := rows.Scan(&user.ID, &user.Fname, &user.Lname, &user.Email, &user.Password, &user.TotalAmount); err != nil {
			return nil
		}
		users = append(users, user)
	}
	return users
}*/
