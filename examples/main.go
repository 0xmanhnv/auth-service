package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Start the server
	port := ":80"
	log.Printf("Starting server on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

	// http://127.0.0.1:8080/auth/telegram?id=833425787&first_name=Nupakachi&username=nup4k4ch1&photo_url=https%3A%2F%2Ft.me%2Fi%2Fuserpic%2F320%2FiDSvQlm3_AaTykJVsPbKxWPTGGdO-Hf6tPYCV9mpIAU.jpg&auth_date=1732209606&hash=af1c33913ea23abd7c540d8641335343313b8549e08b4ef45ffa2c25ba204596
}
