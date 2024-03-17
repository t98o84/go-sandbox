package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title       string
	Content     sql.NullString
	PublishedAt sql.NullTime
}

func main() {
	db, err := gorm.Open(mysql.Open(getDbDsn()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&News{}); err != nil {
		panic("failed to migrate database")
	}

	// Register the route
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/news", func (w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			news := News{Title: "Test News"}
			db.Create(&news)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "test news was created!"}`))
			return
		}

		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowd"))
	})

	// Start the server
	http.ListenAndServe(fmt.Sprintf(":%s", getPort()), nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}



func getPort() string {
	if v, ok := os.LookupEnv("APP_PORT"); ok {
		return v
	}
	return "80"
}

func getDbDsn() string {
	dns := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	fmt.Println(dns)

	return dns
}
