package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "redis-server:6379")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = conn.Cmd("HMSET", "0", "visits", 0).Err
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visits, err := conn.Cmd("HGET", "0", "visits").Str()
		if err != nil {
			panic(err)
		}

		numOfVisits, err := strconv.ParseInt(visits, 10, 64)
		if err != nil {
			panic(err)
		}

		err = conn.Cmd("HMSET", "0", "visits", numOfVisits+1).Err
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, "Visits: ", numOfVisits)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
