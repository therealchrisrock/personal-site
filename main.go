package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"xerus/internal/routes"
)

func main() {
	r := mux.NewRouter()
	_ = godotenv.Load()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/favicon.ico", routes.ServeFavicon)
	r.PathPrefix("/static/").HandlerFunc(routes.ServeStaticFiles)

	r.HandleFunc("/", routes.Home)
	r.HandleFunc("/admin", routes.Admin)
	r.HandleFunc("/md/{title}", routes.Markdown)
	r.HandleFunc("/mantracker", routes.Mantracker)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

}

//
//func main() {
//	log.Println("App Started")
//
//	err := generate.GenerateMain()
//	if err != nil {
//		panic(err)
//	}
//
//	_ = godotenv.Load()
//	r := mux.NewRouter()
//	r.HandleFunc("GET /", routes.Home)
//
//	r.HandleFunc("GET /favicon.ico", routes.ServeFavicon)
//	r.HandleFunc("GET /static/", routes.ServeStaticFiles)
//
//	r.HandleFunc("GET /mantracker", routes.Mantracker)
//	fmt.Printf("server is running on port %s\n", os.Getenv("PORT"))
//	server := &http.Server{
//		Handler:      r,
//		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//	log.Fatal(server.ListenAndServe())
//}
