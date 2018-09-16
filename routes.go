package main

import (
	"github.com/gorilla/mux"
	"github.com/jaylevin/TMN-API/controllers"
	"github.com/jaylevin/TMN-API/database"
)

func initRoutes(metaDb *database.MetaDb, router *mux.Router) {
	router.Handle("/", metaDb.Handler(controllers.GetRoot)).Methods("GET")

	initUserRoutes(metaDb, router)
}

func initUserRoutes(metaDb *database.MetaDb, router *mux.Router) {
	router.Handle("/users/{id}", metaDb.Handler(controllers.ShowUser)).Methods("GET")
	router.Handle("/users", metaDb.Handler(controllers.GetUsers)).Methods("GET")
	router.Handle("/users", metaDb.Handler(controllers.CreateUser)).Methods("POST")
	router.Handle("/users/{id}", metaDb.Handler(controllers.UpdateUser)).Methods("PUT")
	router.Handle("/users/{id}", metaDb.Handler(controllers.DeleteUser)).Methods("DELETE")
}
