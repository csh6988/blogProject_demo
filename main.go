package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/**
 * @Author: cheney
 * @Date: 2022/11/13 3:09 PM
 * @Desc:
 */

type indexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	index := indexData{
		Title: "go项目",
		Desc:  "第一次来咯",
	}
	jsonStr, _ := json.Marshal(index)
	w.Write(jsonStr)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
