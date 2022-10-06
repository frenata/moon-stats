package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/frenata/moon-stats/lib"
)

func main() {
	log_file := os.Args[1]
	games := lib.ReadLog(log_file)
	for _, game := range games {
		fmt.Println(game)

		bs, err := json.Marshal(game)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post("http://localhost:8080/game", "applicaton/json", bytes.NewBuffer(bs))
		log.Println(resp, err)
	}
}
