package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/rajath002/bookings/pkg/config"
	"github.com/rajath002/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplateDynamicCache(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	// Create a template cache
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateDynamicCache()
	}

	// get a requested template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template from the template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateDynamicCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		log.Println("Something went wrong! 98")
		log.Println(err, pages)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			log.Println("Something went wrong!108")
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			log.Println("Something went wrong!115")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println("Something went wrong!122")
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
