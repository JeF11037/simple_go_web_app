package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/JeF11037/simple_web_goapp/database"
	"github.com/JeF11037/simple_web_goapp/models"
)

type PageData struct {
	Title string
	Users []models.User
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var paths []string
	for _, tmpl := range []string{
		"page.tmpl",
		"head.html",
		"header.html",
		"content.html",
		"footer.html"} {
		paths = append(paths, "./front/templates/"+tmpl)
	}
	template := template.Must(template.ParseFiles(paths...))
	users := database.GetUsers()
	data := PageData{
		Title: "GoApp",
		Users: users,
	}
	template.ExecuteTemplate(w, "page", data)
}

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.BirthDate = r.FormValue("birth_date")
	user.Gender = r.FormValue("gender")
	user.Email = r.FormValue("email")
	user.Address = r.FormValue("address")
	switch r.FormValue("action") {
	case "add":
		database.AddUser(user)
	case "update":
		var err error
		user.ID, err = strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err)
		}
		database.UpdateUser(user)
	case "delete":
		var err error
		user.ID, err = strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err)
		}
		database.DeleteUser(user)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	if database.Ping() {
		if !database.CheckUserTable() {
			database.CreateUserTable()
			database.FillSeedData()
		}
		http.HandleFunc("/", IndexHandler)
		http.HandleFunc("/action", ActionHandler)
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("front/css"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("front/js"))))
		http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("front/icons"))))
		http.Handle("/translation/", http.StripPrefix("/translation/", http.FileServer(http.Dir("front/translation"))))
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}

}
