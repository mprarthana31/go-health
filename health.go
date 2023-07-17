package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Port = os.Getenv("PORT")

var BuildTimestamp string

type Response struct {
	GitSHA  string
	BuildAt string
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		res := &Response{
			GitSHA:  os.Getenv("RENDER_GIT_COMMIT"),
			BuildAt: BuildTimestamp,
		}
		e, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, string(e))
	})

	fmt.Printf("start web server at port %s\n", Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", Port), nil); err != nil {
		log.Fatal(err)
	}

}
