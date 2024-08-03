package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rajath002/bookings/internal/config"
	"github.com/rajath002/bookings/internal/driver"
	"github.com/rajath002/bookings/internal/handlers"
	"github.com/rajath002/bookings/internal/helpers"
	"github.com/rajath002/bookings/internal/models"
	"github.com/rajath002/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("starting mail listener.")
	listenForMail()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what I'm going to put in session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})
	gob.Register(map[string]int{})

	// Create a channel and assign it to app reporsitory for external usage
	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	helpers.NewHelpers(&app)

	//read flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbhost := flag.String("dbhost", "localhost", "Database host")
	dbname := flag.String("dbname", "", "Database name")
	dbuser := flag.String("dbuser", "", "Database user")
	dbpass := flag.String("dbpass", "", "Database password")
	dbport := flag.String("dbport", "5432", "Database port")
	dbssl := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

	flag.Parse()

	if *dbname == "" || *dbuser == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	// change this to true when in Production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connect to Database
	log.Println("Connecting to Database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbhost, *dbport, *dbname, *dbuser, *dbpass, *dbssl)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to Database!")

	tc, err := render.CreateTemplateDynamicCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = true

	// db is the driver which we are passing to NewRepo, so it will be avaialble to all the handlers
	repo := handlers.NewRepo(&app, db)

	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	return db, nil
}
