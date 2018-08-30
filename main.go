package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	err = client.Set("visits", 0, 0).Err()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visits, err := client.Get("visits").Result()
		if err != nil {
			panic(err)
		}

		numOfVisits, err := strconv.ParseInt(visits, 10, 64)
		if err != nil {
			panic(err)
		}

		err = client.Set("visits", numOfVisits+1, 0).Err()
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, "Visits: ", visits)
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	log.Fatal(http.ListenAndServe(":80", nil))
}
