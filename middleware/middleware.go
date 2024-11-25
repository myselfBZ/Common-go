package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware interface{
    Do(http.ResponseWriter, *http.Request) (error, int)
}

type Logger struct{

}

func (s *Logger) Do(w http.ResponseWriter, r *http.Request) (error, int) {
    log.Printf("{ ip: %s }", r.RemoteAddr)
    return nil, 0
}


func middlewareFunc(next http.Handler, middlwareFuncs...Middleware) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        for _, m := range middlwareFuncs{
            // stop the chain if one of middleware functions return an error
            if err, s := m.Do(w, r); err != nil{
                http.Error(w, err.Error(), s)
                return 
            }
        }

        defer func(){
            log.Printf("Took: %f", time.Since(start).Seconds())
        }()
		next.ServeHTTP(w, r)
	})
}

