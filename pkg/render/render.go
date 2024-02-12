package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rajath002/bookings/pkg/config"
	"github.com/rajath002/bookings/pkg/models"
)

// -------------  Version #1 -----------
// below code is deprecated
func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template ", err)
		return
	}
}

// -------------  Version #2 -----------

var tc = make(map[string]*template.Template)

func RenderTemplateV1(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache

	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("Creating a template and adding to template")
		err = CreateTemplate(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("Using the template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

}

func CreateTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}

// ------------- New Optimized Version #3 -----------

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplateDynamicCache(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Create a template cache
	// tc, err := CreateTemplateDynamicCache()
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateDynamicCache()
	}

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// get a requested template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template from the template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
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
