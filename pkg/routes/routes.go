package routes

import (
	"github.com/Pratham-Karmalkar/Tracker/pkg/controllers"
	"github.com/gorilla/mux"
)

var TrackerRoutes = func(router *mux.Router) {
	//router.HandleFunc("/login/{user}", controllers.LoginHandler).Methods("GET")
	router.HandleFunc("/sign-up/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login/user", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/transaction", controllers.DoTransaction).Methods("POST")
	router.HandleFunc("/totalamount", controllers.TotalAmount).Methods("GET")

	//router.HandleFunc("/history/{user}", controllers.TransactionHistory).Methods("POST")
	//router.HandleFunc("/logout/{user}", controllers.LogoutUser).Methods("POST")

}
