package main

import (
	"fmt"
	"github.com/claudioontheweb/go-web-structure/middleware"
	"github.com/claudioontheweb/go-web-structure/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {

	// ================================================================
	// Get Config variables
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.user")
	dbPass := viper.GetString("db.pass")
	dbName := viper.GetString("db.name")
	dbType := viper.GetString("db.type")

	// ================================================================
	// Connect to Repo

	var userRepo user.UserRepository

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	switch dbType {
	case "mysql":
		userRepo = user.NewMysqlUserRepository(connectMysql(connection))

	default:
		panic("Unknown database")

	}


	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// ================================================================
	// START HTTP SERVER

	r := mux.NewRouter().StrictSlash(true)

	// Init Middleware
	mw := middleware.InitMiddleware()
	r.Use(mw.CORS)
	r.Use(mw.COMMON)

	// Init Endpoints
	r.HandleFunc("/user", userHandler.FindAll).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.Find).Methods("GET")
	r.HandleFunc("/user", userHandler.Create).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.Delete).Methods("DELETE")

	// Start Server
	fmt.Println("Server listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	return nil
}

// ================================================================
// If dbType = "myql"
func connectMysql(connection string) *gorm.DB {
	db, err := gorm.Open(`mysql`, connection)
	if err != nil {
		panic(err)
	}

	return db
}