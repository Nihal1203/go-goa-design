package main

import (
	"log"
	"net/http"

	goahttp "goa.design/goa/v3/http"

	hello "github.com/Nihal1203/go-goa-design/gen/hello"
	hellohttp "github.com/Nihal1203/go-goa-design/gen/http/hello/server"
	myhello "github.com/Nihal1203/go-goa-design/hello"

	userhttp "github.com/Nihal1203/go-goa-design/gen/http/user/server"
	user "github.com/Nihal1203/go-goa-design/gen/user"
	myuser "github.com/Nihal1203/go-goa-design/user"
)

func main() {
	mux := goahttp.NewMuxer()

	// ✅ CORS Middleware
	handler := corsMiddleware(mux)

	// HELLO
	helloSvc := &myhello.Service{}
	helloEndpoints := hello.NewEndpoints(helloSvc)
	helloHandler := hellohttp.New(helloEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	hellohttp.Mount(mux, helloHandler)

	// USER
	userSvc := &myuser.Service{}
	userEndpoints := user.NewEndpoints(userSvc)
	userHandler := userhttp.New(userEndpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	userhttp.Mount(mux, userHandler)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler)) // ✅ use handler not mux
}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
