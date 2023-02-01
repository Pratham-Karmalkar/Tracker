package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Pratham-Karmalkar/Tracker/pkg/models"
	"github.com/Pratham-Karmalkar/Tracker/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

/*---------------------------------------------------------------------------*/

func DoTransaction(w http.ResponseWriter, r *http.Request) {

	//Transaction Model Call
	newTransaction := &models.Transaction{}
	utils.ParseBody(r, newTransaction)

	cookieResp, err := r.Cookie("bluemarker")

	if err != nil {
		fmt.Print("LOGIN TO ACCESS DATA")
	} else {
		mapClaim, status := extractToken(cookieResp.Value)
		if status != true {
			fmt.Println("JWT Status false ")
		}

		id := mapClaim["iss"].(string)

		transaction, err := newTransaction.DoTransaction(id)
		if err != nil {
			fmt.Println("Error on transaction")
		}

		res, _ := json.Marshal(transaction)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

/*---------------------------------------------------------------------------*/

func TotalAmount(w http.ResponseWriter, r *http.Request) {
	//Transaction Model Call
	newTransaction := &models.Transaction{}

	cookieResp, err := r.Cookie("bluemarker")
	if err != nil {
		fmt.Println("LOGIN TO ACCESS DATA")
	} else {

		mapClaim, status := extractToken(cookieResp.Value)
		if status != true {
			fmt.Println("JWT Status false")
		}

		id := mapClaim["iss"].(string)
		transaction := newTransaction.TotalAmount(id)
		res, _ := json.Marshal(transaction)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

/*---------------------------------------------------------------------------*/

//function to extract jwt form cookies
func extractToken(tokenStr string) (jwt.MapClaims, bool) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading env")
	}

	var secretKeyString = os.Getenv("SECRET")
	secret := []byte(secretKeyString)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println("Invalid JWT Token")
		return nil, false
	}
}
