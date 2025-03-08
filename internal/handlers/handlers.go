package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"handlers/database"
)

type BulletinForPostform struct {
	Title       string
	Id          int
	Name        string
	Description string
	Price       string
	Contacts    string
}

type Bulletin struct {
	Id          int
	Name        string
	Description string
	Price       string
	Contacts    string
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "main_page.html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	title := []string{"Main page"}
	err = tmpl.ExecuteTemplate(w, "main_page.html", title[0])
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "not_found.html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	title := []string{"Not Found :("}
	err = tmpl.ExecuteTemplate(w, "not_found.html", title[0])
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func AddFormHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "form.html"),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	title := []string{"Add"}
	err = tmpl.ExecuteTemplate(w, "form.html", title[0])
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func GetFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	contacts := r.FormValue("contacts")

	data := BulletinForPostform{
		Title:       "Send form",
		Name:        name,
		Description: description,
		Price:       price,
		Contacts:    contacts,
	}

	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "send_form.html"),
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		log.Println("Template parsing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = database.CreateTable("bulletins", []string{"name", "description", "price", "contacts"}, filepath.Join("..", "..", "internal", "database", "db", "sqlite3.db"))

	if err != nil {
		log.Println("Database table creation error:", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = database.AddToDB("bulletins", []string{"name", "description", "price", "contacts"}, []string{name, description, price, contacts}, filepath.Join("..", "..", "internal", "database", "db", "sqlite3.db"))

	if err != nil {
		log.Println("Database addition error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "send_form.html", data)
	if err != nil {
		log.Println("Template execution error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func DeleteBulletinHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "delete_bulletin.html"),
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Atoi error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = database.DeleteFromDB("bulletins", id, filepath.Join("..", "..", "internal", "database", "db", "sqlite3.db"))
	if err != nil {
		log.Println("Database removing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Template parsing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	data := []string{"Remove bulletin"}

	err = tmpl.ExecuteTemplate(w, "delete_bulletin.html", data[0])
	if err != nil {
		log.Println("Template executing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func BulletinHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "bulletins.html"),
	}

	data, err := database.GetDataFromDB("bulletins", filepath.Join("..", "..", "internal", "database", "db", "sqlite3.db"))
	if err != nil {
		log.Println("Database getting error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Template parsing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	var bulletins []Bulletin
	for _, d := range data {
		bulletins = append(bulletins, Bulletin{
			Id:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			Price:       d.Price,
			Contacts:    d.Contacts,
		})
	}

	err = tmpl.ExecuteTemplate(w, "bulletins.html", map[string]interface{}{"bulletins": bulletins})
	if err != nil {
		log.Println("Template executing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func UpdateBulletinHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		filepath.Join("..", "..", "web", "templates", "layout", "layout.html"),
		filepath.Join("..", "..", "web", "templates", "update_bulletin.html"),
	}
	new_data := r.URL.Query().Get("update")
	field := r.URL.Query().Get("field")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Atoi error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	field = strings.TrimSpace(strings.ToLower(field))

	err = database.UpdateDB("bulletins", new_data, field, id, filepath.Join("..", "..", "internal", "database", "db", "sqlite3.db"))
	if err != nil {
		log.Println("Database updating error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("template parsing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	title := []string{"Update"}

	err = tmpl.ExecuteTemplate(w, "update_bulletin.html", title[0])
	if err != nil {
		log.Println("Template executing error: ", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
