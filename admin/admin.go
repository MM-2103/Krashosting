package admin

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Title string
	// Add more variables to pass to your admin page here
}

func admin() {
	http.HandleFunc("/admin", AdminPageHandler)

	fs := http.FileServer(http.Dir("static/")) // Assuming you have a "static" directory for css, js, etc.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting web server on port :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// AdminPageHandler handles the admin page access.
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {

	// Check if user is authenticated here.
	// You can use session cookies, JWT, or other methods to check.

	// If not authenticated, redirect to login page or return a 403 Forbidden status
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	// return

	pageVars := PageVariables{
		Title: "Admin Dashboard",
	}

	tmpl, err := template.ParseFiles("templates/admin.html") // Parse the .html file
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = tmpl.Execute(w, pageVars) // Execute the template and pass the variables to the template
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
