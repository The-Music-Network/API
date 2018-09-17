package main

import (
	/* Inner package dependencies */
	"github.com/jaylevin/TMN-API/config"
	"github.com/jaylevin/TMN-API/controllers/middleware"
	"github.com/jaylevin/TMN-API/models"

	/* Golang SDK packages */
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	/* Remote package dependencies */
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	conf, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading configuration file: ", err)
	}

	var db *gorm.DB
	switch conf.DbDriver {
	case "sqlite":
		db, err = gorm.Open("sqlite3", conf.SQLite.FilePath)
	case "mysql":
		db, err = gorm.Open("mysql", conf.MySQL.ConnectionString())
	}

	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	defer db.Close()

	if db == nil {
		log.Fatal(`Invalid server configuration! Check value of DbDriver value in tmn-api.json!`)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// MetaDB is passed to our route controllers for database manipulation on http request.
	metaDb := &middleware.MetaDb{Db: db, Recorder: httptest.NewRecorder()}

	// Creates tables based off of the User & Track structs in the 'models' package
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Track{})
	r := mux.NewRouter()

	// Initialize our routes, and set up handler functions (controllers)
	initRoutes(metaDb, r)

	// Bind to a port and pass our router in
	log.Print("TMN-API is running and using sql driver:" + conf.DbDriver)

	addr := fmt.Sprintf(":%d", 8080)
	log.Fatal(http.ListenAndServe(addr, r))
}

func loadConfig() (*config.Config, error) {
	/* Command line arguments */
	var (
		init       *bool   = flag.Bool("init", false, "Set the init flag to true to generate the default config.")
		configPath *string = flag.String("config", "tmn-api.json", "The path of the config file you want to use. If the file does not exist, the default will be created at that path.")
	)
	flag.Parse()

	if *init {
		log.Println("A new config is being generated for you.")
		err := config.WriteDefault(*configPath)
		if err != nil {
			return nil, errors.New("An error was encountered while generating your configuration: " + err.Error())
		}

		log.Fatal("Please restart with the new configuration file.")
	}

	conf, err := config.Load(*configPath)
	if err != nil {
		return nil, errors.New("Error encountered loading config file: " + err.Error())
	} else if conf == nil {
		return nil, errors.New("Config file not found! Generate a new one by setting -init=true command line flag!")
	}

	log.Printf("Using config file: %s", *configPath)
	config.Log(conf)

	return conf, nil
}
