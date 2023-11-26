package renderer

import (
	"bytes"
	"html/template"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/models"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for thetemplate package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(tmpl_data *models.TemplateData) *models.TemplateData {
	return tmpl_data
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl_name string, tmpl_data *models.TemplateData) {
	var tmpl_cache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tmpl_cache = app.TemplateCache
	} else {
		tmpl_cache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	tmpl, ok := tmpl_cache[tmpl_name]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	tmpl_data = AddDefaultData(tmpl_data)

	_ = tmpl.Execute(buf, tmpl_data)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from /templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl_set, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			tmpl_set, err = tmpl_set.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = tmpl_set
	}

	return myCache, nil
}
