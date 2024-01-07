package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name                 string `json:"name"`
	Age                  uint16 `json:"age"`
	Money                int16
	Avg_grades, Happines float64
}

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
	
	//var bob User
	//bob := User{name: "Bob", age: 25, money: -50, avg_grades: 33, happines: 0.5}
	db, err :=sql.Open("mysql","slide:slide@tcp(127.0.0.1:3306)/golang")
	if err !=nil {
		panic(err)
	}
	defer db.Close()
	//Установка базы
	// insert, err :=db.Query("INSERT INTO `user` (`name`,`age`) VALUES ('Evg',34)")
	// if err!=nil {
	// 	panic(err)
	// }
	// defer insert.Close()
	//handleRequest()
	//Выборка 
	res, err :=db.Query("SELECT `name`, `age` FROM `user`")
	if err != nil {
		
	}
	for res.Next(){
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
		
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}
	

}
