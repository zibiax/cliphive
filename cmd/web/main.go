package main

import (
    "log"
    "net/http"
    "strings"
    "flag"
    "fmt"
    "os"
    "database/sql"
    "html/template"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/go-playground/form"
    "github.com/zibiax/cliphive/internal/models"
    "github.com/alexedwards/scs/v2"
    "github.com/alexedwards/scs/mysqlstore"
)


type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
    clip *models.ClipModel
    templateCache map[string]*template.Template
    formDecoder *form.Decoder
    sessionManager *scs.SessionManager
}


func main() {
    // set port address to webserver. Default is 4000
    port := flag.String("port", "4000", "HTTP network portess")
    dbUser := flag.String("dbuser", "web", "Database user")
    dbName := flag.String("dbname", "cliphive", "Database name")

    dbPass := os.Getenv("DB_PASSWORD")
    if dbPass == "" {
        log.Fatal("DB_PASSWORD environment variable not set")
    }

    flag.Parse()

    dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", *dbUser, dbPass, *dbName,)

    // Check if port number starts with ":"
    if !strings.HasPrefix(*port, ":") {
        *port = fmt.Sprintf(":%s", *port)
    }

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()

    templateCache, err := newTemplateCache()
    if err != nil {
        errorLog.Fatal(err)
    }


    formDecoder :=  form.NewDecoder()

    sessionManager := scs.New()
    sessionManager.Store = mysqlstore.New(db)
    sessionManager.Lifetime = 12 * time.Hour

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        clip: &models.ClipModel{DB: db},
        templateCache: templateCache,
        formDecoder: formDecoder,
        sessionManager: sessionManager,
    }


    // Http.Server struct, so that it uses same configuration that we set.
    srv := &http.Server{
        Addr: *port,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }

    infoLog.Printf("Starting server on %s", *port)

    // Error handling
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
