package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
	//name := r.FormValue("name")
	//company := r.FormValue("company")
	//email := r.FormValue("email")

	//fmt.Println(name, company, email)

	//fmt.Fprintf(w, `received %s %s %s`, name, company, email)

	//method-02
	r.ParseForm()

	for key, val := range r.Form {
		fmt.Println(key, val)
	}
	fmt.Fprintf(w, `received`)
}
