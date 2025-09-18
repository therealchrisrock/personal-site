package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"xerus/internal/middleware"
	"xerus/internal/view"
	"xerus/internal/view/page"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	head := view.DefaultHead()
	head.PageInfo.Title = "Cloud Animation"
	head.PageInfo.Description = "Interactive Three.js cloud animation"
	head.Version = strconv.FormatInt(time.Now().Unix(), 10)
	middleware.Chain(w, r, page.Clouds(head))
}
