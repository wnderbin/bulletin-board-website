package main

import (
	"fmt"
	"log"
	"main/handlers"
	"net/http"
	"os"
	"path/filepath"
)

var PORT string = ":8080"

func main() {
	if len(os.Args) == 2 {
		PORT = ":" + os.Args[1]
	}
	fmt.Printf("[PORT] %s\n", PORT)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("..", "..", "web", "static")))))

	http.HandleFunc("/", handlers.NotFoundHandler)
	http.HandleFunc("/main/", handlers.MainPageHandler)
	http.HandleFunc("/bulletin/", handlers.BulletinHandler)
	http.HandleFunc("/bulletin/add/", handlers.AddFormHandler)
	http.HandleFunc("/bulletin/add/postform", handlers.GetFormHandler)
	http.HandleFunc("/bulletin/delete", handlers.DeleteBulletinHandler)
	http.HandleFunc("/bulletin/update", handlers.UpdateBulletinHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
