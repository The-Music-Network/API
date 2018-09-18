package main

import (
	"github.com/gorilla/mux"

	// inner package dependencies
	"github.com/The-Music-Network/TMN-API/middleware"
	"github.com/The-Music-Network/TMN-API/components/users"
)

func initRoutes(metaDb *middleware.MetaDb, router *mux.Router) {
	initUserRoutes(metaDb, router)
}

func initUserRoutes(metaDb *middleware.MetaDb, router *mux.Router) {
	router.Handle("/users/{id}", metaDb.Handler(users.ShowUser)).Methods("GET")
	router.Handle("/users", metaDb.Handler(users.GetUsers)).Methods("GET")
	router.Handle("/users", metaDb.Handler(users.CreateUser)).Methods("POST")
	router.Handle("/users/{id}", metaDb.Handler(users.UpdateUser)).Methods("PUT")
	router.Handle("/users/{id}", metaDb.Handler(users.DeleteUser)).Methods("DELETE")
}
