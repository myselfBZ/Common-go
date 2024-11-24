package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware interface{
    Do(http.ResponseWriter, *http.Request)
}

type Logger struct{

}

func (s *Logger) Do(w http.ResponseWriter, r *http.Request) {
    log.Printf("{ ip: %s }", r.RemoteAddr)
}


func middlewareFunc(next http.Handler, middlwareFuncs...Middleware) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        for _, m := range middlwareFuncs{
            m.Do(w, r)
        }

        defer func(){
            log.Printf("Took: %f", time.Since(start).Seconds())
        }()
		next.ServeHTTP(w, r)
	})
}

