package main

import (
	"fmt"
	"net/http"

	"github.com/ermurachi/web-dev-calhoun/controllers"
	"github.com/ermurachi/web-dev-calhoun/templates"
	"github.com/ermurachi/web-dev-calhoun/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not Found", http.StatusNotFound) })

	fmt.Println("web server is starting on port :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
