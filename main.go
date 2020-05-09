package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/syads321/starter/controller"
	"log"
	"net/http"
)

//Query sfjd
type Query struct {
	Query string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t Query
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	result := controller.ExecuteQuery(t.Query, request)

	json.NewEncoder(rw).Encode(result)
}

func main() {
	godotenv.Load()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	// http.HandleFunc("/tentang-pondok-tahfidz-pktabah", controller.About)
	// http.HandleFunc("/hubungi", controller.Contact)
	// http.HandleFunc("/notfound", controller.RenderNotFound)

	// http.HandleFunc("/kegiatan-pondok-tahfidz-pktabah", controller.Activities)

	// homepage := http.HandlerFunc(controller.Home)
	// http.Handle("/", middleware.HTTPNotFound(homepage))

	// homepage := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// })

	http.HandleFunc("/graphql", parseGhPost)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
