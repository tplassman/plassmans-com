package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

func getEnv() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		panic(err)
	}
	log.Println("Loaded environment file")

	return env
}

func getTemplate(name string, fm template.FuncMap) *template.Template {
	funcMap := template.FuncMap{
		"now": func() int {
			return time.Now().Year()
		},
	}
	// Merge custom funcMap
	for k, v := range fm {
		funcMap[k] = v
	}

	t, err := template.New("main.html").Funcs(funcMap).ParseFiles(
		"app/templates/_layouts/main.html",
		"app/templates/_meta/data.html",
		"app/templates/_meta/favicons.html",
		fmt.Sprintf("app/templates/%s.html", name),
	)
	if err != nil {
		panic(err)
	}
	log.Printf("Parsed template: %s", name)

	return t
}

func handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := "index"
		title := "plassmans.com"

		// Handle 404 routes
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			name = "404"
			title = "Uh oh"
		}

		t := getTemplate(name, nil)

		viewData := struct {
			Title string
			Env   map[string]string
		}{title, getEnv()}
		t.Execute(w, viewData)
	}
}

func handleTrevor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := getTemplate("trevor", nil)

		type workProject struct {
			Title       string
			URL         string
			Description string
		}
		type data struct {
			WorkProjects []workProject
		}
		var projects data
		if _, err := toml.DecodeFile("app/data/trevor.toml", &projects); err != nil {
			fmt.Println(err)
			return
		}

		viewData := struct {
			Title        string
			Env          map[string]string
			WorkProjects []workProject
		}{"Trevor", getEnv(), projects.WorkProjects}
		t.Execute(w, viewData)
	}
}

func main() {
	// Routes
	http.HandleFunc("/", handleIndex())
	http.HandleFunc("/trevor", handleTrevor())
	// Start server
	http.ListenAndServe(":8080", nil)
}
