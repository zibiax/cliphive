package main

import (
    "log"
    "net/http"
    "strings"
    "flag"
    "fmt"
    "os"
)


type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}


func main() {
    // set port address to webserver. Default is 4000
    port := flag.String("port", ":4000", "HTTP network portess")

    flag.Parse()

    // Check if port number starts with ":"
    if !strings.HasPrefix(*port, ":") {
        *port = fmt.Sprintf(":%s", *port)
    }

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    mux := http.NewServeMux()


    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/clip/create", app.cliphiveCreate)
    mux.HandleFunc("/clip/view", app.cliphiveView)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

    // Http.Server struct, so that it uses same configuration that we set.
    srv := &http.Server{
        Addr: *port,
        ErrorLog: errorLog,
        Handler: mux,
    }

    infoLog.Printf("Starting server on %s", *port)

    // Error handling
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
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
