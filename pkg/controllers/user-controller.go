package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Pratham-Karmalkar/Tracker/pkg/models"
	"github.com/Pratham-Karmalkar/Tracker/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var NewUser models.User

// User to be redirected to login/signup page
//JWT token not genereated here.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	//userExists := mewUser.FindUserByEmail()
	userCreate, stat := newUser.CreateUser()

	if stat == 1 {
		str := "Username already exists"
		resp, _ := json.Marshal(str)
		w.Write(resp)
	} else if stat == 2 {
		str := "Password must more than 7 characters long"
		resp, _ := json.Marshal(str)
		w.Write(resp)
	} else {
		res, _ := json.Marshal(userCreate)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

//

//JWT token is generated in this function
func LoginUser(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading env")
	}

	var SecretKey = os.Getenv("SECRET")
	user := &models.User{}
	utils.ParseBody(r, user)

	id, email, password := user.LoginUser(user.Email, user.Password)

	if email != "" && password != false {

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    id,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := claims.SignedString([]byte(SecretKey))
		if err != nil {
			fmt.Println(err.Error())
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "bluemarker",
			Value:    token,
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 24),
			Path:     "/",
		})

	} else {
		errStr := "Wrong Email or Password"
		fmt.Println(errStr)
	}

}

// func GetAUser(w http.ResponseWriter, r *http.Request) {
// 	newUser := &models.User{}
// 	utils.ParseBody(r, newUser)
// 	vars := mux.Vars(r)
// 	val := vars["user"]
// 	val = string(val)
// 	//userName := newUser.FindUserId(val)
// 	fmt.Println(newUser.Password)
// 	verPwd := newUser.VerifyPassword(val, newUser.Password)
// 	res, _ := json.Marshal(verPwd)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }
