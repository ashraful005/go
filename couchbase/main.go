package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:m10203049@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()

	// insert, err := db.Query("INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Ashraful Alam', 'Tiger IT', 'emdashrafulalam@gmail.com', '1');")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	fmt.Println("database connected")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/request", request)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `Welcome`)

}

func features(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Fprintf(w, "Welcome")

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Fprintf(w, "Welcome")

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func request(w http.ResponseWriter, r *http.Request) {
	//method-01
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	//fmt.Println(name, company, email)

	//fmt.Fprintf(w, `received %s %s %s`, name, company, email)

	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"

	sql := fmt.Sprintf(qs, name, company, email)

	insert, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	//method-02
	//r.ParseForm()

	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }
	fmt.Fprintf(w, `received`)
}
