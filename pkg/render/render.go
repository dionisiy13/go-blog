package render

import (
	"bytes"
	"github.com/dionisiy13/go-web/pkg/config"
	"github.com/dionisiy13/go-web/pkg/models"
	"html/template"
	"log"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(tmpl string, data *models.TemplateData) string {
	// create a template cache
	app := config.GetAppConfig()

	if len(app.TemplateCache) < 1 {
		cache, err := CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
		app.TemplateCache = cache
	}

	t, ok := app.TemplateCache[tmpl]
	if !ok {
		log.Fatal("The template is not existed")
	}

	buf := &bytes.Buffer{}

	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}

	return buf.String()
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
