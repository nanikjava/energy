package main

import (
	"context"
	"github.com/nanikjava/cdsenergy/v1/storage"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"

	"github.com/nanikjava/cdsenergy/v1/impl"
	energyserver "github.com/nanikjava/cdsenergy/v1/pkg/pkg"
)

func main() {
	log.Printf("Server started")

	r := storage.NewRedisBackend(redis.UniversalOptions{
		Addrs: []string{"127.0.0.1:6379"},
	})

	seedData(r)
	if err := r.Ping(context.Background()); err != nil {
		log.Fatal("redis issues")
	}

	DefaultAPIService := impl.NewAPIService(r)
	DefaultAPIController := energyserver.NewDefaultAPIController(DefaultAPIService)

	router := energyserver.NewRouter(DefaultAPIController)

	// use our built-in middleware
	router.Use(logMW)
	// router.Use(authMW)

	log.Fatal(http.ListenAndServe(":8888", router))
}

func seedData(r *storage.RedisBackend) {
	r.Client.Set(context.TODO(), "test", "Nanik_Tolaram", -1)
}

func logMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Nanik %s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

		// compare the return-value to the authMW
		next.ServeHTTP(w, r)
	})
}

// func authMW(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// read basic auth information
// 		usr, _, ok := r.BasicAuth()

// 		// if there is no basic auth (no matter which credentials)
// 		if !ok {
// 			errMsg := "Authentication error!"
// 			// return a 403 forbidden
// 			http.Error(w, errMsg, http.StatusForbidden)
// 			log.Println(errMsg)

// 			// stop processing route
// 			return
// 		}

// 		// let's assume we check the user against a database to get
// 		// his admin-right and put this to the request context

// 		// else continue processing
// 		log.Printf("User %s logged in.", usr)
// 		next.ServeHTTP(w, r)
// 	})
// }
