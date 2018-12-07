package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("verifyCode: ", r.Form["verifycode"])
	}
}

func registe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("registe.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		fmt.Println("phoneNumber: ", r.Form["phone"])
		fmt.Println("Email: ", r.Form["email"])
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", login)
	http.HandleFunc("/registe", registe)
	http.ListenAndServeTLS(":443", "1_www.hawtech.cn_bundle.crt", "2_www.hawtech.cn.key", nil)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
