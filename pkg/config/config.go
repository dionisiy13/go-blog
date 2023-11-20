package config

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

var App *AppConfig

func GetAppConfig() *AppConfig {
	if App == nil {
		lock.Lock()
		defer lock.Unlock()
		if App == nil {
			fmt.Println("Creating single instance now.")

			App = &AppConfig{}

			// set up the session
			sessionManager := scs.New()
			sessionManager.Lifetime = 24 * time.Hour
			sessionManager.Cookie.Persist = true
			sessionManager.Cookie.SameSite = http.SameSiteLaxMode
			sessionManager.Cookie.Secure = false

			App.Session = sessionManager
			App.UseCache = true
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return App
}
