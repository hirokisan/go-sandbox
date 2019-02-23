package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func hello_world(w http.ResponseWriter, r *http.Request) {
	key := "black"
	e := `"` + key + `"`
	w.Header().Set("Etag", e)
	//w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	w.Header().Set("Access-Control-Allow-Origin", "*")
	now := time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")
	w.Header().Set("Content-Type", "application/json")

	sinceText := r.Header.Get("If-Modified-Since")
	if sinceText != "" {
		log.Println("Last-Modified")
	} else {
		w.Header().Set("Last-Modified", now)
	}

	if match := r.Header.Get("If-None-Match"); match != "" {
		log.Println("called")
		if strings.Contains(match, e) {
			log.Println("matched")
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	log.Println("world")
	ping := Ping{http.StatusOK, "ok"}
	res, _ := json.Marshal(ping)
	w.Write(res)
}

func main() {
	http.HandleFunc("/", hello_world)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
