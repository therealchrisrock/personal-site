package routes

import (
	"fmt"
	"net/http"
	"xerus/internal/middleware"
	"xerus/internal/view/page"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	middleware.Chain(w, r, page.Home("Chris Rock | Home"))
}
