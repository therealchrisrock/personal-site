package routes

import (
	"net/http"
	"xerus/internal/middleware"
	"xerus/internal/view"
	"xerus/internal/view/page"
)

func Clouds(w http.ResponseWriter, r *http.Request) {
	head := view.DefaultHead()
	head.PageInfo.Title = "Cloud Animation"
	head.PageInfo.Description = "Interactive Three.js cloud animation"
	middleware.Chain(w, r, page.Clouds(head))
}