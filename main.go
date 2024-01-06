package main

import (
	"fmt"
	"html/template"
	"net/http"
)

	type User struct{
		Name string
		Age uint16
		Money int16
		Avg_grades, Happines float64
	}
//
//	func (u User) getAllInfo () string{
//		return fmt.Sprintf("User name is %s -- %d -- %d", u.name, u.age, u.money)
//	}
//
//	func(u *User) setNewName(newName string){
//		u.name = newName
//	}
func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"BOb", 25, -50, 33.3, 0.5}
	//fmt.Fprintf(page, "Имя юзера ")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, bob)
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Страница контактов это...")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":4567", nil)
}
func main() {
	handleRequest()
	//var bob User
	//bob := User{name: "Bob", age: 25, money: -50, avg_grades: 33, happines: 0.5}

}
