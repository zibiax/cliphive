package main

import (
    "log"
    "net/http"
    "strings"
    "flag"
    "fmt"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}


func main() {
    // set port address to webserver. Default is 4000
    port := flag.String("port", ":4000", "HTTP network portess")

    dsn := flag.String("dsn", "web:pass@/cliphive?parseTime=true", "MySQL data source name")

    flag.Parse()

    // Check if port number starts with ":"
    if !strings.HasPrefix(*port, ":") {
        *port = fmt.Sprintf(":%s", *port)
    }

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }




    // Http.Server struct, so that it uses same configuration that we set.
    srv := &http.Server{
        Addr: *port,
        ErrorLog: errorLog,
        Handler: app.routes(),
    }

    infoLog.Printf("Starting server on %s", *port)

    // Error handling
    err := srv.ListenAndServe()
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

//This is run if there is a trailing '/', so that static files isn't accessed inappropriately. http not found is run
func neuter(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }
        next.ServeHTTP(w, r)
    })
}
