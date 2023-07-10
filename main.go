package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var r = chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		var msg = map[string]string{"message": "Hello, Railway!"}

		data, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	var port = envPortOr("3000")

	fmt.Println("starting server on port " + port[1:])
	log.Fatal(http.ListenAndServe(port, r))
}

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
