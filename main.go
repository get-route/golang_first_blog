package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)
type Article struct{
	Id uint16
	Title, Anons, FullText string
}
var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request){
	t,err :=template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err!=nil {
		fmt.Fprintf(w, err.Error())
	
	}
	db, err :=sql.Open("mysql","slide:slide@tcp(127.0.0.1:3306)/golang")
	if err !=nil {
		panic(err)
	}
	defer db.Close()
	//Выборка 
	res, err :=db.Query("SELECT * FROM `articles`")
	if err != nil {
		
	}
	posts = []Article{}
	for res.Next(){
		var articles Article
		err = res.Scan(&articles.Id, &articles.Title,&articles.Anons,&articles.FullText)
		if err != nil {
	
			
		}
		
		posts = append(posts, articles)
		//fmt.Println(fmt.Sprintf("User: %s with age %d", articles.Title, articles.Id))
	}
	t.ExecuteTemplate(w, "index", posts)
}
func create(w http.ResponseWriter, r *http.Request){
	t,err :=template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err!=nil {
		fmt.Fprintf(w, err.Error())
	
	}
	t.ExecuteTemplate(w, "create", nil)
}
func save_article(w http.ResponseWriter, r *http.Request){
	  
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	if title =="" || anons =="" || full_text =="" {
		fmt.Fprintf(w,"Данные заполнены некорректно")
	}else{
		db, err :=sql.Open("mysql","slide:slide@tcp(127.0.0.1:3306)/golang")
	if err !=nil {
		panic(err)
	}
	defer db.Close()

	insert, err :=db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`,`anons`, `full_text`) VALUES ('%s','%s','%s')", title, anons, full_text))
	if err!=nil {
		panic(err)
	}
	defer insert.Close()
	http.Redirect(w,r,"/", http.StatusSeeOther)	
	}

}
func show_post(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	t,err :=template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err!=nil {
		fmt.Fprintf(w, err.Error())
	
	}
	db, err :=sql.Open("mysql","slide:slide@tcp(127.0.0.1:3306)/golang")
	if err !=nil {
		panic(err)
	}
	defer db.Close()
	//Выборка 
	res, err :=db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		
	}
	showPost = Article{}
	for res.Next(){
		var articles Article
		err = res.Scan(&articles.Id, &articles.Title,&articles.Anons,&articles.FullText)
		if err != nil {
			panic(err)
		}
		showPost = articles
		//fmt.Println(fmt.Sprintf("User: %s with age %d", articles.Title, articles.Id))
	}
	t.ExecuteTemplate(w, "show", showPost)

}
func handleFunc(){
	rtr :=mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	http.Handle("/",rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))	
	http.ListenAndServe(":7878", nil)
}

func main(){
	handleFunc()
}