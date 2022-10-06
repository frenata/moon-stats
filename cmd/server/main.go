package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/frenata/moon-stats/lib"
)

type App struct {
	db *sql.DB
}

func (app *App) InjectGameRecord(w http.ResponseWriter, req *http.Request) {
	var game lib.Game
	err := json.NewDecoder(req.Body).Decode(&game)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(game)

	app.upsertPlayer(game.Player.Id, game.Player.Name)
	app.upsertGame(game.Player.Id, game.Id, game.Victory(), game.Units)
}

func main() {
	app := App{db: boot()}
	http.HandleFunc("/game", app.InjectGameRecord)

	http.ListenAndServe(":8080", nil)
}
