package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  int
}

//var albums []Album
var db *sql.DB

func main() {
	// Capture connection properties.
	// Get a database handle.
	/*var err error
	db, err = sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS sample")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec("CREATE DATABASE sample")
	fmt.Println("created database")

	//Table creation
	_, err = db.Exec("USE sample")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB selected successfully..")
	}

	cre, err := db.Prepare("create table album(id int auto_increment not null,title varchar(50),artist varchar(50),price int,primary key(id));")
	if err != nil {
		log.Fatal(err)
	}

	_, err = cre.Exec()
	if err != nil {
		log.Fatal(err)
	}

	ins, _ := db.Prepare("insert into sample.album(title,artist,price)values('The natural','logesh',20),('The book','logesh',20),('Sleep','mash',750),('sed','hun',1000)")
	_, err = ins.Exec()
	if err != nil {
		panic(err)
	}*/

	albums, err := albumsByArtist(1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}

func albumsByArtist(id int) (Album, error) {
	var err error
	var c error
	db, err := sql.Open("mysql", "root:Root@12345@tcp(127.0.0.1:3306)/sample")
	var alb Album
	row, c := db.Query("Delete from sample.album WHERE id = ?", id)
	if c != nil {
		log.Fatal(c)
	}
	row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

	return alb, err
}
